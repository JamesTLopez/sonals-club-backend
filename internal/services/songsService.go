package services

import (
	"context"
	"time"
)


type GetAllSongsResponse struct {
	Song
	DisplayName string `json:"display_name"`
}
func (s *Song) GetAllSongs(id string) ([]*GetAllSongsResponse,error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	
	// sql, _, err := psql.Select("songs.id","display_name","song_name","labels","description","duration","color","songs.created_at").From("songs").Join("users ON users.id = songs.user_id").ToSql()
	sql, _, err := psql.Select("songs.id","display_name","song_name","labels","description","duration","color","songs.created_at").From("songs").Join("users ON users.spotify_id = songs.user_id").Where("songs.user_id IN ($1)",id).ToSql()

	if err != nil {
		return nil,err
	}
	
	rows,err := db.QueryContext(ctx, sql,id)

	if err != nil {
		
		return nil, err
	}

	var songs []*GetAllSongsResponse

	for rows.Next() {
		var data GetAllSongsResponse 
		err := rows.Scan(
			&data.Song.ID,
			&data.DisplayName,
			&data.Song.Name,
			&data.Song.Labels,
			&data.Song.Description,
			&data.Song.Duration,
			&data.Song.Color,
			&data.Song.CreatedAt,
		)

		if err != nil {
			return nil,err
		}

		songs = append(songs, &data)
	}

	return songs,nil
}

type GetSongByIdResponse struct {
	Song
	DisplayName string `json:"display_name"`
}
func (s *Song) GetSongById(id string) (*GetSongByIdResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	sql, _, err := psql.Select("songs.id","display_name","song_name","labels","description","duration","color","songs.created_at").From("songs").Join("users ON users.spotify_id = songs.user_id").Where("songs.id IN ($1)").ToSql()

	if err != nil {
		return nil,err
	}
	
	var song GetSongByIdResponse

	row :=  db.QueryRowContext(ctx,sql,id)
	
	err = row.Scan(
		&song.ID,
		&song.DisplayName,
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


func (s *Song) CreateSong(id string, song Song) (*Song,error) {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)

	defer cancel()
	query := `
		INSERT INTO songs (user_id, song_name, labels, description, duration, color) 
		VALUES ($1,$2,$3,$4,$5,$6) returning *
	`

	 
	_, err := db.ExecContext(
		ctx,
		query,
		id, // TODO: when auth is implemented with jwt, this is where we would put it
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
