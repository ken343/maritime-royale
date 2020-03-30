#!/bin/bash

# Kubeadm prepares worker node
#passing the kube token to master setting up cluster
kubeadm join $(sudo jq -r '.masteripp' mastertoken.json) --token $(sudo jq -r '.token' mastertoken.json) \
	--discovery-token-ca-cert-hash $(sudo jq -r '.discovery_token' mastertoken.json)

