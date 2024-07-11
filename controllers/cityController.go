package controllers

import (
	"app/models"
	"app/pb/cities"
	"context"
	"database/sql"
	"log"
)

// City struct
type City struct {
	DB  *sql.DB
	Log *log.Logger
	cities.UnimplementedCitiesServiceServer
}

// GetCity function
func (s *City) GetCity(ctx context.Context, in *cities.Id) (*cities.City, error) {
	var cityModel models.City
	cityModel.Log = s.Log
	err := cityModel.Get(ctx, s.DB, in)
	return &cityModel.Pb, err
}
