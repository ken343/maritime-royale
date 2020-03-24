#Creates an ami for our master, worker.

cd ../image
terraform init
terraform apply --auto-approve
echo "$(jq ".imageid=$(terraform output -json image_id)" ../var.json)" > ../var.json
cd ami
terraform init
terraform apply --auto-approve
echo "$(jq ".imageami=$(terraform output -json image_ami)" ../../var.json)" > ../../var.json
cd ..
terraform destroy --auto-approve



# base ubuntu ami
# {
#      "ami" :"ami-0fc20dd1da406780b"
# }