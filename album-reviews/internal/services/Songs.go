package services

import (
	"context"
	"time"
)


func (s *Song) GetAllSongs() ([]*Song,error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()


	query := `SELECT id, name, description, created_at FROM songs`

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
			&song.Description,
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
		SELECT id, name, description, created_at FROM songs WHERE id = $1
	`

	var song Song

	row :=  db.QueryRowContext(ctx,query,id)
	
	err := row.Scan(
		&song.ID,
		&song.Name,
		&song.Description,
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
		INSERT INTO songs (user_id, name, description) 
		VALUES ($1,$2,$3) returning *
	`
	 
	_, err := db.ExecContext(
		ctx,
		query,
		1, // TODO: when auth is implemented with jwt, this is where we would put it
		song.Name,
		song.Description)
	
	if err != nil {
		return nil, err
	}
	

	return &song,nil

}


func (s *Song) UpdateSong(id string, body Song) (*Song,error) {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)

	defer cancel()
	query := `
		UPDATE songs
		SET
			name = $1,
			description = $2,
			updated_at = $3
		WHERE id = $4
		returning *

	`

	 
	_, err := db.ExecContext(
		ctx,
		query,
		body.Name,
		body.Description,
		time.Now(),
		id,
	)
	
	if err != nil {
		return nil, err
	}
	

	return &body, nil

}
