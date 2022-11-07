package args

import (
	"github.com/bilgehannal/harbctl/internal/pkg/errors"
	"log"
	"strings"
)

const (
	orderOfVerbValue   = 0
	orderOfObjectValue = 1
)

func getVerb(args []string) (Verb, error) {
	if len(args) < 2 {
		return Verb{}, errors.VerbMissingArgumentsError{}
	}
	objectValue := args[orderOfObjectValue]
	isObjectCombined := false
	secondObjectValue := ""
	if strings.Contains(objectValue, ":") {
		isObjectCombined = true
		secondObjectValue = strings.Split(objectValue, ":")[1]
	}
	currentVerb := Verb{
		Value: args[orderOfVerbValue],
		Object: Object{
			Value: objectValue,
		},
		IsObjectCombined: isObjectCombined,
		SecondObject: Object{
			Value: secondObjectValue,
		},
	}
	return validateVerb(currentVerb)
}

func addFlagToList(flags []Flag, element string) ([]Flag, error) {
	flag := Flag{Value: element}
	flag, err := validateFlag(flag)
	if err != nil {
		return flags, err
	}
	flags = append(flags, flag)
	return flags, nil
}

func getFlags(args []string) ([]Flag, error) {
	var flags []Flag
	var err error
	for i := 2; i < len(args); i++ {
		if strings.HasPrefix(args[i], "-") {
			if i != len(args)-1 {
				if strings.HasPrefix(args[i+1], "-") {
					flags, err = addFlagToList(flags, args[i])
				}
			} else {
				flags, err = addFlagToList(flags, args[i])
			}
			if err != nil {
				return flags, err
			}
		}
	}
	return flags, nil
}

func addParameterToList(parameters []Parameter, typeOfParameter string, valueOfPatameter string) ([]Parameter, error) {
	parameter := Parameter{TypeOfParameter: typeOfParameter, ValueOfParameter: valueOfPatameter}
	parameter, err := validateParameter(parameter)
	if err != nil {
		return parameters, err
	}
	parameters = append(parameters, parameter)
	return parameters, nil
}

func getParameters(args []string) ([]Parameter, error) {
	var parameters []Parameter
	var err error
	for i := 2; i < len(args)-1; i++ {
		if strings.HasPrefix(args[i], "-") {
			if i != len(args)-1 {
				if !strings.HasPrefix(args[i+1], "-") {
					parameters, err = addParameterToList(parameters, args[i], args[i+1])
					if err != nil {
						return parameters, err
					}
				}
			}
		}
	}
	return parameters, nil
}

func GetArgs(args []string) Args {
	flags, err := getFlags(args)
	if err != nil {
		log.Fatal("Error in getting flags: ", err)
	}
	verb, err := getVerb(args)
	if err != nil {
		log.Fatal("Error in getting verb: ", err)
	}
	params, err := getParameters(args)
	if err != nil {
		log.Fatal("Error in getting parameters: ", err)
	}
	return Args{
		Verb:   verb,
		Flags:  flags,
		Params: params,
	}
}
