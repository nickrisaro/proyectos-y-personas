package solucion_test

import (
	"math"
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

	generadorDeSoluciones := solucion.NewGenerador(empresa, &algoritmoConstante{})
	unaSolucion := generadorDeSoluciones.ObtenerSolucion()

	assert.Equal(t, []int{0, 1}, unaSolucion.Configuracion(), "No se obtuvo la solución esperada")
	assert.Equal(t, 2.0, unaSolucion.Fitness(), "No se obtuvo el fitness esperado")
}

type algoritmoConstante struct{}

func (a *algoritmoConstante) GenerarNuevaSolucion() *solucion.Solucion {
	return solucion.New([]int{0, 1}, 2.0)
}

func (a *algoritmoConstante) NuevaGeneracionDeSoluciones(soluciones []*solucion.Solucion) []*solucion.Solucion {
	return soluciones
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

func TestElAlgoritmoGeneticoReemplazaLaPeorSolucionConUnaCombinacionDeLasMejores(t *testing.T) {

	algoritmo := solucion.NewAlgoritmoGenetico(2, 2)
	soluciones := make([]*solucion.Solucion, 4)
	soluciones[0] = solucion.New([]int{0, 0}, 0.1)
	soluciones[1] = solucion.New([]int{0, 1}, 0.3)
	soluciones[2] = solucion.New([]int{1, 0}, 0.7)
	soluciones[3] = solucion.New([]int{1, 1}, 0.9)

	nuevasSoluciones := algoritmo.NuevaGeneracionDeSoluciones(soluciones)

	assert.Len(t, nuevasSoluciones, 4, "Se obtuvieron menos soluciones de las esperadas")
	assert.Equal(t, 0.9, nuevasSoluciones[0].Fitness(), "El fitness no es el esperado")
	assert.Equal(t, 0.7, nuevasSoluciones[1].Fitness(), "El fitness no es el esperado")
	assert.Equal(t, 0.3, nuevasSoluciones[2].Fitness(), "El fitness no es el esperado")
	assert.Equal(t, math.Inf(-1), nuevasSoluciones[3].Fitness(), "El fitness no es el esperado")
}

func TestElAlgoritmoGeneticoFuncionaCon3Soluciones(t *testing.T) {

	algoritmo := solucion.NewAlgoritmoGenetico(2, 2)
	soluciones := make([]*solucion.Solucion, 3)
	soluciones[0] = solucion.New([]int{0, 0}, 0.0)
	soluciones[1] = solucion.New([]int{0, 1}, 0.3)
	soluciones[2] = solucion.New([]int{1, 0}, 0.7)

	nuevasSoluciones := algoritmo.NuevaGeneracionDeSoluciones(soluciones)

	assert.Len(t, nuevasSoluciones, 3, "Se obtuvieron menos soluciones de las esperadas")
}

func TestElAlgoritmoGeneticoFuncionaCon9Soluciones(t *testing.T) {

	algoritmo := solucion.NewAlgoritmoGenetico(2, 2)
	soluciones := make([]*solucion.Solucion, 9)
	soluciones[0] = solucion.New([]int{0, 0}, 0.0)
	soluciones[1] = solucion.New([]int{0, 1}, 0.3)
	soluciones[2] = solucion.New([]int{1, 0}, 0.7)
	soluciones[3] = solucion.New([]int{1, 0}, 0.2)
	soluciones[4] = solucion.New([]int{1, 0}, 0.5)
	soluciones[5] = solucion.New([]int{1, 0}, 0.9)
	soluciones[6] = solucion.New([]int{1, 0}, 0.6)
	soluciones[7] = solucion.New([]int{1, 0}, 0.8)
	soluciones[8] = solucion.New([]int{1, 0}, 0.1)

	nuevasSoluciones := algoritmo.NuevaGeneracionDeSoluciones(soluciones)

	assert.Len(t, nuevasSoluciones, 9, "Se obtuvieron menos soluciones de las esperadas")
}
