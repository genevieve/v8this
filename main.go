package main

import (
	"fmt"
	"io/ioutil"

	"rogchap.com/v8go"
)

func main() {
	script, err := ioutil.ReadFile("script.js")
	if err != nil {
		panic(err)
	}

	iso, err := v8go.NewIsolate()
	if err != nil {
		panic(err)
	}

	ctx, err := v8go.NewContext(iso)
	if err != nil {
		panic(err)
	}

	val, err := ctx.RunScript(string(script), "script.js")
	if err != nil {
		panic(err)
	}

	respObj, err := val.AsObject()
	if err != nil {
		panic(err)
	}

	textVal, err := respObj.Get("text")
	if err != nil {
		panic(err)
	}

	textFn, err := textVal.AsFunction()
	if err != nil {
		panic(err)
	}

	text, err := textFn.Call(respObj)
	if err != nil {
		panic(err)
	}
	fmt.Println(text.String())

	globalText, err := ctx.Global().Get("text")
	if err != nil {
		panic(err)
	}
	globalTextFn, err := globalText.AsFunction()
	if err != nil {
		panic(err)
	}
	text, err = globalTextFn.Call(ctx.Global())
	if err != nil {
		panic(err)
	}
	fmt.Println(text.String())
}
