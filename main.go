package main

import (
	"github.com/tobiahrex/golang-sandbox/interfaces/decorator"
	dependencyInject_iface "github.com/tobiahrex/golang-sandbox/interfaces/dependency_inject"
	"github.com/tobiahrex/golang-sandbox/interfaces/factory"
	"github.com/tobiahrex/golang-sandbox/interfaces/middleware"
)

func main() {
	decorator_iface.Main()
	dependencyInject_iface.Main()
	factory_iface.Main()
	middleware_iface.Main()
}