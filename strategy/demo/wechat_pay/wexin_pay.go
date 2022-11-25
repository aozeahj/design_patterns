package wechat_pay

import (
	"fmt"
	"github.com/aozeahj/design_patterns/strategy/demo/definition"
)

type WeChatPayer struct {
}

func (w WeChatPayer) Pay(ctx *definition.Context) *definition.Result {
	fmt.Println("weChat pay execute")
	return nil
}



