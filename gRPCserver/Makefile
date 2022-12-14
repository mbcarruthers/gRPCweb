BIN_DIR = bin
PROTO_DIR = proto
SERVER_DIR = server
CLIENT_DIR = client

ifeq ($(OS), Windows_NT)
	SHELL := powershell.exe
	.SHELLFLAGS := -NoProfile -Command
	SHELL_VERSION = $(shell (Get-Host | Select-Object Version | Format-Table -HideTableHeaders | Out-String).Trim())
	OS = $(shell "{0} {1}" -f "windows", (Get-ComputerInfo -Property OsVersion, OsArchitecture | Format-Table -HideTableHeaders | Out-String).Trim())
	PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])
	CHECK_DIR_CMD = if (!(Test-Path $@)) { $$e = [char]27; Write-Error "$$e[31mDirectory $@ doesn't exist$${e}[0m" }
	HELP_CMD = Select-String "^[a-zA-Z_-]+:.*?\#\# .*$$" "./Makefile" | Foreach-Object { $$_data = $$_.matches -split ":.*?\#\# "; $$obj = New-Object PSCustomObject; Add-Member -InputObject $$obj -NotePropertyName ('Command') -NotePropertyValue $$_data[0]; Add-Member -InputObject $$obj -NotePropertyName ('Description') -NotePropertyValue $$_data[1]; $$obj } | Format-Table -HideTableHeaders @{Expression={ $$e = [char]27; "$$e[36m$$($$_.Command)$${e}[0m" }}, Description
	RM_F_CMD = Remove-Item -erroraction silentlycontinue -Force
	RM_RF_CMD = ${RM_F_CMD} -Recurse
	SERVER_BIN = ${SERVER_DIR}.exe
	CLIENT_BIN = ${CLIENT_DIR}.exe
else
	SHELL := bash
	SHELL_VERSION = $(shell echo $$BASH_VERSION)
	UNAME := $(shell uname -s)
	VERSION_AND_ARCH = $(shell uname -rm)
	ifeq ($(UNAME),Darwin)
		OS = macos ${VERSION_AND_ARCH}
	else ifeq ($(UNAME),Linux)
		OS = linux ${VERSION_AND_ARCH}
	else
    $(error OS not supported by this Makefile)
	endif
	PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
	CHECK_DIR_CMD = test -d $@ || (echo "\033[31mDirectory $@ doesn't exist\033[0m" && false)
	HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	RM_F_CMD = rm -f
	RM_RF_CMD = ${RM_F_CMD} -r
	SERVER_BIN = ${SERVER_DIR}
	CLIENT_BIN = ${CLIENT_DIR}
endif



project := gRPCserver


gRPCserver: $@help

build_protoc: ## build proto files
	protoc \
	-IgRPCserver/proto \
	--go_out=gRPCserver/proto \
	--go_opt=module=github.com/mbcarruthers/gRPCserver \
	--go-grpc_out=gRPCserver/proto \
	--go-grpc_opt=module=github.com/mbcarruthers/gRPCserver gRPCserver/proto/gRPCserver.proto

clean_protoc: ## clean grpc generated Go files
	rm $(project)/${PROTO_DIR}/*pb.go


build_protoc_grpc_web: ## build grpc-web protoc files for the client-side. ../app/src/web is the build directory.
	protoc \
    	-IgRPCserver/proto gRPCserver/proto/gRPCserver.proto \
    	--js_out=import_style=commonjs,binary:../app/src/Web/ \
    	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:../app/src/Web/

clean_web_grpc_protoc: ## clean Web directory of proto files
	rm ../app/src/Web/*

build_server: ## Builds server in bin/gRPCserver
	go build -o ${BIN_DIR}/$(project)/${SERVER_BIN} ./$(project)/${SERVER_DIR}

build_client: ## Builds client in bin/gRPCserver. Note: Client is just for testing at the moment
	go build -o ${BIN_DIR}/$(project)/${CLIENT_BIN} ./$(project)/${CLIENT_DIR}

about: ## Display info related to the build
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL} ${SHELL_VERSION}"
	@echo "Protoc version: $(shell protoc --version)"
	@echo "Go version: $(shell go version)"
	@echo "Go package: ${PACKAGE}"
	@echo "Openssl version: $(shell openssl version)"

help: ## Show this help
	@${HELP_CMD}
