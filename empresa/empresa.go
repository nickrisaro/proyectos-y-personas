package empresa

import (
	"fmt"

	"github.com/nickrisaro/proyectos-y-personas/persona"
	"github.com/nickrisaro/proyectos-y-personas/proyecto"
)

// Empresa es la representación de una empresa, tiene empleados y proyectos
type Empresa struct {
	empleados []*persona.Persona
	proyectos []*proyecto.Proyecto
}

// New construye una nueva empresa
func New() *Empresa {
	return new(Empresa)
}

// DarDeAltaEmpleado registra a la persona como empleada de la empresa
func (e *Empresa) DarDeAltaEmpleado(persona *persona.Persona) {
	e.empleados = append(e.empleados, persona)
}

// DarDeAltaProyecto registra el proyecto en la empresa
func (e *Empresa) DarDeAltaProyecto(proyecto *proyecto.Proyecto) {
	e.proyectos = append(e.proyectos, proyecto)
}

// Empleados retorna todos los empleados de la empresa
func (e *Empresa) Empleados() []*persona.Persona {
	return e.empleados
}

// Proyectos retorna todos los proyectos de la empresa
func (e *Empresa) Proyectos() []*proyecto.Proyecto {
	return e.proyectos
}

// ArmarEquiposDeTrabajo busca para cada proyecto el mejor grupo de personas posible
func (e *Empresa) ArmarEquiposDeTrabajo() {

	unProyecto := e.Proyectos()[0]

	combinacionesPosibles := make(map[*proyecto.Proyecto]float64)

	fitness, _ := unProyecto.Fitness()
	combinacionesPosibles[unProyecto] = fitness

	clon1 := unProyecto.Clonar()
	clon2 := unProyecto.Clonar()
	clon3 := unProyecto.Clonar()

	clon1.AsignarPersona(e.empleados[0])
	fitnessClon1, _ := clon1.Fitness()
	combinacionesPosibles[clon1] = fitnessClon1

	clon2.AsignarPersona(e.empleados[1])
	fitnessClon2, _ := clon2.Fitness()
	combinacionesPosibles[clon2] = fitnessClon2

	clon3.AsignarPersona(e.empleados[0])
	clon3.AsignarPersona(e.empleados[1])
	fitnessClon3, _ := clon3.Fitness()
	combinacionesPosibles[clon3] = fitnessClon3 // Este tiene mejor fitness, pero debería ser peor, porque tiene más gente y debería sobrepasar el presupuesto

	fmt.Println(combinacionesPosibles)
}
