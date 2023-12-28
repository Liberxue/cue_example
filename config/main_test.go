package main

import (
	"reflect"
	"testing"
)

func TestCueToStruct(t *testing.T) {
	expected := []*Config{
		{
			ServerName:    "liber_test",
			ServerTimeout: "30s",
			Endpoints:     []string{"0.0.0.0:9090", "127.0.0.1:5672"},
		},
		{
			ServerName:    "liber_test_server",
			ServerTimeout: "30s",
			Endpoints:     []string{"0.0.0.0:9091", "127.0.0.1:5672"},
		},
	}

	result, err := cueToStruct()
	if err != nil {
		t.Fatalf("cueToStruct returned an error: %v", err)
	}

	if len(expected) != len(result) {
		t.Fatalf("Expected %d results, got %d", len(expected), len(result))
	}

	for i, config := range expected {
		if !compareConfig(*config, *result[i]) {
			t.Errorf("Expected %+v, got %+v at index %d", *config, *result[i], i)
		}
	}
}

func compareConfig(a, b Config) bool {
	if a.ServerName != b.ServerName || a.ServerTimeout != b.ServerTimeout {
		return false
	}
	return reflect.DeepEqual(a.Endpoints, b.Endpoints)
}
