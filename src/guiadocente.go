package model

import (
	"errors"
)

type GuiaDocente struct{
	Anio int
	Asignatura string
	
	BibliografiaBasica []libro	    // Lista de libros básicos
	BibliografiaComplementaria []libro  // Lista de libros complementarios
}


// Crea una nueva guía docente con datos válidos
func NuevaGuiaDocente(anio int, asignatura string, basica []Libro, complementaria []Libro) (*GuiaDocente, error){
	anioValido := EsPositivo(anio)

	// Ambos datos inválidos
	if !anioValido {
		return nil, errors.New("El año debe ser positivo")
	}

	// Datos válidos: se crea la guía docente sin devolver errores
	return &GuiaDocente{anio, asignatura, basica, complementaria}, nil
}
