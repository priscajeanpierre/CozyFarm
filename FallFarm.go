package main

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/lafriks/go-tiled"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"log"
)

//go:embed CozyFarmProject/*
var CozyFarmProject embed.FS
var currentState GameState = CharacterSelection
var Game = mapGame{}
var cozyMaps = []string{"CozyFarm.tmx", "CozyFarmInterior.tmx", "WinterFarm.tmx"}
var currentDialogue string

// transition from map 1 to 2
const transitionAreaX = 10
const transitionAreaY = 30
const transitionAreaWidth = 50
const transitionAreaHeight = 50

// transition from map 2 to 1
const transitionArea1X = 50
const transitionArea1Y = 450
const transitionArea1Width = 50
const transitionArea1Height = 50

const MaxGrowthStage = 5
const SOUND_SAMPLE_RATE = 48000

var playerInventory Inventory

var player = Player{
	X: 100,
	Y: 100,
}

var NPCs = map[string]NPC{
	"NPC1": {
		Name: "NPC1",
		Dialogs: []string{
			"Hello Farmer you need to feed your animals!",
		},
		X: 450,
		Y: 70,
	},
	"NPC2": {
		Name: "NPC2",
		Dialogs: []string{
			"Hi, Neighbor, make sure you dont forget to harvest your crops!",
		},
		X: 50,
		Y: 450,
	},
}

type NPC struct {
	Name    string
	Dialogs []string
	X, Y    int
}
type Player struct {
	X, Y             int
	playerX, playerY float64
	speed            float64
	Frames           []*ebiten.Image
}

type Plant struct {
	GrowthStage   int
	Harvestable   bool
	HarvestedOnce bool
}
type Inventory struct {
	Crops int
	Items map[string]int
}
type mapGame struct {
	Level              *tiled.Map
	Level1             *tiled.Map
	Level2             *tiled.Map
	harvestSoundPlayer *audio.Player
	tileHash           map[uint32]*ebiten.Image
	tileHash1          map[uint32]*ebiten.Image
	tileHash2          map[uint32]*ebiten.Image
	Frame              int
	FrameDelay         int
	Count              int
	currentMap         int
	isMovingUp         bool
	isMovingDown       bool
	isMovingLeft       bool
	isMovingRight      bool
	xloc               int
	yloc               int
	xlocNPC1           int
	ylocNPC1           int
	xlocNPC2           int
	ylocNPC2           int
	xlocBunny          int
	ylocBunny          int
	xlocChick          int
	ylocChick          int
	xlocCow            int
	ylocCow            int
	xlocGoat           int
	ylocGoat           int
	xlocPiggy          int
	ylocPiggy          int
	xlocTurk           int
	ylocTurk           int
	xlocCow1           int
	ylocCow1           int
	FrameNPC1          int
	FrameDelayNPC1     int
	FrameNPC2          int
	FrameDelayNPC2     int
	FrameBUNNY         int
	FrameDelayBUNNY    int
	FrameCHICK         int
	FrameDelayCHICK    int
	FrameCOW           int
	FrameDelayCOW      int
	FrameGOAT          int
	FrameDelayGOAT     int
	FramePIGGY         int
	FrameDelayPIGGY    int
	FrameTURK          int
	FrameDelayTURK     int
	FrameCow1          int
	FrameDelayCow1     int

	playerSprite    []*ebiten.Image
	npcOne          []*ebiten.Image
	npcTwo          []*ebiten.Image
	bunny           []*ebiten.Image
	chicken         []*ebiten.Image
	cow             []*ebiten.Image
	goat            []*ebiten.Image
	piggy           []*ebiten.Image
	turkey          []*ebiten.Image
	cow1            []*ebiten.Image
	tree            []*ebiten.Image
	Plants          []Plant
	displayDialogue bool
	dialogueText    string
	audioContext    *audio.Context
	counter         int
	playerInventory Inventory
	font            font.Face
	IsCrafting      bool
	CraftingMessage string
}

