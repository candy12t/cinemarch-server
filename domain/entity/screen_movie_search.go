package entity

import "github.com/candy12t/cinemarch-server/lib"

type ScreenMovieSearchCondition map[ScreenMovieSearchKey]string

type ScreenMovieSearchKey string

var searchKeys = map[string]string{"screen_type": "", "translate_type": "", "three_d": "", "prefecture": "東京都", "screen_date": lib.Today()}

const ()
