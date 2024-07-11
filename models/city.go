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

func (u *City) Get(ctx context.Context, db *sql.DB, in *cities.Id) error {
	query := `SELECT id, name FROM cities WHERE id = $1`
	err := db.QueryRowContext(ctx, query, in.Id).Scan(&u.Pb.Id, &u.Pb.Name)
	if err != nil {
		u.Log.Println("Error on query", err)
		return err
	}
	return nil
}
