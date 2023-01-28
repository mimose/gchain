package lib

type Command interface {
	Execute(ctx *Context) (bool, error)
	//PostExecute(ctx *Context) error
}
