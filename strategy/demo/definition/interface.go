package definition

type PayStrategy interface {
	Pay(ctx *Context) *Result
}

type Result struct {

}

type Context struct {
	params map[string]interface{}
	payStrategy PayStrategy
}

func (c *Context)SetStrategy(strategy PayStrategy)  {
	c.payStrategy = strategy
}

func (c *Context)Pay() *Result {
	return c.payStrategy.Pay(c)
}
