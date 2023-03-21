package database

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/candy12t/cinemarch-server/domain/entity"
	"github.com/candy12t/cinemarch-server/domain/repository"
	"github.com/candy12t/cinemarch-server/lib"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ScreenMovieRepository struct {
	db *sqlx.DB
}

var _ repository.ScreenMovie = (*ScreenMovieRepository)(nil)

func NewScreenMovieRepository(db *sqlx.DB) *ScreenMovieRepository {
	return &ScreenMovieRepository{
		db: db,
	}
}
func (r *ScreenMovieRepository) FindByUniqKey(ctx context.Context, cinemaID, movieID entity.UUID, screenType entity.ScreenType, translateType entity.TranslateType, threeD bool) (*entity.ScreenMovie, error) {
	dto := new(screenMovieDTO)
	query := `
		SELECT
			screen_movies.id AS id,
			screen_movies.cinema_id AS cinema_id,
			screen_movies.movie_id AS movie_id,
			screen_movies.screen_type AS screen_type,
			screen_movies.translate_type AS translate_type,
			screen_movies.three_d AS three_d,
			CAST(concat('[',
				GROUP_CONCAT(JSON_OBJECT('id', schedules.id, 'screen_movie_id', schedules.screen_movie_id, 'start_time', CAST(schedules.start_time AS CHAR), 'end_time', CAST(schedules.end_time AS CHAR))),
				']') AS JSON) AS schedules
		FROM
			screen_movies
		INNER JOIN screen_schedules AS schedules
			ON screen_movies.id = schedules.screen_movie_id
		WHERE
				cinema_id = ?
			AND movie_id = ?
			AND screen_type = ?
			AND translate_type = ?
			AND three_d = ?
		GROUP BY cinema_id, movie_id, screen_type, translate_type, three_d
	`
	if err := r.db.GetContext(ctx, dto, query, cinemaID, movieID, screenType, translateType, threeD); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrScreenMovieNotFound
		}
		return nil, err
	}
	return dtoToScreenMovie(dto)
}

func (r *ScreenMovieRepository) Search(ctx context.Context) (entity.ScreenMovies, error) {
	return nil, nil
}

func (r *ScreenMovieRepository) CreateScreenMovie(ctx context.Context, screenMovie *entity.ScreenMovie) error {
	dto := screenMovieToDTO(screenMovie)
	query := `INSERT INTO screen_movies (id, cinema_id, movie_id, screen_type, translate_type, three_d) VALUES (:id, :cinema_id, :movie_id, :screen_type, :translate_type, :three_d)`
	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr); mysqlErr.Number == mysqlDuplicateEntryErrorCode {
			return entity.ErrScreenMovieAlreadyExisted
		}
		return err
	}
	return nil
}

func (r *ScreenMovieRepository) CreateScreenSchedules(ctx context.Context, screenSchedules entity.ScreenSchedules) error {
	query := `INSERT INTO screen_schedules (id, screen_movie_id, start_time, end_time) VALUES (:id, :screen_movie_id, :start_time, :end_time)`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}
	for _, schedule := range screenSchedules {
		if _, err := stmt.ExecContext(ctx, screenScheduleToDTO(schedule)); err != nil {
			var mysqlErr *mysql.MySQLError
			if errors.As(err, &mysqlErr); mysqlErr.Number == mysqlDuplicateEntryErrorCode {
				return entity.ErrScreenScheduleAlreadyExisted
			}
			return err
		}
	}
	return nil
}

type screenMovieDTO struct {
	ID            string        `db:"id"`
	CinemaID      string        `db:"cinema_id"`
	MovieID       string        `db:"movie_id"`
	ScreenType    string        `db:"screen_type"`
	TranslateType string        `db:"translate_type"`
	ThreeD        bool          `db:"three_d"`
	Schedules     schedulesDTOs `db:"schedules"`
}

func dtoToScreenMovie(dto *screenMovieDTO) (*entity.ScreenMovie, error) {
	schedules := make(entity.ScreenSchedules, 0, len(dto.Schedules))
	for _, schedule := range dto.Schedules {
		startTime, err := lib.ParseDateTime(schedule.StartTime)
		if err != nil {
			return nil, err
		}
		endTime, err := lib.ParseDateTime(schedule.EndTime)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, &entity.ScreenSchedule{
			ID:            entity.UUID(schedule.ID),
			ScreenMovieID: entity.UUID(schedule.ScreenMovieID),
			StartTime:     startTime,
			EndTime:       endTime,
		})
	}

	return &entity.ScreenMovie{
		ID:              entity.UUID(dto.ID),
		CinemaID:        entity.UUID(dto.CinemaID),
		MovieID:         entity.UUID(dto.MovieID),
		ScreenType:      entity.ScreenType(dto.ScreenType),
		TranslateType:   entity.TranslateType(dto.TranslateType),
		TreeD:           dto.ThreeD,
		ScreenSchedules: schedules,
	}, nil
}

func screenMovieToDTO(screenMovie *entity.ScreenMovie) *screenMovieDTO {
	return &screenMovieDTO{
		ID:            screenMovie.ID.String(),
		CinemaID:      screenMovie.CinemaID.String(),
		MovieID:       screenMovie.MovieID.String(),
		ScreenType:    screenMovie.ScreenType.String(),
		TranslateType: screenMovie.TranslateType.String(),
		ThreeD:        screenMovie.TreeD,
	}
}

type scheduleDTO struct {
	ID            string `db:"id" json:"id"`
	ScreenMovieID string `db:"screen_movie_id" json:"screen_movie_id"`
	StartTime     string `db:"start_time" json:"start_time"`
	EndTime       string `db:"end_time" json:"end_time"`
}

func screenScheduleToDTO(screenSchedule *entity.ScreenSchedule) *scheduleDTO {
	return &scheduleDTO{
		ID:            screenSchedule.ID.String(),
		ScreenMovieID: screenSchedule.ScreenMovieID.String(),
		StartTime:     lib.FormatDateTime(screenSchedule.StartTime),
		EndTime:       lib.FormatDateTime(screenSchedule.EndTime),
	}
}

type schedulesDTOs []*scheduleDTO

func (s *schedulesDTOs) Scan(val any) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &s)
		return nil
	case string:
		json.Unmarshal([]byte(v), &s)
		return nil
	default:
		return fmt.Errorf("Unsupported type: %T", v)
	}
}

func (s *schedulesDTOs) Value() (driver.Value, error) {
	return json.Marshal(s)
}
