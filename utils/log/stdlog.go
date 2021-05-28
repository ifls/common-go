package log

import (
	"fmt"
	"os"
)

type stdlog struct {
}

func (l *stdlog) Print(args ...interface{}) {
	content := fmt.Sprintln(args)
	os.Stdout.Write([]byte(content))
}
