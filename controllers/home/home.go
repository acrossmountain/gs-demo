package home

import "github.com/go-spring/spring-core/web"

type Controller struct {
}

func (c *Controller) Home(ctx web.Context) {
	ctx.String("Ok! \n")
}
