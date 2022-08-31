.PHONY: clean test

ehs:
	@go build -o ehs -ldflags "-w -s" cmd/main.go

clean:
	@rm ehs

test:
	@go test -v
