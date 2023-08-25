package models

import (
	"time"
)

type Vaccine struct {
	ID        int 				 `sql:"id,omitempty"`
	Date	  time.Time          `sql:"fecha,omitempty"`
	Name      string             `sql:"nombreVacuna,omitempty"`
	Dog       int				 `sql:"id_perro,omitempty"`
	CreatedAt time.Time          `sql:"created_at,omitempty"`
	UpdatedAt time.Time          `sql:"updated_at,omitempty"`
}