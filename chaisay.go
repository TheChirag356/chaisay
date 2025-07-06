package chaisay

import (
	"fmt"

	"github.com/TheChirag356/chaisay/art"
)

type ChaiSay struct {
	Text     string
	ArtStyle art.ArtStyle
}

func New(text string, artstyle art.ArtStyle) ChaiSay {
	return ChaiSay{
		Text:     text,
		ArtStyle: artstyle,
	}
}

func (c *ChaiSay) Print() {
	PrintTextBox(c.Text)
	fmt.Println(c.ArtStyle)
}
