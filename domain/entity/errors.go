package entity

import "errors"

var (
	ErrInvalidLengthCinemaName = errors.New("cinema name must be 1-64 characters")
	ErrInvalidLengthAddress    = errors.New("address must be 1-255 characters")
	ErrInvalidLengthWebSiteURL = errors.New("web site url must be 1-255 characters")

	ErrInvalidLengthMovieTitle          = errors.New("movie title must be 1-64 characters")
	ErrInvalidReleaseStatus             = errors.New("invalid release status")
	ErrNotAllowChangeMovieReleaseStatus = errors.New("not allow change release status")

	ErrInvalidScreenType        = errors.New("invalid screen type")
	ErrInvalidSubtitleOrDubbing = errors.New("invalid")

	ErrMovieNotFound       = errors.New("movie not found")
	ErrMovieAlreadyExisted = errors.New("movie has already existed")

	ErrCinemaNotFound       = errors.New("cinema not found")
	ErrCinemaAlreadyExisted = errors.New("cinema has already existed")

	ErrScreenMovieNotFound       = errors.New("screen movie not found")
	ErrScreenMovieAlreadyExisted = errors.New("screen movie has already existed")

	ErrScreenTypeNotFound       = errors.New("screen type not found")
	ErrScreenTypeAlreadyExisted = errors.New("screen type has already existed")
)
