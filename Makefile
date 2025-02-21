help: ## Help Dialog.
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -F -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'


install-requirements: ## Generate go.mod & go.sum Files. Also Install Additional Requirements
	go mod tidy
	go get github.com/walle/lll/...
	go install github.com/cespare/reflex@latest
	go install github.com/mdempsky/gocode@latest
	go install github.com/segmentio/golines@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

clean-packages: ## Clean Packages
	go clean -modcache


include makefiles/local.mk
