.PHONY: server
server: ## run the server
	go run cmd/server/main.go

.PHONY: run
run: ## run the server with bazel
	bazel run //cmd/server:server

.PHONY: prod
prod: ## run the production server with bazel
	bazel run --action_env=GIN_MODE=release //cmd/server:server

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
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies

.PHONE: doc
doc: ## update swagger document
	swag init --parseDependency --parseInternal  -g cmd/server/main.go