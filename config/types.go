package config

const (
	DeployMethodFunc   string = "func"
	DeployMethodAzFunc string = "azfunc"
	DeployMethodAzZip  string = "azzip"
)

type DeployMethod = string

type DeployFuncInfo struct {
	FuncName   string
	ProjectDir string
	ShouldRun  bool
}

type DeploySet struct {
	Name              string
	ResourceGroupName string
	FuncInfos         []DeployFuncInfo
}

var _deployMethods []DeployMethod = []DeployMethod{DeployMethodAzFunc, DeployMethodAzZip, DeployMethodFunc}

type DeployConfig struct {
	ConfigJsonLocation string `json:"-"`
	Method             DeployMethod
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
