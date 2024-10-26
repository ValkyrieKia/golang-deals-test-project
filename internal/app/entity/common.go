package entity

import "time"

type GetQueryParameter struct {
	Start *time.Time
	End   *time.Time
}
