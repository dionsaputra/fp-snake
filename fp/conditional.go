package fp

func DoIf(condition bool, action func()) {
	if condition {
		action()
	}
}

