package helper

import "fmt"

func MapTokenAuthorizationHeader(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}
