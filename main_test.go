package main

import "testing"

func TestHello(t *testing.T) {
	stCore := &SupertokensCore{}
	backends := []BackendConfig{
		BackendConfig{
			"localhost", 3568,
		},
		BackendConfig{
			"localhost", 3567,
		},
	}
	deviceDriverInfo := &DeviceDriverInfo{
		frontendSDKList: []FrontendSDK{
			FrontendSDK{
				"vuejs", "1.1",
			},
			FrontendSDK{
				"react", "1.0",
			},
		},
	}
	stCore.init(backends, deviceDriverInfo)
	resp := stCore.hello()
	got := resp.StatusCode
	want := 200
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestHandshake(t *testing.T) {
	stCore := &SupertokensCore{}
	backends := []BackendConfig{
		BackendConfig{
			"localhost", 3568,
		},
		BackendConfig{
			"localhost", 3567,
		},
	}
	deviceDriverInfo := &DeviceDriverInfo{
		frontendSDKList: []FrontendSDK{
			FrontendSDK{
				"vuejs", "1.1",
			},
			FrontendSDK{
				"react", "1.0",
			},
		},
	}
	stCore.init(backends, deviceDriverInfo)
	resp := stCore.handshake()
	got := resp.StatusCode
	want := 200
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestInit(t *testing.T) {
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
			FrontendSDK{
				"react", "1.0",
			},
		},
	}
	stCore.init(backends, deviceDriverInfo)
	got := stCore.backends[0].port
	want := 3567
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
