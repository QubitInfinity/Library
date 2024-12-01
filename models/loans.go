package models

import (
	"time"
)

type Loans struct {
	Id           int64
	BookId       int64
	MemberId     int64
	BorrowedDate *time.Time `json:"BorrowedDate,omitempty"`
	ReturnedDate *time.Time `json:"ReturnedDate,omitempty"`
}
