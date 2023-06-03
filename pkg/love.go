package pkg

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"tank/src"
)

type Love struct {
	Pos				fyne.Position
	CellSize		float32
}

// Draw 优化：love对象改成变量
func (l *Love)Draw() fyne.CanvasObject {
	image := canvas.NewImageFromResource(src.ResourceLoveSvg)
	image.Resize(fyne.NewSize(l.CellSize*2, l.CellSize*2))
	image.Move(l.Pos)
	return image
}
