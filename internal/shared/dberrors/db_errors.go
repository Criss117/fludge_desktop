package dberrors

import "errors"

func ErrDatabaseConnection() error {
	return errors.New("error de conexión a la base de datos")
}

func ErrDatabaseQuery() error {
	return errors.New("error de consulta a la base de datos")
}
