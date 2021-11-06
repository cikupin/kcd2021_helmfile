APP_NAME              = kcd2021-helmfile
DOCKER_REGISTRY       = cikupin

.PHONY: help
help:
	@echo 'Commands for ${APP_NAME}:'
	@echo
	@echo '    make docker-ctx-minikube              Use minikube as docker environment.'
	@echo '    make docker-ctx-local                 Use local computer as docker environment.'
	@echo
	@echo '    make build                            Compile app'
	@echo '    make package                          Build docker image'
	@echo '    make push                             Push docker image to registry'
	@echo

.PHONY: docker-env-minikube
docker-env-minikube:
	eval $(minikube docker-env)

.PHONY: docker-env-local
docker-env-local:
	eval $(minikube docker-env -u)

.PHONY: build
build:
	@echo "Building ${APP_NAME}"
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-w" -o bin/${APP_NAME}

.PHONY: package
package: build
	@echo "Building image ${APP_NAME}"
	docker build -t ${DOCKER_REGISTRY}/${APP_NAME}:latest .

.PHONY: push
push: package
	@echo "Pushing Docker image ${APP_NAME} to registry"
	docker push ${DOCKER_REGISTRY}/${APP_NAME}:latest

.PHONY: run-vault
run-vault:
	@echo "Running vault"
	export VAULT_ADDR=http://127.0.0.1:8200
	export VAULT_TOKEN=toor
	docker run --cap-add=IPC_LOCK -d --name=dev-vault -p8200:8200 vault:1.8.4 server -dev -dev-root-token-id=toor
