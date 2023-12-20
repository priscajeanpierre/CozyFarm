package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math/rand"
)

func displayNPCDialogue(npc NPC) string {
	selectedDialogue := npc.Dialogs[rand.Intn(len(npc.Dialogs))]
	return selectedDialogue
}

func renderDialogue(screen *ebiten.Image) {
	// Render dialogue text on the screen
	screen.Fill(color.Black)
	ebitenutil.DebugPrint(screen, currentDialogue)
}
