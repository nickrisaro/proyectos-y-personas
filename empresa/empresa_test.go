package empresa_test

import (
	"testing"

	"github.com/nickrisaro/proyectos-y-personas/empresa"
	"github.com/nickrisaro/proyectos-y-personas/persona"
	"github.com/nickrisaro/proyectos-y-personas/proyecto"
	"github.com/stretchr/testify/assert"
)

func TestUnaEmpresaTieneEmpleados(t *testing.T) {

	ana := persona.New("Ana", 1.0, persona.Senior, persona.Desarrollo, persona.Negociacion)
	juan := persona.New("Juan", 1.0, persona.Junior, persona.Operaciones, persona.Investigacion)
	empresa := empresa.New()

	empresa.DarDeAltaEmpleado(ana)
	empresa.DarDeAltaEmpleado(juan)

	personas := []*persona.Persona{ana, juan}
	assert.Equal(t, personas, empresa.Empleados(), "Los empleados de la empresa no son correctos")
}

func TestUnaEmpresaTieneProyectos(t *testing.T) {

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	unProyecto := proyecto.New("Proyecto uno", personasRequeridas, 1.0)
	otroProyecto := proyecto.New("Proyecto dos", personasRequeridas, 1.0)
	empresa := empresa.New()

	empresa.DarDeAltaProyecto(unProyecto)
	empresa.DarDeAltaProyecto(otroProyecto)

	proyectos := []*proyecto.Proyecto{unProyecto, otroProyecto}
	assert.Equal(t, proyectos, empresa.Proyectos(), "Los proyectos de la empresa no son correctos")
}

func TestLaEmpresaPuedeBuscarLaMejorCombinacionDePersonasParaUnProyecto(t *testing.T) {

	empresa := empresa.New()

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	unProyecto := proyecto.New("Proyecto uno", personasRequeridas, 1.0)
	empresa.DarDeAltaProyecto(unProyecto)
	ana := persona.New("Ana", 1.0, persona.Senior, persona.Desarrollo, persona.Negociacion)
	juan := persona.New("Juan", 1.0, persona.Junior, persona.Operaciones, persona.Investigacion)
	empresa.DarDeAltaEmpleado(ana)
	empresa.DarDeAltaEmpleado(juan)

	empresa.ArmarEquiposDeTrabajo()
}
