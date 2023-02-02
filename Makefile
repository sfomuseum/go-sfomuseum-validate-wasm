GOROOT=$(shell go env GOROOT)

rebuild:
	@make wasmjs
	@make wasm

wasm:
	GOOS=js GOARCH=wasm go build -mod vendor -ldflags="-s -w" -o static/wasm/sfomuseum_validate_feature.wasm cmd/validate_feature/main.go

wasmjs:
	cp "$(GOROOT)/misc/wasm/wasm_exec.js" static/javascript/
