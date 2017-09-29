package console

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/uis"
)

type UI struct {
	Game          core.Game
	BoardRenderer uis.BoardRender
}

func MakeConsoleUI(game core.Game, render uis.BoardRender) UI {
	return UI{
		Game:          game,
		BoardRenderer: render,
	}
}

//RenderMessage prints string to console
func (ui UI) RenderMessage(str string) {
	fmt.Println(str)
}

func (ui UI) GetMove() error {
	userMove := ui.Game.GetMove()
	move, err := ValidateMove(userMove)
	if err != nil {
		fmt.Println(fmt.Sprintf("%s is Invalid Input \n", userMove))
		return err
	}
	if err := ui.Game.MakeMove(move); err != nil {
		switch err {
		case core.ErrOutOfBounds:
			fmt.Println(fmt.Sprintf("%s is Out of Bounds \n", userMove))
			return err
		case core.ErrSquareOccupied:
			fmt.Println(fmt.Sprintf("%s is Occupied \n", userMove))
			return err
		default:
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (ui UI) RenderBoard() {
	board := ui.Game.GameBoard()
	players := ui.Game.GamePlayers()
	ui.clearConsole()
	fmt.Println("\n\n" + ui.BoardRenderer.RenderBoard(board, players))
}

func (ui UI) winMessage(marker string) string {
	return fmt.Sprintf("%s's win!", marker)
}

func (ui UI) tieMessage() string {
	return fmt.Sprintf("Tie...")
}

func (ui UI) clearConsole() {
	fmt.Print("\033c")
}

func (ui UI) NextGameState() (message string, endGame bool) {
	if ui.Game.IsWin() {
		message := ui.winMessage(ui.Game.InActivePlayerMarker())
		return message, false
	}
	if ui.Game.IsTie() {
		message := ui.tieMessage()
		return message, false
	}
	return "", true
}