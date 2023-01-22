package hive

import (
	"strconv"
	time2 "time"
)

// PlayerResponse is a struct containing all the information about a player's statistics.
type PlayerResponse struct {
	UUID                   string `json:"UUID"`
	XP                     int    `json:"xp"`
	Played                 int    `json:"played"`
	Victories              int    `json:"victories"`
	FirstPlayed            int    `json:"first_played"`
	Kills                  int    `json:"kills"`
	MysteryChestsDestroyed int    `json:"mystery_chests_destroyed"`
	OresMined              int    `json:"ores_mined"`
	SpellsUsed             int    `json:"spells_used"`
	FirstPlayedTime        time2.Time
}

// firstPlayedTime converts unix time into time.Time
func firstPlayedTime(r PlayerResponse) time2.Time {
	time := strconv.Itoa(r.FirstPlayed)
	i, err := strconv.ParseInt(time, 10, 64)
	if err != nil {
		return time2.Time{}
	}
	return time2.Unix(i, 0)
}
