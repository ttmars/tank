package pkg

import (
	"fmt"
	"fyne.io/fyne/v2"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"
)

// GetTwoPD 获取两点的距离
func GetTwoPD(p1 fyne.Position, p2 fyne.Position) float32 {
	return float32(math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))))
}

func GetRandomString(n int) (result string) {
	rand.Seed(time.Now().UnixNano())
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

func GetRandomNum(N,M int) (result int)  {
	if M<N {
		log.Fatalln("N必须小于等于M!")
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(M-N+1)
	return n+N
}

var Red = color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
