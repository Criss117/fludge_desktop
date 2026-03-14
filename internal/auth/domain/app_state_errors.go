package domain

import "errors"

func ErrAppStateNotFound() error {
	return errors.New("El estado de la aplicación no se encuentra")
}

func ErrAppStateNotUpdated() error {
	return errors.New("No se ha actualizado el estado de la aplicación")
}

func ErrAppStateNotSetOperator() error {
	return errors.New("No se ha establecido el operador activo")
}
