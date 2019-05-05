package plog

// initialize the logger
//
// Every concret logger calls this method
// to initialize onself
func initLogger(l loggerImpl) {
	l.setImpl(l)
	l.SetLevel(Info)
}

type loggerImpl interface {
	SetLevel(LogLevel)
	setImpl(logImpl)
	log(LogMessage)
}

func (self *logger) setImpl(impl logImpl) {
	self.logImpl = impl
}
