package vinyldns

import "testing"

func TestGetUserAgent(t *testing.T) {
	testCases := []struct {
		name       string
		version    string
		expectedUA string
	}{
		{
			name:       "empty version",
			version:    "",
			expectedUA: "terraform-provider-vinyldns",
		},
		{
			name:       "version set",
			version:    "0.9.3",
			expectedUA: "terraform-provider-vinyldns/0.9.3",
		},
		{
			name:       "version set",
			version:    "1.0.1-dev",
			expectedUA: "terraform-provider-vinyldns/1.0.1-dev",
		},
	}

	for _, testCase := range testCases {
		SetVersion(testCase.version)

		t.Run(testCase.name, func(t *testing.T) {
			useragent := GetUserAgent()
			if useragent != testCase.expectedUA {
				t.Fatalf("expected user-agent to be '%s', got '%s'", testCase.expectedUA, useragent)
			}
		})
	}
}
