package middleware

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/globalxtreme/gobaseconf/data"
	"github.com/globalxtreme/gobaseconf/response/error"
	"net/http"
	"strings"
)

func EmployeeIdentifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("IDENTIFIER")
		if len(token) == 0 {
			error.ErrUnauthenticated("IDENTIFIER not found")
		}

		tokenDecode, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			error.ErrUnauthenticated("Unable to decode token!!")
		}

		iv := tokenDecode[0:16]
		tokenExplode := strings.Split(string(tokenDecode[16:]), "-:-")

		secret := []byte(tokenExplode[1][0:tokenDecode[0]])
		identifierData := []byte(tokenExplode[1][0:])

		block, err := aes.NewCipher(secret)
		if err != nil {
			fmt.Println("Failed to initialize cipher: ", err)
			return
		}
		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(identifierData, identifierData)

		err = json.Unmarshal(identifierData, &data.Employee)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
