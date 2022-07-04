.PHONY: server
server: ## run the server
	go run cmd/server/main.go

.PHONY: run
run: ## run the server with bazel
	bazel run //cmd/server:main
	
.PHONY: build
build: ## update dependency and build using bazel
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
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
dependency: ## update all bazel file wtih necessary depedency
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies