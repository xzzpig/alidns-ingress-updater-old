package utils_test

import (
	"errors"
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

func TestGetDnsUtils(t *testing.T) {
	fmt.Println(GetDnsUtils("kapp.fae2ly.com"))
}

func GetDnsUtils(host string) (util *utils.AliDnsUtils, rr string, err error) {
	home := homedir.HomeDir()
	config := filepath.Join(home, ".kube", "config")
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", config)
	accounts, err := utils.ListAliDNSAccounts(kubeConfig)
	if err != nil {
		return nil, "", err
	}
	for _, account := range accounts.Items {
		domainName := "." + account.Spec.DomainName
		if strings.HasSuffix(host, domainName) {
			util, err := utils.NewAliDnsUtils(account.Spec)
			if err != nil {
				return nil, "", err
			}
			return util, strings.ReplaceAll(host, domainName, ""), nil
		}
	}
	return nil, "", errors.New("No Account Found")
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
		if account.Name != "fae2ly-alidns-account" {
			continue
		}
		dnsUtils, err := utils.NewAliDnsUtils(account.Spec)
		if err != nil {
			t.Error(err)
		}
		t.Log(dnsUtils.GetRecordCount())
		records, err := dnsUtils.ListRecords()
		if err != nil {
			t.Error(err)
		}
		for _, record := range records {
			fmt.Println(*record.RR)
			if strings.HasPrefix(*record.RR, "_dnsauth") {
				fmt.Println("Delete ", *record.RR)
				err = dnsUtils.DeleteRecord(*record.RecordId)
				if err != nil {
					t.Error(err)
				}
			}
		}

		// recordId, err := dnsUtils.CreateRecord("monitor-lan", "192.168.0.141", "A")
		// t.Log("CreateRecord", recordId)
		// recordId, err = dnsUtils.CreateRecord("code-lan", "192.168.0.100", "A")
		// t.Log("CreateRecord", recordId)
		// recordId, err = dnsUtils.CreateRecord("nas-lan", "192.168.0.166", "A")
		// t.Log("CreateRecord", recordId)

		// err = dnsUtils.UpdateRecord(recordId, "test", "1.1.1.2", "A")
		// if err != nil {
		// 	t.Error(err)
		// }
		// t.Log("UpdateRecord", err)
		// err = dnsUtils.DeleteRecord(recordId)
		// if err != nil {
		// 	t.Error(err)
		// }
		// t.Log("DeleteRecord", err)
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
		for _, rule := range ingress.Spec.Rules {
			t.Log(rule.Host)
		}
	}
}

func TestGetIp(t *testing.T) {
	fmt.Println(utils.GetPublicIP())
}
