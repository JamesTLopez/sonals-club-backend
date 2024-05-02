package services

import (
	"context"
	"time"
)


func (sample *Sample) GetAllSamples() ([]*Sample, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, name, labels, description, duration, created_at FROM songs`

	var arrayOfSamples []*Sample

	_,err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	return arrayOfSamples,nil
}


func (sample *Sample) CreateSample(sampleData Sample) (*Sample,error) {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)
	defer cancel()
	
	query := `
		INSERT INTO Samples (user_id, name, created_at, updated_at) 
		VALUES ($1, $2, $3, $4) returning *
	`

	_,err := db.ExecContext(
		ctx,
		query,
		1,
		sampleData.Name,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return nil, err
	}


	return &sampleData,nil
}	