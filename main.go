package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// SupertokensCore main
type SupertokensCore struct {
	backends         []BackendConfig
	handshakeInfo    *HandshakeInfo
	deviceDriverInfo *DeviceDriverInfo
	client           *http.Client
	isInitialized    bool
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

// FrontendSDK Details about Frontend SDKs
type FrontendSDK struct {
	name    string
	version string
}

// BackendConfig Details about available supertokens-core backend instances
type BackendConfig struct {
	hostname string
	port     int
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

func (st *SupertokensCore) doRoundRobin(method string, path string, body io.Reader) (*http.Response, error) {
	for _, backend := range st.backends {
		url := fmt.Sprintf("http://%s:%d%s", backend.hostname, backend.port, path)
		req, err := http.NewRequest(method, url, body)
		resp, err := st.client.Do(req)
		if err == nil {
			return resp, nil
		}
	}
	return nil, errors.New("none of the backends are active")
}

func (st *SupertokensCore) hello() (*http.Response, error) {
	if !st.isInitialized {
		return nil, errors.New("driver has not yet been initialized")
	}

	resp, err := st.doRoundRobin("GET", "/hello", nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (st *SupertokensCore) init(backends []BackendConfig, deviceDriverInfo *DeviceDriverInfo) {
	st.client = &http.Client{}
	deviceDriverInfo.driver = &Driver{
		"supertokens-go", "0.0",
	}
	st.deviceDriverInfo = deviceDriverInfo
	st.backends = backends
	st.isInitialized = true

	// Perform Handshake
	if st.handshakeInfo == nil {
		st.handshake()
	}
}

func (st *SupertokensCore) handshake() error {
	if !st.isInitialized {
		return errors.New("driver has not yet been initialized")
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(st.deviceDriverInfo)

	resp, err := st.doRoundRobin("POST", "/handshake", buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	st.handshakeInfo = &HandshakeInfo{}
	err = json.NewDecoder(resp.Body).Decode(st.handshakeInfo)
	if err != nil {
		return err
	}

	log.Println(st.handshakeInfo)
	return nil
}

