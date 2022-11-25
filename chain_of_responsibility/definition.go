package chain_of_responsibility

import "fmt"

/*
	责任链需要处理的对象
 */
type Client struct {

}
/*
	责任链处理者抽象的接口
 */
type Handler interface {
	Handle(c *Client)
	SetNext(handler Handler) Handler
}

type ConcreteHandler struct {
	Next Handler
}

func (c *ConcreteHandler)Handle(client *Client)  {
	fmt.Println(" 执行")

	c.Next.Handle(client)
}

func (c *ConcreteHandler)SetNext(handler Handler) Handler {
	c.Next = handler
	return c
}