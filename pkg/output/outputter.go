package output

import (
	"github.com/fatih/color"
)

type defaultOutputter struct {
	styles outputStyles
}

type outputStyles struct {
	vvv, vv, v, m, o *color.Color
}

func NewDefaultOutputter() Outputter {
	styles := outputStyles{
		vvv: color.RGB(64, 64, 64),
		vv:  color.RGB(128, 128, 128),
		v:   color.RGB(192, 192, 192),
		m:   color.New(color.FgWhite),
		o:   color.New(color.FgGreen),
	}
	return &defaultOutputter{
		styles: styles,
	}
}

func (d *defaultOutputter) VVVMessage(msg string) {
	d.styles.vvv.Println(msg)
}
func (d *defaultOutputter) VVMessage(msg string) {
	d.styles.vv.Println(msg)
}
func (d *defaultOutputter) VMessage(msg string) {
	d.styles.v.Println(msg)
}
func (d *defaultOutputter) Message(msg string) {
	d.styles.m.Println(msg)
}
func (d *defaultOutputter) Output(o any) {

	d.styles.o.Println(o)
}
