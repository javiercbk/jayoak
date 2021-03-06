GIT_DATE := $(firstword $(shell git --no-pager show --date=short --format="%ai" --name-only))
GIT_VERSION := $(shell git rev-parse HEAD)
BIN_VERSION := $(GIT_VERSION)|$(GIT_DATE)
MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CUR_DIR := $(patsubst %/,%,$(dir $(MKFILE_PATH)))

.PHONY: server

# re-builds the sql boiler models
sql-boiler: clean-sql-boiler
	cd ~/go/bin;\
	~/go/bin/sqlboiler --struct-tag-casing camel --no-tests -c $(CUR_DIR)/database.toml -o $(CUR_DIR)/models psql

# removes the autogenerated sql boiler files
clean-sql-boiler:
	rm -rf models/*

# remove unused dependencies and tidy up modules
mod-tydy:
	go mod tidy

# lints the project
lint:
	~/go/bin/golangci-lint run

# builds the server binary
server: lint
	go build -tags=jsoniter -ldflags "-X http.response.jayOakVersion=$(BIN_VERSION)" cmd/server/server.go

# outputs the current version
version:
	@echo "$(BIN_VERSION)"

# run the audio frequencies storage in DB benchmarks
bench-sound:
	go test -benchmem -run=^$ github.com/javiercbk/jayoak/sound -bench ^Benchmark.*$ -benchtime=20s

frontend:
	cd frontend
	elm make src/Main.elm --output=dist/main.js

frontend-prod:
	cd frontend
	./optimize.sh src/Main.elm