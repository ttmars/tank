package pkg

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

// Bullet 炮弹
type Bullet struct {
	TankName		string
	Direction  		fyne.KeyName
	Pos				fyne.Position
	Color			color.Color
	CellSize		float32
}

func (b *Bullet)Draw() fyne.CanvasObject {
	t := canvas.NewCircle(b.Color)
	t.Resize(fyne.NewSize(b.CellSize, b.CellSize))
	t.Move(b.Pos)
	return t
}
