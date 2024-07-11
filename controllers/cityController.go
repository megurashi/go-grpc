package controllers

import (
<<<<<<< HEAD
	"app/pb/cities"
=======
	"app/models"
	"app/pb/cities"
	"context"
>>>>>>> 099be60759990b60004f0fdf1e2c579740433b14
	"database/sql"
	"log"
)

// City struct
type City struct {
	DB  *sql.DB
	Log *log.Logger
	cities.UnimplementedCitiesServiceServer
}
<<<<<<< HEAD
=======

// GetCity function
func (s *City) GetCity(ctx context.Context, in *cities.Id) (*cities.City, error) {
	var cityModel models.City
	cityModel.Log = s.Log
	err := cityModel.Get(ctx, s.DB, in)
	return &cityModel.Pb, err
}
>>>>>>> 099be60759990b60004f0fdf1e2c579740433b14
