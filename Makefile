SRC_MAIN = cmd/main.go
SRC_AUTMCS = cmd/authorization/auth.go
SRC_COMMCS = cmd/comment/comt.go

server: go-run
autmcs: go-run-autmcs
commcs: go-run-commcs

test: go-test go-tool
	cat cover | grep -v "mock" | grep -v  "easyjson" | grep -v "proto" > cover.out

clean:
	rm -f /tmp/session_*

.PHONY: server test clean

go-run: ${SRC_MAIN}
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run ${SRC_MAIN}

go-run-autmcs: ${SRC_AUTMCS}
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run ${SRC_AUTMCS}

go-run-commcs: ${SRC_COMMCS}
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run ${SRC_COMMCS}

go-test: 
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -coverpkg=./... -coverprofile=cover ./...

go-tool:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go tool cover -func=cover.out

# Команда для запуска тестов: go test -coverpkg=./... -coverprofile=cover ./... && cat cover | grep -v "mock" | grep -v  "easyjson" | grep -v "proto" > cover.out && go tool cover -func=cover.out
