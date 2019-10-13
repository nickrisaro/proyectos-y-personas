package proyecto

import (
	"errors"

	"github.com/nickrisaro/proyectos-y-personas/persona"
)

const coeficientePersonas float64 = 1.0
const coeficientePresupuesto float64 = 0.5
const coeficienteSeniority float64 = 0.4

// PersonasRequeridasPorSkill indica cuántas personas se requieren de cada skill
type PersonasRequeridasPorSkill map[persona.HardSkill]int

// NewPersonasRequeridasPorSkill construye una nueva instancia de la configuración
func NewPersonasRequeridasPorSkill() PersonasRequeridasPorSkill {
	return PersonasRequeridasPorSkill(make(map[persona.HardSkill]int))
}

// Desarrollo modifica la cantidad de personas con ese skill
func (p PersonasRequeridasPorSkill) Desarrollo(cantidad int) {
	p[persona.Desarrollo] = cantidad
}

func (p PersonasRequeridasPorSkill) sinPersonasAsignadas() bool {
	return len(p) == 0
}

func (p PersonasRequeridasPorSkill) cantidadDePersonasRequeridas() int {
	cantidad := 0
	for _, cantidadPersonas := range p {
		cantidad += cantidadPersonas
	}
	return cantidad
}

// Proyecto contiene toda la información relativa a un proyecto.
type Proyecto struct {
	personasRequeridas PersonasRequeridasPorSkill
	presupuesto        float64
	personasAsignadas  []*persona.Persona
}

// New construye un nuevo proyecto
func New(cantidadDePersonasRequeridas PersonasRequeridasPorSkill, presupuesto float64) *Proyecto {
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
	if p.personasRequeridas.sinPersonasAsignadas() {
		return 0.0, errors.New("El proyecto debe tener personas requeridas para calcular el fitness")
	}

	if p.presupuesto == 0.0 {
		return 0.0, errors.New("El proyecto debe tener un presupuesto para calcular el fitness")
	}

	sueldos := 0.0
	seniorities := 0
	for _, unaPersona := range p.personasAsignadas {
		sueldos += unaPersona.Sueldo()
		seniorities += int(unaPersona.Seniority())
	}

	fitness := coeficientePersonas*float64(len(p.personasAsignadas)-p.personasRequeridas.cantidadDePersonasRequeridas()) +
		coeficientePresupuesto*(p.presupuesto-sueldos) + coeficienteSeniority*float64(seniorities)
	return fitness, nil
}
