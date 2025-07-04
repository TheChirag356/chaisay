package chaisay

import (
	"github.com/TheChirag356/chaisay/art"
	"github.com/TheChirag356/chaisay/textbox"
)

type ChaiSay struct {
	Text     string
	ArtStyle string
}

func New(text string, artstyle string) ChaiSay {
	return ChaiSay{
		Text:     text,
		ArtStyle: artstyle,
	}
}

func (c *ChaiSay) Print() {
	textbox.PrintTextBox(c.Text)
	art.Art(c.ArtStyle)
}
