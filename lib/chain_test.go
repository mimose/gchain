package lib

import (
	"fmt"
	"testing"
	"time"
)

func TestExecute(t *testing.T) {
	chain := NewChain()
	err := chain.AddCommand(commandOne{}, commandTwo{})
	if err != nil {
		t.Error(err)
	}

	context := NewContext()
	param := make(map[string]interface{})
	param["param1"] = "paramValue2"
	param["param2"] = 123
	context.AddMultiParam(param)

	_, err2 := chain.Execute(&context)
	if err2 != nil {
		t.Error(err2)
	}
}

type commandOne struct{}

func (one commandOne) Execute(ctx *Context) (bool, error) {
	fmt.Println("execute one...")
	time.Sleep(time.Second * 1)
	ctx.AddSingleData("data1", "dataValue1")
	err := ctx.AddSingleParam("param3", "paramValue3")
	if err != nil {
		fmt.Println(err)
	}
	return false, nil
}

type commandTwo struct{}

func (two commandTwo) Execute(ctx *Context) (bool, error) {
	fmt.Println("execute two...")
	fmt.Println(ctx.Summary())
	time.Sleep(time.Second * 1)
	return false, nil
}
