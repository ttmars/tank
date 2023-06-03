package pkg

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"sync"
)

// Tank 坦克
type Tank struct {
	Pos 			fyne.Position			// 左上角第一个坐标
	BodyPos			map[int]fyne.Position	// 坦克所有坐标，6
	Direction  		fyne.KeyName			// 方向
	Color 			color.Color				// 颜色
	MoveSpeed		float32					// 移速，单次移动的像素点
	BulletSpeed		float32					// 炮弹速度，单次移动的像素点
	BodySize		float32					// 大小
	BulletNumbers	int						// 载弹数
	HP				int						// 生命值
	CellSize		float32
	WindowWidth		float32
	Mutex			sync.Mutex
}

func (tank *Tank)Draw() fyne.CanvasObject {
	pos := tank.Pos
	if tank.Direction == fyne.KeyUp || tank.Direction == fyne.KeyW {
		tank.Mutex.Lock()
		tank.BodyPos[1] = fyne.Position{X: pos.X+tank.CellSize, Y: pos.Y}
		tank.BodyPos[2] = fyne.Position{X: pos.X, Y: pos.Y+tank.CellSize}
		tank.BodyPos[3] = fyne.Position{X: pos.X+tank.CellSize, Y: pos.Y+tank.CellSize}
		tank.BodyPos[4] = fyne.Position{X: pos.X+tank.CellSize*2, Y: pos.Y+tank.CellSize}
		tank.BodyPos[5] = fyne.Position{X: pos.X, Y: pos.Y+tank.CellSize*2}
		tank.BodyPos[6] = fyne.Position{X: pos.X+tank.CellSize*2, Y: pos.Y+tank.CellSize*2}
		tank.Mutex.Unlock()
	}else if tank.Direction == fyne.KeyDown || tank.Direction == fyne.KeyS {
		tank.Mutex.Lock()
		tank.BodyPos[1] = fyne.Position{X: pos.X + tank.CellSize, Y: pos.Y + tank.CellSize*2}
		tank.BodyPos[2] = fyne.Position{X: pos.X, Y: pos.Y + tank.CellSize}
		tank.BodyPos[3] = fyne.Position{X: pos.X + tank.CellSize, Y: pos.Y + tank.CellSize}
		tank.BodyPos[4] = fyne.Position{X: pos.X + tank.CellSize*2, Y: pos.Y + tank.CellSize}
		tank.BodyPos[5] = fyne.Position{X: pos.X, Y: pos.Y}
		tank.BodyPos[6] = fyne.Position{X: pos.X + tank.CellSize*2, Y: pos.Y}
		tank.Mutex.Unlock()
	}else if tank.Direction == fyne.KeyLeft || tank.Direction == fyne.KeyA{
		tank.Mutex.Lock()
		tank.BodyPos[1] = fyne.Position{X: pos.X + tank.CellSize, Y: pos.Y}
		tank.BodyPos[2] = fyne.Position{X: pos.X, Y: pos.Y+tank.CellSize}
		tank.BodyPos[3] = fyne.Position{X: pos.X+tank.CellSize, Y: pos.Y+tank.CellSize}
		tank.BodyPos[4] = fyne.Position{X: pos.X+tank.CellSize*2, Y: pos.Y}
		tank.BodyPos[5] = fyne.Position{X: pos.X+tank.CellSize, Y: pos.Y+tank.CellSize*2}
		tank.BodyPos[6] = fyne.Position{X: pos.X+tank.CellSize*2, Y: pos.Y+tank.CellSize*2}
		tank.Mutex.Unlock()
	}else if tank.Direction == fyne.KeyRight || tank.Direction == fyne.KeyD{
		tank.Mutex.Lock()
		tank.BodyPos[1] = fyne.Position{X: pos.X+tank.CellSize, Y: pos.Y}
		tank.BodyPos[2] = fyne.Position{X: pos.X+tank.CellSize*2, Y: pos.Y+tank.CellSize}
		tank.BodyPos[3] = fyne.Position{X: pos.X+tank.CellSize, Y: pos.Y+tank.CellSize}
		tank.BodyPos[4] = fyne.Position{X: pos.X, Y: pos.Y}
		tank.BodyPos[5] = fyne.Position{X: pos.X+tank.CellSize, Y: pos.Y+tank.CellSize*2}
		tank.BodyPos[6] = fyne.Position{X: pos.X, Y: pos.Y+tank.CellSize*2}
		tank.Mutex.Unlock()
	}
	var tankBody []fyne.CanvasObject
	tank.Mutex.Lock()
	for _,v := range tank.BodyPos {
		t := canvas.NewRectangle(tank.Color)
		t.Resize(fyne.NewSize(tank.CellSize, tank.CellSize))
		t.Move(v)
		tankBody = append(tankBody, t)
	}
	tank.Mutex.Unlock()

	// 坦克状态
	HPInfo := canvas.NewText(fmt.Sprintf("HP:%v", tank.HP), tank.Color)
	HPInfo.TextStyle = fyne.TextStyle{Italic: true, Bold:true}
	if tank.Color == Red {
		HPInfo.Alignment = fyne.TextAlignLeading
		HPInfo.Move(fyne.Position{X: 0, Y: 0})
	}else{
		HPInfo.Alignment = fyne.TextAlignTrailing
		HPInfo.Move(fyne.Position{X: tank.WindowWidth-20, Y: 0})
	}
	HPInfo.TextSize = tank.CellSize*2
	tankBody = append(tankBody, HPInfo)
	c := container.NewWithoutLayout(tankBody...)
	return c
}
