package routines

func DoIf(condition bool, action func()) {
	if condition {
		action()
	}
}
