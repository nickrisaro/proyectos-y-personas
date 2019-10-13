package proyecto_test

import (
	"testing"

	"github.com/nickrisaro/proyectos-y-personas/persona"
	"github.com/nickrisaro/proyectos-y-personas/proyecto"
	"github.com/stretchr/testify/assert"
)

func TestUnProyectoSinPersonasRequeridasNoTieneFitness(t *testing.T) {

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	unProyecto := proyecto.New(personasRequeridas, 1.0)

	_, err := unProyecto.Fitness()

	assert.Error(t, err, "Se esperaba un error")
}

func TestUnProyectoSinPersonasAsignadasTieneFitnessNegativo(t *testing.T) {

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	unProyecto := proyecto.New(personasRequeridas, 1.0)

	fitness, _ := unProyecto.Fitness()

	assert.Less(t, fitness, 0.0, "El fitness del proyecto debería ser negativo")
}

func TestUnProyectoConIgualCantidadDePersonasRequeridasQueAsignadasTieneFitnessMayorOIgualQueCero(t *testing.T) {

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	unProyecto := proyecto.New(personasRequeridas, 1.0)
	unaPersona := persona.New(1.0, persona.Senior, persona.Desarrollo)
	unProyecto.AsignarPersona(unaPersona)

	fitness, _ := unProyecto.Fitness()

	assert.GreaterOrEqual(t, fitness, 0.0, "El fitness del proyecto debería ser mayor o igual que cero")
}

func TestUnProyectoSinPresupuestoAsignadoNoTieneFitness(t *testing.T) {

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	unProyecto := proyecto.New(personasRequeridas, 0.0)

	_, err := unProyecto.Fitness()

	assert.Error(t, err, "Se esperaba un error")
}

func TestUnProyectoQueSeExcedeDelPresupuestoTieneFitnessNegativo(t *testing.T) {

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	unProyecto := proyecto.New(personasRequeridas, 1.0)
	unaPersona := persona.New(2.0, persona.Junior, persona.Desarrollo)
	unProyecto.AsignarPersona(unaPersona)

	fitness, _ := unProyecto.Fitness()

	assert.Less(t, fitness, 0.0, "El fitness del proyecto debería ser negativo")
}

func TestUnProyectoConMenorGastoDeSueldosTieneMejorFitnessQueUnoConMayorGastoDeSueldos(t *testing.T) {

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	proyectoQueGastaMenos := proyecto.New(personasRequeridas, 1.0)
	unaPersona := persona.New(0.7, persona.Senior, persona.Desarrollo)
	proyectoQueGastaMenos.AsignarPersona(unaPersona)
	proyectoQueGastaMas := proyecto.New(personasRequeridas, 1.0)
	otraPersona := persona.New(0.9, persona.Senior, persona.Desarrollo)
	proyectoQueGastaMas.AsignarPersona(otraPersona)

	fitnessDelProyectoQueGastaMenos, _ := proyectoQueGastaMenos.Fitness()
	fitnessDelProyectoQueGastaMas, _ := proyectoQueGastaMas.Fitness()

	assert.Greater(t, fitnessDelProyectoQueGastaMenos, fitnessDelProyectoQueGastaMas, "El fitness del proyecto más barato debería ser mayor que el del más caro")

}

func TestUnProyectoConPersonasDeMasSeniorityTieneMejorFitnessQueUnoConPersonasDeMenorSeniority(t *testing.T) {

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	proyectoConMenosSeniority := proyecto.New(personasRequeridas, 1.0)
	unaPersona := persona.New(0.9, persona.Junior, persona.Desarrollo)
	proyectoConMenosSeniority.AsignarPersona(unaPersona)
	proyectoConMAyorSeniority := proyecto.New(personasRequeridas, 1.0)
	otraPersona := persona.New(0.9, persona.Senior, persona.Desarrollo)
	proyectoConMAyorSeniority.AsignarPersona(otraPersona)

	fitnessDelProyectoConMenorSeniority, _ := proyectoConMenosSeniority.Fitness()
	fitnessDelProyectoConMayorSeniority, _ := proyectoConMAyorSeniority.Fitness()

	assert.Less(t, fitnessDelProyectoConMenorSeniority, fitnessDelProyectoConMayorSeniority, "El fitness del proyecto de más seniority debería ser mayor que el del menos")

}

func TestUnProyectoQueCubreTodosLosHardSkillsRequeridosTieneMejorFitnessQueUnoQueNoCubreTodos(t *testing.T) {

	unDesarrolladorJunior := persona.New(0.9, persona.Junior, persona.Desarrollo)
	unaDesarrolladoraJunior := persona.New(0.9, persona.Junior, persona.Desarrollo)
	otroDesarrolladorJunior := persona.New(0.9, persona.Junior, persona.Desarrollo)
	unaDiseñadoraJunior := persona.New(0.9, persona.Junior, persona.Diseño)
	unaOperadoraJunior := persona.New(0.9, persona.Junior, persona.Operaciones)

	personasRequeridas := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas.Desarrollo(1)
	personasRequeridas.Diseño(1)
	personasRequeridas.Operaciones(1)
	proyectoQueCubreTodosLosHardSkills := proyecto.New(personasRequeridas, 1.0)
	proyectoQueNoCubreTodosLosHardSkills := proyecto.New(personasRequeridas, 1.0)

	proyectoQueCubreTodosLosHardSkills.AsignarPersona(unDesarrolladorJunior)
	proyectoQueCubreTodosLosHardSkills.AsignarPersona(unaDiseñadoraJunior)
	proyectoQueCubreTodosLosHardSkills.AsignarPersona(unaOperadoraJunior)

	proyectoQueNoCubreTodosLosHardSkills.AsignarPersona(unDesarrolladorJunior)
	proyectoQueNoCubreTodosLosHardSkills.AsignarPersona(unaDesarrolladoraJunior)
	proyectoQueNoCubreTodosLosHardSkills.AsignarPersona(otroDesarrolladorJunior)

	fitnessDelProyectoQueCubreTodosLosHardSkills, _ := proyectoQueCubreTodosLosHardSkills.Fitness()
	fitnessDelProyectoQueNoCubreTodosLosHardSkills, _ := proyectoQueNoCubreTodosLosHardSkills.Fitness()

	assert.Less(t, fitnessDelProyectoQueNoCubreTodosLosHardSkills, fitnessDelProyectoQueCubreTodosLosHardSkills, "El fitness del proyecto que cubre todos los  HardSkills debería ser mayor que el del que no los cubre")
}
