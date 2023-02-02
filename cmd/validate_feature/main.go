package main

import (
	"fmt"
	"log"
	"syscall/js"

	"github.com/sfomuseum/go-sfomuseum-validate"	
)

func ExportFunc(opts *validate.Options) js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		geojson_data := args[0].String()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			go func() {

				err := validate.ValidateWithOptions([]byte(geojson_data), opts)

				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to export data, %v", err))
					return
				}

				resolve.Invoke()
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func main() {

	opts := validate.DefaultValidateOptions()

	export_func := ExportFunc(opts)
	defer export_func.Release()

	js.Global().Set("sfomuseum_validate_feature", export_func)

	c := make(chan struct{}, 0)

	log.Println("SFO Museum validate_feature WASM binary initialized")
	<-c
}
