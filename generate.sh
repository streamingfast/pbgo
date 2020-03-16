#!/bin/bash

ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

current_dir="`pwd`"
trap "cd \"$current_dir\"" EXIT
pushd "$ROOT" &> /dev/null

# Service definitions
SERVICES=${1:-../service-definitions}

protoc -I$SERVICES grpc/health/v1/health.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/funnel/v1/funnel.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/abicodec/eosio/v1/abicodec.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/abicodec/eth/v1/abicodec.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/blockmeta/v1/blockmeta.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/graphql/v1/graphql.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/headinfo/v1/headinfo.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/merger/v1/merger.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/search/v1/search.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/tokenmeta/v1/tokenmeta.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/trxstatetracker/v1/trxstatetracker.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/bstream/v1/bstream.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/codecs/deos/deos.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/codecs/deth/deth.proto --go_out=plugins=grpc,paths=source_relative:.
protoc -I$SERVICES dfuse/vms/geth/evm/v1/executor.proto --go_out=plugins=grpc,paths=source_relative:.

echo "generate.sh - `date` - `whoami`" > last_generate.txt
echo -n "service-definitions revision: " >> last_generate.txt
GIT_DIR=$SERVICES/.git  git rev-parse HEAD >> last_generate.txt
