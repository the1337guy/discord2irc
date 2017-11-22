package main

// HandleFatal handles fatal errors, simple.
func HandleFatal(err error) {
	if err != nil {
		panic(err)
	}
}
