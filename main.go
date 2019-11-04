package main

import (
	"github.com/nickrisaro/proyectos-y-personas/empresa"
	"github.com/nickrisaro/proyectos-y-personas/rest"
)

func main() {

	nuevaEmpresa := empresa.New()
	api := rest.New(nuevaEmpresa)

	api.Start()
}
