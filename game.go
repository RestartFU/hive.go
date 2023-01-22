package hive

type Game struct {
	game game
}

func (g Game) String() string { return string(g.game) }

type game string

func GameSkyWars() Game { return Game{game: "sky"} }

func GameHideAndSeek() Game { return Game{game: "hide"} }

func GameSurvivalGame() Game { return Game{game: "sg"} }

func GameMurderMystery() Game { return Game{game: "murder"} }

func GameTreasureWars() Game { return Game{game: "wars"} }

func GameDeathRun() Game { return Game{game: "dr"} }
