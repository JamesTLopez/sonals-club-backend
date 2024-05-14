package services

import (
	"context"
	"time"
)


func (s *Song) GetAllSongs() ([]*Song,error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()


	query := `SELECT id, song_name, labels, description, duration, color, created_at FROM songs`

	rows,err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var songs []*Song
	for rows.Next() {
		var song Song 
		err := rows.Scan(
			&song.ID,
			&song.Name,
			&song.Labels,
			&song.Description,
			&song.Duration,
			&song.Color,
			&song.CreatedAt,
		)

		if err != nil {
			return nil,err
		}

		songs = append(songs, &song)
	}

	return songs,nil
}

func (s *Song) GetSongById(id string) (*Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT id, song_name, labels, description, duration, color, created_at FROM songs WHERE id = $1
	`

	var song Song

	row :=  db.QueryRowContext(ctx,query,id)
	
	err := row.Scan(
		&song.ID,
		&song.Name,
		&song.Labels,
		&song.Description,
		&song.Duration,
		&song.Color,
		&song.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &song, err
}


func (s *Song) CreateSong(song Song) (*Song,error) {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)

	defer cancel()
	query := `
		INSERT INTO songs (user_id, song_name, labels, description, duration, color) 
		VALUES ($1,$2,$3,$4,$5,$6) returning *
	`

	 
	_, err := db.ExecContext(
		ctx,
		query,
		1, // TODO: when auth is implemented with jwt, this is where we would put it
		song.Name,
		song.Labels,
		song.Description,
		song.Duration,
		song.Color,
	)
	
	if err != nil {
		return nil, err
	}
	

	return &song,nil

}

// TODO: added update labels functionality
func (s *Song) UpdateSong(id string, body Song) (*Song,error) {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)

	defer cancel()
	query := `
		UPDATE songs
		SET
			song_name = $1,
			labels = $2,
			description = $3,
			color = $4,
			updated_at = $5
		WHERE id = $6
		returning *

	`
	_, err := db.ExecContext(
		ctx,
		query,
		body.Name,
		body.Labels,
		body.Description,
		body.Color,
		time.Now(),
		id,
	)
	
	if err != nil {
		return nil, err
	}
	

	return &body, nil

}


func (s *Song) DeleteSong(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)

	defer cancel()
	query := `
		DELETE FROM songs WHERE id = $1
	`
	_, err := db.ExecContext(
		ctx,
		query,
		id,
	)
	
	if err != nil {
		return err
	}
	

	return nil

}
