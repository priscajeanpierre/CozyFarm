package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"path"
)

func loadPlayer1() []*ebiten.Image {

	allFrames := make([]*ebiten.Image, 12, 16)
	animationList := []string{"p0.png", "p1.png", "p2.png", "p3.png", "p4.png", "p5.png",
		"p6.png", "p7.png", "p8.png", "p9.png", "p10.png", "p11.png",
	}

	for index, animation := range animationList {
		framePict := LoadEmbeddedImage("playerOne", animation)
		allFrames[index] = framePict
	}

	return allFrames
}
func loadCharOne() []*ebiten.Image {

	allFrames := make([]*ebiten.Image, 12, 16)
	animationList := []string{"charOne.png", "charOne1.png", "charOne2.png", "charOne3.png", "charOne4.png", "charOne5.png",
		"charOne6.png", "charOne7.png", "charOne8.png", "charOne9.png", "charOne10.png", "charOne11.png",
	}

	for index, animation := range animationList {
		framePict := LoadEmbeddedImage("character One", animation)
		allFrames[index] = framePict
	}

	return allFrames
}
func loadCharTwo() []*ebiten.Image {

	allFrames := make([]*ebiten.Image, 12, 16)
	animationList := []string{"charTwo0.png", "charTwo1.png", "charTwo2.png", "charTwo3.png", "charTwo4.png", "charTwo5.png",
		"charTwo6.png", "charTwo7.png", "charTwo8.png", "charTwo9.png", "charTwo10.png", "charTwo11.png",
	}

	for index, animation := range animationList {
		framePict := LoadEmbeddedImage("character One", animation)
		allFrames[index] = framePict
	}

	return allFrames
}
func loadCharThree() []*ebiten.Image {

	allFrames := make([]*ebiten.Image, 12, 16)
	animationList := []string{"charThree0.png", "charThree1.png", "charThree2.png", "charThree3.png", "charThree4.png", "charThree5.png",
		"charThree6.png", "charThree7.png", "charThree8.png", "charThree9.png", "charThree10.png", "charThree11.png",
	}

	for index, animation := range animationList {
		framePict := LoadEmbeddedImage("character One", animation)
		allFrames[index] = framePict
	}

	return allFrames
}

