#!/bin/bash

set -x

GOPATH=$(go env GOPATH)
PACKAGE_NAME=kubeform.dev/kubeform
REPO_ROOT="$GOPATH/src/$PACKAGE_NAME"
DOCKER_REPO_ROOT="/go/src/$PACKAGE_NAME"
DOCKER_CODEGEN_PKG="/go/src/k8s.io/code-generator"

pushd $REPO_ROOT

goimports -w $REPO_ROOT/apis

for provider in $(find $REPO_ROOT/apis -maxdepth 1 -mindepth 1 -type d -printf '%f '); do
    rm -rf "$REPO_ROOT"/apis/${provider}/v1alpha1/*.generated.go
done
mkdir -p "$REPO_ROOT"/config/api-rules
mkdir -p "$REPO_ROOT"/config/crd

apiGroups=$(find $REPO_ROOT/apis -maxdepth 1 -mindepth 1 -type d -printf '%f:v1alpha1 ')
docker run --rm -ti -u $(id -u):$(id -g) \
  -v "$REPO_ROOT":"$DOCKER_REPO_ROOT" \
  -w "$DOCKER_REPO_ROOT" \
  appscode/gengo:release-1.14 "$DOCKER_CODEGEN_PKG"/generate-groups.sh "deepcopy,client,informer,lister" \
  kubeform.dev/kubeform/client \
  kubeform.dev/kubeform/apis \
  "${apiGroups}" \
  --go-header-file "$DOCKER_REPO_ROOT/hack/gengo/boilerplate.go.txt"

# Generate openapi
for gv in $(find $REPO_ROOT/apis -maxdepth 1 -mindepth 1 -type d -printf '%f/v1alpha1 '); do
  docker run --rm -ti -u $(id -u):$(id -g) \
    -v "$REPO_ROOT":"$DOCKER_REPO_ROOT" \
    -w "$DOCKER_REPO_ROOT" \
    appscode/gengo:release-1.14 openapi-gen \
    --v 1 --logtostderr \
    --go-header-file "hack/gengo/boilerplate.go.txt" \
    --input-dirs "$PACKAGE_NAME/apis/${gv},k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/api/resource,k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/util/intstr,k8s.io/apimachinery/pkg/version,k8s.io/api/core/v1,k8s.io/api/apps/v1,k8s.io/api/rbac/v1" \
    --output-package "$PACKAGE_NAME/apis/${gv}" \
    --report-filename config/api-rules/violation_exceptions.list
done

# Generate crds.yamls
docker run --rm -ti -u $(id -u):$(id -g) \
  -v /tmp:/.cache \
  -v "$REPO_ROOT":"$DOCKER_REPO_ROOT" \
  -w "$DOCKER_REPO_ROOT" \
  appscode/gengo:release-1.14 controller-gen crd:trivialVersions=true paths="./apis/..." output:crd:artifacts:config=config/crd

popd
