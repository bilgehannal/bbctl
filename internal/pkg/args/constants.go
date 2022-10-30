package args

const (
	VerbLogin = "login"

	ParameterUser1     = "-u"
	ParameterUser2     = "--user"
	ParameterPassword1 = "-p"
	ParameterPassword2 = "--password"

	FlagHelp1 = "--help"
	FlagHelp2 = "-h"
)

func GetAvailableVerbValues() []string {
	return []string{
		VerbLogin,
	}
}

func getAvailableObjectValues() []string {
	return []string{}
}

func getAvailableFlagValues() []string {
	return []string{
		FlagHelp1, FlagHelp2,
	}
}

func getAvailableParameterTypes() []string {
	return []string{
		ParameterUser1, ParameterUser2,
		ParameterPassword1, ParameterPassword2,
	}
}
