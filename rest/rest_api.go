package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nickrisaro/proyectos-y-personas/empresa"
	"github.com/nickrisaro/proyectos-y-personas/persona"
	"github.com/nickrisaro/proyectos-y-personas/proyecto"
	"github.com/nickrisaro/proyectos-y-personas/solucion"
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
	router.PUT("/persona/:id", a.modificacionPersona)

	router.GET("/proyectos", a.listarProyectos)
	router.POST("/proyectos", a.altaProyectos)
	router.PUT("/proyecto/:id", a.modificacionProyecto)

	router.POST("/solucionar", a.solucionar)

	router.Run()
}

func (a *API) listarPersonas(c *gin.Context) {
	c.JSON(http.StatusOK, a.empresa.Empleados())
}

func (a *API) altaPersonas(c *gin.Context) {
	personas := make([]*persona.Persona, 0)
	c.ShouldBindJSON(&personas)

	for _, persona := range personas {
		a.empresa.DarDeAltaEmpleado(persona)
	}
	c.JSON(http.StatusOK, struct{ Message string }{fmt.Sprintf("%d personas dadas de alta", len(personas))})
}

func (a *API) modificacionPersona(c *gin.Context) {
	nuevaPersona := persona.Persona{}
	c.ShouldBindJSON(&nuevaPersona)

	IDPersona := c.GetInt("id")
	a.empresa.ModificarPersona(IDPersona, &nuevaPersona)

	c.JSON(http.StatusOK, struct{ Message string }{fmt.Sprintf("Persona %d modificada", IDPersona)})
}

func (a *API) altaProyectos(c *gin.Context) {
	proyectos := make([]*proyecto.Proyecto, 0)
	c.ShouldBindJSON(&proyectos)

	for _, proyecto := range proyectos {
		a.empresa.DarDeAltaProyecto(proyecto)
	}
	c.JSON(http.StatusOK, struct{ Message string }{fmt.Sprintf("%d proyectos dados de alta", len(proyectos))})
}

func (a *API) listarProyectos(c *gin.Context) {
	c.JSON(http.StatusOK, a.empresa.Proyectos())
}

func (a *API) modificacionProyecto(c *gin.Context) {
	nuevoProyecto := proyecto.Proyecto{}
	c.ShouldBindJSON(&nuevoProyecto)

	IDProyecto := c.GetInt("id")
	a.empresa.ModificarProyecto(IDProyecto, &nuevoProyecto)

	c.JSON(http.StatusOK, struct{ Message string }{fmt.Sprintf("Proyecto %d modificado", IDProyecto)})
}

func (a *API) solucionar(c *gin.Context) {
	solucionador := solucion.NewGenerador(a.empresa,
		solucion.NewAlgoritmoGenetico(len(a.empresa.Empleados()),
			len(a.empresa.Proyectos())))

	solucionador.ObtenerSolucion()

	c.JSON(http.StatusOK, a.empresa.ResumenDeProyectos())
}
