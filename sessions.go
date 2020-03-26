package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// SessionResponse session response containing tokens
type SessionResponse struct {
	Session                       *stSession
	AccessToken                   *StToken
	RefreshToken                  *StToken
	IDRefreshToken                *StToken `json:"idRefreshToken"`
	AntiCsrfToken                 string
	JwtSigningPublicKey           string
	JwtSigningPublicKeyExpiryTime int64
}

type stSession struct {
	Handle        string
	UserID        string `json:"userId"`
	UserDataInJWT map[string]interface{}
}

// StToken Token Object
type StToken struct {
	Token        string
	Expiry       int
	CreatedTime  int
	CookiePath   string
	CookieSecure bool
	Domain       string
}

func (st *SupertokensCore) createSession(userID string, jwtPayload *map[string]interface{}, sessionData *map[string]interface{}) (*SessionResponse, error) {
	if !st.isInitialized {
		return nil, errors.New("driver has not yet been initialized")
	}

	sessionInput := &struct {
		UserID      string                  `json:"userId"`
		JwtPayload  *map[string]interface{} `json:"userDataInJWT"`
		SessionData *map[string]interface{} `json:"userDataInDatabase"`
	}{
		userID,
		jwtPayload,
		sessionData,
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(sessionInput)
	resp, err := st.doRoundRobin("POST", "/session", buf)
	fmt.Println()
	if err != nil {
		return nil, err
	}

	sessionResponse := &SessionResponse{}
	err = json.NewDecoder(resp.Body).Decode(sessionResponse)
	if err != nil {
		return nil, err
	}

	st.handshakeInfo.JwtSigningPublicKey = sessionResponse.JwtSigningPublicKey
	st.handshakeInfo.JwtSigningPublicKeyExpiryTime = sessionResponse.JwtSigningPublicKeyExpiryTime

	return sessionResponse, nil
}
