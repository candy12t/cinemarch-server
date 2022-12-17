package entity

import "errors"

var (
	ErrInvalidLengthMovieTitle = errors.New("movie title must be 1-255 characters")
	ErrInvalidReleaseStatus    = errors.New("invalid release status")

	ErrInvalidLengthCinemaName = errors.New("cinema name must be 1-255 characters")
	ErrInvalidLengthPrefecture = errors.New("prefecture must be 1-255 characters")
	ErrInvalidLengthAddress    = errors.New("address name must be 1-255 characters")
	ErrInvalidLengthCinemaURL  = errors.New("cinema url must be 1-255 characters")
)
