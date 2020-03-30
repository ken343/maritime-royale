##Local values pulled from var.json
locals {
  json_data = jsondecode(file("../config/var.json"))
  json_secrets= jsondecode(file("../config/keys/creds.json"))
}
##Public IPs for Master EC2, 1 ip.
output "master_ip" {
  value = aws_instance.master.public_ip
  description = "The Private IP address of the server instance"
}

##AWS Login Settings and Setup
provider "aws" {
  access_key = local.json_secrets.access_key
  secret_key = local.json_secrets.secret_key
  region     = "us-east-2"
}
##SSH LOGIN KEYS
resource "aws_key_pair" "deployer" {
  ##key_name	  = "Key_Master"
  public_key	= file("../config/keys/public.pub")
}
##EC2for MASTER 
resource "aws_instance" "master" {
  key_name = aws_key_pair.deployer.key_name
  ami           = local.json_data.imageami
  instance_type = "t2.medium"
  security_groups = [aws_security_group.SSH.name]
  connection {
    user = "ubuntu"
    type = "ssh"
    private_key = file("../config/keys/private.pem")
    host =  self.public_ip
    timeout = "4m"
  }
##Setup Directories for Master
  provisioner "remote-exec" {
    inline = [
    "mkdir -p terradir/keys",
    "mkdir pkg",
    "mkdir cmd",

    ]
  }
##Core Script
  provisioner "file" {
    source      = "../config/scripts/core.sh"
    destination = "/tmp/core.sh"
  }
##Master Script
  provisioner "file" {
    source      = "../config/scripts/master.sh"
    destination = "/tmp/master.sh"
  }
##Slave main Script 
  provisioner "file" {
    source      = "../config/scripts/worker.sh"
    destination = "/home/ubuntu/terradir/worker.sh"
  }

##Terraform worker tf file
  provisioner "file" {
    source      = "../worker/worker.tf"
    destination = "/home/ubuntu/terradir/worker.tf"
  }

##Place creds, and keys into secrets directory
  provisioner "file" {
    source      = "../config/keys/"
    destination = "/home/ubuntu/terradir/keys"
  }
##Place varraibles json into terradir directory
   provisioner "file" {
    source      = "../config/var.json"
    destination = "/home/ubuntu/terradir/var.json"
  }
  ##Place varraibles json into terradir directory
   provisioner "file" {
    source      = "../config/terraform"
    destination = "/home/ubuntu/terraform"
  }

##Exicute Script
  provisioner "remote-exec" {
    inline = [
      "sudo /bin/bash /tmp/master.sh"
    ]
  } 
}
##Secuirty Group Allow SSH
resource "aws_security_group" "SSH" {

  description = "Allow SSH traffic"
  ingress {
    from_port   = 0 
    to_port     = 0
    protocol =   "-1"

    cidr_blocks =  ["0.0.0.0/0"]
  }
  egress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    cidr_blocks     = ["0.0.0.0/0"]
  }
}

