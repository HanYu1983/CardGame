package hello

import (
    "net/http"
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
    CardIds []int64
}

type ICardRepository interface {
    IRepository
}

type ICardSuitRepository interface {
    IRepository
}

type IApp interface {
    GetCardSuitRepository (r *http.Request) ICardSuitRepository
    GetCardRepository (r *http.Request) ICardRepository
}