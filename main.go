package main

import (
	"fmt"
	"io/ioutil"
	"log"

	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
    wasmBytes, err := ioutil.ReadFile("c++-wasm/add.wasm")

	if err != nil {
		log.Fatalf("Unable to read file %v \n",err)
	}

    engine := wasmer.NewEngine()
    store := wasmer.NewStore(engine)

    // Compiles the module
    module, err := wasmer.NewModule(store, wasmBytes)

	if err != nil {
		log.Fatalf("Unable to get new module %v \n",err)
	}

    // Instantiates the module
    importObject := wasmer.NewImportObject()
    instance, err := wasmer.NewInstance(module, importObject)

	if err != nil {
		log.Fatalf("Unable to get instance %v \n",err)
	}

    // Gets the `add` exported function from the WebAssembly instance.
    add, err := instance.Exports.GetFunction("Add")

	if err != nil {
		log.Fatalf("Unable to get function %v \n",err)
	}

    // Calls that exported function with Go standard values. The WebAssembly
    // types are inferred and values are casted automatically.
   if add != nil {
		result, err := add(5, 37)
		if err != nil {
			log.Fatalf("Unable to call function %v \n",err)
		}
		fmt.Println(result) // 42!
   }

    
}