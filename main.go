package main

import (
	"github.com/penguinn/penguin/component/router"
	"github.com/penguinn/penguin-test/controllers"
	"github.com/penguinn/penguin/component/server"
)

func main() {
	router.RegisterControllerGroup(controllers.NewHelloController(), "api", /*middleware*/)
	server.Serve()
}
