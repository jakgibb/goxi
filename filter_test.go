package goxi

import (
	"fmt"
	"testing"
)

func TestToParam(t *testing.T) {
	mock := HostFilter{
		Name:    []string{"server-a", "server-b"},
		Address: []string{"127.0.0.1"},
		Records: "10",
	}

	result := toParam(mock)
	expected := "&name=in:server-a,server-b&address=in:127.0.0.1&records=10"

	if result != expected {
		t.Error(fmt.Sprintf("Expected the result to be %s but instead got %s", result, expected))
	}

}
