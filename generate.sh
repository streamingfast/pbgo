#!/bin/bash
# Copyright 2019 dfuse Platform Inc.
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

ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Protobuf definitions
PROTO=${PROTO:-"$ROOT/../proto"}

function main() {
  checks

  current_dir="`pwd`"
  trap "cd \"$current_dir\"" EXIT
  pushd "$ROOT/pb" &> /dev/null

  generate "dfuse/blockmeta/v1/blockmeta.proto"
  generate "dfuse/bstream/v1/bstream.proto"
  generate "dfuse/fluxdb/v1/fluxdb.proto"
  generate "dfuse/graphql/v1/graphql.proto"
  generate "dfuse/headinfo/v1/headinfo.proto"
  generate "dfuse/merger/v1/merger.proto"
  generate "dfuse/search/v1/search.proto"
  generate "grpc/health/v1/health.proto"

  echo "generate.sh - `date` - `whoami`" > $ROOT/last_generate.txt
  echo "streamingfast/proto revision: `GIT_DIR=$PROTO/.git git rev-parse HEAD`" >> $ROOT/last_generate.txt
}

function generate() {
    protoc -I$PROTO $1 --go_out=plugins=grpc,paths=source_relative:.
}

function checks() {
  result_1_4_0_and_later=`printf "" | protoc-gen-go --version 2>&1 | grep -Eo 'unknown argument'`
  if [[ $result_1_4_0_and_later == "unknown argument" ]]; then
     # We are using github.com/golang/protobuf/protoc-gen-go@v1.4.0+ it's the correct version we want here
     return
  fi

  # The old `protoc-gen-go` did not accept any flags. Just using `protoc-gen-go --version` in this
  # version waits forever. So we pipe some wrong input to make it exit fast. This in the new version
  # which supports `--version` correctly print the version anyway and discard the standard input
  # so it's good with both version.
  result_1_3_5_and_older=`printf "" | protoc-gen-go --version 2>&1 | grep -Eo v[0-9\.]+`
  if [[ "$result_1_3_5_and_older" == "" ]]; then
    echo "Your version of 'protoc-gen-go' (at `which protoc-gen-go`) is not recent enough."
    echo ""
    echo "To fix your problem, perform those commands:"
    echo ""
    echo "  pushd /tmp"
    echo "    go install github.com/golang/protobuf/protoc-gen-go@v1.5.2"
    echo "  popd"
    echo ""
    echo "If everything is working as expetcted, the command:"
    echo ""
    echo "  protoc-gen-go --version"
    echo ""
    echo "Should print 'protoc-gen-go: unknown argument "--version" (this program should be run by protoc, not directly)'"
    exit 1
  fi

  if [[ "$result_1_3_5_and_older" != "" ]]; then
    echo "Your version of 'protoc-gen-go' is **too** recent!"
    echo ""
    echo "This repository requires a strict gRPC version not higher than v1.29.1 however"
    echo "the newer protoc-gen-go versions generates code compatible with v1.32 at the minimum."
    echo ""
    echo "To keep the compatibility until the transitive dependency TiKV is updated (through streamingfast/kvdb)"
    echo "you must ue the older package which is hosted at 'github.com/golang/protobuf/protoc-gen-go' (you most"
    echo "probably have 'google.golang.org/protobuf/cmd/protoc-gen-go')."
    echo ""
    echo "To fix your problem, perform those commands:"
    echo ""
    echo "  pushd /tmp"
    echo "    go install github.com/golang/protobuf/protoc-gen-go@v1.5.2"
    echo "  popd"
    echo ""
    echo "If everything is working as expected, the command:"
    echo ""
    echo "  protoc-gen-go --version"
    echo ""
    echo "Should print 'protoc-gen-go: unknown argument "--version" (this program should be run by protoc, not directly)'"
    exit 1
  fi
}


main "$@"