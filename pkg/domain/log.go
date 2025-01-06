package audit

import (
	"errors"
	"github.com/biletskyi-mykyta/crud-logger/pkg/domain/audit"
	"time"
)

const (
	ENTITY_USER = "USER"
	ENTITY_BOOK = "BOOK"

	ACTION_CREATE   = "CREATE"
	ACTION_UPDATE   = "UPDATE"
	ACTION_GET      = "GET"
	ACTION_DELETE   = "DELETE"
	ACTION_REGISTER = "REGISTER"
	ACTION_LOGIN    = "LOGIN"
)

var (
	entities = map[string]audit.LogRequest_Entities{
		ENTITY_USER: audit.LogRequest_USER,
		ENTITY_BOOK: audit.LogRequest_BOOK,
	}

	actions = map[string]audit.LogRequest_Actions{
		ACTION_CREATE:   audit.LogRequest_CREATE,
		ACTION_UPDATE:   audit.LogRequest_UPDATE,
		ACTION_GET:      audit.LogRequest_GET,
		ACTION_DELETE:   audit.LogRequest_DELETE,
		ACTION_REGISTER: audit.LogRequest_REGISTER,
		ACTION_LOGIN:    audit.LogRequest_LOGIN,
	}
)

type LogItem struct {
	Entity    string    `bson:"entity"`
	Action    string    `bson:"action"`
	EntityID  int64     `bson:"entity_id"`
	Timestamp time.Time `bson:"timestamp"`
}

func ToPbEntity(entity string) (audit.LogRequest_Entities, error) {
	val, ex := entities[entity]
	if !ex {
		return 0, errors.New("invalid entity")
	}

	return val, nil
}

func ToPbAction(action string) (audit.LogRequest_Actions, error) {
	val, ex := actions[action]
	if !ex {
		return 0, errors.New("invalid action")
	}

	return val, nil
}
