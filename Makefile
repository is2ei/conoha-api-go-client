
# COLORS
GREEN = $(shell printf "\33[32m")
YELLOW = $(shell printf "\33[33m")
RESET = $(shell printf "\33[0m")

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