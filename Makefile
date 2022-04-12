SRC = cmd/main.go

start: go-run

test: go-test go-tool
	cat cover | grep -v "mock" | grep -v  "easyjson" | grep -v "proto" > cover.out

.PHONY: start test

go-run: ${SRC}
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run ${SRC}

go-test: 
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -coverpkg=./... -coverprofile=cover ./...

go-tool:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go tool cover -func=cover.out

