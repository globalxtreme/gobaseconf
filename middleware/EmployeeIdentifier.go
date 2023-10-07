package middleware

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/globalxtreme/gobaseconf/data"
	errResponse "github.com/globalxtreme/gobaseconf/response/error"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func EmployeeIdentifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("IDENTIFIER")
		if len(token) == 0 {
			errResponse.ErrXtremeUnauthenticated("IDENTIFIER not found")
		}

		tokenDecode, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			errResponse.ErrXtremeUnauthenticated("Unable to decode token!!")
		}

		iv := tokenDecode[0:16]
		tokenExplode := strings.Split(string(tokenDecode[16:]), "-:-")

		secretLength, _ := strconv.Atoi(tokenExplode[0])
		secret := []byte(tokenExplode[1][0:secretLength])
		identifierData := []byte(tokenExplode[1][secretLength:])

		block, err := aes.NewCipher(secret)
		if err != nil {
			errResponse.ErrXtremeUnauthenticated(fmt.Sprintf("Unable to decode token!! %s", err))
		}

		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(identifierData, identifierData)
		identifierData = unpadPKCS7(identifierData)

		identifierData, err = DecompressZlib(identifierData)
		if err != nil {
			errResponse.ErrXtremeUnauthenticated(fmt.Sprintf("Unable to decompress data: %s", err))
		}

		err = json.Unmarshal(identifierData, &data.Employee)
		if err != nil {
			errResponse.ErrXtremeUnauthenticated(fmt.Sprintf("Unable to decode data json: %s", err))
		}

		next.ServeHTTP(w, r)
	})
}

func unpadPKCS7(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func DecompressZlib(data []byte) (resData []byte, err error) {
	b := bytes.NewBuffer(data)

	var r io.Reader
	r, err = gzip.NewReader(b)
	if err != nil {
		return
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		return
	}

	resData = resB.Bytes()

	return
}
