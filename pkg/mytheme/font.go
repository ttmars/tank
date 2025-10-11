package mytheme

import (
	"github.com/flopp/go-findfont"
	"log"
	"os"
)

//https://github.com/flopp/go-findfont

type MyFont string

const (
	STKAITI MyFont = "STKAITI.TTF"			// 正楷
	Deng MyFont = "Deng.ttf"				// 隶书
	Dengb MyFont = "Dengb.ttf"
	STXIHEI MyFont = "STXIHEI.TTF"
	STFANGSO MyFont = "STFANGSO.TTF"
	Dengl MyFont = "Dengl.ttf"
	STSONG MyFont = "STSONG.TTF"
	STZHONGS MyFont = "STZHONGS.TTF"

	FZSTK MyFont = "FZSTK.TTF"				// 草书
)

func InitFont(font MyFont)  {
	fontPath, err := findfont.Find(string(font))
	if err != nil {
		log.Fatalln("系统不支持该字体！", err)
	}
	os.Setenv("FYNE_FONT", fontPath)
}
