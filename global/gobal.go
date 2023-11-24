package global

import (
	"fmt"
	"os"
)

var (
	Home, _   = os.UserHomeDir()
	KcsHome   = fmt.Sprintf("%s/.kcs", Home)
	KcsConfig = fmt.Sprintf("%s/config", KcsHome)
	KcsData   = fmt.Sprintf("%s/data", KcsHome)

	KubeHome = fmt.Sprintf("%s/.kube", Home)

	// 添加变量 name
	Name       string
	Kubeconfig string
	// 查看变量
	Server bool
)
