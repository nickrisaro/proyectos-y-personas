package proyecto_test

import (
	"testing"

	"github.com/nickrisaro/proyectos-y-personas/persona"
	"github.com/nickrisaro/proyectos-y-personas/proyecto"
	"github.com/stretchr/testify/assert"
)

func TestUnProyectoSinPersonasRequeridasNoTieneFitness(t *testing.T) {

	unProyecto := proyecto.New(0, 1.0)

	_, err := unProyecto.Fitness()

	assert.Error(t, err, "Se esperaba un error")
}

func TestUnProyectoSinPersonasAsignadasTieneFitnessNegativo(t *testing.T) {

	unProyecto := proyecto.New(1, 1.0)

	fitness, _ := unProyecto.Fitness()

	assert.Less(t, fitness, 0.0, "El fitness del proyecto debería ser negativo")
}

func TestUnProyectoConIgualCantidadDePersonasRequeridasQueAsignadasTieneFitnessMayorOIgualQueCero(t *testing.T) {

	unProyecto := proyecto.New(1, 1.0)
	unaPersona := persona.New(1.0, persona.Senior)
	unProyecto.AsignarPersona(unaPersona)

	fitness, _ := unProyecto.Fitness()

	assert.GreaterOrEqual(t, fitness, 0.0, "El fitness del proyecto debería ser mayor o igual que cero")
}

func TestUnProyectoSinPresupuestoAsignadoNoTieneFitness(t *testing.T) {

	unProyecto := proyecto.New(1, 0.0)

	_, err := unProyecto.Fitness()

	assert.Error(t, err, "Se esperaba un error")
}

func TestUnProyectoQueSeExcedeDelPresupuestoTieneFitnessNegativo(t *testing.T) {

	unProyecto := proyecto.New(1, 1.0)
	unaPersona := persona.New(2.0, persona.Junior)
	unProyecto.AsignarPersona(unaPersona)

	fitness, _ := unProyecto.Fitness()

	assert.Less(t, fitness, 0.0, "El fitness del proyecto debería ser negativo")
}

func TestUnProyectoConMenorGastoDeSueldosTieneMejorFitnessQueUnoConMayorGastoDeSueldos(t *testing.T) {

	proyectoQueGastaMenos := proyecto.New(1, 1.0)
	unaPersona := persona.New(0.7, persona.Senior)
	proyectoQueGastaMenos.AsignarPersona(unaPersona)
	proyectoQueGastaMas := proyecto.New(1, 1.0)
	otraPersona := persona.New(0.9, persona.Senior)
	proyectoQueGastaMas.AsignarPersona(otraPersona)

	fitnessDelProyectoQueGastaMenos, _ := proyectoQueGastaMenos.Fitness()
	fitnessDelProyectoQueGastaMas, _ := proyectoQueGastaMas.Fitness()

	assert.Greater(t, fitnessDelProyectoQueGastaMenos, fitnessDelProyectoQueGastaMas, "El fitness del proyecto más barato debería ser mayor que el del más caro")

}

func TestUnProyectoConPersonasDeMasSeniorityTieneMejorFitnessQueUnoConPersonasDeMenorSeniority(t *testing.T) {

	proyectoConMenosSeniority := proyecto.New(1, 1.0)
	unaPersona := persona.New(0.9, persona.Junior)
	proyectoConMenosSeniority.AsignarPersona(unaPersona)
	proyectoConMAyorSeniority := proyecto.New(1, 1.0)
	otraPersona := persona.New(0.9, persona.Senior)
	proyectoConMAyorSeniority.AsignarPersona(otraPersona)

	fitnessDelProyectoConMenorSeniority, _ := proyectoConMenosSeniority.Fitness()
	fitnessDelProyectoConMayorSeniority, _ := proyectoConMAyorSeniority.Fitness()

	assert.Less(t, fitnessDelProyectoConMenorSeniority, fitnessDelProyectoConMayorSeniority, "El fitness del proyecto de más seniority debería ser mayor que el del menos")

}
