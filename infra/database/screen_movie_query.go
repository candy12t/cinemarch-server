package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/candy12t/cinemarch-server/query/dto"
	"github.com/candy12t/cinemarch-server/query/service"
	"github.com/jmoiron/sqlx"
)

type ScreenMovieQueryService struct {
	db *sqlx.DB
}

var _ service.ScreenMovie = (*ScreenMovieQueryService)(nil)

func NewScreenMovieQueryService(db *sqlx.DB) *ScreenMovieQueryService {
	return &ScreenMovieQueryService{
		db: db,
	}
}

func (qs *ScreenMovieQueryService) Search(ctx context.Context, condition dto.ScreenMovieSearchCondition) (dto.ScreenMovies, error) {
	screenMovies := dto.ScreenMovies{}
	query := `
		SELECT
			screen_movies.id AS id,
			cinemas.name AS cinema_name,
			movies.title AS movie_title,
			screen_movies.screen_type AS screen_type,
			screen_movies.translate_type AS translate_type,
			screen_movies.three_d AS three_d,
			CAST(concat('[',
				GROUP_CONCAT(
					JSON_OBJECT(
						'id', schedules.id,
						'screen_movie_id', schedules.screen_movie_id,
						'start_time', CAST(schedules.start_time AS CHAR),
						'end_time', CAST(schedules.end_time AS CHAR)
					)
				),
			']') AS JSON) AS schedules
		FROM
			screen_movies
		INNER JOIN screen_schedules AS schedules
			ON screen_movies.id = schedules.screen_movie_id
		INNER JOIN cinemas
			ON screen_movies.cinema_id = cinemas.id
		INNER JOIN movies
			ON screen_movies.movie_id = movies.id
		WHERE
				movie_id = ?
			AND cinemas.prefecture = ?
	`
	// AND ? <= start_time AND start_time < ?
	query = fmt.Sprintf("%s %s", query, `GROUP BY id, cinema_id, movie_id, screen_type, translate_type, three_d`)
	if err := qs.db.SelectContext(ctx, &screenMovies, query, condition[dto.MovieID], condition[dto.Prefecture]); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("screen movie not found")
		}
		return nil, err
	}

	return screenMovies, nil
}
