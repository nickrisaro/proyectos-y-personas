package solucion

import (
	"math/rand"
	"time"

	"github.com/nickrisaro/proyectos-y-personas/empresa"
)

// GeneradorDeSoluciones es el encargado de buscar distribuciones posibles de
// personas en proyectos para la empresa
type GeneradorDeSoluciones struct {
	laEmpresa *empresa.Empresa
}

// NewGenerador construye un nuevo generador de soluciones
func NewGenerador(unaEmpresa *empresa.Empresa) *GeneradorDeSoluciones {
	return &GeneradorDeSoluciones{unaEmpresa}
}

// ObtenerSolucion nos da una distribución de personas en proyectos
func (g *GeneradorDeSoluciones) ObtenerSolucion() *Solucion {
	configuracion := []int{0, 1}
	fitness := g.laEmpresa.EvaluarSolucion(configuracion)
	return &Solucion{configuracion, fitness}
}

// Solucion es una configuración posible de personas en proyectos con el fitness de esa configuración
type Solucion struct {
	configuracion []int
	fitness       float64
}

// Configuracion devuelve la distribución de personas en proyectos
func (s *Solucion) Configuracion() []int {
	return s.configuracion
}

// Fitness devuelve el fitness de la solución
func (s *Solucion) Fitness() float64 {
	return s.fitness
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
		personas[i] = rand.Intn(2)
	}

	return &Solucion{personas, 0.0}
}
