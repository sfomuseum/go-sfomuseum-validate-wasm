# go-sfomuseum-validate-wasm

Go package for compiling the `Validate` method of the go-sfomuseum-validate package to a JavaScript-compatible WebAssembly (wasm) binary. It also provides a net/http middleware packages for appending the necessary static assets and HTML resources to use the wasm binary in web applications.

## Build

To build the `validate_feature` WebAssembly binary for use in your applications run the following command:

```
GOOS=js GOARCH=wasm go build -mod vendor -o validate_feature.wasm cmd/validate_feature/main.go
```

## Use

To use the `validate_feature` WebAssembly binary in your applications a JSON-encoded GeoJSON string to the `validate_feature` function.

The function returns a JavaScript `Promise` that will return a JSON-encoded Who's On First (WOF) GeoJSON string on success or an error message if there was a problem.

```
	var str_f = '{"type": "Feature" ... }'	// A valid GeoJSON Feature
	
	sfomuseum_validate_feature().then(rsp => {
	    console.log("WOF feature is valid.");
	}).catch(err => {
	    console.log("Failed to validate feature: ", err);
	});
```

In order to load the `validate_feature` function you will need to include the `wasm_exec.js` and `sfomuseum.validate.feature.js` JavaScript files, or functional equivalents. Both scripts are bundled with this package in the [static/javascript](static/javascript) folder.

## Middleware

The `go-sfomuseum-validate-wasm/http` package provides methods for appending static assets and HTML resources to existing web applications to facilitate the use of the `validate_feature` WebAssembly binary. For example:

```
package main

import (
	"embed"
	wasm "github.com/sfomuseum/go-sfomuseum-validate-wasm/http"
	"log"
	"net/http"
)

//go:embed index.html example.*
var FS embed.FS

func main() {

	mux := http.NewServeMux()

	wasm.AppendAssetHandlers(mux)

	http_fs := http.FS(FS)
	example_handler := http.FileServer(http_fs)

	wasm_opts := wasm.DefaultWASMOptions()
	wasm_opts.EnableWASMExec()

	example_handler = wasm.AppendResourcesHandler(example_handler, wasm_opts)

	mux.Handle("/", example_handler)

	addr := "localhost:8080"
	log.Printf("Listening for requests on %s\n", addr)

	http.ListenAndServe(addr, mux)
}
```

_Error handling omitted for brevity._

## See also

* https://github.com/sfomuseum/go-sfomuseum-validate