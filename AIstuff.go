package main

type Enemy struct {
	X, Y   int
	Health int
	Active bool
}

var enemies []*Enemy

func initializeEnemies() {
	enemies = append(enemies, &Enemy{X: 150, Y: 100, Health: 100, Active: true})

}

func (m *mapGame) attackEnemy(playerX, playerY int) {
	for _, enemy := range enemies {
		if isPlayerNearEnemy(playerX, playerY, enemy) && enemy.Active {
			enemy.Health -= 10
			if enemy.Health <= 0 {
				enemy.Active = false
			}
			break
		}
	}
}

func isPlayerNearEnemy(x int, y int, enemy *Enemy) bool {

	return true
}
