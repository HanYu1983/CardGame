package hello

import (
    "net/http"
	"lib/tool"
)

type CardPO struct {
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

type CardSuitPO struct {
    Key int64
    Name string
    Description string
    CardIds []int64
}

type ICardRepository interface {
    tool.IDataAccessObject
}

type ICardSuitRepository interface {
    tool.IDataAccessObject
}

type IApp interface {
    GetCardSuitRepository (r *http.Request) ICardSuitRepository
    GetCardRepository (r *http.Request) ICardRepository
}