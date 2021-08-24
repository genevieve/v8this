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
	fmt.Printf("%v\n", &val)

	respObj, _ := val.AsObject()
	fmt.Printf("%v\n", &respObj)

	bodyVal, _ := respObj.Get("body")
	fmt.Println(bodyVal.String())
	fmt.Println(respObj)

	err := respObj.Set("this", respObj)
	if err != nil {
		panic(err)
	}

	textVal, _ := respObj.Get("text")
	fmt.Println(textVal)

	textObj, _ := textVal.AsObject()
	fmt.Println(textObj)

	textFn, _ := textObj.AsFunction()
	fmt.Println(textFn)

	bindVal, _ := textObj.Get("bind")
	fmt.Println(bindVal)

	bindFn, _ := bindVal.AsFunction()
	fmt.Println(bindFn)

	v, e := textFn.Call(bindFn, respObj)
	if e != nil {
		panic(e)
	}
	fmt.Println(v)

	// value, err := bindFn.Call(respObj)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(value)

	// textFn, _ := value.AsFunction()
	// fmt.Println(textFn)

	body, err := textFn.Call()
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
}
