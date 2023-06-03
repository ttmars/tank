package pkg

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"gitee.com/ttmasx/cob/player"
	"image/color"
	"sync"
	"tank/src/mp3"
	"time"
)


var (
	cellSize float32 = 10
	WindowWidth float32 = 800
	WindowHigh float32 = 800
	fps = 10
)

// Game 整个游戏状态
type Game struct {
	WindowWidth	float32
	WindowHigh	float32
	CellSize	float32
	RedTank		*Tank					// 红方坦克
	BlueTank	*Tank					// 蓝方坦克
	BulletMap	map[string]Bullet		// 炮弹
	LoveMap		map[string]Love			// 爱心
	MaxLoveNum	int						// 最大爱心数
	GunVoice    *player.Player				// 游戏声音
	Mutex		sync.Mutex
	FPS         int						// 页面渲染间隔，默认10毫秒
	BGM			bool					// 是否播放背景音乐
}

// InitGame 初始化game
func InitGame() *Game {
	return &Game{
		BGM: true,
		WindowHigh: WindowHigh,
		WindowWidth: WindowWidth,
		CellSize: cellSize,
		RedTank:  &Tank{
			Pos:       fyne.Position{X: 200, Y: 500},
			Direction: fyne.KeyW,
			Color: Red,
			MoveSpeed: cellSize,
			BodySize: cellSize*4,
			BulletSpeed: cellSize,
			HP: 10,
			BodyPos: make(map[int]fyne.Position, 6),
			CellSize: cellSize,
			WindowWidth: WindowWidth,

		},
		BlueTank: &Tank{
			Pos:       fyne.Position{X: 500, Y: 200},
			Direction: fyne.KeyUp,
			Color: color.Black,
			MoveSpeed: cellSize,
			BodySize: cellSize*4,
			BulletSpeed: cellSize,
			HP: 10,
			BodyPos: make(map[int]fyne.Position, 6),
			CellSize: cellSize,
			WindowWidth: WindowWidth,
		},
		BulletMap: make(map[string]Bullet),
		LoveMap: make(map[string]Love),
		MaxLoveNum: 3,
		FPS: fps,
		GunVoice: player.NewPlayerByEmbed(mp3.EFile, []string{"start.mp3", "98k.mp3","gunshot.mp3","hit1.mp3", "hit2.mp3" ,"hit3.mp3", "laser.mp3", "夺回苏雅.mp3"}),
	}
}

func (g *Game)Draw() fyne.CanvasObject {
	readTank := g.RedTank.Draw()		// 红坦
	blueTank := g.BlueTank.Draw()		// 蓝坦

	var collect []fyne.CanvasObject
	// 炮弹
	for _,v := range g.BulletMap {
		collect = append(collect, v.Draw())
	}
	// 爱心
	for _,v := range g.LoveMap {
		collect = append(collect, v.Draw())
	}

	c := container.NewWithoutLayout(readTank, blueTank, container.NewWithoutLayout(collect...))
	return c
}

