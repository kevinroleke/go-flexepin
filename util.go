package flexepin

import (
	"io"
	"fmt"
	"time"
	"bytes"
	"net/http"
	"strconv"
	"io/ioutil"
	"math/rand"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/md5"
	"encoding/hex"
)

func GetNonce() string {
	h := md5.New()
	now := time.Now().UnixMicro()
	io.WriteString(h, strconv.FormatInt(now, 10))
	io.WriteString(h, strconv.FormatInt(rand.Int63(), 10))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func getHmac(payload string, secret string) string {
	sig := hmac.New(sha256.New, []byte(secret))
	sig.Write([]byte(payload))

	return hex.EncodeToString(sig.Sum(nil))
}

func authReq(method string, url string, siteKey string, sig string, nonce string, body []byte) (string, error) {
	client := &http.Client{}
	
	var err error
	var req *http.Request
	if len(body) == 0 {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(body))
	}

	if err != nil {
		return "", err
	}

	// The API request authentication is done by specifiying the hashed + signed object containing the nonce, http request type, 
	// request path, and body, alongside the siteKey/publicKey and same none present inside of the signed object.
	
	req.Header.Set("AUTHENTICATION", fmt.Sprintf("HMAC %s:%s:%s", siteKey, sig, nonce))

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), err
}