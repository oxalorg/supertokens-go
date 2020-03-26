package main

import "testing"

func TestGetSession(t *testing.T) {
	stCore := &SupertokensCore{}
	backends := []BackendConfig{
		BackendConfig{
			"localhost", 3567,
		},
		BackendConfig{
			"localhost", 3568,
		},
	}
	deviceDriverInfo := &DeviceDriverInfo{
		frontendSDKList: []FrontendSDK{
			FrontendSDK{
				"vuejs", "1.1",
			},
		},
	}
	stCore.init(backends, deviceDriverInfo)

	jwtPayload := &map[string]interface{}{
		"userId": "User1",
		"name":   "spooky action at a distance",
	}

	sessionData := &map[string]interface{}{
		"awesomeThings": []string{
			"ox", "oxalorg", "programming", "supertokens",
		},
	}
	_, err := stCore.createSession("User1", jwtPayload, sessionData)
	if err != nil {
		t.Errorf("createSession() failed with errors %v", err)
	}
}
