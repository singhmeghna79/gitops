#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

CODEGEN_PKG=${GOPATH}/src/github.com/mayadata-io/gittrack

${CODEGEN_PKG}/vendor/k8s.io/code-generator/generate-groups.sh all \
  github.com/mayadata-io/gittrack/pkg/client github.com/mayadata-io/gittrack/pkg/apis \
  mayadata.io:v1alpha1 
