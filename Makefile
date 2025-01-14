PWD           := $(shell pwd)
MYSELF        := $(shell id -u)
MY_GROUP      := $(shell id -g)
DEV_NAMESPACE := dev

.PHONY: ensure_yamlfmt_installed
ensure_yamlfmt_installed:
ifeq (, $(shell which yamlfmt))
	$(shell go install github.com/google/yamlfmt/cmd/yamlfmt@latest)
endif

.PHONY: setup
setup: ensure_yamlfmt_installed
	(cd backend && $(MAKE) setup)
	(cd frontend && $(MAKE) setup)

.PHONY: format
format: format_yaml
	(cd backend && $(MAKE) format)
	(cd frontend && $(MAKE) format)

.PHONY: terraformat
terraformat:
	(cd backend && $(MAKE) terraformat)
	(cd frontend && $(MAKE) terraformat)

.PHONY: format_yaml
format_yaml: ensure_yamlfmt_installed
	yamlfmt -conf .yamlfmt.yaml

.PHONY: lint
lint:
	(cd backend && $(MAKE) lint)
	(cd frontend && $(MAKE) lint)

.PHONY: test
test:
	(cd backend && $(MAKE) test)
	(cd frontend && $(MAKE) test)

.PHONY: openapi-clients
openapi-clients:
	(cd backend && $(MAKE) openapi-client)
	(cd frontend && $(MAKE) openapi-client)

.PHONY: openapi-lint
openapi-lint:
	npx @stoplight/spectral-cli@v6.13.1 lint openapi_spec.yamls

.PHONY: regit
regit:
	cd ../
	git clone git@github.com:dinnerdonebetter/dinnerdonebetter tempdir
	@if [ -n "$(BRANCH)" ]; then \
	  (cd tempdir && git checkout $(BRANCH)); \
	fi
	cp -rf tempdir/.git .
	rm -rf tempdir

.PHONY: deploy_dev
deploy_dev:
	kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.16.2/cert-manager.yaml
	skaffold run --filename=skaffold.yaml --build-concurrency 1 --profile $(DEV_NAMESPACE)

.PHONY: nuke_dev
nuke_dev:
	kubectl delete deployments,cronjobs,configmaps,services,secrets --namespace $(DEV_NAMESPACE) --selector='managed_by!=terraform'
