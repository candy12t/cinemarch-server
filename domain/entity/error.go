package entity

import "errors"

var (
	ErrInvalidLengthMovieTitle = errors.New("movie title must be 1-255 characters")
	ErrInvalidReleaseStatus    = errors.New("invalid release status")
)
