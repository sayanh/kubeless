#!/usr/bin/env bash
make controller-image
kubectl delete -n kubeless po -l kubeless=controller
