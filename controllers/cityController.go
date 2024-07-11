package controllers

import (
	"app/pb/cities"
	"context"
	"database/sql"
)

// City struct
type City struct {
	DB *sql.DB
	cities.UnimplementedCitiesServiceServer
}

// GetCity function
func (s *City) GetCity(ctx context.Context, in *cities.City) (*cities.City, error) {
	return &cities.City{Id: 1, Name: "Jakarta"}, nil
}

// GetCity function
/*func (s *City) GetCity(ctx context.Context, in *cities.Id) (*cities.City, error) {
	var cityModel models.City
	err := cityModel.Get(ctx, s.DB, in)
	return &cityModel.Pb, err
}*/
