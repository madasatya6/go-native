package env 

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

type EnvironmentMethod interface {
	setPort(port string)
}

type Environment struct {
	Port string
	WriteTimeout int 
	ReadTimeout int 
	Database string
}

func Init() *Environment {
	var method EnvironmentMethod
	var en Environment
	method = &en
	method.setPort("9090")
	return &en
}

func (e *Environment) setPort(port string) {
	var p = kingpin.Arg("port","Web Service Port").Default(port).String()
	e.Port = fmt.Sprintf(":%s", *p)
	if e.Port == ":" {
		e.Port = fmt.Sprintf(":%s", port)
	}
}