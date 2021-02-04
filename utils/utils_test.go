package utils_test

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/xzzpig/alidns-ingress-updater/utils"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func TestListAliDNSAccounts(t *testing.T) {
	home := homedir.HomeDir()
	config := filepath.Join(home, ".kube", "config")
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", config)
	if err != nil {
		panic(err.Error())
	}
	accounts, err := utils.ListAliDNSAccounts(kubeConfig)
	if err != nil {
		panic(err)
	}
	for _, account := range accounts.Items {
		t.Log(account.Name)
	}
}

func TestDnsUtils(t *testing.T) {
	home := homedir.HomeDir()
	config := filepath.Join(home, ".kube", "config")
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", config)
	if err != nil {
		panic(err.Error())
	}
	accounts, err := utils.ListAliDNSAccounts(kubeConfig)
	if err != nil {
		panic(err)
	}
	for _, account := range accounts.Items {
		if account.Name != "xzzpig-alidns-account" {
			continue
		}
		dnsUtils, err := utils.NewAliDnsUtils(account.Spec)
		if err != nil {
			t.Error(err)
		}
		t.Log(dnsUtils.GetRecordCount())
		recordId, err := dnsUtils.CreateRecord("test", "1.1.1.1", "A")
		t.Log("CreateRecord", recordId)
		err = dnsUtils.UpdateRecord(recordId, "test", "1.1.1.2", "A")
		if err != nil {
			t.Error(err)
		}
		t.Log("UpdateRecord", err)
		err = dnsUtils.DeleteRecord(recordId)
		if err != nil {
			t.Error(err)
		}
		t.Log("DeleteRecord", err)
	}
}

func TestUpdateRecord(t *testing.T) {
	home := homedir.HomeDir()
	config := filepath.Join(home, ".kube", "config")
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", config)
	if err != nil {
		panic(err.Error())
	}
	accounts, err := utils.ListAliDNSAccounts(kubeConfig)
	if err != nil {
		panic(err)
	}
	for _, account := range accounts.Items {
		if account.Name != "xzzpig-alidns-account" {
			continue
		}
		dnsUtils, err := utils.NewAliDnsUtils(account.Spec)
		if err != nil {
			t.Error(err)
		}
		err = dnsUtils.UpdateRecord("11111", "test", "1.1.1.2", "A")
		if err != nil {
			fmt.Println(strings.Contains(err.Error(), "DomainRecordNotBelongToUser"))
		}
		t.Log("UpdateRecord", err)
	}
}

func TestListIngresses(t *testing.T) {
	home := homedir.HomeDir()
	config := filepath.Join(home, ".kube", "config")
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", config)
	if err != nil {
		panic(err)
	}
	ingresses, err := utils.ListIngresses(kubeConfig)
	if err != nil {
		panic(err)
	}
	for _, ingress := range ingresses.Items {
		t.Log(ingress.Name)
	}
}
