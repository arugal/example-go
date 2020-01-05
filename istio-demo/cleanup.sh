#!/bin/bash

CURRENT_DIR="$(cd "$(dirname $0)"; pwd)"
EXAMPLE_HOME=${CURRENT_DIR}/../

kubectl delete -f ${CURRENT_DIR}/istio-demo.yaml
kubectl delete -f ${CURRENT_DIR}/istio-demo-gateway.yaml


docker rmi project-a:1.0
docker rmi project-b:1.0

docker system prune