package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
)

type GameState = int

const (
	screenWidth                  = 640
	screenHeight                 = 480
	CharacterSelection GameState = iota
	MainGame           GameState = iota
	InteriorState      GameState = iota
	WinterState        GameState = iota
)

var characterImages []*ebiten.Image
var selectedCharacterIndex = -1

// var currentState GameState = CharacterSelection
func init() {

	// Load character images
	characterImages = append(characterImages, LoadEmbeddedImage("characters", "c1.png"))
	characterImages = append(characterImages, LoadEmbeddedImage("characters", "c2.png"))
	characterImages = append(characterImages, LoadEmbeddedImage("characters", "c3.png"))
	characterImages = append(characterImages, LoadEmbeddedImage("characters", "c4.png"))

}

type CharacterSelectionState struct {
}

func (g *CharacterSelectionState) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{202, 154, 107, 1})

	// Display character selection options
	for i, img := range characterImages {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(i*100+50), 100)
		screen.DrawImage(img, op)
	}

	ebitenutil.DebugPrint(screen, "Select a character by clicking on it")
}

func (g *CharacterSelectionState) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func run() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Character Selection")

	gameState := &CharacterSelectionState{}

	if err := ebiten.RunGame(gameState); err != nil {
		log.Fatal(err)
	}
}

func (g *CharacterSelectionState) Update() error {
	// Detect mouse clicks
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()

		// Iterate through character images and check if mouse click inside any character bounds
		for i, img := range characterImages {
			characterX := i*100 + 50
			characterY := 100

			characterWidth := img.Bounds().Max.X
			characterHeight := img.Bounds().Max.Y

			// Check if mouse clicks character bounds
			if mouseX >= characterX && mouseX <= characterX+characterWidth &&
				mouseY >= characterY && mouseY <= characterY+characterHeight {
				selectedCharacterIndex = i
				log.Printf("Character %d clicked!", i)
				break

			}
		}
	}

	if selectedCharacterIndex != -1 {
		currentState = MainGame
	}
	characterSelected = true
	fmt.Println("Character Selected:", selectedCharacterIndex)

	return nil
}
