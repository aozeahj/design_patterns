package doctor

import (
	"fmt"
	"github.com/aozeahj/design_patterns/chain_of_responsibility"
)

type Handler struct {
	Next chain_of_responsibility.Handler
}

func (h *Handler) Handle(c *chain_of_responsibility.Client) {
	fmt.Println("医生看病")

	if h.Next != nil {
		h.Next.Handle(c)
	}
}

func (h *Handler) SetNext(handler chain_of_responsibility.Handler) chain_of_responsibility.Handler{
	h.Next = handler
	return h
}



