package solucion_test

import (
	"testing"

	"github.com/nickrisaro/proyectos-y-personas/empresa"
	"github.com/nickrisaro/proyectos-y-personas/persona"
	"github.com/nickrisaro/proyectos-y-personas/proyecto"
	"github.com/nickrisaro/proyectos-y-personas/solucion"
	"github.com/stretchr/testify/assert"
)

func TestSeGeneraUnaSolucionParaUnaEmpresa(t *testing.T) {

	empresa := empresa.New()
	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	unProyecto := proyecto.New("Proyecto uno", personasRequeridas, 1.0)
	empresa.DarDeAltaProyecto(unProyecto)
	otroProyecto := proyecto.New("Proyecto dos", personasRequeridas, 1.0)
	empresa.DarDeAltaProyecto(otroProyecto)
	ana := persona.New("Ana", 1.0, persona.Senior, persona.Desarrollo, persona.Negociacion)
	empresa.DarDeAltaEmpleado(ana)
	juan := persona.New("Juan", 1.0, persona.Junior, persona.Operaciones, persona.Investigacion)
	empresa.DarDeAltaEmpleado(juan)

	generadorDeSoluciones := solucion.NewGenerador(empresa)
	unaSolucion := generadorDeSoluciones.ObtenerSolucion()

	assert.Equal(t, []int{0, 1}, unaSolucion.Configuracion(), "No se obtuvo la solución esperada")
	assert.Equal(t, 2.0, unaSolucion.Fitness(), "No se obtuvo el fitness esperado")
}

func TestElAlgoritmoGeneticoGeneraSolucionesAleatorias(t *testing.T) {

	algoritmo := solucion.NewAlgoritmoGenetico(2, 2)

	unaSolucion := algoritmo.GenerarNuevaSolucion()
	configuracion := unaSolucion.Configuracion()

	assert.Equal(t, 2, len(configuracion), "La longitud de la solución no es la esperada")
	assert.GreaterOrEqual(t, configuracion[0], 0, "El valor está fuera del rango")
	assert.LessOrEqual(t, configuracion[0], 1, "El valor está fuera del rango")
	assert.GreaterOrEqual(t, configuracion[1], 0, "El valor está fuera del rango")
	assert.LessOrEqual(t, configuracion[1], 1, "El valor está fuera del rango")
}
