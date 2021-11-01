package main

import (
	"github.com/satori/go.uuid"
)

func GetId() uuid.UUID {

  userUuid := uuid.NewV4()

  return userUuid
}


type Album struct {
    Id     uuid.UUID    `json:"id"`
    Title  string  		`json:"title"`
    Artist string  		`json:"artist"`
    Price  float64 		`json:"price"`
}

var Albums = []Album{
    {Id: GetId(), Title: "Don't Know", Artist: "Eminiem", Price: 49.99},
    {Id: GetId(), Title: "Love you Like song", Artist: "Selene Gomez", Price: 17.99},
    {Id: GetId(), Title: "Ente Eih", Artist: "Nancy Ajaram", Price: 39.99},
}