func FallFarm() {
	var cozyMaps = []string{"CozyFarm.tmx", "CozyFarmInterior.tmx", "WinterFarm.tmx"}
	var tileMaps = make([]tiled.Map, 3, 3)
	maxWidth, maxHeight := 0, 0

	for i, mapPath := range cozyMaps {

		gameMap, err := tiled.LoadFile(mapPath)

		if err != nil {
			fmt.Printf("Error loading map %s: %s\n", mapPath, err.Error())
			continue
		}
		tileMaps[i] = *gameMap
	}
	ebitenImageMap := makeEbiteImagesFromMap(tileMaps[0])
	playerFrames := loadPlayer1()
	npcOneFrames := npc1()
	npcTwoFrames := npc2()
	bunnyFrames := bunny()
	chickenFrames := chicken()
	cowFrames := cow()
	cow1Frames := cow1()
	goatFrames := goat()
	piggyFrames := piggy()
	turkeyFrames := turkey()
	treeFrame := tree()
	loadCropSprites()
	initializeCrops()
	initializeTrees()
	loadTreeSprites()
	initializeEnemies()
	soundContext := audio.NewContext(SOUND_SAMPLE_RATE)

	Game := mapGame{
		Level:              &tileMaps[0],
		Level1:             &tileMaps[1],
		Level2:             &tileMaps[2],
		tileHash:           ebitenImageMap,
		tileHash1:          makeEbiteImagesFromMap(tileMaps[1]),
		tileHash2:          makeEbiteImagesFromMap(tileMaps[2]),
		playerSprite:       playerFrames,
		Frame:              0,
		FrameDelay:         0,
		yloc:               90,
		xloc:               25,
		npcOne:             npcOneFrames,
		npcTwo:             npcTwoFrames,
		bunny:              bunnyFrames,
		chicken:            chickenFrames,
		cow:                cowFrames,
		cow1:               cow1Frames,
		goat:               goatFrames,
		piggy:              piggyFrames,
		turkey:             turkeyFrames,
		xlocNPC1:           450,
		ylocNPC1:           70,
		xlocNPC2:           50,
		ylocNPC2:           450,
		xlocBunny:          10,
		ylocBunny:          300,
		xlocChick:          10,
		ylocChick:          350,
		xlocCow:            30,
		ylocCow:            330,
		xlocGoat:           40,
		ylocGoat:           300,
		xlocPiggy:          475,
		ylocPiggy:          50,
		xlocTurk:           405,
		ylocTurk:           370,
		xlocCow1:           70,
		ylocCow1:           450,
		FrameNPC1:          0,
		FrameDelayNPC1:     0,
		FrameNPC2:          0,
		FrameDelayNPC2:     0,
		FrameBUNNY:         0,
		FrameDelayBUNNY:    0,
		FrameCHICK:         0,
		FrameDelayCHICK:    0,
		FrameCOW:           0,
		FrameDelayCOW:      0,
		FrameGOAT:          0,
		FrameDelayGOAT:     0,
		FramePIGGY:         0,
		FrameDelayPIGGY:    0,
		FrameTURK:          0,
		FrameDelayTURK:     0,
		FrameCow1:          0,
		FrameDelayCow1:     0,
		tree:               treeFrame,
		harvestSoundPlayer: loadSoundFile(soundContext),
		counter:            20,
		playerInventory:    Inventory{},
		font:               basicfont.Face7x13,
	}

	if Game.Level.Width*Game.Level.TileWidth > maxWidth {
		maxWidth = Game.Level.Width * Game.Level.TileWidth
	}
	if Game.Level.Height*Game.Level.TileHeight > maxHeight {
		maxHeight = Game.Level.Height * Game.Level.TileHeight
	}
	if err := ebiten.RunGame(&Game); err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(maxWidth, maxHeight)
	ebiten.SetWindowTitle("Cozy Farm <3")

}

