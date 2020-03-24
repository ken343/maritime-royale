# Kubeadm initialization
# Flannel uses 10.244.0.0/16 as the pod network CIDR
kubeadm init --pod-network-cidr=10.244.0.0/16

# To make kubectl work for non-root user
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# Container Network Interface (CNI) Flannel installation
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/2140ac876ef134e0ed5af15c65e414cf26827915/Documentation/kube-flannel.yml

# IPtables configuration required by Container Network Interface (CNI), Flannel
sysctl net.bridge.bridge-nf-call-iptables=1

# Set environment variables
export kubever=$(kubectl version | base64 | tr -d '\n')

var10=$(sudo kubeadm token create --print-join-command)

arrvar10=(${var10// / })

masteripp=$(echo ${arrvar10[2]})
token=$(echo ${arrvar10[4]})
discovery_token=$(echo ${arrvar10[6]})
json_data="{
      \"masteripp\" : \"${masteripp}\",
      \"token\" : \"${token}\",
      \"discovery_token\": \"${discovery_token}\"
      }"

touch /home/ubuntu/terradir/mastertoken.json
sudo chmod 777 /home/ubuntu/terradir/mastertoken.json 
echo $json_data | cat > /home/ubuntu/terradir/mastertoken.json

# IPtables setting
iptables -P FORWARD ACCEPT
sudo sysctl net.bridge.bridge-nf-call-iptables=1

#Terraform install
sudo chmod 777 terraform
sudo mv -f terraform /bin

#Create Worker
cd terradir
terraform init
terraform apply --auto-approve