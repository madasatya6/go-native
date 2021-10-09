package env 

import (
	"fmt"
	"time"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/madasatya6/go-native/helpers/utils"
)

type EnvironmentMethod interface {
	setEnv(port string, write int, read int)
}

type Environment struct {
	Port string
	WriteTimeout time.Duration 
	ReadTimeout time.Duration
}

func Init() *Environment {
	var method EnvironmentMethod
	var en Environment

	method = &en
	method.setEnv("9090", 15, 15)

	return &en
}

func (e *Environment) setEnv(port string, write int, read int) {

	var argPort = kingpin.Flag("port","Web Service Port").Default(port).String()
	var argWrite = kingpin.Flag("write_timeout","Write Time Out").Default(utils.IntToString(write)).Int()
	var argRead = kingpin.Flag("read_timeout","Read Time Out").Default(utils.IntToString(read)).Int()
	kingpin.Parse()

	//set default port 
	e.Port = fmt.Sprintf(":%v", *argPort)
	if e.Port == ":" {
		e.Port = fmt.Sprintf(":%s", port)
	}

	//set default write timeout
	var strWrite = fmt.Sprintf("%v", *argWrite)
	if strWrite == "" {
		e.WriteTimeout = time.Duration(write) * time.Second
	} else {
		wt, _ := strconv.Atoi(strWrite)
		e.WriteTimeout = time.Duration(wt) * time.Second
	}

	//set default read timeout 
	var strRead = fmt.Sprintf("%d", *argRead)
	if strRead == "" {
		e.ReadTimeout = time.Duration(read) * time.Second
	} else {
		rt, _ := strconv.Atoi(strRead)
		e.ReadTimeout = time.Duration(rt) * time.Second
	}
}
