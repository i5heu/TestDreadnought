package fmtLog

import "github.com/fatih/color"

func Log(a ...interface{}) (n int, err error) {
	clr := color.New(color.FgCyan)
	inter := []interface{}{"	Extention:"}
	args := append(inter, a...)
	return clr.Println(args...)
}
