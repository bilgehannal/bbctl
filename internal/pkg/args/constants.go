package args

const (
	VerbLogin = "login"

	ObjectProject = "project"

	FlagHelp1 = "--help"
	FlagHelp2 = "-h"
)

type ParameterType struct {
	User     []string
	Password []string
	File     []string
}

func GetParameterTypes() ParameterType {
	return ParameterType{
		User:     []string{"-u", "--user"},
		Password: []string{"-p", "--password"},
		File:     []string{"--file", "-f"},
	}
}

func GetAvailableVerbValues() []string {
	return []string{"login", "get"}
}

func getAvailableObjectValues() []string {
	return []string{"projects"}
}

func getAvailableFlagValues() []string {
	return []string{
		FlagHelp1, FlagHelp2,
	}
}

func getAvailableParameterTypes() []string {
	return []string{
		"-u", "--user",
		"-p", "--password",
	}
}
