package router

import (
	"regexp"
)

func ValidatePath(path string) (bool, string) {
	// Ensure the path starts with a single slash
	if path[0] != '/' {
			path = "/" + path
	}

	// Regular expression to validate the path
	validationRex := regexp.MustCompile(`^\/[a-zA-Z0-9]+(\/[a-zA-Z0-9]+)*(\/)?(\?[a-zA-Z0-9=&]*)?$`)

	// Validate the path
	isValid := validationRex.MatchString(path)

	return isValid, path
}

func ValidateHttpMethod(method string) bool {
	if method == "GET" || method == "POST" || method == "PUT" || method == "DELETE" {
		return true
	}

	return false
}
