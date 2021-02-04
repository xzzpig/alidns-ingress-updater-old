package utils

import (
	"context"

	v1 "github.com/xzzpig/alidns-ingress-updater/pkg/apis/alidnsaccount/v1"
	accountClientSet "github.com/xzzpig/alidns-ingress-updater/pkg/client/clientset/versioned"
	netv1 "k8s.io/api/networking/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListAliDNSAccounts(kubeConfig *rest.Config) (*v1.AliDnsAccountList, error) {
	accountClient, err := accountClientSet.NewForConfig(kubeConfig)
	if err != nil {
		return nil, err
	}
	return accountClient.XzzpigV1().AliDnsAccounts().List(context.TODO(), metav1.ListOptions{})
}

func ListIngresses(kubeConfig *rest.Config) (*netv1.IngressList, error) {
	coreClient, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, err
	}
	return coreClient.NetworkingV1().Ingresses("").List(context.TODO(), metav1.ListOptions{})
}
