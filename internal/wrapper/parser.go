package wrapper

import (
	"fmt"
	"regexp"

	"github.com/fatih/color"
	"github.com/z3orc/ender-cli/internal/logger"
)

const (
	INFO  string = "INFO"
	ERROR string = "ERROR"
	WARN  string = "WARN"
	FATAL string = "FATAL"
)

var contentRegex = regexp.MustCompile(`: (.*)`)
var typeRegex = regexp.MustCompile(`(INFO|ERROR|WARN|FATAL)`)

func Parse(line string) {
	contentMatches := contentRegex.FindStringSubmatch(line)
	typeMatches := typeRegex.FindStringSubmatch(line)
	if len(contentMatches) > 1 && len(typeMatches) > 1 {
		if typeMatches[0] == INFO {
			logger.Info.Println(contentMatches[1])
		} else if typeMatches[0] == WARN {
			logger.Warn.Println(contentMatches[1])
		} else {
			logger.Error.Println(contentMatches[1] + ". Check the logs for more info, or turn on verbose mode.")
		}
	} else if Verbose {
		color.Set(color.FgYellow)
		defer color.Unset()
		fmt.Print(line)
	}
}
