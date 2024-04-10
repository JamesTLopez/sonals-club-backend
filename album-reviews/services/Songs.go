package services

import (
	"context"
	"time"
)

type Song struct {
	ID string `json:"id"`
	User_id string `json:"user_id"` // TODO
	Name string `json:"name"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"` // TODO
}

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


func (s *Song) CreateSong(song Song) (*Song,error) {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)

	defer cancel()
	query := `
		INSERT INTO songs (user_id, name,description) 
		VALUES ($1,$2,$3) returning *
	`
	 
	_, err := db.ExecContext(
		ctx,
		query,
		song.User_id,
		song.Name,
		song.Description)
	
	if err != nil {
		return nil, err
	}
	

	return &song,nil

}