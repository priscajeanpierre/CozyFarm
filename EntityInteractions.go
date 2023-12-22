package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image"
	"image/color"
	"math"
	"math/rand"
)

const InteractionThreshold = 50

func (m *mapGame) playerNearNPC() *NPC {
	for _, npc := range NPCs {
		if distance(player.X, player.Y, npc.X, npc.Y) < InteractionThreshold {
			return &npc
		}
	}
	return nil
}
func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

func displayNPCDialogue(npc NPC) string {
	selectedDialogue := npc.Dialogs[rand.Intn(len(npc.Dialogs))]
	return selectedDialogue
}

// chatgpt
func (m *mapGame) renderDialogue(screen *ebiten.Image) {
	if m.font == nil {
		fmt.Println("Font is not initialized!")

		return
	}

	w, h := screen.Size()
	dialogueBoxHeight := 100

	dialogueBox := image.Rect(0, h-dialogueBoxHeight, w, h)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(h-dialogueBoxHeight))
	op.ColorM.Scale(0, 0, 0, 0.8)
	dialogueBoxImage := ebiten.NewImageFromImage(image.NewRGBA(dialogueBox))
	screen.DrawImage(dialogueBoxImage, op)

	text.Draw(screen, currentDialogue, m.font, 10, h-10, color.White)
	fmt.Println("Rendering dialogue:", currentDialogue)
}
