OS = $(shell uname)

# COLORS
RED    = $(shell printf "\33[31m")
GREEN  = $(shell printf "\33[32m")
WHITE  = $(shell printf "\33[37m")
YELLOW = $(shell printf "\33[33m")
RESET  = $(shell printf "\33[0m")

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  HELP
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

.DEFAULT: help

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  SETUP
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

install:
	@echo "${YELLOW}Installing dependencies${RESET}"
	go list -f '{{range .Imports}}{{.}} {{end}}' ./... | xargs go get ${install_flags}
	@echo "${GREEN}✔ successfully installed dependencies${RESET}\n"

update:
	@echo "${YELLOW}Updating dependencies${RESET}"
	go list -f '{{range .Imports}}{{.}} {{end}}' ./... | xargs go get -u ${install_flags}
	@echo "${GREEN}✔ successfully updated dependencies${RESET}\n"

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  TEST
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

test_lib: 
	@echo "${YELLOW}Running lib tests${RESET}"
	@go test -v ./conoha/.
	@echo "${GREEN}✔ Lib tests successfully passed${RESET}\n"

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  CLI
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

cli_build:
	@echo "${YELLOW}Building cli${RESET}"
	@cd cli && go build -o conoha
	@echo "${GREEN}✔ successfully built CLI${RESET}\n"

cli_build_all:
	@echo "${YELLOW}Building CLI for various platforms${RESET}"
	cd cli && GOOS=darwin GOARCH=amd64 go build -o build/conoha-darwin-amd64
	cd cli && GOOS=linux  GOARCH=amd64 go build -o build/conoha-linux-amd64
	cd cli && GOOS=linux  GOARCH=386   go build -o build/conohalinux-386
	cd cli && GOOS=linux  GOARCH=arm   go build -o build/conoha-linux-arm
	cd cli && GOOS=linux  GOARCH=arm64 go build -o build/conoha-linux-arm64
	@echo "${GREEN}✔ successfully built CLI flavors${RESET}\n"