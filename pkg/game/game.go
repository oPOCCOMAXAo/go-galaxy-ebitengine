package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/opoccomaxao/go-galaxy-experiment/pkg/emath"
	"github.com/opoccomaxao/go-galaxy-experiment/pkg/physics"
	"github.com/pkg/errors"
)

var _ ebiten.Game = &Game{}

type Game struct {
	size    image.Point
	buffer  *ebiten.Image
	buffer2 *ebiten.Image
	physics *physics.Engine
}

func New() *Game {
	size := image.Pt(640, 480)

	return &Game{
		size:    size,
		buffer:  ebiten.NewImage(size.X, size.Y),
		buffer2: ebiten.NewImage(size.X, size.Y),
		physics: newPhysics(emath.Rect2D{
			Min: emath.Vector2D{X: 0, Y: 0},
			Max: emath.FromImagePoint(size),
		}),
	}
}

func (g *Game) Update() error {
	g.physics.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.physics.DrawInto(screen, emath.Rect2D{
		Min: emath.Vector2D{X: 0, Y: 0},
		Max: emath.FromImagePoint(g.size),
	})

	// g.physics.DrawInto(g.buffer2, emath.Rect2D{
	// 	Min: emath.Vector2D{X: 0, Y: 0},
	// 	Max: emath.FromImagePoint(g.size),
	// })

	// g.buffer, g.buffer2 = g.buffer2, g.buffer

	// screen.DrawImage(g.buffer, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.size.X, g.size.Y
}

func (g *Game) Run() error {
	ebiten.SetWindowSize(g.size.X, g.size.Y)
	ebiten.SetWindowTitle("Galaxy")

	err := ebiten.RunGame(g)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
