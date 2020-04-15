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

current_dir="`pwd`"
trap "cd \"$current_dir\"" EXIT
pushd "$ROOT" &> /dev/null

# Service definitions
SERVICES=${1:-../service-definitions}

protoc -I$SERVICES grpc/health/v1/health.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/blockmeta/v1/blockmeta.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/bstream/v1/bstream.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/graphql/v1/graphql.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/headinfo/v1/headinfo.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/merger/v1/merger.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/search/v1/search.proto --go_out=plugins=grpc,paths=source_relative:.

echo "generate.sh - `date` - `whoami`" > last_generate.txt
echo -n "service-definitions revision: " >> last_generate.txt
GIT_DIR=$SERVICES/.git  git rev-parse HEAD >> last_generate.txt
