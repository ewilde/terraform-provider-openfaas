## How to run this example?

1. Deploy faas using these [instructions](https://docs.openfaas.com/deployment/docker-swarm/) 

```bash
git clone https://github.com/openfaas/faas && \
  cd faas && \
  ./deploy_stack.sh
```

2. Save password from output and create a file `terraform.tfvars`:

```bash
cat <<EOF >>terraform.tfvars
provider_password = "72b97dd9abe096b91478df91b6549f2998aee27dbdde9a4bbec4182801d6c398"
EOF 
```

3. Clean up and remove the deployed stack and any running function (caution removes all running docker containers)
```bash
docker stack rm func
docker rm -f (docker ps -aq) 
docker secret remove basic-auth-user
```