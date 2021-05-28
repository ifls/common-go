package log

type Printer interface {
	Print(args ...interface{})
}

type Logger interface {
	Debugf(fmt string, args ...interface{})
	Infof(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
	Panicf(fmt string, args ...interface{})
}

func NewLogger() Logger {

}

func Debugf(fmt string, args ...interface{}) {

}
func Infof(fmt string, args ...interface{}) {

}

func Errorf(fmt string, args ...interface{}) {

}

func Panicf(fmt string, args ...interface{}) {

}