func npc1() []*ebiten.Image {

	allFrames := make([]*ebiten.Image, 12, 15)
	animationList := []string{"npcOne0.png", "npcOne1.png", "npcOne2.png", "npcOne3.png", "npcOne4.png", "npcOne5.png", "npcOne6.png",
		"npcOne7.png", "npcOne8.png", "npcOne9.png", "npcOne10.png", "npcOne11.png"}

	for index, animation := range animationList {
		framePict := LoadEmbeddedImage("npc One", animation)
		allFrames[index] = framePict
	}

	return allFrames
}
func npc2() []*ebiten.Image {

	allFrames := make([]*ebiten.Image, 12, 15)
	animationList := []string{"npcTwo0.png", "npcTwo1.png", "npcTwo2.png", "npcTwo3.png", "npcTwo4.png", "npcTwo5.png", "npcTwo6.png",
		"npcTwo7.png", "npcTwo8.png", "npcTwo9.png", "npcTwo10.png", "npcTwo11.png"}

	for index, animation := range animationList {
		framePict := LoadEmbeddedImage("npc Two", animation)
		allFrames[index] = framePict
	}

	return allFrames
}
func bunny() []*ebiten.Image {

	allFrames := make([]*ebiten.Image, 12, 15)
	animationList := []string{"bun0.png", "bun1.png", "bun2.png", "bun3.png", "bun4.png", "bun5.png", "bun6.png",
		"bun7.png", "bun8.png", "bun9.png", "bun10.png", "bun11.png"}

	for index, animation := range animationList {
		framePict := LoadEmbeddedImage("bunny", animation)
		allFrames[index] = framePict
	}

	return allFrames
}
func chicken() []*ebiten.Image {

	allFrames := make([]*ebiten.Image, 12, 15)
	animationList := []string{"chick1.png", "chick2.png", "chick3.png", "chick4.png", "chick5.png", "chick6.png",
		"chick7.png", "chick8.png", "chick9.png", "chick10.png", "chick11.png", "chick12.png"}

	for index, animation := range animationList {
		framePict := LoadEmbeddedImage("chicken", animation)
		allFrames[index] = framePict
	}

	return allFrames
}
func cow() []*ebiten.Image {

	allFrames := make([]*ebiten.Image, 12, 15)
	animationList := []string{"cow0.png", "cow1.png", "cow2.png", "cow3.png", "cow4.png", "cow5.png", "cow6.png",
		"cow7.png", "cow8.png", "cow9.png", "cow10.png", "cow11.png"}

	for index, animation := range animationList {
		framePict := LoadEmbeddedImage("cow", animation)
		allFrames[index] = framePict
	}

	return allFrames
}
func cow1() []*ebiten.Image {

	allFrames := make([]*ebiten.Image, 12, 15)
	animationList := []string{"cow0.png", "cow1.png", "cow2.png", "cow3.png", "cow4.png", "cow5.png", "cow6.png",
		"cow7.png", "cow8.png", "cow9.png", "cow10.png", "cow11.png"}

	for index, animation := range animationList {
		framePict := LoadEmbeddedImage("cow", animation)
		allFrames[index] = framePict
	}

	return allFrames
}
func tree() []*ebiten.Image {

	allFrames := make([]*ebiten.Image, 1, 3)
	animationList := []string{"farmtrees5.png"}

	for index, animation := range animationList {
		framePict := LoadEmbeddedImage("FallFarmTiles", animation)
		allFrames[index] = framePict
	}

	return allFrames
}
func goat() []*ebiten.Image {

	allFrames := make([]*ebiten.Image, 12, 15)
	animationList := []string{"goat0.png", "goat1.png", "goat2.png", "goat3.png", "goat4.png", "goat5.png", "goat6.png",
		"goat7.png", "goat8.png", "goat9.png", "goat10.png", "goat11.png"}

	for index, animation := range animationList {
		framePict := LoadEmbeddedImage("goat", animation)
		allFrames[index] = framePict
	}

	return allFrames
}
func piggy() []*ebiten.Image {

	allFrames := make([]*ebiten.Image, 12, 15)
	animationList := []string{"piggy0.png", "piggy1.png", "piggy2.png", "piggy3.png", "piggy4.png", "piggy5.png", "piggy6.png",
		"piggy7.png", "piggy8.png", "piggy9.png", "piggy10.png", "piggy12.png"}

	for index, animation := range animationList {
		framePict := LoadEmbeddedImage("piggy", animation)
		allFrames[index] = framePict
	}

	return allFrames
}
func turkey() []*ebiten.Image {

	allFrames := make([]*ebiten.Image, 12, 15)
	animationList := []string{"turk0.png", "turk1.png", "turk2.png", "turk3.png", "turk4.png", "turk5.png", "turk6.png",
		"turk7.png", "turk8.png", "turk9.png", "turk10.png", "turk11.png"}

	for index, animation := range animationList {
		framePict := LoadEmbeddedImage("turkey", animation)
		allFrames[index] = framePict
	}

	return allFrames
}

func LoadEmbeddedImage(folderName string, imageName string) *ebiten.Image {
	fmt.Printf("DEBUG CozyFarmProject embed:", CozyFarmProject)
	embeddedFile, err := CozyFarmProject.Open(path.Join("CozyFarmProject", folderName, imageName))
	if err != nil {
		log.Fatal("failed to load embedded image ", imageName, err)
	}
	ebitenImage, _, err := ebitenutil.NewImageFromReader(embeddedFile)
	if err != nil {
		fmt.Println("Error loading tile image:", imageName, err)
	}
	return ebitenImage
}
