#!/usr/bin/env bash
make controller-image
kubectl apply -f kubeless-rbac-v0.2.4-diff-namespace.yaml
kubectl delete -n kubeless-new po -l kubeless=controller
