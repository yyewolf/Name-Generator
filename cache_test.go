package namegenerator

import (
	"testing"
)

//////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////// START OF TESTS

// TestNames tests the name cache. This should not return an error.
func TestNames(t *testing.T) {
	session := NewSession()
	for i := 0; i < 25; i++ {
		_, err := session.GetName()
		if err != nil {
			t.Errorf("Problem requesting names: %+v", err)
		}
	}
}

// TestCityNames tests the city name cache. This should not return an error.
func TestCityNames(t *testing.T) {
	session := NewSession()
	for i := 0; i < 25; i++ {
		_, err := session.GetCityName()
		if err != nil {
			t.Errorf("Problem requesting names: %+v", err)
		}
	}
}
