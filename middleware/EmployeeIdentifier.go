package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/globalxtreme/gobaseconf/data"
	"github.com/globalxtreme/gobaseconf/response/error"
	"net/http"
)

func EmployeeIdentifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("IDENTIFIER")
		if len(token) == 0 {
			error.ErrUnauthenticated("IDENTIFIER not found")
		}

		jsonDecode, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			error.ErrUnauthenticated("Unable to decode token!!")
		}

		err = json.Unmarshal(jsonDecode, &data.Employee)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
