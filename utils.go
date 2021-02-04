package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/xzzpig/alidns-ingress-updater/utils"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

var kubeConfig *rest.Config

func init() {
	cfg, err := clientcmd.BuildConfigFromFlags("", appConfig.kubeconfig)
	if err != nil {
		cfg, err = rest.InClusterConfig()
		if err != nil {
			panic(err)
		}
	}
	kubeConfig = cfg
}

func GetDnsUtils(host string) (util *utils.AliDnsUtils, rr string, err error) {
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

func SetRecord(host string, ip string) error {
	dnsUtils, rr, err := GetDnsUtils(host)
	if err != nil {
		return err
	}
	record, err := dnsUtils.FindRecordByRR(rr)
	if err != nil {
		return err
	}
	if record == nil {
		_, err = dnsUtils.CreateRecord(rr, ip, "A")
		klog.Info("Record ", host, " created as ", ip)
		return err

	}
	if *record.Type == "A" && *record.Value == ip {
		return nil
	}
	klog.Info("Record ", host, " updated from ", *record.Value, " to ", ip)
	return dnsUtils.UpdateRecord(*record.RecordId, rr, ip, "A")

}

func DelRecord(host string) error {
	dnsUtils, rr, err := GetDnsUtils(host)
	if err != nil {
		return err
	}
	record, err := dnsUtils.FindRecordByRR(rr)
	if err != nil {
		return err
	}
	if record == nil {
		klog.Info("Record ", host, " deleted")
		return dnsUtils.DeleteRecord(*record.RecordId)
	}
	return err
}

func UpdateAllRecord(ip string) error {
	ingresses, err := utils.ListIngresses(kubeConfig)
	if err != nil {
		return err
	}
	for _, ingress := range ingresses.Items {
		if ingress.ObjectMeta.Annotations["xzzpig.com/alidns-ignore"] == "true" {
			continue
		}
		if ingress.Spec.Rules == nil {
			continue
		}
		for _, rule := range ingress.Spec.Rules {
			host := rule.Host
			if host == "" {
				continue
			}
			err = SetRecord(host, ip)
			if err.Error() == "No Account Found" {
				//TODO
				continue
			} else {
				fmt.Println(err) //TODO
			}
		}
	}
	return nil
}
