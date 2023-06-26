package logger

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var Info *log.Logger
var Warn *log.Logger
var Error *log.Logger

func init() {

	infoText := color.New(color.FgWhite).SprintFunc()
	warnText := color.New(color.FgYellow).SprintFunc()
	errorText := color.New(color.FgRed).SprintFunc()

	Info = log.New(os.Stdout, infoText("[info] "), log.Ltime)
	Warn = log.New(os.Stdout, warnText("[warn] "), log.Ltime)
	Error = log.New(os.Stderr, errorText("[error] "), 0)
}
