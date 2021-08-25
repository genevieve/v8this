package main

import (
	"fmt"
	"io/ioutil"

	"rogchap.com/v8go"
)

func main() {
	script, _ := ioutil.ReadFile("script.js")
	iso, _ := v8go.NewIsolate() // creates a new JavaScript VM
	ctx, _ := v8go.NewContext(iso)
	val, _ := ctx.RunScript(string(script), "script.js")
	respObj, _ := val.AsObject()
	// bodyVal, _ := respObj.Get("body")
	textVal, _ := respObj.Get("text")
	textFn, _ := textVal.AsFunction()

	// body, err := textFn.Call(bodyVal)
	body, err := textFn.Call()
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
}
