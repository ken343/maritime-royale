image:
	cd ./build/config/scripts &&\
	cd ../image &&\
	terraform init &&\
	terraform apply --auto-approve &&\
	echo "$$(jq ".imageid=$$(terraform output -json image_id)" ../var.json)" > ../var.json &&\
	cd ami &&\
	terraform init &&\
	terraform apply --auto-approve &&\
	echo "$$(jq ".imageami=$$(terraform output -json image_ami)" ../../var.json)" > ../../var.json &&\
	cd .. &&\
	terraform destroy --auto-approve

master: 
	cd ./build/master && \
	terraform init && \
	terraform apply --auto-approve && \
	export masterip=$$(terraform output master_ip) && \
	echo "$$masterip" &&\
	ssh -i ../config/keys/private.pem ubuntu@$$masterip

ssh:
	cd ./build/master && \
	export masterip=$$(terraform output master_ip) && \
	ssh -i ../config/keys/private.pem ubuntu@$$masterip

destroy: ## Tear down whole dev env
	cd ./build/master && \
	export masterip=$$(terraform output master_ip) && \
	ssh -i ../config/keys/private.pem ubuntu@$$masterip 'cd terradir && terraform destroy --auto-approve'
	cd ./build/master && \
	terraform destroy --auto-approve
