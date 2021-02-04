package main

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
	"k8s.io/utils/env"
)

type config struct {
	kubeconfig       string
	endpoint         string
	autoDeleteRecord bool
}

var appConfig = config{
	endpoint:         getConfigString("endpoint", "(optional) aliyun dns endpoint", "ENDPOINT", "dns.aliyuncs.com"),
	autoDeleteRecord: getConfigString("autoDeleteRecord", "(optional) will auto delete dns record when ingress is deleted", "AUTO_DELETE_RECORD", "false") == "true",
}

func init() {
	if home := homedir.HomeDir(); home != "" {
		appConfig.kubeconfig = getConfigString("kubeconfig", "(optional) absolute path to the kubeconfig file", "KUBE_CONFIG", filepath.Join(home, ".kube", "config"))
	} else {
		appConfig.kubeconfig = getConfigString("kubeconfig", "absolute path to the kubeconfig file", "KUBE_CONFIG", "")
	}
}

func getConfigString(flagKey string, flagDescribe string, envKey string, defaultValue string) string {
	var value *string
	value = flag.String(flagKey, env.GetString(envKey, defaultValue), flagDescribe)
	flag.Parse()
	return *value
}
