build-all: build-wasm-add
	go build main.go

run:
	go run main.go

build-wasm-add:
	(cd c++-wasm; make build-add)