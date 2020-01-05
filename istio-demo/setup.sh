#!/bin/bash

CURRENT_DIR="$(cd "$(dirname $0)"; pwd)"
EXAMPLE_HOME=${CURRENT_DIR}../

docker build -f ${CURRENT_DIR}/projectA/Dockerfile -t project-a:1.0 ${EXAMPLE_HOME}
docker build -f ${CURRENT_DIR}/projectB/Dockerfile -t project-b:1.0 ${EXAMPLE_HOME}

kubectl apply -f ${CURRENT_DIR}/istio-demo.yaml
kubectl apply -f ${CURRENT_DIR}/istio-demo-gateway.yaml