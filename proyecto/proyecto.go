package proyecto

import (
	"errors"
	"math"

	"github.com/nickrisaro/proyectos-y-personas/persona"
)

const coeficientePersonas float64 = 1.0
const coeficientePresupuesto float64 = 0.5
const coeficienteSeniority float64 = 0.4
const coeficienteHardSkills float64 = 0.4
const coeficienteSoftSkills float64 = 0.4

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

// Diseño modifica la cantidad de personas con ese skill
func (p PersonasRequeridasPorSkill) Diseño(cantidad int) {
	p[persona.Diseño] = cantidad
}

// Operaciones modifica la cantidad de personas con ese skill
func (p PersonasRequeridasPorSkill) Operaciones(cantidad int) {
	p[persona.Operaciones] = cantidad
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
	Anombre             string                     `json:"nombre"`
	ApersonasRequeridas PersonasRequeridasPorSkill `json:"personasRequeridas"`
	Apresupuesto        float64                    `json:"presupuesto"`
	ApersonasAsignadas  []*persona.Persona         `json:"personasAsignadas"`
}

// Resumen contiene un resumen de la información del proyecto
type Resumen struct {
	Nombre             string
	PersonasRequeridas PersonasRequeridasPorSkill
	Presupuesto        float64
	Sueldos            float64
	HardSkills         map[persona.HardSkill]int
	SoftSkills         map[persona.SoftSkill]int
}

// New construye un nuevo proyecto
func New(nombre string, cantidadDePersonasRequeridas PersonasRequeridasPorSkill, presupuesto float64) *Proyecto {
	return &Proyecto{
		Anombre:             nombre,
		ApersonasRequeridas: cantidadDePersonasRequeridas,
		Apresupuesto:        presupuesto,
	}
}

// AsignarPersona agrega una persona al proyecto
func (p *Proyecto) AsignarPersona(unaPersona *persona.Persona) {
	p.ApersonasAsignadas = append(p.ApersonasAsignadas, unaPersona)
}

// Nombre indica el nombre del proyecto
func (p *Proyecto) Nombre() string {
	return p.Anombre
}

//Clonar realiza una copia de las características del proyecto, pero no de las personas asignadas a él
func (p *Proyecto) Clonar() *Proyecto {
	return New(p.Anombre, p.ApersonasRequeridas, p.Apresupuesto)
}

func (p *Proyecto) sueldos() float64 {
	sueldos := 0.0
	for _, unaPersona := range p.ApersonasAsignadas {
		sueldos += unaPersona.Sueldo()
	}
	return sueldos
}

func (p *Proyecto) seniorities() int {
	seniorities := 0
	for _, unaPersona := range p.ApersonasAsignadas {
		seniorities += int(unaPersona.Seniority())
	}
	return seniorities
}

func (p *Proyecto) cantidadDeHardSkillsFaltantes() int {
	faltantes := 0

	for skill, cantidad := range p.ApersonasRequeridas {
		faltantesSkill := cantidad
		for _, persona := range p.ApersonasAsignadas {
			if persona.HardSkill() == skill && faltantesSkill > 0 {
				faltantesSkill--
			}
		}
		faltantes += faltantesSkill
	}
	return faltantes
}

func (p *Proyecto) hardSkills() map[persona.HardSkill]int {
	hardSkills := make(map[persona.HardSkill]int)

	for _, persona := range p.ApersonasAsignadas {
		hardSkills[persona.HardSkill()] = hardSkills[persona.HardSkill()] + 1
	}

	return hardSkills
}

func (p *Proyecto) cantidadDeSoftSkillsDiferentes() int {

	softSkills := make(map[persona.SoftSkill]bool)

	for _, persona := range p.ApersonasAsignadas {
		softSkills[persona.SoftSkill()] = true
	}
	return len(softSkills)
}

func (p *Proyecto) softSkills() map[persona.SoftSkill]int {
	softSkills := make(map[persona.SoftSkill]int)

	for _, persona := range p.ApersonasAsignadas {
		softSkills[persona.SoftSkill()] = softSkills[persona.SoftSkill()] + 1
	}

	return softSkills
}

// Fitness evalúa cuan bien está este proyecto
// es una medida para comparar contra otro proyecto u otras "versiones" del mismo proyecto
func (p *Proyecto) Fitness() (float64, error) {
	if p.ApersonasRequeridas.sinPersonasAsignadas() {
		return 0.0, errors.New("El proyecto debe tener personas requeridas para calcular el fitness")
	}

	if p.Apresupuesto == 0.0 {
		return 0.0, errors.New("El proyecto debe tener un presupuesto para calcular el fitness")
	}

	fitness := coeficientePersonas*float64(len(p.ApersonasAsignadas)-p.ApersonasRequeridas.cantidadDePersonasRequeridas()) +
		coeficienteSeniority*float64(p.seniorities()) +
		coeficienteSoftSkills*float64(p.cantidadDeSoftSkillsDiferentes()) -
		coeficienteHardSkills*float64(p.cantidadDeHardSkillsFaltantes())

	if p.Apresupuesto-p.sueldos() < 0 {
		fitness = math.Inf(-1)
	}
	return fitness, nil
}

// ObtenerResumen devuelve un resumen de la información del proyecto
func (p *Proyecto) ObtenerResumen() Resumen {
	resumen := Resumen{
		Nombre:             p.Anombre,
		PersonasRequeridas: p.ApersonasRequeridas,
		Presupuesto:        p.Apresupuesto,
		Sueldos:            p.sueldos(),
		HardSkills:         p.hardSkills(),
		SoftSkills:         p.softSkills(),
	}
	return resumen
}
