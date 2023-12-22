package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
	"math"
	"os"
)

type Crop struct {
	X, Y  int
	Grown bool
}

func initializeCrops() {
	// crop initialization
	crops = append(crops, &Crop{X: 100, Y: 200, Grown: true})
	crops = append(crops, &Crop{X: 100, Y: 220, Grown: true})
	crops = append(crops, &Crop{X: 100, Y: 240, Grown: true})
	crops = append(crops, &Crop{X: 80, Y: 200, Grown: true})
	crops = append(crops, &Crop{X: 80, Y: 220, Grown: true})
	crops = append(crops, &Crop{X: 80, Y: 240, Grown: true})
	crops = append(crops, &Crop{X: 60, Y: 200, Grown: true})
	crops = append(crops, &Crop{X: 60, Y: 220, Grown: true})
	crops = append(crops, &Crop{X: 60, Y: 240, Grown: true})
	crops = append(crops, &Crop{X: 120, Y: 200, Grown: false})
	crops = append(crops, &Crop{X: 120, Y: 220, Grown: false})
	crops = append(crops, &Crop{X: 120, Y: 240, Grown: false})

}

var crops []*Crop
var (
	spriteCropGrown     *ebiten.Image
	spriteCropHarvested *ebiten.Image
)

func isPlayerNearCrop(playerX, playerY int, crop *Crop) bool {
	const InteractionThreshold = 30

	dx := playerX - crop.X
	dy := playerY - crop.Y
	distance := math.Sqrt(float64(dx*dx + dy*dy))

	return distance <= InteractionThreshold
}

// HARVEST STUFF
func playerIsTryingToInteract() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		fmt.Println("E key pressed")
		return false
	}
	return ebiten.IsKeyPressed(ebiten.KeyE)
}

func getPlayerPosition() (int, int) {
	return player.X, player.Y
}
func loadCropSprites() {

	img, _, err := ebitenutil.NewImageFromFile("tomatoGrown.png")
	if err != nil {
		log.Fatal(err)
	}
	spriteCropGrown = img

	img, _, err = ebitenutil.NewImageFromFile("harvestedPlot.png")
	if err != nil {
		log.Fatal(err)
	}
	spriteCropHarvested = img
}

func loadSoundFile(context *audio.Context) *audio.Player {

	harvestFile, err := os.Open("harvestSound.wav")
	if err != nil {
		fmt.Println("Error Loading sound: ", err)
	}
	harvestSound, err := wav.DecodeWithoutResampling(harvestFile)
	if err != nil {
		fmt.Println("Error interpreting sound file: ", err)
	}
	soundPlayer, err := context.NewPlayer(harvestSound)
	if err != nil {
		fmt.Println("Couldn't create sound player: ", err)
	}
	return soundPlayer

}
func (m *mapGame) harvestCropAt(x, y int) {
	for _, crop := range crops {
		if crop.X == x && crop.Y == y && crop.Grown {
			crop.Grown = false
			playerInventory.Crops++
			break
		}
	}
}
