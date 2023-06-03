package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"gitee.com/ttmasx/cob/mytheme"
	"image/color"
	"log"
	"tank/pkg"
)

func main() {
	log.SetFlags(log.Ldate|log.Lshortfile)
	mytheme.InitFont(mytheme.STKAITI)		// 设置中文字体
	myApp := app.New()
	myWindow := myApp.NewWindow("坦克大战")
	myWindow.CenterOnScreen()
	myWindow.SetMaster()
	myApp.Settings().SetTheme(mytheme.NewMyTheme(color.NRGBA{R: 0xd3, G: 0xd3, B: 0xd3, A: 0xff}, nil))		// 设置背景为浅灰色

	// 初始化game
	game := pkg.InitGame()
	myWindow.SetContent(game.Draw())
	myWindow.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		game.ListenKeyEvent(event, myWindow)
	})
	myWindow.Resize(fyne.NewSize(pkg.WindowWidth, pkg.WindowHigh))
	go game.UpdateGameStatus(myWindow)
	go game.GenerateProps()
	if game.BGM {
		game.GunVoice.Play("start.mp3")
	}
	myWindow.ShowAndRun()
}
