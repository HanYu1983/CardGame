package app

import (
    "net/http"
	"lib/tool"
)

type CardEntity struct {
    Key int64
    Name string
    Action string
    ActionContent string
    ActionCost string
    ActionType string
    AtkType string
    Content string
    Level int
    Magic int
    Power int
    Tag string
    Type string
    Weight int
}

type CardSuitEntity struct {
    Key int64
    Name string
    Description string
    CardIds []int64
}

type ICardDAO interface {
    tool.IDataAccessObject
}

type ICardSuitDAO interface {
    tool.IDataAccessObject
}

type IApp interface {
    GetCardSuitDAO(r *http.Request) ICardSuitDAO
    GetCardDAO(r *http.Request) ICardDAO
}