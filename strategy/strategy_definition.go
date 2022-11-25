package strategy

/*
	策略抽象接口
 */
type Strategy interface {
	Algorithm(ctx  *Context) *Result
}

/*
	策略算法的返回
 */
type Result struct {
}


/*
	策略的上下文，
 */
type Context struct {
	params map[string]interface{}
	strategy Strategy
}

/*
	具体策略接口的实现
 */
type ConcreteStrategy struct {
}

func (c ConcreteStrategy) Algorithm(ctx *Context) *Result {
	return nil
}



