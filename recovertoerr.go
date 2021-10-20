package goutils

func RecoverToErr(err *error) {
	if er := recover(); er != nil {
		if e, ok := er.(error); ok {
			*err = e
			return
		} else {
			panic(er)
		}
	}
}
