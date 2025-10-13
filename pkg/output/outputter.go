package output

import (
	"context"

	"github.com/fatih/color"
)

type CtxKey int

const OutputterKey CtxKey = 1

type defaultOutputter struct {
	styles outputStyles
	level  int
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

func GetOutputter(ctx context.Context) Outputter {
	return ctx.Value(OutputterKey).(Outputter)
}

func SetOutputter(ctx context.Context, outputter Outputter) context.Context {
	return context.WithValue(ctx, OutputterKey, outputter)
}

func (d *defaultOutputter) SetLevel(level int) {
	d.level = level
}

func (d *defaultOutputter) VVVMessage(msg string) {
	if d.level >= 3 {
		d.styles.vvv.Println(msg)
	}
}
func (d *defaultOutputter) VVMessage(msg string) {
	if d.level >= 2 {
		d.styles.vv.Println(msg)
	}
}
func (d *defaultOutputter) VMessage(msg string) {
	if d.level >= 1 {
		d.styles.v.Println(msg)
	}
}
func (d *defaultOutputter) Message(msg string) {
	if d.level >= 0 {
		d.styles.m.Println(msg)
	}
}
func (d *defaultOutputter) Output(o any) {
	d.styles.o.Println(o)
}
