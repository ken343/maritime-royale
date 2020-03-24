#!/bin/bash

# This script will set up the core part for both master and worker node
# Downloads and Installs:
# jq 
# golang
# docker engine
# kubectl, kubelet, kubeadm
echo Setting up Kubernetes Node...

# Update the apt package index
apt-get update

# Docker installation per guidance from docs.docker.com
# Install packages to allow apt to use a repository over HTTPS
apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common 

sudo apt install -y jq
sudo snap install go --classic

# Add Docker’s official GPG key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# Set up the stable repository
add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

# Install the latest version of Docker Engine - Community and containerd
apt-get install -y docker-ce docker-ce-cli containerd.io

# Setup daemon
cat > /etc/docker/daemon.json <<EOF
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}
EOF

mkdir -p /etc/systemd/system/docker.service.d

# Restart docker
systemctl daemon-reload
systemctl restart docker

# Kubernetes installation per guidance from kubernetes.io
# Update repo list with kubernetes tools and install the 3 universal tools
# all nodes are expected to have.
sudo apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl

# Mark up kublet, kubeadm and kubectl to prevent from auto updating
sudo apt-mark hold kubelet kubeadm kubectl

# Restart the kubelet
systemctl daemon-reload
systemctl restart kubelet