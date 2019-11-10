package empresa

import (
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

// ModificarPersona modifica la persona del ID indicado con los datos de la persona recibida
func (e *Empresa) ModificarPersona(ID int, persona *persona.Persona) {
	e.empleados[ID] = persona
}

// DarDeAltaProyecto registra el proyecto en la empresa
func (e *Empresa) DarDeAltaProyecto(proyecto *proyecto.Proyecto) {
	e.proyectos = append(e.proyectos, proyecto)
}

// ModificarProyecto modifica el proyecto del ID indicado con los datos del proyecto recibido
func (e *Empresa) ModificarProyecto(ID int, proyecto *proyecto.Proyecto) {
	e.proyectos[ID] = proyecto
}

// Empleados retorna todos los empleados de la empresa
func (e *Empresa) Empleados() []*persona.Persona {
	return e.empleados
}

// Proyectos retorna todos los proyectos de la empresa
func (e *Empresa) Proyectos() []*proyecto.Proyecto {
	return e.proyectos
}

// EvaluarSolucion indica cuán buena es una solución
// solución contiene para cada persona a qué proyecto se la asigna
func (e *Empresa) EvaluarSolucion(solucion []int) float64 {

	proyectosClonados := make([]*proyecto.Proyecto, 0)
	for _, proyecto := range e.proyectos {
		proyectosClonados = append(proyectosClonados, proyecto.Clonar())
	}

	for persona, proyecto := range solucion {
		proyectosClonados[proyecto].AsignarPersona(e.empleados[persona])
	}

	fitness := 0.0

	for _, proyecto := range proyectosClonados {
		fitnessProyecto, _ := proyecto.Fitness()
		fitness += fitnessProyecto
	}

	return fitness
}

// AplicarSolucion asigna personas a los proyectos  según lo que dice la solución
func (e *Empresa) AplicarSolucion(configuracion []int) {

	for i, proyecto := range e.proyectos {
		e.proyectos[i] = proyecto.Clonar()
	}

	for persona, proyecto := range configuracion {
		e.proyectos[proyecto].AsignarPersona(e.empleados[persona])
	}

}

// ResumenDeProyectos devuelve un resumen de todos los proyectos de la empresa
func (e *Empresa) ResumenDeProyectos() []proyecto.Resumen {
	resumenes := make([]proyecto.Resumen, len(e.proyectos))

	for i, proyecto := range e.proyectos {
		resumenes[i] = proyecto.ObtenerResumen()
	}

	return resumenes
}
