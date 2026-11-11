#!/usr/bin/env bash
# TEMPORARY experiment: skaffold custom builder. Builds the single "gardener" binary via ko (identical layers for all
# components) and only differs in the image entrypoint. Usage: hack-build-gardener.sh <component>
set -euo pipefail
component="$1"
cd "$(dirname "$0")/.."
repo="$(echo "$IMAGE" | cut -d/ -f1)/gardener-single"

# Inject version info like the regular skaffold ko builds do ($LD_FLAGS is exported by the Makefile).
ld_flags="${LD_FLAGS:-$(hack/get-build-ld-flags.sh k8s.io/component-base VERSION Gardener)}"
ko_cfg="$(mktemp -d)"
trap 'rm -rf "$ko_cfg"' EXIT
{
  echo "builds:"
  echo "- id: gardener"
  echo "  main: ./cmd/gardener"
  echo "  ldflags:"
  for f in $ld_flags; do
    [[ "$f" == "-X" ]] && continue
    # fixed build date so that all invocations (operator/seed) produce byte-identical binaries
    [[ "$f" == *buildDate=* ]] && f="k8s.io/component-base/version.buildDate=2000-01-01T00:00:00+00:00"
    echo "  - -X $f"
  done
} > "$ko_cfg/.ko.yaml"

ref=$(KO_CONFIG_PATH="$ko_cfg" KO_DOCKER_REPO="$repo" ko build --bare --platform=linux/arm64 ./cmd/gardener 2>/dev/null | tail -1)
crane mutate --insecure "$ref" --entrypoint "/ko-app/gardener,$component" -t "$IMAGE" >/dev/null
