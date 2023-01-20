package entity

import "errors"

var (
	ErrInvalidLengthMovieTitle   = errors.New("movie title must be 1-255 characters")
	ErrInvalidMovieReleaseStatus = errors.New("invalid release status")

	ErrInvalidLengthCinemaName    = errors.New("cinema name must be 1-255 characters")
	ErrInvalidLengthCinemaAddress = errors.New("address name must be 1-255 characters")
	ErrInvalidLengthCinemaURL     = errors.New("cinema url must be 1-255 characters")

	ErrInvalidLengthScreeningTypeName = errors.New("screening type name must be 1-255 characters")
)
