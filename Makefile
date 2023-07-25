# Build configuration
# -------------------

APP_NAME = 221b
APP_VERSION = 0.0.3
GIT_REVISION = `git rev-parse HEAD`
# Introspection targets
# ---------------------

.PHONY: all
all: compile


.PHONY: compile
compile: ## compile the project
	@echo "[ *** build 221b binary *** ]"
	@go build -o $(APP_NAME) .
	@echo "[ *** build successful, use -h *** ]"

.PHONY: clean
clean: ## clean up the project directory
	@rm -vf $(APP_NAME)

