package main

import (
	"github.com/cegorah/auth_service/api"
)

func main(){
	cl := api.Client{}
	cl.RunServer()
}
