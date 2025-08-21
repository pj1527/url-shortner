package utils

import "fmt"

func GenerateURL(httpScheme string, domain string, shortKey string) string {
	return fmt.Sprintf("%s://%s/%s", httpScheme, domain, shortKey)
}
