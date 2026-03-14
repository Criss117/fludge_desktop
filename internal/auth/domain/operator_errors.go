package domain

import "errors"

func ErrOperatorAlreadyExists() error {
	return errors.New("El nombre de usuario ya está en uso")
}

func ErrEmailAlreadyExists() error {
	return errors.New("El correo electrónico ya está en uso")
}

func ErrUsernameAlreadyExists() error {
	return errors.New("El nombre de usuario ya está en uso")
}

func ErrOperatorNotFoundById() error {
	return errors.New("El operador no se encuentra")
}

func ErrOperatorNotFoundByEmail() error {
	return errors.New("El correo electrónico no se encuentra")
}

func ErrOperatorNotFoundByUsername() error {
	return errors.New("El usuario no se encuentra")
}

func ErrOperatorNotFound() error {
	return errors.New("El operador no se encuentra")
}

func ErrOperatorNotCreated() error {
	return errors.New("El operador no se pudo crear")
}

func ErrInvalidCredentials() error {
	return errors.New("Credenciales inválidas")
}
