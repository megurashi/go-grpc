package models

import (
	"app/pb/cities"
	"context"
	"database/sql"
	"fmt"
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

// CreateCity
func (i *City) Create(ctx context.Context, db *sql.DB, in *cities.CityInput) error {
	query := `INSERT INTO cities (name) VALUES ($1) RETURNING id;`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	err = stmt.QueryRowContext(ctx, in.Name).Scan(&i.Pb.Id)
	if err != nil {
		return err
	}

	i.Pb.Name = in.Name

	return nil

}

// updateCity
func (i *City) Update(ctx context.Context, db *sql.DB, in *cities.City) error {
	query := `UPDATE cities SET name = $1 WHERE id = $2`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, in.Name, in.Id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("data not found!")
	}

	i.Pb.Id = in.Id
	i.Pb.Name = in.Name

	return nil
}
