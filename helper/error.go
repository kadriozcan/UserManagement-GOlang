package helper

func ErrorPanic(err error) {
	// if there is an error
	if err != nil {
		// it stops the program
		panic(err)
	}
}
