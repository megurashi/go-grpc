package models

import (
	"app/pb/cities"
	"context"
	"database/sql"
	"log"
)

type City struct {
	Pb  cities.City
	Log *log.Logger
}

// getCity
func (i *City) Get(ctx context.Context, db *sql.DB, in *cities.Id) error {
	query := `SELECT id, name FROM cities WHERE id = $1`
	err := db.QueryRowContext(ctx, query, in.Id).Scan(&i.Pb.Id, &i.Pb.Name)

	if err != nil {
		i.Log.Println("Error on query", err)
		return err
	}

	return nil
}
