#!/usr/bin/env bash

set -e

cd $(dirname $0)

PROTOC_VERSION="22.0"
PROTOC_OS="$(uname -s)"
PROTOC_ARCH="$(uname -m)"
case "${PROTOC_OS}" in
  Darwin) PROTOC_OS="osx" ;;
  Linux) PROTOC_OS="linux" ;;
  *)
    echo "Invalid value for uname -s: ${PROTOC_OS}" >&2
    exit 1
esac

# This is for macs with M1 chips. Precompiled binaries for osx/amd64 are not available for download, so for that case
# we download the x86_64 version instead. This will work as long as rosetta2 is installed.
if [ "$PROTOC_OS" = "osx" ] && [ "$PROTOC_ARCH" = "arm64" ]; then
  PROTOC_ARCH="x86_64"
fi

PROTOC="${PWD}/protoc/bin/protoc"

if [[ "$(${PROTOC} --version 2>/dev/null)" != "libprotoc 3.${PROTOC_VERSION}" ]]; then
  rm -rf ./protoc
  mkdir -p protoc
  curl -L "https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-${PROTOC_OS}-${PROTOC_ARCH}.zip" > protoc/protoc.zip
  cd ./protoc && unzip protoc.zip && cd ..
fi

go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
go install github.com/jhump/protoreflect/desc/sourceinfo/cmd/protoc-gen-gosrcinfo

# Output directory will effectively be GOPATH/src.
outdir="."
${PROTOC} "--go_out=paths=source_relative:$outdir" "--gosrcinfo_out=paths=source_relative,debug:$outdir" -I. *.proto
${PROTOC} "--go_out=paths=source_relative:$outdir" "--gosrcinfo_out=paths=source_relative,debug:$outdir" -I. nopkg/*.proto
${PROTOC} "--go_out=paths=source_relative:$outdir" "--gosrcinfo_out=paths=source_relative,debug:$outdir" -I. pkg/*.proto
${PROTOC} "--go_out=paths=source_relative:$outdir" "--gosrcinfo_out=paths=source_relative,debug:$outdir" -I. proto3_optional/desc_test_proto3_optional.proto
${PROTOC} "--go_out=paths=source_relative:$outdir" "--go-grpc_out=paths=source_relative:$outdir" "--gosrcinfo_out=paths=source_relative,debug:$outdir" -I. grpc/*.proto

# And make descriptor set (with source info) for several files
${PROTOC} --descriptor_set_out=./desc_test1.protoset --include_source_info --include_imports -I. desc_test1.proto
${PROTOC} --descriptor_set_out=./desc_test_comments.protoset --include_source_info --include_imports -I. desc_test_comments.proto
${PROTOC} --descriptor_set_out=./desc_test_complex.protoset -I. desc_test_complex.proto
${PROTOC} --descriptor_set_out=./desc_test_complex_source_info.protoset --include_source_info --include_imports -I. desc_test_complex.proto
${PROTOC} --descriptor_set_out=./descriptor.protoset --include_source_info --include_imports -I./protoc/include/ google/protobuf/descriptor.proto
${PROTOC} --descriptor_set_out=./duration.protoset -I./protoc/include/ google/protobuf/duration.proto
${PROTOC} --descriptor_set_out=./proto3_optional/desc_test_proto3_optional.protoset --include_source_info --include_imports -I. proto3_optional/desc_test_proto3_optional.proto

# The grpc library doesn't yet have generated code for v1 of the reflection service.
# https://github.com/grpc/grpc-go/issues/5684
# So, for now, we generate a copy into an internal package so that our client can support that version.
cd ../../
${PROTOC} --go_opt=Mgrpcreflect/internal/grpc_reflection_v1/reflection.proto=github.com/jhump/protoreflect/grpcreflect/internal/grpc_reflection_v1 \
    --go_opt=paths=source_relative \
    --go_out=. \
    --go-grpc_opt=Mgrpcreflect/internal/grpc_reflection_v1/reflection.proto=github.com/jhump/protoreflect/grpcreflect/internal/grpc_reflection_v1 \
    --go-grpc_opt=paths=source_relative \
    --go-grpc_out=. \
    grpcreflect/internal/grpc_reflection_v1/reflection.proto
# We've changed the package so the above generated code plays nicely if/when there is canonical
# generated code in the grpc repo. So now we have to change it back in the generated code so
# that the RPC service is advertised with the correct package name.
sed -I .bk 's/jhump\.protoreflect\.//g' grpcreflect/internal/grpc_reflection_v1/reflection_grpc.pb.go
rm grpcreflect/internal/grpc_reflection_v1/reflection_grpc.pb.go.bk
