##Local values pulled from var.json
locals {
  json_data = jsondecode(file("../var.json"))
  json_secrets= jsondecode(file("../keys/creds.json"))

}
##Public IPs for Master EC2, 1 ip.
output "image_ip" {
  value = aws_instance.image.public_ip
  description = "The Private IP address of the core server instance"
}
##Image ID 
output "image_id" {
  value = aws_instance.image.id
  description = "The Instance ID of the core server"
}

#AWS Login Settings and Setup
provider "aws" {
  access_key = local.json_secrets.access_key
  secret_key = local.json_secrets.secret_key
  region     = "us-east-2"
}
##SSH LOGIN KEYS
resource "aws_key_pair" "deployer" {
  key_name	  = "Key_Image"
  public_key	= file("../keys/public.pub")
}
##EC2for MASTER 
resource "aws_instance" "image" {
  key_name = aws_key_pair.deployer.key_name
  ##base ubuntu image
  ami           = local.json_data.ami
  instance_type = "t2.medium"
  connection {
    user = "ubuntu"
    type = "ssh"
    private_key = file("../keys/private.pem")
    host =  self.public_ip
    timeout = "1m"
  }
    ##Core Script
  provisioner "file" {
    source      = "../scripts/core.sh"
    destination = "/home/ubuntu/core.sh"
  }
    ##Exicute Script
  provisioner "remote-exec" {
    inline = [
      "sudo /bin/bash /home/ubuntu/core.sh",
    ]
  } 
}