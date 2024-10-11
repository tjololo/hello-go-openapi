
SWAG ?= $(LOCALBIN)/swag
SWAG_VERSION ?= v1.16.3
OASDIFF ?= $(LOCALBIN)/oasdiff
OASDIFF_VERSION ?= v1.10.26

.PHONY: build
build: gen-swagger
	bazelisk build //...

.PHONY: run
run: gen-swagger
	bazelisk run //:app

.PHONY: gazelle
gazelle:
	bazelisk run //:gazelle

.PHONY: gen-swagger
gen-swagger: swag
	@$(SWAG) init > /dev/null

.PHONE: check-openapi
check-openapi: gen-swagger oasdiff
	@$(OASDIFF) breaking --fail-on WARN https://raw.githubusercontent.com/tjololo/hello-go-openapi/refs/heads/main/docs/swagger.yaml docs/swagger.yaml

LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

.PHONY: swag
swag: $(SWAG) ## Download controller-gen locally if necessary.
$(SWAG): $(LOCALBIN)
	$(call go-install-tool,$(SWAG),github.com/swaggo/swag/cmd/swag,$(SWAG_VERSION))


.PHONY: oasdiff
oasdiff: $(OASDIFF) ## Download controller-gen locally if necessary.
$(OASDIFF): $(LOCALBIN)
	$(call go-install-tool,$(OASDIFF),github.com/tufin/oasdiff,$(OASDIFF_VERSION))
# go-install-tool will 'go install' any package with custom target and name of binary, if it doesn't exist
# $1 - target path with name of binary
# $2 - package url which can be installed
# $3 - specific version of package
define go-install-tool
@[ -f "$(1)-$(3)" ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
rm -f $(1) || true ;\
GOBIN=$(LOCALBIN) go install $${package} ;\
mv $(1) $(1)-$(3) ;\
} ;\
ln -sf $(1)-$(3) $(1)
endef
