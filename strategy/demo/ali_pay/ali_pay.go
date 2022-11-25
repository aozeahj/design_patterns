package ali_pay

import (
	"fmt"
	"github.com/aozeahj/design_patterns/strategy/demo/definition"
)

type AliPayer struct {
}

func (a *AliPayer) Pay(ctx *definition.Context) *definition.Result {
	fmt.Println("ali pay execute")
	return nil
}



