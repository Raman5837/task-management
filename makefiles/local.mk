run-local: ## Run the app locally
	go run main.go

watch-local: ## Run app with auto reload
	reflex -s -r '\.go$$' make run-local
