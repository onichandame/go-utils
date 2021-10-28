package goutils

func Try(fn func()) error {
	var err error
	func() {
		defer RecoverToErr(&err)
		fn()
	}()
	return err
}
