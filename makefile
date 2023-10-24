.PHONY: run
run:
	go run ./main.go

.PHONY: test
test:
	go test ./... -v --race

.PHONY: compile
compile:
	go generate ./...

.PHONY: graph
graph:
	go run -mod=mod ariga.io/entviz ./db/ent/schema
