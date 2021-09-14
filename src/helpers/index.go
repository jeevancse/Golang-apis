package helpers

func responseMethod(message string, code int, success bool, data string) {
	type Person struct {
		message string
		code    int
		success bool
		// data    interface{}
	}
	// Handle ...

}

func responseObject(message string, code int, success bool) (map[string]interface{}, error) {
	return map[string]interface{}{
			"message": message,
			"code":    code,
			"body":    success,
		},
		nil
}