// ListenKeyEvent 处理键盘事件
func (g *Game)ListenKeyEvent(event *fyne.KeyEvent, window fyne.Window)  {
	switch event.Name {
	case fyne.KeyUp:
		if g.BlueTank.Direction == event.Name {
			if g.BlueTank.Pos.Y - g.BlueTank.MoveSpeed >= 0 {
				g.BlueTank.Pos.Y -= g.BlueTank.MoveSpeed
			}else{
				g.BlueTank.Pos.Y = 0
			}
		}else{
			g.BlueTank.Direction = fyne.KeyUp
		}
	case fyne.KeyDown:
		if g.BlueTank.Direction == event.Name {
			if g.BlueTank.Pos.Y + g.BlueTank.BodySize + g.BlueTank.MoveSpeed <= g.WindowHigh {
				g.BlueTank.Pos.Y += g.BlueTank.MoveSpeed
			}else{
				g.BlueTank.Pos.Y = g.WindowHigh-g.BlueTank.BodySize
			}
		}else{
			g.BlueTank.Direction = fyne.KeyDown
		}
	case fyne.KeyLeft:
		if g.BlueTank.Direction == event.Name {
			if g.BlueTank.Pos.X-g.BlueTank.MoveSpeed > 0 {
				g.BlueTank.Pos.X -= g.BlueTank.MoveSpeed
			}else{
				g.BlueTank.Pos.X = 0
			}
		}else{
			g.BlueTank.Direction = fyne.KeyLeft
		}
	case fyne.KeyRight:
		if g.BlueTank.Direction == event.Name {
			if g.BlueTank.Pos.X + g.BlueTank.BodySize + g.BlueTank.MoveSpeed <= g.WindowWidth {
				g.BlueTank.Pos.X += g.BlueTank.MoveSpeed
			}else{
				g.BlueTank.Pos.X = g.WindowWidth-g.BlueTank.BodySize
			}
		}else{
			g.BlueTank.Direction = fyne.KeyRight
		}

	case fyne.KeyW:
		if g.RedTank.Direction == event.Name {
			if g.RedTank.Pos.Y - g.RedTank.MoveSpeed >= 0 {
				g.RedTank.Pos.Y -= g.RedTank.MoveSpeed
			}else{
				g.RedTank.Pos.Y = 0
			}
		}else{
			g.RedTank.Direction = fyne.KeyW
		}
	case fyne.KeyS:
		if g.RedTank.Direction == event.Name {
			if g.RedTank.Pos.Y + g.RedTank.BodySize + g.RedTank.MoveSpeed <= g.WindowHigh {
				g.RedTank.Pos.Y += g.RedTank.MoveSpeed
			}else{
				g.RedTank.Pos.Y = g.WindowHigh-g.RedTank.BodySize
			}
		}else{
			g.RedTank.Direction = fyne.KeyS
		}
	case fyne.KeyA:
		if g.RedTank.Direction == event.Name {
			if g.RedTank.Pos.X-g.RedTank.MoveSpeed > 0 {
				g.RedTank.Pos.X -= g.RedTank.MoveSpeed
			}else{
				g.RedTank.Pos.X = 0
			}
		}else{
			g.RedTank.Direction = fyne.KeyA
		}
	case fyne.KeyD:
		if g.RedTank.Direction == event.Name {
			if g.RedTank.Pos.X + g.RedTank.BodySize + g.RedTank.MoveSpeed <= g.WindowWidth {
				g.RedTank.Pos.X += g.RedTank.MoveSpeed
			}else{
				g.RedTank.Pos.X = g.WindowWidth-g.RedTank.BodySize
			}
		}else{
			g.RedTank.Direction = fyne.KeyD
		}
	case fyne.KeySpace:			// 红方发射炮弹
		g.GunVoice.Play("laser.mp3")
		// 计算炮口坐标
		var mouthPos fyne.Position
		switch g.RedTank.Direction{
		case fyne.KeyUp,fyne.KeyW:
			mouthPos = fyne.Position{X: g.RedTank.Pos.X + g.CellSize, Y: g.RedTank.Pos.Y,}
		case fyne.KeyDown,fyne.KeyS:
			mouthPos = fyne.Position{X: g.RedTank.Pos.X + g.CellSize, Y: g.RedTank.Pos.Y+g.CellSize*2,}
		case fyne.KeyLeft,fyne.KeyA:
			mouthPos = fyne.Position{X: g.RedTank.Pos.X, Y: g.RedTank.Pos.Y+g.CellSize,}
		case fyne.KeyRight,fyne.KeyD:
			mouthPos = fyne.Position{X: g.RedTank.Pos.X+g.CellSize*2, Y: g.RedTank.Pos.Y+g.CellSize,}
		}
		g.BulletMap[GetRandomString(10)] = Bullet{
			TankName:  "red",
			Direction: g.RedTank.Direction,
			Pos:       mouthPos,
			Color:     g.RedTank.Color,
			CellSize: g.CellSize,
		}
	case fyne.KeyReturn:			// 蓝方发射炮弹
		g.GunVoice.Play("laser.mp3")

		var mouthPos fyne.Position
		switch g.BlueTank.Direction{
		case fyne.KeyUp,fyne.KeyW:
			mouthPos = fyne.Position{X: g.BlueTank.Pos.X + g.CellSize, Y: g.BlueTank.Pos.Y,}
		case fyne.KeyDown,fyne.KeyS:
			mouthPos = fyne.Position{X: g.BlueTank.Pos.X + g.CellSize, Y: g.BlueTank.Pos.Y+g.CellSize*2,}
		case fyne.KeyLeft,fyne.KeyA:
			mouthPos = fyne.Position{X: g.BlueTank.Pos.X, Y: g.BlueTank.Pos.Y+g.CellSize,}
		case fyne.KeyRight,fyne.KeyD:
			mouthPos = fyne.Position{X: g.BlueTank.Pos.X+g.CellSize*2, Y: g.BlueTank.Pos.Y+g.CellSize,}
		}
		g.Mutex.Lock()
		g.BulletMap[GetRandomString(10)] = Bullet{
			TankName:  "blue",
			Direction: g.BlueTank.Direction,
			Pos:       mouthPos,
			Color:     g.BlueTank.Color,
			CellSize: g.CellSize,
		}
		g.Mutex.Unlock()
	}
	//fmt.Println("blue:",g.BlueTank.Pos)
	window.SetContent(g.Draw())
}

