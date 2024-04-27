package errors

import "errors"

var (
	ErrFileNotFound     = errors.New("berkas tidak dapat ditemukan")
	ErrFileDeleteFailed = errors.New("gagal menghapus file")
)
