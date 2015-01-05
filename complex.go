package assert

type Complex struct {
	logFacade *logFacade
	actual    complex128
}

func (a *Complex) IsZero() *Complex {
	return a.isTrue(a.actual == 0,
		"Expected zero, but was <%v>.", a.actual)
}

func (a *Complex) IsNonZero() *Complex {
	return a.isTrue(a.actual != 0,
		"Expected nonzero, but was zero.")
}

func (a *Complex) isTrue(condition bool, format string, args ...interface{}) *Complex {
	logIfFalse(a.logFacade, condition, format, args...)
	return a
}