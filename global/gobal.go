package global

import (
	"fmt"
	"os"
)

var (
	Home, _  = os.UserHomeDir()
	KvmHome = fmt.Sprintf("%s/.kvm",Home)
	KubeHome = fmt.Sprintf("%s/.kube",Home)

	// 添加变量 name
	Name       string
	Kubeconfig string
)