func (m *mapGame) Update() error {

	if currentState == CharacterSelection {

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

	if currentState == MainGame {

		if playerInteractsWithTransitionPoint(player.X, player.Y) {
			currentState = InteriorState

		}
		if playerIsTryingToInteract() {
			playerX, playerY := getPlayerPosition()
			for _, crop := range crops {
				if isPlayerNearCrop(playerX, playerY, crop) && crop.Grown {
					m.harvestCropAt(crop.X, crop.Y)

					break
				}
			}
		}

		if playerIsTryingToInteract() {
			npc := m.playerNearNPC()
			if npc != nil {
				fmt.Printf("Interacting with NPC at (%d, %d)\n", npc.X, npc.Y)
				currentDialogue = displayNPCDialogue(*npc)
				fmt.Println("Dialogue set to:", currentDialogue)
			}
		}
		if npc := m.playerNearNPC(); npc == nil {
			currentDialogue = ""
		}

		m.FrameDelay += 47
		if m.FrameDelay%37 == 0 {
			m.Frame += 1
			fmt.Println("updating frame:", m.Frame)
			if m.Frame >= len(m.playerSprite) {
				m.Frame = 0
				fmt.Println("resetting Frame:", m.Frame)
			}
		}
		m.FrameDelayNPC1 += 47
		if m.FrameDelayNPC1%37 == 0 {
			m.FrameNPC1 += 1
			if m.FrameNPC1 >= len(m.npcOne) {
				m.FrameNPC1 = 0
			}
		}
		m.FrameDelayNPC2 += 43
		if m.FrameDelayNPC2%37 == 0 {
			m.FrameNPC2 += 1
			if m.FrameNPC2 >= len(m.npcTwo) {
				m.FrameNPC2 = 0
			}
		}

		m.FrameDelayBUNNY += 1
		if m.FrameDelayBUNNY%5 == 0 {
			m.FrameBUNNY += 1
			if m.FrameBUNNY >= len(m.bunny) {
				m.FrameBUNNY = 0
			}
		}

		m.FrameDelayCHICK += 25
		if m.FrameDelayCHICK%31 == 0 {
			m.FrameCHICK += 1
			if m.FrameCHICK >= len(m.chicken) {
				m.FrameCHICK = 0
			}
		}
		m.FrameDelayCOW += 29
		if m.FrameDelayCOW%37 == 0 {
			m.FrameCOW += 1
			if m.FrameCOW >= len(m.cow) {
				m.FrameCOW = 0
			}
		}
		m.FrameDelayGOAT += 25
		if m.FrameDelayGOAT%27 == 0 {
			m.FrameGOAT += 1
			if m.FrameGOAT >= len(m.goat) {
				m.FrameGOAT = 0
			}
		}
		m.FrameDelayPIGGY += 31
		if m.FrameDelayPIGGY%39 == 0 {
			m.FramePIGGY += 1
			if m.FramePIGGY >= len(m.piggy) {
				m.FramePIGGY = 0
			}
		}
		m.FrameDelayTURK += 21
		if m.FrameDelayTURK%19 == 0 {
			m.FrameTURK += 1
			if m.FrameTURK >= len(m.turkey) {
				m.FrameTURK = 0
			}
		}
		m.FrameDelayCow1 += 29
		if m.FrameDelayCow1%37 == 0 {
			m.FrameCow1 += 1
			if m.FrameCow1 >= len(m.cow1) {
				m.FrameCow1 = 0
			}
		}

		m.processPlayerInput()
		PlayerSpeed := 4

		if m.isMovingUp {
			player.Y -= PlayerSpeed
		}
		if m.isMovingDown {
			player.Y += PlayerSpeed
		}
		if m.isMovingLeft {
			player.X -= PlayerSpeed
		}
		if m.isMovingRight {
			player.X += PlayerSpeed
		}

	}

	if currentState == InteriorState {

		if playerInteractsWithTransitionPoint1(player.X, player.Y) {

			currentState = WinterState

		}

		if inpututil.IsKeyJustPressed(ebiten.KeyC) {
			recipeToCraft := recipes[0]
			if playerInventory.CraftRecipe(recipeToCraft) {
				m.IsCrafting = true
				m.CraftingMessage = "Crafting " + recipeToCraft.Name + "..."

			} else {
				fmt.Println("Not enough ingredients to craft the recipe")
			}
		}

		m.processPlayerInput()
		PlayerSpeed := 4

		if m.isMovingUp {
			player.Y -= PlayerSpeed
		}
		if m.isMovingDown {
			player.Y += PlayerSpeed
		}
		if m.isMovingLeft {
			player.X -= PlayerSpeed
		}
		if m.isMovingRight {
			player.X += PlayerSpeed
		}

		m.FrameDelay += 47
		if m.FrameDelay%37 == 0 {
			m.Frame += 1
			fmt.Println("updating frame:", m.Frame)
			if m.Frame >= len(m.playerSprite) {
				m.Frame = 0
				fmt.Println("resetting Frame:", m.Frame)
			}
		}
		m.FrameDelayNPC1 += 47
		if m.FrameDelayNPC1%37 == 0 {
			m.FrameNPC1 += 1
			if m.FrameNPC1 >= len(m.npcOne) {
				m.FrameNPC1 = 0
			}
		}
		m.FrameDelayNPC2 += 43
		if m.FrameDelayNPC2%37 == 0 {
			m.FrameNPC2 += 1
			if m.FrameNPC2 >= len(m.npcTwo) {
				m.FrameNPC2 = 0
			}
		}

	}

	if currentState == WinterState {

		if playerIsTryingToInteract() {
			playerX, playerY := getPlayerPosition()
			for _, tree := range trees {
				if isPlayerNearTree(playerX, playerY, tree) && tree.SnowCovered {
					m.clearSnowAt(tree.X, tree.Y)
					break
				}
			}
		}
	}

	m.FrameDelay += 47
	if m.FrameDelay%37 == 0 {
		m.Frame += 1
		fmt.Println("updating frame:", m.Frame)
		if m.Frame >= len(m.playerSprite) {
			m.Frame = 0
			fmt.Println("resetting Frame:", m.Frame)
		}
	}
	m.FrameDelayNPC1 += 47
	if m.FrameDelayNPC1%37 == 0 {
		m.FrameNPC1 += 1
		if m.FrameNPC1 >= len(m.npcOne) {
			m.FrameNPC1 = 0
		}
	}
	m.FrameDelayNPC2 += 43
	if m.FrameDelayNPC2%37 == 0 {
		m.FrameNPC2 += 1
		if m.FrameNPC2 >= len(m.npcTwo) {
			m.FrameNPC2 = 0
		}
	}

	m.FrameDelayBUNNY += 1
	if m.FrameDelayBUNNY%5 == 0 {
		m.FrameBUNNY += 1
		if m.FrameBUNNY >= len(m.bunny) {
			m.FrameBUNNY = 0
		}
	}

	m.FrameDelayCHICK += 25
	if m.FrameDelayCHICK%31 == 0 {
		m.FrameCHICK += 1
		if m.FrameCHICK >= len(m.chicken) {
			m.FrameCHICK = 0
		}
	}
	m.FrameDelayCOW += 29
	if m.FrameDelayCOW%37 == 0 {
		m.FrameCOW += 1
		if m.FrameCOW >= len(m.cow) {
			m.FrameCOW = 0
		}
	}
	m.FrameDelayGOAT += 25
	if m.FrameDelayGOAT%27 == 0 {
		m.FrameGOAT += 1
		if m.FrameGOAT >= len(m.goat) {
			m.FrameGOAT = 0
		}
	}
	m.FrameDelayPIGGY += 31
	if m.FrameDelayPIGGY%39 == 0 {
		m.FramePIGGY += 1
		if m.FramePIGGY >= len(m.piggy) {
			m.FramePIGGY = 0
		}
	}
	m.FrameDelayTURK += 21
	if m.FrameDelayTURK%19 == 0 {
		m.FrameTURK += 1
		if m.FrameTURK >= len(m.turkey) {
			m.FrameTURK = 0
		}
	}
	m.FrameDelayCow1 += 29
	if m.FrameDelayCow1%37 == 0 {
		m.FrameCow1 += 1
		if m.FrameCow1 >= len(m.cow1) {
			m.FrameCow1 = 0
		}
	}

	m.processPlayerInput()
	PlayerSpeed := 4

	if m.isMovingUp {
		player.Y -= PlayerSpeed
	}
	if m.isMovingDown {
		player.Y += PlayerSpeed
	}
	if m.isMovingLeft {
		player.X -= PlayerSpeed
	}
	if m.isMovingRight {
		player.X += PlayerSpeed
	}

	return nil
}

func (m mapGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	return outsideWidth, outsideHeight
}

func makeEbiteImagesFromMap(tiledMap tiled.Map) map[uint32]*ebiten.Image {
	idToImage := make(map[uint32]*ebiten.Image)
	for _, tile := range tiledMap.Tilesets[0].Tiles {
		ebitenImageTile, _, err :=
			ebitenutil.NewImageFromFile(tile.Image.Source)
		if err != nil {
			fmt.Println("Error loading tile image:",
				tile.Image.Source, err)
		}
		idToImage[tile.ID] = ebitenImageTile
	}
	return idToImage
}

func (m *mapGame) Draw(screen *ebiten.Image) {
	drawOptions := ebiten.DrawImageOptions{}

	if currentState == CharacterSelection {
		screen.Fill(color.RGBA{202, 154, 107, 1})

		// Display character selection options
		for i, img := range characterImages {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*100+50), 100)
			screen.DrawImage(img, op)
		}

		ebitenutil.DebugPrint(screen, "Select a character by clicking on it")
		return
	}

	if currentState == MainGame {

		for _, layer := range m.Level.Layers {

			for tileY := 0; tileY < m.Level.Height; tileY++ {
				for tileX := 0; tileX < m.Level.Width; tileX++ {
					tileToDraw := layer.Tiles[tileY*m.Level.Width+tileX]
					if tileToDraw.ID == 0 {
						continue
					}
					drawOptions.GeoM.Reset()
					TileXpos := float64(m.Level.TileWidth * tileX)
					TileYpos := float64(m.Level.TileHeight * tileY)
					drawOptions.GeoM.Translate(TileXpos, TileYpos)

					ebitenTileToDraw := m.tileHash[tileToDraw.ID]
					screen.DrawImage(ebitenTileToDraw, &drawOptions)
				}
			}
		}
		//fmt.Println("Drawxloc:", m.xloc, "Drawyloc:", m.yloc)
		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(player.X), float64(player.Y))
		frameToDraw := m.playerSprite[0]
		screen.DrawImage(frameToDraw, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocNPC1), float64(m.ylocNPC1))
		frameToDraw1 := m.npcOne[m.FrameNPC1]
		screen.DrawImage(frameToDraw1, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocNPC2), float64(m.ylocNPC2))
		frameToDraw2 := m.npcTwo[m.FrameNPC2]
		screen.DrawImage(frameToDraw2, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocBunny), float64(m.ylocBunny))
		frameToDraw3 := m.bunny[m.FrameBUNNY]
		screen.DrawImage(frameToDraw3, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocChick), float64(m.ylocChick))
		frameToDraw4 := m.chicken[m.FrameCHICK]
		screen.DrawImage(frameToDraw4, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocCow), float64(m.ylocCow))
		frameToDraw5 := m.cow[m.FrameCOW]
		screen.DrawImage(frameToDraw5, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocCow1), float64(m.ylocCow1))
		frameToDraw9 := m.cow1[m.FrameCow1]
		screen.DrawImage(frameToDraw9, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocGoat), float64(m.ylocGoat))
		frameToDraw6 := m.goat[m.FrameGOAT]
		screen.DrawImage(frameToDraw6, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocPiggy), float64(m.ylocPiggy))
		frameToDraw7 := m.piggy[m.FramePIGGY]
		screen.DrawImage(frameToDraw7, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocTurk), float64(m.ylocTurk))
		frameToDraw8 := m.turkey[m.FrameTURK]
		screen.DrawImage(frameToDraw8, &drawOptions)

		for _, crop := range crops {
			var sprite *ebiten.Image
			if crop.Grown {
				sprite = spriteCropGrown
			} else {
				sprite = spriteCropHarvested
			}

			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(crop.X), float64(crop.Y))
			screen.DrawImage(sprite, opts)
		}

		text.Draw(screen, fmt.Sprintf("Harvested Crops: %d", playerInventory.Crops), m.font, 10, 20, color.White)

		if currentDialogue != "" {
			m.renderDialogue(screen)
		}
	}
	if currentState == InteriorState {
		for _, layer := range m.Level1.Layers {

			for tileY := 0; tileY < m.Level1.Height; tileY++ {
				for tileX := 0; tileX < m.Level1.Width; tileX++ {
					tileToDraw := layer.Tiles[tileY*m.Level1.Width+tileX]
					if tileToDraw.ID == 0 {
						continue
					}
					drawOptions.GeoM.Reset()
					TileXpos := float64(m.Level1.TileWidth * tileX)
					TileYpos := float64(m.Level1.TileHeight * tileY)
					drawOptions.GeoM.Translate(TileXpos, TileYpos)

					ebitenTileToDraw := m.tileHash1[tileToDraw.ID]
					screen.DrawImage(ebitenTileToDraw, &drawOptions)
				}
			}
		}

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(player.X), float64(player.Y))
		frameToDraw := m.playerSprite[0]
		screen.DrawImage(frameToDraw, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocNPC1), float64(m.ylocNPC1))
		frameToDraw1 := m.npcOne[m.FrameNPC1]
		screen.DrawImage(frameToDraw1, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocNPC2), float64(m.ylocNPC2))
		frameToDraw2 := m.npcTwo[m.FrameNPC2]
		screen.DrawImage(frameToDraw2, &drawOptions)

		text.Draw(screen, "Press C to craft a recipe", m.font, 10, 20, color.White)
		text.Draw(screen, fmt.Sprintf("Crops: %d", playerInventory.Crops), m.font, 10, 40, color.White)
		if m.IsCrafting {
			text.Draw(screen, m.CraftingMessage, m.font, 10, 40, color.White)

		}

	}

	if currentState == WinterState {

		for _, layer := range m.Level2.Layers {

			for tileY := 0; tileY < m.Level2.Height; tileY++ {
				for tileX := 0; tileX < m.Level2.Width; tileX++ {
					tileToDraw := layer.Tiles[tileY*m.Level2.Width+tileX]
					if tileToDraw.ID == 0 {
						continue
					}
					drawOptions.GeoM.Reset()
					TileXpos := float64(m.Level2.TileWidth * tileX)
					TileYpos := float64(m.Level2.TileHeight * tileY)
					drawOptions.GeoM.Translate(TileXpos, TileYpos)

					ebitenTileToDraw := m.tileHash2[tileToDraw.ID]
					screen.DrawImage(ebitenTileToDraw, &drawOptions)
				}
			}
		}
		//fmt.Println("Drawxloc:", m.xloc, "Drawyloc:", m.yloc)
		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(player.X), float64(player.Y))
		frameToDraw := m.playerSprite[0]
		screen.DrawImage(frameToDraw, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocNPC1), float64(m.ylocNPC1))
		frameToDraw1 := m.npcOne[m.FrameNPC1]
		screen.DrawImage(frameToDraw1, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocNPC2), float64(m.ylocNPC2))
		frameToDraw2 := m.npcTwo[m.FrameNPC2]
		screen.DrawImage(frameToDraw2, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocBunny), float64(m.ylocBunny))
		frameToDraw3 := m.bunny[m.FrameBUNNY]
		screen.DrawImage(frameToDraw3, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocChick), float64(m.ylocChick))
		frameToDraw4 := m.chicken[m.FrameCHICK]
		screen.DrawImage(frameToDraw4, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocCow), float64(m.ylocCow))
		frameToDraw5 := m.cow[m.FrameCOW]
		screen.DrawImage(frameToDraw5, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocCow1), float64(m.ylocCow1))
		frameToDraw9 := m.cow1[m.FrameCow1]
		screen.DrawImage(frameToDraw9, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocGoat), float64(m.ylocGoat))
		frameToDraw6 := m.goat[m.FrameGOAT]
		screen.DrawImage(frameToDraw6, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocPiggy), float64(m.ylocPiggy))
		frameToDraw7 := m.piggy[m.FramePIGGY]
		screen.DrawImage(frameToDraw7, &drawOptions)

		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(float64(m.xlocTurk), float64(m.ylocTurk))
		frameToDraw8 := m.turkey[m.FrameTURK]
		screen.DrawImage(frameToDraw8, &drawOptions)

		for _, tree := range trees {
			var sprite *ebiten.Image
			if tree.SnowCovered {
				sprite = spriteTreeSnowCovered
			} else {
				sprite = spriteTreeCleared
			}

			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(tree.X), float64(tree.Y))
			screen.DrawImage(sprite, opts)
		}

	}
}

func playerInteractsWithTransitionPoint(playerX, playerY int) bool {

	return playerX >= transitionAreaX &&
		playerX <= transitionAreaX+transitionAreaWidth &&
		playerY >= transitionAreaY &&
		playerY <= transitionAreaY+transitionAreaHeight
}

func playerInteractsWithTransitionPoint1(playerX, playerY int) bool {

	return playerX >= transitionArea1X &&
		playerX <= transitionArea1X+transitionArea1Width &&
		playerY >= transitionArea1Y &&
		playerY <= transitionArea1Y+transitionArea1Height
}

func (m *mapGame) processPlayerInput() {

	//fmt.Println("xloc:", m.xloc, "yloc:", m.yloc)
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		m.isMovingUp = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		m.isMovingDown = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) {
		m.isMovingUp = false
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) {
		m.isMovingDown = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		m.isMovingLeft = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		m.isMovingRight = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) || inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) {
		m.isMovingLeft = false
		m.isMovingRight = false
	}

}
