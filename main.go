package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// SupertokensCore main
type SupertokensCore struct {
	backends         *[]backendConfig
	handshakeInfo    *HandshakeInfo
	deviceDriverInfo *DeviceDriverInfo
}

// DeviceDriverInfo info about device and driver
type DeviceDriverInfo struct {
	frontendSDKList *[]frontendSDK
	driver          *driver
}

type driver struct {
	name    string
	version string
}
type frontendSDK struct {
	name    string
	version string
}

type backendConfig struct {
	hostname string
	port     int
}

// Querier allows transparently calling multiple backends
type Querier struct {
}

// HandshakeInfo singleton
type HandshakeInfo struct {
	jwtSigningPublicKey            string
	cookieDomain                   string
	cookieSecure                   bool
	accessTokenPath                string
	refreshTokenPath               string
	enableAntiCsrf                 bool
	accessTokenBlacklistingEnabled bool
	jwtSigningPublicKeyExpiryTime  int
}

func main() {
	stCore := &SupertokensCore{}
	stCore.backends = &[]backendConfig{
		backendConfig{
			"localhost", 3567,
		},
		backendConfig{
			"localhost", 3568,
		},
	}
	stCore.deviceDriverInfo = &DeviceDriverInfo{
		frontendSDKList: &[]frontendSDK{
			frontendSDK{
				"vuejs", "1.1",
			},
			frontendSDK{
				"react", "1.0",
			},
		},
		driver: &driver{
			"supertokens-go", "0.0",
		},
	}
	stCore.init()
}

func (st *SupertokensCore) hello() *http.Response {
	resp, err := http.Get("http://localhost:3567/hello")
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

func (st *SupertokensCore) init() {
	for _, backend := range *st.backends {
		log.Println(backend)
	}
	// Perform Handshake
	if st.handshakeInfo == nil {
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(st.deviceDriverInfo)
		resp, err := http.Post("http://localhost:3567/handshake", "application/json", buf)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(body))
	}
}

func handshake() {
}
