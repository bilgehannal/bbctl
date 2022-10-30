package args

import "testing"

func Test_GetAvailableVerbValues(t *testing.T) {
	testValues := []string{"login"}
	values := GetAvailableVerbValues()

	if len(testValues) != len(values) {
		t.Errorf("length is not matched!")
	}
	for i := 0; i < len(values); i++ {
		if values[i] != testValues[i] {
			t.Errorf("elements are not matched!")
		}
	}
}

func Test_getAvailableObjectValues(t *testing.T) {
	testValues := []string{"project"}
	values := getAvailableObjectValues()

	if len(testValues) != len(values) {
		t.Errorf("length is not matched!")
	}
	for i := 0; i < len(values); i++ {
		if values[i] != testValues[i] {
			t.Errorf("elements are not matched!")
		}
	}
}

func Test_getAvailableFlagValues(t *testing.T) {
	testValues := []string{"--help", "-h"}
	values := getAvailableFlagValues()

	if len(testValues) != len(values) {
		t.Errorf("length is not matched!")
	}
	for i := 0; i < len(values); i++ {
		if values[i] != testValues[i] {
			t.Errorf("elements are not matched!")
		}
	}
}

func Test_getAvailableParameterTypes(t *testing.T) {
	testValues := []string{"-u", "--user", "-p", "--password"}
	values := getAvailableParameterTypes()

	if len(testValues) != len(values) {
		t.Errorf("length is not matched!")
	}
	for i := 0; i < len(values); i++ {
		if values[i] != testValues[i] {
			t.Errorf("elements are not matched!")
		}
	}
}
