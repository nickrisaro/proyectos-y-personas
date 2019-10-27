# Proyectos y Personas

Es un trabajo práctico para la materia Modelos y Optimización III de la FIUBA.

La idea es asistir en el armado de equipos de trabajo para diversos proyectos.

Cada persona tiene una habilidad blanda, una habilidad dura, un seniority y un sueldo.

Cada proyecto requiere distintas cantidades de personas con distintos skills y tiene un presupuesto.

El objetivo es armar el mejor equipo posible para cada proyecto sin excederse del presupuesto, con la mayor cantidad de skills duros y blandos y el mayor seniority posible.

## Modelo

Modelé el problema como si cada proyecto fuera una mochila y las personas los elementos que pueden ir dentro de esa mochila.

El presupuesto del proyecto es el peso máximo que soporta la mochila-proyecto.

El sueldo es el peso de cada elemento-personas

Los skills de la persona y las necesidades del proyecto determinan el valor que aporta el elemento-persona a la mochila-proyecto

Para generar las posibles soluciones se utiliza un algoritmo genético que se inicia con unas soluciones aleatorias y en cada ciclo reemplaza el 25% de las soluciones (las de menor fitness) con una combinación de soluciones generadas a partir del 50% de las soluciones (las de mayor fitness). Luego se obtiene de todas las soluciones que se fueron generando la de mayor fitness.

## Limitaciones/Posibilidades a futuro

Las personas tienen un sólo skill duro y un sólo skill blando, y no se hace distinción de seniority entre ellos.

Los skills no son configurables.

## TODO
* Generar una solución bastante buena
* Guardar la solución en un archivo
* Experimentar con otras maneras de armar las soluciones hijas
* Explorar con mutaciones en las soluciones
* Exportar la solución a distintos tipos de archivo (CSV)
* Cargar empresa con personas y proyecto de un archivo
* Cargar los pesos de las características de un archivo
* Permitir que un proyecto se exceda de su presupuesto
* ¿Me interesa ver las soluciones parciales?