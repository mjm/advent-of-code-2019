package day13

import (
	"fmt"
	"sync"

	"github.com/gdamore/tcell"

	"github.com/mjm/advent-of-code-2019/pkg/point"
)

// Game plays a game controlled by an Intcode program
type Game struct {
	canvas    *Canvas
	screen    tcell.Screen
	lock      sync.Mutex
	ballLoc   point.Point2D
	paddleLoc point.Point2D
	score     int
}

// NewGame creates a new game that draws on the given canvas.
func NewGame(c *Canvas) *Game {
	return &Game{
		canvas: c,
	}
}

// Connect connects the game to a terminal screen.
func (g *Game) Connect(screen tcell.Screen) {
	g.screen = screen
}

// Run runs the game according to the instructions from output.
func (g *Game) Run(output chan int) {
	for {
		if ok := g.handleGameUpdate(output); !ok {
			break
		}
	}
}

// Score returns the current score of the game, or the final score if the game
// is over.
func (g *Game) Score() int {
	return g.score
}

func (g *Game) handleGameUpdate(output <-chan int) bool {
	x, ok := <-output
	if !ok {
		return false
	}
	g.lock.Lock()
	defer g.lock.Unlock()

	y, ok := <-output
	if !ok {
		return false
	}
	tile, ok := <-output
	if !ok {
		return false
	}

	if x == -1 && y == 0 {
		g.score = tile
		return true
	}

	u := paintUpdate{
		point: point.Point2D{X: x, Y: y},
		tile:  tile,
	}
	g.handlePaintUpdate(u)

	return true
}

func (g *Game) handlePaintUpdate(u paintUpdate) {
	if u.tile == 4 {
		// 4 = ball
		g.ballLoc = u.point
	} else if u.tile == 3 {
		// 3 = paddle
		g.paddleLoc = u.point
	}
	g.canvas.Paint(u.point, u.tile)
	g.updateScreen()
}

func (g *Game) updateScreen() {
	if g.screen == nil {
		return
	}

	g.screen.Clear()
	g.canvas.Draw(g.drawTile)
	g.drawScore()
	g.screen.Show()
}

func (g *Game) drawTile(x, y, tile int) {
	var style tcell.Style
	var c rune = ' '
	switch tile {
	case 0:
		style = style.Background(tcell.ColorBlack)
	case 1:
		style = style.Background(tcell.ColorWhite)
	case 2:
		c = 'B'
		style = style.Background(tcell.ColorBlue).Foreground(tcell.ColorTeal)
	case 3:
		c = '='
		style = style.Background(tcell.ColorRed).Foreground(tcell.ColorMaroon)
	case 4:
		c = 'o'
		style = style.Background(tcell.ColorBlack).Foreground(tcell.ColorGreen)
	}
	g.screen.SetContent(x, y, c, nil, style)
}

func (g *Game) drawScore() {
	s := fmt.Sprintf(" Score: %d ", g.score)
	y := g.canvas.Height()

	var style tcell.Style
	style = style.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	for x, c := range s {
		g.screen.SetContent(x, y, c, nil, style)
	}
}

type paintUpdate struct {
	point point.Point2D
	tile  int
}
