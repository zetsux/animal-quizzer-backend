package errors

import "errors"

var (
	ErrUsernameAlreadyExists = errors.New("username sudah digunakan")
	ErrUserNotFound          = errors.New("pengguna tidak ditemukan")
	ErrUserNoPicture         = errors.New("pengguna tidak memiliki gambar")
)
