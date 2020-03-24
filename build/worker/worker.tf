##Local values pulled from var.json
locals {
  json_data = jsondecode(file("./var.json"))
  json_secrets= jsondecode(file("./keys/creds.json"))
}
##Public IPs for WORKER EC2's, array of worker ips.
output "instance_ip" {
  value = aws_instance.worker.*.public_ip
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
  public_key	= file("./keys/public.pub")
}
##EC2's for SLAVES
resource "aws_instance" "worker" {
  count   = 1
  key_name = aws_key_pair.deployer.key_name
  ami           = local.json_data.imageami
  instance_type = "t2.micro"
  security_groups = [aws_security_group.SSH_slave.name]
  connection {
    user = "ubuntu"
    type = "ssh"
    private_key = file("./keys/private.pem")
    host =  self.public_ip
    timeout = "4m"
}
##Make file structure pods and services
    provisioner "remote-exec" {
    inline = [
      "mkdir pods",
      "mkdir services",      
    ]
  }

##Place scrpits in temp folder to run and delete after
  provisioner "file" {
    source      = "/home/ubuntu/terradir/worker.sh"
    destination = "/home/ubuntu/worker.sh"
  }
##Place varraibles json into terraform directory
   provisioner "file" {
    source      = "/home/ubuntu/terradir/mastertoken.json"
    destination = "/home/ubuntu/mastertoken.json"
  }

##Run scrpits for slave setup
  provisioner "remote-exec" {
    inline = [
      "sudo /bin/bash /home/ubuntu/worker.sh",
    ]
  }   
}
##Allows SSH
resource "aws_security_group" "SSH_slave" {
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
