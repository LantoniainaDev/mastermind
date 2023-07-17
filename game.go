package main

import (
	"log"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"mastermind.dude/mind"
)

const (
	screenWidth  = 320
	screenHeight = 240
	maxAttemp    = 12
	combinaison  = 4
	nbrDeChoix   = 6
)

type Game struct {
	affichage string
	attemps   []int
	line      int
	secret    []int
	match     [2]int
	playing   bool
	won       bool
}

var nums []ebiten.Key = []ebiten.Key{
	ebiten.Key1,
	ebiten.Key2,
	ebiten.Key3,
	ebiten.Key4,
	ebiten.Key5,
	ebiten.Key6,
}

func (this *Game) Update(screen *ebiten.Image) error {
	if !this.playing {
		if !this.won {
			this.DrawLost(screen)
		} else {
			this.DrawVictory(screen)
		}
		return nil
	}

	if this.line >= maxAttemp {
		this.playing = false
		this.won = false
		return nil
	}

	onPress(ebiten.KeyBackspace, func() {
		l := len(this.attemps)
		if l != 0 {
			this.attemps = this.attemps[:l-1]
		}
	})

	if len(this.attemps) < combinaison {

		for i := 0; i < len(nums); i++ {
			onPress(nums[i], func() {
				num := i + 1
				if num != 0 {
					this.attemps = append(this.attemps, num)
				}
			})
		}

	} else {
		onPress(ebiten.KeyEnter, this.AddLine)
		this.won = this.match[1] == 4
		this.playing = !(this.match[1] == 4)
	}

	this.Draw(screen)
	return nil
}

func (this *Game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}

func (this *Game) Draw(screen *ebiten.Image) error {
	var text string = this.affichage
	text += "\n" + join(this.attemps, " ")
	text += clignote()
	ebitenutil.DebugPrint(screen, text)

	return nil
}

func (game *Game) DrawLost(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, `
	Vous avez epuise vos chances,
	la reponse etait`+join(game.secret[:], " "),
		0, screenHeight*2/3)
	game.waitRestart(screen)
}

func (game *Game) DrawVictory(screen *ebiten.Image) {
	txt := game.affichage
	txt += "\n" + "felicitation vous avez trouve en " + strconv.Itoa(game.line) +
		" tentatives"
	ebitenutil.DebugPrint(screen, txt)
	game.waitRestart(screen)
}

func (game *Game) waitRestart(screen *ebiten.Image) {
	onPress(ebiten.KeyEnter, func() {
		game.init()
	})
}

func (game *Game) init() {
	game.secret = mind.Generate(nbrDeChoix, combinaison)
	game.affichage = " a b c d e|f"
	game.playing = true
	game.attemps = game.attemps[:0]
	game.line = 0
	game.match[1] = 0
}

func main() {
	var game Game
	game.init()
	err := ebiten.RunGame(&game)
	if err != nil {
		log.Fatal("le jeu ne veut pas demarer")
	}
}

func join(array []int, separator string) string {
	var concat string
	for i := 0; i < len(array); i++ {
		el := strconv.Itoa(array[i])
		concat += separator + el
	}

	return concat
}

func onPress(key ebiten.Key, callBack func()) {
	if inpututil.IsKeyJustPressed(key) {
		callBack()
	}
}

func clignote() string {
	var cursor string
	second := time.Now().Nanosecond()
	if (second/500000000)%2 != 1 {
		cursor = "_"
	}

	return cursor
}
func (this *Game) AddLine() {
	this.match = mind.Match(this.secret, this.attemps)
	this.affichage += "\n" + join(this.attemps, " ") + " " + join(this.match[:], "|")
	this.attemps = this.attemps[:0]
	this.line += 1
}
