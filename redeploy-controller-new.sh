#!/usr/bin/env bash
make controller-image
kubectl apply -f kubeless-rbac-v0.3.0-diff-namespace.yaml
kubectl delete -n kubeless-new po -l kubeless=controller
