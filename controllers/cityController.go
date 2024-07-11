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
func (c *City) GetCitylocal(ctx context.Context, in *cities.City) (*cities.City, error) {
	return &cities.City{Id: 1, Name: "Jakarta"}, nil
}

// GetCity from database
func (c *City) GetCity(ctx context.Context, in *cities.Id) (*cities.City, error) {
	var cityModel models.City
	cityModel.Log = c.Log
	err := cityModel.Get(ctx, c.DB, in)
	return &cityModel.Pb, err
}

// CreateCity to database
func (c *City) Create(ctx context.Context, in *cities.CityInput) (*cities.City, error) {
	var cityModel models.City
	err := cityModel.Create(ctx, c.DB, in)
	return &cityModel.Pb, err
}

// UpdateCity to database
func (c *City) Update(ctx context.Context, in *cities.City) (*cities.City, error) {
	var cityModel models.City
	cityModel.Log = c.Log
	err := cityModel.Update(ctx, c.DB, in)
	return &cityModel.Pb, err
}

// DeleteCity from database
func (c *City) Delete(ctx context.Context, in *cities.Id) (*cities.MyBoolean, error) {
	var cityModel models.City
	err := cityModel.Delete(ctx, c.DB, in)
	if err != nil {
		return &cities.MyBoolean{Boolean: false}, err
	}
	return &cities.MyBoolean{Boolean: true}, nil
}
