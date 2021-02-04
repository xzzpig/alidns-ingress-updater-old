module github.com/xzzpig/alidns-ingress-updater

go 1.15

replace k8s.io/client-go => k8s.io/client-go v0.20.0

require (
	github.com/alibabacloud-go/alidns-20150109 v1.0.1
	github.com/alibabacloud-go/tea v1.1.15
	github.com/alibabacloud-go/tea-rpc v1.1.8
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v1.5.1
	k8s.io/klog/v2 v2.5.0
	k8s.io/utils v0.0.0-20210111153108-fddb29f9d009
)
