package entity

import "time"

type Item struct {
	ID             int32
	Name           string
	IdUnit         int32
	IdItemCategory int32
	Visible        int32
	DateMake       *time.Time
	DateUpdate     *time.Time
}
