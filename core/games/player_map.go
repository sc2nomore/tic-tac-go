package games

import (
	"github.com/sc2nomore/tic-tac-go/core"
)

type PlayerMap struct {
	Players map[int]core.Player
}

func (players PlayerMap) PlayerSymbol(player int) (string, error) {
	if player, ok := players.Players[player]; ok {
		return player.Symbol(), nil
	}
	return "", core.ErrPlayerSymbolNotFoundError
}

func (players PlayerMap) Player(player int) core.Player {
	return players.Players[player]
}

func MakePlayers(player1 core.Player, player2 core.Player) core.PlayerMapper {
	return PlayerMap{
		Players: map[int]core.Player{
			-1: player1,
			1:  player2,
		},
	}
}
