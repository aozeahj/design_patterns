package reception

import (
	"fmt"
	"github.com/aozeahj/design_patterns/chain_of_responsibility"
)

/*
	挂号
 */

type Handler struct {
	Next chain_of_responsibility.Handler
}

func (h *Handler) Handle(c *chain_of_responsibility.Client) {
	fmt.Println("前台挂号")

	if h.Next != nil {
		h.Next.Handle(c)
	}
}

func (h *Handler) SetNext(handler chain_of_responsibility.Handler)chain_of_responsibility.Handler {
	h.Next = handler
	return h
}



