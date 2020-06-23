package goping

import (
	"fmt"
	"net/http"
)

// GoPing method pings a given website and returns the result
func GoPing(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	return fmt.Sprintf("%v", resp.StatusCode)
}
