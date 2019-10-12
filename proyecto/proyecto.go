package proyecto

import (
	"errors"

	"github.com/nickrisaro/proyectos-y-personas/persona"
)

const coeficientePersonas float64 = 1.0
const coeficientePresupuesto float64 = 0.5

// Proyecto contiene toda la información relativa a un proyecto.
type Proyecto struct {
	personasRequeridas int
	presupuesto        float64
	personasAsignadas  []*persona.Persona
}

// New construye un nuevo proyecto
func New(cantidadDePersonasRequeridas int, presupuesto float64) *Proyecto {
	return &Proyecto{
		personasRequeridas: cantidadDePersonasRequeridas,
		presupuesto:        presupuesto,
	}
}

// AsignarPersona agrega una persona al proyecto
func (p *Proyecto) AsignarPersona(unaPersona *persona.Persona) {
	p.personasAsignadas = append(p.personasAsignadas, unaPersona)
}

// Fitness evalúa cuan bien está este proyecto
// es una medida para comparar contra otro proyecto u otras "versiones" del mismo proyecto
func (p *Proyecto) Fitness() (float64, error) {
	if p.personasRequeridas == 0 {
		return 0.0, errors.New("El proyecto debe tener personas requeridas para calcular el fitness")
	}

	if p.presupuesto == 0.0 {
		return 0.0, errors.New("El proyecto debe tener un presupuesto para calcular el fitness")
	}

	sueldos := 0.0
	for _, unaPersona := range p.personasAsignadas {
		sueldos += unaPersona.Sueldo()
	}

	fitness := coeficientePersonas*float64(len(p.personasAsignadas)-p.personasRequeridas) +
		coeficientePresupuesto*(p.presupuesto-sueldos)
	return fitness, nil
}
