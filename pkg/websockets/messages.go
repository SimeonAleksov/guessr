package websockets

const (
	CURRENT_USERS = "current-users"
	CREATED       = "game-created"
	JOINED        = "joined-game"
	TICKER        = "ticker"
)

type User struct {
	Username string `json:"username"`
}

type CurrentUsersMessage struct {
	Users []User `json:"users"`
	Type  string `json:"type"`
}

type GameMessage struct {
	Type string `json:"type"`
}

type ProducerMessage struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}
