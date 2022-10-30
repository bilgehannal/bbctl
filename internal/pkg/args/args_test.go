package args

import (
	"testing"
)

func Test_validateVerb(t *testing.T) {
	testDataSuccess := []Verb{
		Verb{Value: "login", Object: Object{Value: "bitbucket.com"}},
		Verb{Value: "login", Object: Object{Value: "bitbucket"}},
		Verb{Value: "login", Object: Object{Value: "project"}},
	}
	testDataErr := []Verb{
		Verb{Value: "loginn", Object: Object{Value: "bitbucket.com"}},
		Verb{Value: "loginn", Object: Object{}},
		Verb{Value: "login", Object: Object{}},
		Verb{Object: Object{Value: "bitbucket.com"}},
		Verb{},
	}
	for _, verb := range testDataSuccess {
		_, err := validateVerb(verb)
		if err != nil {
			t.Error(verb)
		}
	}
	for _, verb := range testDataErr {
		_, err := validateVerb(verb)
		if err == nil {
			t.Error(verb)
		}
	}
}

func Test_isValidObjectForVerbValue(t *testing.T) {
	testDataSuccess := []Verb{
		Verb{Value: "login", Object: Object{Value: "bitbucket.com"}},
		Verb{Value: "login", Object: Object{Value: "bitbucket"}},
	}
	testDataErr := []Verb{
		Verb{Value: "abcd", Object: Object{Value: "bitbucket.com"}},
		Verb{Value: "login", Object: Object{}},
		Verb{Object: Object{Value: "bitbucket.com"}},
		Verb{},
	}
	for _, verb := range testDataSuccess {
		check := isValidObjectForVerbValue(verb)
		if !check {
			t.Error(verb)
		}
	}
	for _, verb := range testDataErr {
		check := isValidObjectForVerbValue(verb)
		if check {
			t.Error(verb)
		}
	}
}

func Test_validateFlag(t *testing.T) {
	testDataSuccess := []Flag{
		Flag{Value: "--help"},
		Flag{Value: "-h"},
	}
	testDataErr := []Flag{
		Flag{Value: "--asd"},
		Flag{Value: "-sdfsdf"},
		Flag{Value: "sad"},
		Flag{Value: ""},
		Flag{},
	}
	for _, flag := range testDataSuccess {
		_, err := validateFlag(flag)
		if err != nil {
			t.Error(flag)
		}
	}
	for _, flag := range testDataErr {
		_, err := validateFlag(flag)
		if err == nil {
			t.Error(flag)
		}
	}
}

func Test_validateParameter(t *testing.T) {
	testDataSuccess := []Parameter{
		Parameter{TypeOfParameter: "-u", ValueOfParameter: "test"},
		Parameter{TypeOfParameter: "--user", ValueOfParameter: "test"},
		Parameter{TypeOfParameter: "-p", ValueOfParameter: "test"},
		Parameter{TypeOfParameter: "--password", ValueOfParameter: "test"},
	}
	testDataErr := []Parameter{
		Parameter{TypeOfParameter: "-failedParameter", ValueOfParameter: "test"},
		Parameter{TypeOfParameter: "-failedParameter"},
		Parameter{ValueOfParameter: "test"},
		Parameter{},
	}
	for _, parameter := range testDataSuccess {
		_, err := validateParameter(parameter)
		if err != nil {
			t.Error(parameter)
		}
	}
	for _, parameter := range testDataErr {
		_, err := validateParameter(parameter)
		if err == nil {
			t.Error(parameter)
		}
	}
}
