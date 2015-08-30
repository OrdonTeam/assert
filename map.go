package assert

type MapType struct {
	logFacade *logFacade
	actual    map[interface{}]interface{}
}

func (a *MapType) IsEmpty() *MapType {
	return a.isTrue(len(a.actual) == 0,
		"Expected len <0>, but was <%d>.", len(a.actual))
}

func (a *MapType) isTrue(condition bool, format string, args ...interface{}) *MapType {
	logIfFalse(a.logFacade, condition, format, args...)
	return a
}