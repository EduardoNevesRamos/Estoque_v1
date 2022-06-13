install:
	@echo "${COLOR_YELLOW}Downloading dependencies...${COLOR_WHITE}"
	@go get
	@echo "${COLOR_YELLOW}Validating dependencies...${COLOR_WHITE}"
	@go mod tidy
	@echo "${COLOR_YELLOW}Creating vendor...${COLOR_WHITE}"
	@go mod vendor
	@echo "${COLOR_YELLOW}Installation completed successfully.${COLOR_WHITE}"

test:
	@echo "Running project tests..."
	@go test -v -cover ./...
	@echo "Running project tests..."


run:
	@echo "${COLOR_YELLOW}Running Application...${COLOR_WHITE}"
	@go run main.go
	@echo "${COLOR_GREEN}[COMPILED]${COLOR_WHITE}"

coverage:
	@echo "${COLOR_YELLOW}Running project coverage...${COLOR_WHITE}\n"
	@go test ./... -v -coverprofile=cover.out
	@go tool cover -html=cover.out -o cover.html
	@echo "${COLOR_GREEN}Coverage completed successfully.${COLOR_WHITE}"