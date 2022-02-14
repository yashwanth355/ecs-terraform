# ecs-terraform

terraform init

terraform plan -var "environment=dev" --var-file .\environments\dev.tfvars

terraform apply -var "environment=dev" --var-file .\environments\dev.tfvars --auto-approve

terraform destroy -var "environment=dev" --var-file .\environments\dev.tfvars --auto-approve
