package dto

import "github.com/candy12t/cinemarch-server/lib"

type ScreenMovieSearchCondition map[ScreenMovieSearchKey]string

var ScreenMovieSearchKeys = map[ScreenMovieSearchKey]string{ScreenType: "", TranslateType: "", TreeD: "", Prefecture: "東京都", ScreenDate: lib.Today()}

type ScreenMovieSearchKey string

func (s ScreenMovieSearchKey) String() string {
	return string(s)
}

const (
	MovieID       ScreenMovieSearchKey = "movie_id"
	Prefecture    ScreenMovieSearchKey = "prefecture"
	ScreenDate    ScreenMovieSearchKey = "screen_date"
	ScreenType    ScreenMovieSearchKey = "screen_type"
	TranslateType ScreenMovieSearchKey = "translate_type"
	TreeD         ScreenMovieSearchKey = "three_d"
)
