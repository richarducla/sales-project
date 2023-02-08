package models

import "time"

type Sale struct {
	PointOfSale string
	Product     string
	Date        time.Time
	Stock       int
}
