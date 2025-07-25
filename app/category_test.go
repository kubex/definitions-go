package app

import "testing"

func TestCategory_Name(t *testing.T) {
	if CategoryHR.Name() != "Human Resources" {
		t.Errorf("Expected 'Human Resources', got '%s'", CategoryHR.Name())
	}

	if CategoryMarketing.Name() != "Marketing" {
		t.Errorf("Expected 'Marketing', got '%s'", CategoryMarketing.Name())
	}
}
