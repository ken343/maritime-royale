##Local values pulled from var.json
locals {
  json_secrets= jsondecode(file("../../keys/creds.json"))
  json_data = jsondecode(file("../../var.json"))
}
#AWS Login Settings and Setup
provider "aws" {
  access_key = local.json_secrets.access_key
  secret_key = local.json_secrets.secret_key
  region     = "us-east-2"
}
##Image AMI
output "image_ami" {
  value = aws_ami_from_instance.image_ami.id
  description = "The Instance Amazon Image of the core server"
}
##Creates AMI 
resource "aws_ami_from_instance" "image_ami" {
  name               = "core_image"
  source_instance_id = local.json_data.imageid
}