// UpdateGameStatus 实时更新游戏状态，碰撞检测
func (game *Game)UpdateGameStatus(window fyne.Window)  {
	dlgRedWin := dialog.NewInformation("游戏结束", "红方胜利！", window)
	dlgRedWin.SetOnClosed(func() {
		game.RedTank.HP = 10
		game.BlueTank.HP = 10
		game.RedTank.Pos = fyne.Position{X: 200, Y: 500}
		game.BlueTank.Pos = fyne.Position{X: 500, Y: 200}
	})
	dlgBlueWin := dialog.NewInformation("游戏结束", "黑方胜利！", window)
	dlgBlueWin.SetOnClosed(func() {
		game.RedTank.HP = 10
		game.BlueTank.HP = 10
		game.RedTank.Pos = fyne.Position{X: 200, Y: 500}
		game.BlueTank.Pos = fyne.Position{X: 500, Y: 200}
	})
	for {
		time.Sleep(time.Millisecond*time.Duration(game.FPS))
		// 爱心检测
		for k,v := range game.LoveMap {
			if GetTwoPD(v.Pos, game.RedTank.Pos) < game.CellSize*3 {
				for _,body := range game.RedTank.BodyPos {
					if v.Pos.X == body.X && v.Pos.Y == body.Y {
						// 吃掉红心
						game.RedTank.HP++
						delete(game.LoveMap, k)
					}
				}
			}

			if GetTwoPD(v.Pos, game.BlueTank.Pos) < game.CellSize*3 {
				for _,body := range game.BlueTank.BodyPos {
					if v.Pos.X == body.X && v.Pos.Y == body.Y {
						// 吃掉红心
						game.BlueTank.HP++
						delete(game.LoveMap, k)
					}
				}
			}
		}

		// 炮弹碰撞检测
		for k,v := range game.BulletMap {
			if v.Pos.X < 0 || v.Pos.Y < 0 || v.Pos.X > game.WindowWidth || v.Pos.Y > game.WindowHigh {
				delete(game.BulletMap, k)
			}else{
				switch v.Direction{
				case fyne.KeyUp,fyne.KeyW:
					v.Pos.Y -= game.BlueTank.BulletSpeed
					game.BulletMap[k] = v
				case fyne.KeyDown,fyne.KeyS:
					v.Pos.Y += game.BlueTank.BulletSpeed
					game.BulletMap[k] = v
				case fyne.KeyLeft,fyne.KeyA:
					v.Pos.X -= game.BlueTank.BulletSpeed
					game.BulletMap[k] = v
				case fyne.KeyRight,fyne.KeyD:
					v.Pos.X += game.BlueTank.BulletSpeed
					game.BulletMap[k] = v
				}
			}

			// 炮弹碰撞检测
			if v.TankName == "red" {
				// 炮弹与炮弹检测
				game.Mutex.Lock()
				for myKey,mybullet := range game.BulletMap {
					if mybullet.TankName == "blue" && v.Pos.X == mybullet.Pos.X && v.Pos.Y == mybullet.Pos.Y {
						delete(game.BulletMap, k)
						delete(game.BulletMap, myKey)
					}
				}
				game.Mutex.Unlock()

				// 炮弹与坦克碰撞检测
				if GetTwoPD(v.Pos, game.BlueTank.Pos) < game.CellSize*3 {
					// body检测
					game.Mutex.Lock()
					for _,body := range game.BlueTank.BodyPos {
						if v.Pos.X == body.X && v.Pos.Y == body.Y {
							if game.BlueTank.HP > 0 {
								game.BlueTank.HP--
								game.GunVoice.Play("hit3.mp3")
							}
							if game.BlueTank.HP <= 0 {
								//fmt.Println("红方胜利！")
								game.GunVoice.Play("hit2.mp3")
								dlgRedWin.Show()
							}
							delete(game.BulletMap, k)	// 删除炮弹
							break
						}
					}
					game.Mutex.Unlock()
				}
			}else{
				// 炮弹与炮弹检测
				game.Mutex.Lock()
				for myKey,mybullet := range game.BulletMap {
					if mybullet.TankName == "red" && v.Pos.X == mybullet.Pos.X && v.Pos.Y == mybullet.Pos.Y {
						delete(game.BulletMap, k)
						delete(game.BulletMap, myKey)
					}
				}
				game.Mutex.Unlock()

				if GetTwoPD(v.Pos, game.RedTank.Pos) < game.CellSize*3 {
					// body检测
					game.Mutex.Lock()
					for _,body := range game.RedTank.BodyPos {
						if v.Pos.X == body.X && v.Pos.Y == body.Y {
							if game.RedTank.HP > 0 {
								game.RedTank.HP--
								game.GunVoice.Play("hit3.mp3")
							}
							if game.RedTank.HP <= 0 {
								//fmt.Println("蓝方胜利！")
								game.GunVoice.Play("hit2.mp3")
								dlgBlueWin.Show()
							}
							delete(game.BulletMap, k)	// 删除炮弹
							break
						}
					}
					game.Mutex.Unlock()
				}
			}

		}
		window.SetContent(game.Draw())
	}
}

// GenerateProps 系统随机生成游戏道具
func (game *Game)GenerateProps()  {
	for {
		time.Sleep(time.Second*3)		// 每三秒更新一轮道具
		if len(game.LoveMap) < game.MaxLoveNum {
			x := float32(GetRandomNum(0,int(game.WindowWidth-game.CellSize*3))/10*10)
			y := float32(GetRandomNum(0,int(game.WindowHigh-game.CellSize*3))/10*10)
			game.LoveMap[GetRandomString(10)] = Love{
				Pos:       fyne.Position{X: x, Y: y},
				CellSize:  game.CellSize,
			}
		}
	}
}
