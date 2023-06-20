package models

import "github.com/google/uuid"

type Event struct {
	Id  uuid.UUID
	Log string
}
