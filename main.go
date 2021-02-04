package main

import (
	"context"
	"time"

	"github.com/xzzpig/alidns-ingress-updater/utils"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

func main() {
	klog.Info("AliDns Ingress Updater Start")
	coreClient, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		panic(err.Error())
	}
	ingressWatch, err := coreClient.NetworkingV1().Ingresses("").Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
	}
	go func() {
		events := ingressWatch.ResultChan()
		klog.Info("Start Watch Ingress")
		for event := range events {
			ingress := event.Object.(*netv1.Ingress)
			if ingress == nil {
				continue
			}
			if ingress.ObjectMeta.Annotations["xzzpig.com/alidns-ignore"] == "true" {
				continue
			}
			ip, err := utils.GetPublicIP()
			if err != nil {
				klog.Error(err)
				klog.Warning("Get Public IP Failed")
				continue
			}
			if ingress.Spec.Rules == nil {
				continue
			}
			for _, rule := range ingress.Spec.Rules {
				switch event.Type {
				case "ADDED":
					fallthrough
				case "MODIFIED":
					host := rule.Host
					if host == "" {
						continue
					}
					err = SetRecord(host, ip)
					switch {
					case err == nil:
					case err.Error() == "No Account Found":
						klog.Warning("No Account Found For Host ", host)
						continue
					default:
						klog.Error(err)
						klog.Warning("Update DNS For Host ", host, " Failed")
					}
				case "DELETED":
					if appConfig.autoDeleteRecord {
						host := rule.Host
						if host == "" {
							continue
						}
						err = DelRecord(host)
						if err.Error() == "No Account Found" {
							klog.Warning("No Account Found For Host ", host)
							continue
						} else {
							klog.Warning("Delete DNS For Host ", host, " Failed")
						}
					}
				}
			}
		}
		klog.Exit("Ingress watch is stoped")
	}()
	go func() {
		ip, err := utils.GetPublicIP()
		if err != nil {
			klog.Exit(err)
		}
		ticker := time.NewTicker(time.Minute)
		for range ticker.C {
			newIP, err := utils.GetPublicIP()
			if err != nil {
				klog.Error(err)
				klog.Warning("Get Public IP Failed")
				continue
			}
			if ip != newIP {
				ip = newIP
				UpdateAllRecord(ip)
			}
		}
	}()

	<-make(chan bool)
}
