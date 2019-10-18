package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/nickrisaro/proyectos-y-personas/persona"
	"github.com/nickrisaro/proyectos-y-personas/proyecto"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("hola")
	inicio := time.Now()
	personasRequeridas1 := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas1.Desarrollo(1)
	personasRequeridas1.Diseño(1)
	personasRequeridas2 := proyecto.NewPersonasRequeridasPorSkill()
	personasRequeridas2.Operaciones(1)
	unProyecto := proyecto.New("Proyecto uno", personasRequeridas1, 2.0)
	otroProyecto := proyecto.New("Proyecto dos", personasRequeridas2, 1.0)
	ana := persona.New("Ana", 1.0, persona.Senior, persona.Desarrollo, persona.Negociacion)
	juan := persona.New("Juan", 1.0, persona.Junior, persona.Operaciones, persona.Investigacion)
	maría := persona.New("María", 1.0, persona.Junior, persona.Diseño, persona.Investigacion)

	solucionExacta(unProyecto, otroProyecto, ana, juan, maría)

	unaSolucion := [3]int{} // 3 personas

	mejorFitness := math.Inf(-1)
	var mejorSolucion [3]int
	for i := 0; i < 3; i++ {
		unaSolucion[i] = rand.Intn(3) // 2 proyectos
	}

	if unaSolucion[0] == 1 { // ana va al proyecto 1
		unProyecto.AsignarPersona(ana)
	}
	if unaSolucion[0] == 2 { // ana va al proyecto 2
		otroProyecto.AsignarPersona(ana)
	}

	if unaSolucion[1] == 1 { // juan va al proyecto 1
		unProyecto.AsignarPersona(juan)
	}
	if unaSolucion[1] == 2 { // juan va al proyecto 2
		otroProyecto.AsignarPersona(juan)
	}

	if unaSolucion[2] == 1 { // maría va al proyecto 1
		unProyecto.AsignarPersona(maría)
	}
	if unaSolucion[2] == 2 { // maría va al proyecto 2
		otroProyecto.AsignarPersona(maría)
	}

	fitnessUno, _ := unProyecto.Fitness()
	fitnessDos, _ := otroProyecto.Fitness()

	if fitnessUno+fitnessDos > mejorFitness {
		mejorFitness = fitnessUno + fitnessDos
		mejorSolucion = unaSolucion
	}

	for pruebas := 0; pruebas < 10; pruebas++ {

		unProyecto = unProyecto.Clonar()
		otroProyecto = otroProyecto.Clonar()

		for i := 0; i < 2; i++ {

			if unaSolucion[i] < 2 {
				unaSolucion[i] = unaSolucion[i] + 1
			} else {
				unaSolucion[i] = 0
			}
		}

		if unaSolucion[0] == 1 { // ana va al proyecto 1
			unProyecto.AsignarPersona(ana)
		}
		if unaSolucion[0] == 2 { // ana va al proyecto 2
			otroProyecto.AsignarPersona(ana)
		}

		if unaSolucion[1] == 1 { // juan va al proyecto 1
			unProyecto.AsignarPersona(juan)
		}
		if unaSolucion[1] == 2 { // juan va al proyecto 2
			otroProyecto.AsignarPersona(juan)
		}

		if unaSolucion[2] == 1 { // maría va al proyecto 1
			unProyecto.AsignarPersona(maría)
		}
		if unaSolucion[2] == 2 { // maría va al proyecto 2
			otroProyecto.AsignarPersona(maría)
		}

		fitnessUno, _ := unProyecto.Fitness()
		fitnessDos, _ := otroProyecto.Fitness()

		if fitnessUno+fitnessDos > mejorFitness {
			mejorFitness = fitnessUno + fitnessDos
			mejorSolucion = unaSolucion
		}
	}

	fmt.Println("Fitness no exacta", mejorFitness)
	fmt.Println("Solución no exacta", mejorSolucion)
	fin := time.Now()
	fmt.Printf("pasaron %v segundos\n", fin.Sub(inicio).Seconds())
}

func solucionExacta(unProyecto, otroProyecto *proyecto.Proyecto, ana, juan, maría *persona.Persona) {
	personas := 3
	proyectos := 2 + 1
	combinaciones := math.Pow(float64(proyectos), float64(personas))
	personasAProyectos := make([][]int, proyectos)
	for i := range personasAProyectos {
		personasAProyectos[i] = make([]int, int(combinaciones))
	}

	valores := [3]int{0, 0, 0}
	for j := 0; j < int(combinaciones); j++ {
		for i := 0; i < personas; i++ {
			personasAProyectos[i][j] = valores[i]
		}
		if valores[2] < 2 {
			valores[2] = valores[2] + 1
		} else {
			valores[2] = 0
			if valores[1] < 2 {
				valores[1] = valores[1] + 1
			} else {
				valores[1] = 0
				if valores[0] < 2 {
					valores[0] = valores[0] + 1
				} else {
					valores[0] = 0
				}
			}
		}
	}
	mejorFitness := math.Inf(-1)
	indiceMejorFitness := 0

	for j := 0; j < int(combinaciones); j++ {
		unProyecto := unProyecto.Clonar()
		otroProyecto := otroProyecto.Clonar()
		for i := 0; i < personas; i++ {
			if personasAProyectos[i][j] == 1 { // va al proyecto 1
				if i == 0 {
					unProyecto.AsignarPersona(ana)
				} else if i == 1 {
					unProyecto.AsignarPersona(juan)
				} else {
					unProyecto.AsignarPersona(maría)
				}
			}
			if personasAProyectos[i][j] == 2 { // va al proyecto 2
				if i == 0 {
					otroProyecto.AsignarPersona(ana)
				} else if i == 1 {
					otroProyecto.AsignarPersona(juan)
				} else {
					otroProyecto.AsignarPersona(maría)
				}
			}
		}
		fitnessUno, _ := unProyecto.Fitness()
		fitnessDos, _ := otroProyecto.Fitness()

		if fitnessUno+fitnessDos > mejorFitness {
			mejorFitness = fitnessUno + fitnessDos
			indiceMejorFitness = j
		}
	}

	fmt.Printf("El mejor fitness es %v en la combinacion %v\n", mejorFitness, indiceMejorFitness)
	fmt.Printf("La persona 1 va al proyecto %v, la 2 al proyecto %v y la 3 al proyecto %v\n", personasAProyectos[0][indiceMejorFitness], personasAProyectos[1][indiceMejorFitness], personasAProyectos[2][indiceMejorFitness])
}
