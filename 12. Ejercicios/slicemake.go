package main

import "fmt"

func main () {

	/*
	- crea un slice con nombres de alumnos, normalmentes los grupos son de 5 personas, pero luego se inscriben más personas, prevé eso. El máximo puede ser de 10.
	- Cuál sería la capacidad del slice?
	- Cuál sería la capacidad del slice?
	- imprime el contenido del slice con sus índices y su valor.
	*/

	fmt.Println("Creador de grupos de 5 alumnos max. 10 alumnos\n")

	people := make([]string, 5, 10)
	var answer string
	var student string

	for i, _ := range people {
		fmt.Println("Escribe un alumno: ")
		fmt.Scan(&people[i])
	}

	// Esta estructura for {} se ejecutará indefinidamente
	// hasta que se encuentre una instrucción 'break'.
	for {
    	// 1. CONDICIÓN DE LÍMITE: Verificar si la capacidad máxima ya fue alcanzada
    	if len(people) == cap(people) {
        	fmt.Println("\n¡AVISO! Se ha alcanzado la capacidad máxima del grupo (10 alumnos).")
        	break
    	}

    	// 2. SOLICITAR ENTRADA
    	fmt.Println("¿Quieres agregar otro alumno? (Sí/No)")
    	fmt.Scan(&answer)

    	// 3. EVALUAR RESPUESTA
    	if answer == "sí" || answer == "si" || answer == "SI" || answer == "Si" {
        	fmt.Println("Cuál es el nombre del alumno:")
        	fmt.Scan(&student)

        // AGREGAR ALUMNO
        	people = append(people, student)

        // Mostrar estado actual (opcional, para debug)
        	fmt.Println(people)
        	fmt.Println("La longitud es:", len(people))
        	fmt.Println("La capacidad es:", cap(people))
    	} else {
        // DETENER: El usuario eligió no continuar
        	fmt.Println("\nDeteniendo la adición de alumnos por solicitud del usuario.")
        	break
    	}
	}

	/*for ; len(people) < cap(people) ; {
	fmt.Println("Quieres agregar algún alumno?")
	fmt.Scan(&answer)
		if answer == "sí" || answer == "si" || answer == "SI" || answer == "Si" {
			fmt.Println("Cuál es el nombre del alumno: ")
			fmt.Scan(&student)

			people = append(people, student)

			fmt.Println(people)
			fmt.Println("La longitud es: ", len(people))
			fmt.Println("La capacidad es: ", cap(people))
		} else {
			break
		}
			
	}*/

	fmt.Println(people)
	fmt.Println("La longitud es: ", len(people))
	fmt.Println("La capacidad es: ", cap(people))

	fmt.Println(people)
	for indice, valor := range people {
		fmt.Println("La persona #", (indice + 1),  "está en el Índice: ", indice, "y su nombre es: ", valor)
	}
}