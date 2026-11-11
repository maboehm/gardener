#!/usr/bin/env bash
# TEMPORARY experiment: measures memory of the Gardener component processes running in the local kind node, to compare
# separate binaries vs. the single "gardener <component>" binary. Works in both states, as all component binaries live
# under /ko-app/. Usage: hack/measure-memory.sh [label]   (writes ./memory-<label>.txt in addition to stdout)
# Env: NODE (default gardener-local-control-plane)
set -euo pipefail

label="${1:-}"
node="${NODE:-gardener-local-control-plane}"

raw="$(docker exec -i "$node" bash -s <<'EOF'
for d in /proc/[0-9]*; do
  pid="${d#/proc/}"
  [ -r "$d/cmdline" ] && [ -r "$d/smaps_rollup" ] || continue
  mapfile -d '' -t args < "$d/cmdline" || continue
  exe="${args[0]:-}"
  [[ "$exe" == /ko-app/* ]] || continue
  name="$(basename "$exe")"
  [[ "$name" == gardener ]] && name="gardener-${args[1]:-}"
  vals="$(awk '/^(Rss|Pss|Pss_Anon|Pss_File):/ {printf "%s ", $2}' "$d/smaps_rollup")"
  echo "proc $pid $name $vals"
done
awk '/^(MemTotal|MemFree|Buffers|Cached|SReclaimable|MemAvailable):/ {print "meminfo", $1, $2}' /proc/meminfo
EOF
)"

report() {
  echo "# $(date -u +%FT%TZ) label=${label:-none}"
  echo "# in-scope components only (no -local, node-agent, gardenadm); values in kB"
  # fields from awk order in smaps_rollup: Rss Pss Pss_Anon Pss_File
  echo "$raw" | awk '
    $1=="proc" && $3 ~ /^(gardener-(operator|resource-manager|apiserver|controller-manager|scheduler|admission-controller|gardenlet)|gardenlet)$/ {
      printf "%-8s %-34s Rss=%-8s Pss=%-8s Pss_Anon=%-8s Pss_File=%-8s\n", $2, $3, $4, $5, $6, $7
      rss+=$4; pss+=$5; anon+=$6; file+=$7; n++
    }
    $1=="meminfo" {m[$2]=$3}
    END {
      printf "\nprocesses=%d\nSUM Rss=%d kB  Pss=%d kB  Pss_Anon=%d kB  Pss_File=%d kB\n", n, rss, pss, anon, file
      used = m["MemTotal:"] - m["MemFree:"] - m["Buffers:"] - m["Cached:"] - m["SReclaimable:"]
      printf "VM meminfo: MemTotal=%d kB MemFree=%d kB Cached=%d kB MemAvailable=%d kB used(excl. cache)=%d kB\n", m["MemTotal:"], m["MemFree:"], m["Cached:"], m["MemAvailable:"], used
    }'
}

if [[ -n "$label" ]]; then
  report | tee "memory-$label.txt"
else
  report
fi
