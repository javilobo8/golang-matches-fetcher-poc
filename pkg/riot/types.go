package riot

type AccountDTO struct {
	PUUID    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type MatchDTO struct {
	Metadata struct {
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		GameId             int    `json:"gameId"`
		GameCreation       int    `json:"gameCreation"`
		GameDuration       int    `json:"gameDuration"`
		GameEndTimestamp   int    `json:"gameEndTimestamp"`
		GameMode           string `json:"gameMode"`
		GameName           string `json:"gameName"`
		GameStartTimestamp int    `json:"gameStartTimestamp"`
		GameType           string `json:"gameType"`
		GameVersion        string `json:"gameVersion"`
		MapId              int    `json:"mapId"`
		PlatformId         string `json:"platformId"`
		QueueId            int    `json:"queueId"`
		EndOfGameResult    string `json:"endOfGameResult"`
		Participants       []struct {
			PUUID   string `json:"puuid"`
			Kills   int    `json:"kills"`
			Deaths  int    `json:"deaths"`
			Assists int    `json:"assists"`
		}
	} `json:"info"`
}
