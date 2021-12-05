package controllers

import (
	"learn/controllers/home"
	"learn/controllers/upload"

	"github.com/go-spring/spring-core/gs"
)

func init() {

	gs.Object(new(home.Controller)).Init(func(c *home.Controller) {
		gs.GetMapping("/", c.Home)
	})

	gs.Object(new(upload.Controller)).Init(func(c *upload.Controller) {
		gs.PostMapping("/upload", c.Upload)
	})
}
