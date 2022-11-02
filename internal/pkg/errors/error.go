package errors

type VerbMissingArgumentsError struct{}

func (e VerbMissingArgumentsError) Error() string {
	return `There are some missing arguments, There should be a verb and object
		ex: harbctl get project
		harbctl [verb] [object]`
}

type VerbUnsoppertedError struct{}

func (e VerbUnsoppertedError) Error() string {
	return `The verb is not unsopperted. You can check the valid verbs using command below:
		$ harbctl --help`
}

type ObjectUnsoppertedError struct{}

func (e ObjectUnsoppertedError) Error() string {
	return `The object is not unsopperted. You can check the valid verbs using command below:
		$ harbctl --help`
}

type FlagUnsoppertedError struct{}

func (e FlagUnsoppertedError) Error() string {
	return `The flag is not unsopperted. You can check the valid verbs using command below:
		$ harbctl --help`
}

type ParameterUnsoppertedTypeError struct{}

func (e ParameterUnsoppertedTypeError) Error() string {
	return `The parameter type is not unsopperted. You can check the valid verbs using command below:
		$ harbctl --help`
}
