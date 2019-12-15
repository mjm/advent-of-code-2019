package day13

import "time"

// AutoPlayer will automatically move the joystick perfectly.
type AutoPlayer struct {
	game *Game
}

// NewAutoPlayer creates a new autoplayer.
func NewAutoPlayer(g *Game) *AutoPlayer {
	return &AutoPlayer{
		game: g,
	}
}

// HandleInput gives the correct way to move the joystick based on the current
// game state. Use this with SetInputFunc on the game's VM.
func (ap *AutoPlayer) HandleInput() int {
	ap.game.lock.Lock()
	defer ap.game.lock.Unlock()

	time.Sleep(5 * time.Millisecond)
	if ap.game.ballLoc.X < ap.game.paddleLoc.X {
		return -1
	} else if ap.game.ballLoc.X > ap.game.paddleLoc.X {
		return 1
	} else {
		return 0
	}
}
