package services

import "context"


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