package logger

type StartEntry struct {
	Event   string   `json:"event"`
	Deck    []int    `json:"deck"`
	Players []string `json:"players"`
}

func NewStartEntry(deck []int, players []string) StartEntry {
	return StartEntry{"start", deck, players}
}

type PlaceTileEntry struct {
	Event    string `json:"event"`
	Player   int    `json:"player"`
	Rotation int    `json:"rotation"`
	Position []int  `json:"position"`
	Meeple   int    `json:"meeple"`
}

func NewPlaceTileEntry(player int, rotation int, position []int, meeple int) PlaceTileEntry {
	return PlaceTileEntry{"place", player, rotation, position, meeple}
}

type EndEntry struct {
	Event  string `json:"event"`
	Scores []int  `json:"scores"`
}

func NewEndEntry(scores []int) EndEntry {
	return EndEntry{"end", scores}
}
