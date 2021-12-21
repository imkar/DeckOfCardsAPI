package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Parses incoming request to provided data model interface.
func requestParser(w http.ResponseWriter, r *http.Request, data interface{}) error {
	fmt.Println(json.NewDecoder(r.Body).Decode(data))
	return json.NewDecoder(r.Body).Decode(data)
}
