package middleware

import (
	"encoding/json"
	"github.com/globalxtreme/gobaseconf/response"
	"log"
	"net/http"
)

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				w.Header().Set("Content-Type", "application/json")
				log.Println(r)

				var res *response.ResponseError
				if panicData, ok := r.(*response.ResponseError); ok {
					res = panicData
				} else {
					res = &response.ResponseError{
						Status: response.Status{
							Code:    http.StatusInternalServerError,
							Message: "An error Occurred.",
						},
					}
				}

				w.WriteHeader(res.Status.Code)

				jsonData, err := json.Marshal(res)
				if err != nil {
					log.Println("Failed to marshal error response:", err)
					return
				}

				w.Write(jsonData)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
