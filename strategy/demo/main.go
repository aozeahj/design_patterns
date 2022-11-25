package main

import (
	"github.com/aozeahj/design_patterns/strategy/demo/ali_pay"
	"github.com/aozeahj/design_patterns/strategy/demo/definition"
	"github.com/aozeahj/design_patterns/strategy/demo/wechat_pay"
)

func main() {
	payCtx := definition.Context{}

	// 执行 阿里支付
	payCtx.SetStrategy(&ali_pay.AliPayer{})
	payCtx.Pay()

	// 执行 微信支付
	payCtx.SetStrategy(&wechat_pay.WeChatPayer{})
	payCtx.Pay()
}
