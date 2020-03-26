package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// SupertokensCore main
type SupertokensCore struct {
	backends         []BackendConfig
	handshakeInfo    *HandshakeInfo
	deviceDriverInfo *DeviceDriverInfo
}

// DeviceDriverInfo info about device and driver
type DeviceDriverInfo struct {
	frontendSDKList []FrontendSDK
	driver          *Driver
}

// Driver details about Driver
type Driver struct {
	name    string
	version string
}

// Details about Frontend SDKs
type FrontendSDK struct {
	name    string
	version string
}

// BackendConfig Details about available supertokens-core backend instances
type BackendConfig struct {
	hostname string
	port     int
}

// Querier allows transparently calling multiple backends
type Querier struct {
}

// HandshakeInfo singleton
type HandshakeInfo struct {
	JwtSigningPublicKey            string
	CookieDomain                   string
	CookieSecure                   bool
	AccessTokenPath                string
	RefreshTokenPath               string
	EnableAntiCsrf                 bool
	AccessTokenBlacklistingEnabled bool
	JwtSigningPublicKeyExpiryTime  int64
}

func main() {
}

func (st *SupertokensCore) hello() *http.Response {
	// TODO: Round Robin
	resp, err := http.Get("http://localhost:3567/hello")
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

func (st *SupertokensCore) init(backends []BackendConfig, deviceDriverInfo *DeviceDriverInfo) {
	for _, backend := range st.backends {
		log.Println(backend)
	}

	deviceDriverInfo.driver = &Driver{
		"supertokens-go", "0.0",
	}
	st.deviceDriverInfo = deviceDriverInfo
	st.backends = backends

	// Perform Handshake
	if st.handshakeInfo == nil {
		st.handshake()
	}
}

func (st *SupertokensCore) handshake() {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(st.deviceDriverInfo)

	// TODO: Round Robin
	resp, err := http.Post("http://localhost:3567/handshake", "application/json", buf)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	st.handshakeInfo = &HandshakeInfo{}
	err = json.NewDecoder(resp.Body).Decode(st.handshakeInfo)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(st.handshakeInfo)
}
