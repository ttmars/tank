package mytheme

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

//https://github.com/lusingander/fyne-theme-generator

type MyTheme struct {
	BackgroundColor		color.Color			// 背景色
	HyperlinkColor		color.Color			// Hyperlink
}
var _ fyne.Theme = (*MyTheme)(nil)

func NewMyTheme(BackgroundColor color.Color, HyperlinkColor color.Color) *MyTheme {
	return &MyTheme{BackgroundColor: BackgroundColor, HyperlinkColor: HyperlinkColor}
}

func (m MyTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNamePrimary && m.HyperlinkColor != nil{
		return m.HyperlinkColor
	}
	if name == theme.ColorNameBackground && m.BackgroundColor != nil{
		return m.BackgroundColor
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m MyTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m MyTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m MyTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
