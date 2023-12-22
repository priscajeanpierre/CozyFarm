package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math"
)

type Tree struct {
	X, Y        int
	SnowCovered bool
}

var trees []*Tree

func initializeTrees() {

	trees = append(trees, &Tree{X: 300, Y: 300, SnowCovered: true})
	trees = append(trees, &Tree{X: 300, Y: 350, SnowCovered: true})
	trees = append(trees, &Tree{X: 300, Y: 390, SnowCovered: true})
	trees = append(trees, &Tree{X: 350, Y: 350, SnowCovered: true})

}

var (
	spriteTreeSnowCovered *ebiten.Image
	spriteTreeCleared     *ebiten.Image
)

func loadTreeSprites() {
	img, _, err := ebitenutil.NewImageFromFile("trees2W.png")
	if err != nil {
		log.Fatal(err)
	}
	spriteTreeSnowCovered = img

	img, _, err = ebitenutil.NewImageFromFile("farmtrees51.png")
	if err != nil {
		log.Fatal(err)
	}
	spriteTreeCleared = img
}

func isPlayerNearTree(playerX, playerY int, tree *Tree) bool {
	const InteractionThreshold = 30

	dx := playerX - tree.X
	dy := playerY - tree.Y
	distance := math.Sqrt(float64(dx*dx + dy*dy))

	return distance <= InteractionThreshold
}

func (m *mapGame) clearSnowAt(x, y int) {
	for _, tree := range trees {
		if tree.X == x && tree.Y == y && tree.SnowCovered {
			tree.SnowCovered = false
			break
		}
	}
}
