.PHONY: server
server: ## run the server
	go run cmd/server/main.go

.PHONY: run
run: ## run the dev server with bazel
	export TARGET_ENV=DEV && bazel run //cmd/server:server

.PHONY: staging
staging: ## run the staging server with bazel
	export TARGET_ENV=STAGING && bazel run //cmd/server:server

.PHONY: prod
prod: ## run the prod server with bazel
	export TARGET_ENV=PROD && export GIN_MODE=release && bazel run //cmd/server:server

.PHONY: build
build: ## update dependency and build using bazel
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
	bazel run //:gazelle
	bazel build //...

.PHONY: test
test: ## run all test with bazel
	bazel test //...

.PHONY: clean
clean: ## use bazel clean to remove all bazel output folders
	bazel clean

.PHONY: gazelle
gazelle: ## run gazelle to add bazel to each directory
	bazel run //:gazelle

.PHONY: dependency
dependency: ## update all bazel file with necessary dependency
	go get -u github.com/Melon-Network-Inc/common
	go mod tidy
	bazel run //:gazelle -- update-repos -from_file=go.mod  -prune=true -build_file_proto_mode=disable_global -to_macro=deps.bzl%go_repositories

.PHONE: doc
doc: ## update swagger document
	python3 ../NestedJsonMerger/merge.py ../account-service/docs/swagger.json ../payment-service/docs/swagger.json ../gateway-service/docs/swagger.json