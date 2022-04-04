package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

const (
	SCREEN_WIDTH, SCREEN_HEIGHT = 640, 480
)

type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

var (
	err        error
	background *ebiten.Image
	spaceShip  *ebiten.Image
	game       *Game
)

type Game struct {
	Player *player
}

// game state updates
func (g *Game) Update(_ *ebiten.Image) error {
	movePlayer()
	return nil
}

// game rendering logic
func (g *Game) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Scale(0.1, 0.1)
	playerOp.GeoM.Translate(game.Player.xPos, game.Player.yPos)
	screen.DrawImage(game.Player.image, playerOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	screenWidth = SCREEN_WIDTH
	screenHeight = SCREEN_HEIGHT
	return
}

func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	spaceShip, _, err = ebitenutil.NewImageFromFile("assets/spaceship.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	game = &Game{
		Player: &player{spaceShip, SCREEN_WIDTH / 2.0, SCREEN_HEIGHT / 2.0, 4},
	}
}

func movePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		game.Player.yPos -= game.Player.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		game.Player.yPos += game.Player.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		game.Player.xPos -= game.Player.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		game.Player.xPos += game.Player.speed
	}
}

func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Space Slay")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
