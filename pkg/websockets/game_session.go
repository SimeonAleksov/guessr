package websockets

import models "guessr.net/models/trivia"

func StartGameSession() {
	models.GetAllTrivia()
}
