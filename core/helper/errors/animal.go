package errors

import "errors"

var (
	ErrAnimalNotFound = errors.New("hewan tidak ditemukan")
	ErrWrongAnimal    = errors.New("hewan ini bukan target anda saat ini")
)
