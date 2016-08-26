package main

import "github.com/RouteInjector/gojector/infrastructure"

func main() {
	conf := &infrastructure.Configuration{}
	conf.Auth = false
	conf.Bind = 40000
}
