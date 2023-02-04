package entity

import "errors"

var (
	ErrInvalidReleaseStatus             = errors.New("invalid release status")
	ErrNotAllowChangeMovieReleaseStatus = errors.New("not allow change release status")

	ErrInvalidLengthMovieTitle = errors.New("movie title must be 1-255 characters")

	ErrInvalidLengthCinemaName    = errors.New("cinema name must be 1-255 characters")
	ErrInvalidLengthCinemaAddress = errors.New("address name must be 1-255 characters")
	ErrInvalidLengthCinemaURL     = errors.New("cinema url must be 1-255 characters")

	ErrInvalidLengthScreenTypeName = errors.New("screen type name must be 1-255 characters")

	ErrMovieNotFound       = errors.New("movie not found")
	ErrMovieAlreadyExisted = errors.New("movie has already existed")

	ErrCinemaNotFound       = errors.New("cinema not found")
	ErrCinemaAlreadyExisted = errors.New("cinema has already existed")

	ErrScreenMovieNotFound       = errors.New("screen movie not found")
	ErrScreenMovieAlreadyExisted = errors.New("screen movie has already existed")

	ErrScreenTypeNotFound       = errors.New("screen type not found")
	ErrScreenTypeAlreadyExisted = errors.New("screen type has already existed")
)
