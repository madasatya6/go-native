package env 

import (
	"fmt"
	"time"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/madasatya6/go-native/helpers/utils"
)

type EnvironmentMethod interface {
	setPort(port string)
	setWriteTimeout(timeout int)
	setReadTimeout(timeout int)
}

type Environment struct {
	Port string
	WriteTimeout time.Duration 
	ReadTimeout time.Duration
	Database string
}

func Init() *Environment {
	var method EnvironmentMethod
	var en Environment

	method = &en
	method.setPort("9090")
	method.setWriteTimeout(15)
	method.setReadTimeout(15)
	kingpin.Parse()

	return &en
}

func (e *Environment) setPort(port string) {
	var p = kingpin.Arg("port","Web Service Port").Default(port).String()
	e.Port = fmt.Sprintf(":%s", *p)
	if e.Port == ":" {
		e.Port = fmt.Sprintf(":%s", port)
	}
}

func (e *Environment) setWriteTimeout(timeout int) {
	var p = kingpin.Arg("write_timeout","Write Time Out").Default(utils.IntToString(timeout)).Int()
	var str = fmt.Sprintf("%v", *p)
	if str == "" {
		e.WriteTimeout = time.Duration(timeout) * time.Second
	} else {
		wt, _ := strconv.Atoi(str)
		e.WriteTimeout = time.Duration(wt) * time.Second
	}
}

func (e *Environment) setReadTimeout(timeout int) {
	var p = kingpin.Arg("read_timeout","Read Time Out").Default(utils.IntToString(timeout)).Int()
	var str = fmt.Sprintf("%d", *p)
	if str == "" {
		e.ReadTimeout = time.Duration(timeout) * time.Second
	} else {
		wt, _ := strconv.Atoi(str)
		e.ReadTimeout = time.Duration(wt) * time.Second
	}
}