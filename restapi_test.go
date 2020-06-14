package namegenerator

import (
	"testing"
)

//////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////// START OF TESTS

// TestAPI tests the API function. This should not return an error.
func TestAPI(t *testing.T) {
	resp, err := request(EndpointLastNames())
	if err != nil {
		t.Errorf("API returned error: %+v", err)
	}
	t.Log(string(resp))
}
