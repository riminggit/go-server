.PHONY: build clean tool lint help

all: build

build:
	@go build -v .

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf go-server-template
	go clean -i .

help:
	@echo "make: 编译包和依赖项"
	@echo "make tool: 运行指定的 go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: 删除目标文件和缓存文件"
