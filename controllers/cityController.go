package controllers

import (
	"app/pb/cities"
	"database/sql"
	"log"
)

// City struct
type City struct {
	DB  *sql.DB
	Log *log.Logger
	cities.UnimplementedCitiesServiceServer
}
