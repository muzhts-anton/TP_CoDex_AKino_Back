SRC_MAIN = cmd/main.go
SRC_AUTMCS = cmd/authorization/auth.go
SRC_COMMCS = cmd/comment/comt.go
SRC_RATMCS = cmd/rating/rtng.go

server: go-run
autmcs: go-run-autmcs
commcs: go-run-commcs
ratmcs: go-run-ratmcs

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

go-run-ratmcs: ${SRC_RATMCS}
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run ${SRC_RATMCS}

go-test: 
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -coverpkg=./... -coverprofile=cover ./...

go-tool:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go tool cover -func=cover.out

# Команда для запуска тестов: go test -coverpkg=./... -coverprofile=cover ./... && cat cover | grep -v "mock" | grep -v  "easyjson" | grep -v "proto" > cover.out && go tool cover -func=cover.out
