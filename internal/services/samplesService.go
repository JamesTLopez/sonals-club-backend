package services

import (
	"context"
	"time"
)


func (sample *Sample) GetAllSamples() ([]*Sample, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, sample_name, labels, description, duration, created_at FROM samples`

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
		INSERT INTO Samples (user_id, sample_name, created_at, updated_at) 
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


func (sample *Sample) UpdateSample(id string, sampleData Sample) (*Sample,error) {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)
	defer cancel()

	query := `
		UPDATE Samples SET 
			sample_name = $1,
			updated_at = $2
			WHERE id = $3
			returning *
	`

	_,err := db.ExecContext(ctx, query,sampleData.Name,time.Now(),id)

	if err != nil{
		return nil,err
	}

	return &sampleData,nil

}

func (sample *Sample) DeleteSample(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)
	defer cancel()

	query := `
		DELETE FROM Samples WHERE id = $1 	
	`

	_,err := db.ExecContext(
		ctx,
		query,
		id,
	)

	if err != nil {
		return  err
	}

	return nil
}