package model

import (
	"errors"
	"strconv"
)

type datosClave struct {
	ISBN int // Código de identificación

	autor string
	titulo string
}

// Crea nuevos datos válidos
func nuevosDatosClave(ISBN int, autor string, titulo string) (*datosClave, error){
	// Comprueba que el ISBN tenga 13 dígitos
	if len(strconv.Itoa(ISBN)) != 13 {
		return nil, errors.New("El ISBN debe tener 13 dígitos")
	}

	// Datos válidos: se crean los datos sin devolver errores
	return &datosClave{ISBN, autor, titulo}, nil
}
