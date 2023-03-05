package entity

import "errors"

var (
	ErrInvalidReleaseStatus    = errors.New("invalid release status")
	ErrNotChangeReleaseStatus  = errors.New("can not change release status")
	ErrMovieNotFound           = errors.New("movie not found")
	ErrMovieAlreadyExisted     = errors.New("movie has already existed")
	ErrInvalidLengthMovieTitle = errors.New("movie title must be 0-128 characters")
)
