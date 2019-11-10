package solucion

import (
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/nickrisaro/proyectos-y-personas/empresa"
)

const cantidadDeSolucionesAGenerar int = 40
const cantidadDeEpocas int = 100

// GeneradorDeSoluciones es el encargado de buscar distribuciones posibles de
// personas en proyectos para la empresa
type GeneradorDeSoluciones struct {
	laEmpresa     *empresa.Empresa
	algoritmo     Algoritmo
	mejorSolucion *Solucion
}

// NewGenerador construye un nuevo generador de soluciones
func NewGenerador(unaEmpresa *empresa.Empresa, algoritmo Algoritmo) *GeneradorDeSoluciones {
	mejorSolucion := New([]int{}, math.Inf(-1))
	return &GeneradorDeSoluciones{unaEmpresa, algoritmo, mejorSolucion}
}

// ObtenerSolucion nos da una distribución de personas en proyectos
func (g *GeneradorDeSoluciones) ObtenerSolucion() *Solucion {

	soluciones := make([]*Solucion, cantidadDeSolucionesAGenerar)
	for i := 0; i < len(soluciones); i++ {

		soluciones[i] = g.algoritmo.GenerarNuevaSolucion()
		soluciones[i].fitness = g.laEmpresa.EvaluarSolucion(soluciones[i].configuracion)
	}

	for i := 0; i < cantidadDeEpocas; i++ {
		for _, solucion := range soluciones {
			if solucion.fitness >= g.mejorSolucion.fitness {
				g.mejorSolucion = solucion
			}
		}
		soluciones = g.algoritmo.NuevaGeneracionDeSoluciones(soluciones)

		for i := 0; i < len(soluciones); i++ {
			soluciones[i].fitness = g.laEmpresa.EvaluarSolucion(soluciones[i].configuracion)
		}
	}

	g.laEmpresa.AplicarSolucion(g.mejorSolucion.Configuracion())

	return g.mejorSolucion
}

// Solucion es una configuración posible de personas en proyectos con el fitness de esa configuración
type Solucion struct {
	configuracion []int
	fitness       float64
}

// New construye una nueva solucion
func New(configuracion []int, fitness float64) *Solucion {
	return &Solucion{configuracion, fitness}
}

// Configuracion devuelve la distribución de personas en proyectos
func (s *Solucion) Configuracion() []int {
	return s.configuracion
}

// Fitness devuelve el fitness de la solución
func (s *Solucion) Fitness() float64 {
	return s.fitness
}

// Algoritmo interfaz para la generación de soluciones
type Algoritmo interface {
	GenerarNuevaSolucion() *Solucion
	NuevaGeneracionDeSoluciones(soluciones []*Solucion) []*Solucion
}

// AlgoritmoGenetico Genera soluciones aleatorias y permite crear un nuevo conjunto de soluciones
// en base a un conjunto previo
type AlgoritmoGenetico struct {
	cantidadDePersonas  int
	cantidadDeProyectos int
}

// NewAlgoritmoGenetico construye una nueva instancia del algoritmo genético
func NewAlgoritmoGenetico(cantidadDePersonas, cantidadDeProyectos int) *AlgoritmoGenetico {
	rand.Seed(time.Now().UnixNano())
	return &AlgoritmoGenetico{cantidadDePersonas, cantidadDeProyectos}
}

// GenerarNuevaSolucion construye una nueva solución con valores aleatorios
// para la cantidad de personas y proyectos que se configuraron
func (a *AlgoritmoGenetico) GenerarNuevaSolucion() *Solucion {

	personas := make([]int, a.cantidadDePersonas)

	for i := range personas {
		personas[i] = rand.Intn(a.cantidadDeProyectos+1) - 1
	}

	return &Solucion{personas, 0.0}
}

// NuevaGeneracionDeSoluciones reemplaza las peores configuraciones de las soluciones
// por una combinación de las mejores
func (a *AlgoritmoGenetico) NuevaGeneracionDeSoluciones(soluciones []*Solucion) []*Solucion {

	nuevasSoluciones := make([]*Solucion, len(soluciones))

	ordenarSolucionesDeMayorAMenor(soluciones)

	cantidadDeSolucionesAReemplazar := len(soluciones) / 4
	cantidadDeMadres := cantidadDeSolucionesAReemplazar * 2

	madres := soluciones[0:cantidadDeMadres]

	for i := 0; i < len(nuevasSoluciones)-cantidadDeSolucionesAReemplazar; i++ {
		nuevasSoluciones[i] = soluciones[i]
	}

	for i := 0; i < cantidadDeSolucionesAReemplazar; i++ {

		nuevaSolucion := nuevaSolucionCon(madres[i], madres[len(madres)-1-i])
		nuevasSoluciones[len(nuevasSoluciones)-cantidadDeSolucionesAReemplazar+i] = nuevaSolucion
	}

	return nuevasSoluciones
}

func ordenarSolucionesDeMayorAMenor(soluciones []*Solucion) {
	sort.SliceStable(soluciones, func(i, j int) bool { return soluciones[i].Fitness() > soluciones[j].Fitness() })
}

func nuevaSolucionCon(unaSolucion *Solucion, otraSolucion *Solucion) *Solucion {
	nuevaConfiguracion := make([]int, len(unaSolucion.Configuracion()))

	for i := 0; i < len(nuevaConfiguracion); i++ {

		moneda := rand.Float64()
		var proyecto int
		if moneda < 0.5 {
			proyecto = unaSolucion.Configuracion()[i]
		} else {
			proyecto = otraSolucion.Configuracion()[i]

		}
		nuevaConfiguracion[i] = proyecto
	}

	nuevaSolucion := &Solucion{nuevaConfiguracion, math.Inf(-1)}
	return nuevaSolucion
}
