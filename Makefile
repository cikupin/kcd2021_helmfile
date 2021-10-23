APP_NAME = kcd2021-helmfile

.PHONY: help
help:
	@echo 'Commands for ${APP_NAME}:'
	@echo
	@echo '    make docker-ctx-minikube              Use minikube as docker environment.'
	@echo '    make docker-ctx-local                 Use local computer as docker environment.'
	@echo

.PHONY: docker-env-minikube
docker-env-minikube:
	eval $(minikube docker-env)

.PHONY: docker-env-local
docker-env-local:
	eval $(minikube docker-env -u)
