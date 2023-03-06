package entity

import "errors"

var (
	ErrInvalidReleaseStatus    = errors.New("invalid release status")
	ErrNotChangeReleaseStatus  = errors.New("can not change release status")
	ErrMovieNotFound           = errors.New("movie not found")
	ErrMovieAlreadyExisted     = errors.New("movie has already existed")
	ErrInvalidLengthMovieTitle = errors.New("movie title must be 0-128 characters")

	ErrInvalidLengthCinemaName    = errors.New("cinema name must be 0-128 characters")
	ErrInvalidLengthCinemaAddress = errors.New("cinema address must be 0-128 characters")
	ErrInvalidLengthCinemaWebSite = errors.New("cinema web site must be 0-128 characters")
	ErrInvalidFormatCinemaWebSite = errors.New("invalid format cinema web site")
	ErrCinemaNotFound             = errors.New("cinema not found")
	ErrCinemaAlreadyExisted       = errors.New("cinema has already existed")
)
