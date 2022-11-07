package args

import (
	"github.com/bilgehannal/harbctl/internal/pkg/errors"
	"github.com/bilgehannal/harbctl/internal/pkg/utils"
	"strings"
)

type Args struct {
	Verb   Verb
	Flags  []Flag
	Params []Parameter
}

type Verb struct {
	Value            string
	Object           Object
	IsObjectCombined bool
	SecondObject     Object
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

func (args Args) GetParameter(parameterTypes []string) (Parameter, error) {
	for _, param := range args.Params {
		if utils.StringSliceContains(parameterTypes, param.TypeOfParameter) {
			return param, nil
		}
	}
	return Parameter{}, errors.ParameterNotFoundError{}
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
	objectValue := verb.Object.Value
	if verb.IsObjectCombined {
		objectValue = strings.Split(objectValue, ":")[0]
	}
	if utils.StringSliceContains([]string{VerbLogin}, verb.Value) {
		if len(verb.Object.Value) == 0 {
			return false
		}
		return true
	} else if utils.StringSliceContains(getAvailableObjectValues(), objectValue) {
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
