package config

type DeployFuncInfo struct {
	FuncName   string
	ProjectDir string
	ShouldRun  bool
}

type DeploySet struct {
	Name      string
	FuncInfos []DeployFuncInfo
}

type DeployConfig struct {
	ConfigJsonLocation string `json:"-"`
	SubscriptionId     string
	ResourceGroupName  string
	Sets               []DeploySet
}
