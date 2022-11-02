package args

import (
	"github.com/bilgehannal/harbctl/internal/pkg/errors"
	"github.com/bilgehannal/harbctl/internal/pkg/utils"
)

type Args struct {
	Verb   Verb
	Flags  []Flag
	Params []Parameter
}

type Verb struct {
	Value  string
	Object Object
}

type Object struct {
	Value string
}

type Flag struct {
	Value string
}

type Parameter struct {
	TypeOfParameter  string
	ValueOfParameter string
}

func validateVerb(verb Verb) (Verb, error) {
	if !utils.StringSliceContains(GetAvailableVerbValues(), verb.Value) {
		return verb, errors.VerbUnsoppertedError{}
	}
	if !isValidObjectForVerbValue(verb) {
		return verb, errors.ObjectUnsoppertedError{}
	}
	return verb, nil
}

func isValidObjectForVerbValue(verb Verb) bool {
	if utils.StringSliceContains([]string{VerbLogin}, verb.Value) {
		if len(verb.Object.Value) == 0 {
			return false
		}
		return true
	}
	return false
}

func validateFlag(flag Flag) (Flag, error) {
	if !utils.StringSliceContains(getAvailableFlagValues(), flag.Value) {
		return flag, errors.FlagUnsoppertedError{}
	}
	return flag, nil
}

func validateParameter(parameter Parameter) (Parameter, error) {
	if !utils.StringSliceContains(getAvailableParameterTypes(), parameter.TypeOfParameter) {
		return parameter, errors.ParameterUnsoppertedTypeError{}
	}
	return parameter, nil
}
