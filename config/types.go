package config

type DeployMethod = string

type DeployFuncInfo struct {
	FuncName   string
	ProjectDir string
	ShouldRun  bool
}

type DeploySet struct {
	Name      string
	FuncInfos []DeployFuncInfo
}

var _deployMethods []DeployMethod = []DeployMethod{__deployMethodZip, __deployMethodFunc}

type DeployConfig struct {
	ConfigJsonLocation string `json:"-"`
	Method             DeployMethod
	SubscriptionId     string
	ResourceGroupName  string
	Sets               []DeploySet
	CurrentSet         int
}

func GetDeploymentMethods() []DeployMethod {
	return _deployMethods
}

func GetDeploymentMethodIndex(m string) int {
	for i, method := range _deployMethods {
		if method == m {
			return i
		}
	}
	return -1
}
