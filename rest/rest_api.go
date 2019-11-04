package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nickrisaro/proyectos-y-personas/empresa"
	"github.com/nickrisaro/proyectos-y-personas/persona"
)

// API API REST para administrar la empresa
type API struct {
	empresa *empresa.Empresa
}

// New construye una nueva API REST para administrar la empresa
func New(empresa *empresa.Empresa) *API {
	api := new(API)
	api.empresa = empresa
	return api
}

// Start inicia la API, queda escuchando en todas las interfaces en el puerto 8080
func (a *API) Start() {
	router := gin.Default()

	router.GET("/personas", a.listarPersonas)
	router.POST("/personas", a.altaPersonas)

	router.Run()
}

func (a *API) altaPersonas(c *gin.Context) {
	personas := make([]persona.Persona, 0)
	c.ShouldBindJSON(&personas)

	for _, persona := range personas {
		a.empresa.DarDeAltaEmpleado(&persona)
	}
	c.JSON(http.StatusOK, struct{ Message string }{fmt.Sprintf("%d personas dadas de alta", len(personas))})
}

func (a *API) listarPersonas(c *gin.Context) {
	c.JSON(http.StatusOK, a.empresa.Empleados())
}
