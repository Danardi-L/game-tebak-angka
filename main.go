package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	targetNum int
	guessNum  int
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("game ended")
	}

	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		if guessNum == targetNum {
			targetNum = rand.Intn(100)
			guessNum = 0
		}
	}

	if ebiten.IsKeyPressed(ebiten.Key0) {
		guessNum = guessNum*10 + 0
	}
	if ebiten.IsKeyPressed(ebiten.Key1) {
		guessNum = guessNum*10 + 1
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		guessNum = guessNum*10 + 2
	}
	if ebiten.IsKeyPressed(ebiten.Key3) {
		guessNum = guessNum*10 + 3
	}
	if ebiten.IsKeyPressed(ebiten.Key4) {
		guessNum = guessNum*10 + 4
	}
	if ebiten.IsKeyPressed(ebiten.Key5) {
		guessNum = guessNum*10 + 5
	}
	if ebiten.IsKeyPressed(ebiten.Key6) {
		guessNum = guessNum*10 + 6
	}
	if ebiten.IsKeyPressed(ebiten.Key7) {
		guessNum = guessNum*10 + 7
	}
	if ebiten.IsKeyPressed(ebiten.Key8) {
		guessNum = guessNum*10 + 8
	}
	if ebiten.IsKeyPressed(ebiten.Key9) {
		guessNum = guessNum*10 + 9
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Guess the number: %d", guessNum))

	if guessNum == targetNum {
		ebitenutil.DebugPrint(screen, "You win! Press Enter to play again.")
	}

	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	targetNum = rand.Intn(100)

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Guess the Number"); err != nil {
		fmt.Println(err)
	}
}
