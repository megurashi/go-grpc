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
func (m *City) Get(ctx context.Context, db *sql.DB, in *cities.Id) error {
	query := `SELECT id, name FROM cities WHERE id = $1`
	err := db.QueryRowContext(ctx, query, in.Id).Scan(&m.Pb.Id, &m.Pb.Name)

	if err != nil {
		m.Log.Println("Error on query", err)
		return err
	}

	return nil
}

// CreateCity
func (m *City) Create(ctx context.Context, db *sql.DB, in *cities.CityInput) error {
	query := `INSERT INTO cities (name) VALUES ($1) RETURNING id;`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	err = stmt.QueryRowContext(ctx, in.Name).Scan(&m.Pb.Id)
	if err != nil {
		return err
	}

	m.Pb.Name = in.Name

	return nil

}

// updateCity
func (m *City) Update(ctx context.Context, db *sql.DB, in *cities.City) error {
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

	m.Pb.Id = in.Id
	m.Pb.Name = in.Name

	return nil
}

// deleteCity
func (m *City) Delete(ctx context.Context, db *sql.DB, in *cities.Id) error {
	query := `DELETE FROM cities WHERE id = $1`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	rs, err := stmt.ExecContext(ctx, in.Id)
	if err != nil {
		return err
	}

	affected, err := rs.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("DATA NOT FOUND")
	}

	return nil

}
