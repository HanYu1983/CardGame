package hello

import (
    "net/http"
)

type IRepository interface {
    Create(po interface{}) int64
    Update(key int64, po interface{})
    Read(key int64) interface{}
    Delete(key int64)
    GetAll() []interface{}
}

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

type ICardRepository interface {
    IRepository
    /*
    Create(card CardPO) int64
    Update(key int64, card CardPO)
    Read(key int64) CardPO
    Delete(key int64)
    GetAll() []CardPO
    */
}

type IApp interface {
    GetCardRepository (r *http.Request) IRepository
}