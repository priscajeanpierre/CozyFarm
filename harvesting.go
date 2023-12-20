package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"os"
)

// HARVEST STUFF
func playerIsTryingToInteract() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		fmt.Println("E key pressed")
		return false
	}
	return ebiten.IsKeyPressed(ebiten.KeyE)
}
func (m *mapGame) getTileID(x, y int) int {

	tileX := x / 32
	tileY := y / 32

	// Iterate through each layer to find the tile at player's position
	for _, layer := range m.Level.Layers {

		index := tileY*m.Level.Width + tileX

		// Check if index is within the bounds of tile array
		if index < 0 || index >= len(layer.Tiles) {
			continue
		}

		tile := layer.Tiles[index]

		if tile.ID != 0 {
			return int(tile.ID)
		}
	}

	return 0

}
func (m *mapGame) isTileInteractable(x, y int) bool {
	tileID := m.getTileID(x, y)
	fmt.Printf("Interacting with Tile ID: %d\n", tileID)

	// list of interactable tile IDs
	interactableTiles := map[int]bool{
		43: true, //dirt holes
		1:  true, //farmcrops
		2:  true, //farmcrops
		3:  true, //farmcrops

	}

	// Check if tileID is in list of interactable tiles
	_, isInteractable := interactableTiles[tileID]
	return isInteractable
}

func (m *mapGame) handleTileInteraction(tileID, x, y int) {
	switch tileID {
	case 54:

		m.harvestCropAt(x, y)
	case 1:
		m.harvestCropAt(x, y)
	case 2:
		m.harvestCropAt(x, y)
	case 3:
		m.harvestCropAt(x, y)
	}
}
func (m *mapGame) getPlayerTilePosition() (int, int) {
	tileX := player.X / 32
	tileY := player.Y / 32
	fmt.Printf("Player Tile Position: %d, %d\n", tileX, tileY)

	return tileX, tileY
}

func (m *mapGame) harvestCropAt(x, y int) {
	fmt.Printf("Harvesting at coordinates: %d, %d\n", x, y)
	tileX := x / 32
	tileY := y / 32

	// Iterate through each layer to find and update the crop tile
	for _, layer := range m.Level.Layers {
		index := tileY*m.Level.Width + tileX

		if index < 0 || index >= len(layer.Tiles) {
			continue
		}

		tile := &layer.Tiles[index]

		// Check if the tile is a crop and update it
		if isCropTile(int((*tile).ID)) {
			// Change the tile ID to represent the harvested state
			(*tile).ID = 1 // or another ID representing the harvested state

			// Update player inventory
			m.playerInventory.CropsHarvested++

			// play harvest sound effect
			if m.harvestSoundPlayer != nil {
				err := m.harvestSoundPlayer.Rewind()
				if err != nil {
					return
				}
				m.harvestSoundPlayer.Play()
			}

			break
		}
	}
}

func isCropTile(tileID int) bool {
	// Define which tile IDs represent crops
	cropTiles := map[int]bool{
		43: true,
		1:  true,
		2:  true,
		3:  true,
	}

	_, isCrop := cropTiles[tileID]
	return isCrop
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
