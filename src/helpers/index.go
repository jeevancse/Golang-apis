package helpers

func responseMethod(message string, code int, success bool, data string) {
	type Person struct {
		message string
		code    int
		success bool
		data    interface{}
	}

	p1 := Person{message: message, code: code, success: success, data: data}
	return p1
}
