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

// GetCity hardcode
func (i *City) GetCitylocal(ctx context.Context, in *cities.City) (*cities.City, error) {
	return &cities.City{Id: 1, Name: "Jakarta"}, nil
}

// GetCity from database
func (i *City) GetCity(ctx context.Context, in *cities.Id) (*cities.City, error) {
	var cityModel models.City
	cityModel.Log = i.Log
	err := cityModel.Get(ctx, i.DB, in)
	return &cityModel.Pb, err
}

// CreateCity to database
func (i *City) Create(ctx context.Context, in *cities.CityInput) (*cities.City, error) {
	var cityModel models.City
	err := cityModel.Create(ctx, i.DB, in)
	return &cityModel.Pb, err
}

// UpdateCity to database
func (i *City) Update(ctx context.Context, in *cities.City) (*cities.City, error) {
	var cityModel models.City
	cityModel.Log = i.Log
	err := cityModel.Update(ctx, i.DB, in)
	return &cityModel.Pb, err
}
