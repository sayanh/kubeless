#!/usr/bin/env bash
make controller-image
kubectl apply -f kubeless-rbac-v0.2.4.yaml
kubectl delete -n kubeless po -l kubeless=controller
