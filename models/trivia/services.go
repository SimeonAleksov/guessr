package models

import (
	"context"
	"database/sql"
	"guessr.net/pkg/database"
	"log"
	"time"
)

type TriviaService interface {
	GetOrCreateGameSession(code string, triviaID int64) (Gamesession, bool)
}

type S string

func (se S) GetOrCreateGameSession(code string, triviaID int64) (Gamesession, bool) {
	queries := New(database.DB)
	var gameSession = Gamesession{}
	gameSession, err := queries.GetGameSession(context.Background(), code)
	created := false
	if err != nil {
		log.Println(err)
	}
	if (Gamesession{}) == gameSession {
		log.Println("GameSession not found, creating new.")
		gameSession, err = queries.CreateGameSession(context.Background(), CreateGameSessionParams{
			Code:       code,
			TriviaID:   triviaID,
			FinishedAt: sql.NullTime{Time: time.Now()},
		})
		created = true
		if err != nil {
			log.Println(err)
		}
	}
	return gameSession, created
}

func GetAllTrivia() []GetAllTriviaRow {
	queries := New(database.DB)
	tr, err := queries.GetAllTrivia(context.Background())
	if err != nil {
		log.Println(tr)
	}
	return tr
}
