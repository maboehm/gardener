#!/bin/bash
#
# Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e

usage() {
  echo "Usage:"
  echo "> compare-kcm-controllers.sh [ -h | <old version> <new version> ]"
  echo
  echo ">> For example: compare-kcm-controllers.sh 1.22 1.23"

  exit 0
}

if [ "$1" == "-h" ] || [ "$#" -ne 2 ]; then
  usage
fi

versions=("$1" "$2")

out_dir=dev/temp
mkdir -p "${out_dir}"

for version in "${versions[@]}"; do
  rm -rf "${out_dir}/kubernetes-${version}"
  rm -f "${out_dir}/kcm-controllers-${version}.txt"

  git clone --depth 1 --filter=blob:none --sparse https://github.com/kubernetes/kubernetes -b "release-${version}" "${out_dir}/kubernetes-${version}"
  pushd "${out_dir}/kubernetes-${version}" > /dev/null
  git sparse-checkout set "cmd/kube-controller-manager/app"
  popd > /dev/null

  cat "${out_dir}/kubernetes-${version}/cmd/kube-controller-manager/app/"*.go |\
    sed -rn "s/.*[Client|Config]OrDie\(\"(.*)\"\).*/\1/p" |\
    grep -vE "informers|discovery" |\
    sort |\
    uniq > "${out_dir}/kcm-controllers-${version}.txt"
done

echo
echo "kube-controller-manager controllers added in $2 compared to $1:"
diff "${out_dir}/kcm-controllers-$1.txt" "${out_dir}/kcm-controllers-$2.txt" | grep '>' | awk '{print $2}'
echo
echo "kube-controller-manager controllers removed in $2 compared to $1:"
diff "${out_dir}/kcm-controllers-$1.txt" "${out_dir}/kcm-controllers-$2.txt" | grep '<' | awk '{print $2}'
