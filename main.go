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

	val, err = respObj.Call("text")
	if err != nil {
		panic(err)
	}

	fmt.Println(val.String())
}
