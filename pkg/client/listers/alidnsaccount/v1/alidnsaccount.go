/**/
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/xzzpig/alidns-ingress-updater/pkg/apis/alidnsaccount/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AliDnsAccountLister helps list AliDnsAccounts.
// All objects returned here must be treated as read-only.
type AliDnsAccountLister interface {
	// List lists all AliDnsAccounts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AliDnsAccount, err error)
	// Get retrieves the AliDnsAccount from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.AliDnsAccount, error)
	AliDnsAccountListerExpansion
}

// aliDnsAccountLister implements the AliDnsAccountLister interface.
type aliDnsAccountLister struct {
	indexer cache.Indexer
}

// NewAliDnsAccountLister returns a new AliDnsAccountLister.
func NewAliDnsAccountLister(indexer cache.Indexer) AliDnsAccountLister {
	return &aliDnsAccountLister{indexer: indexer}
}

// List lists all AliDnsAccounts in the indexer.
func (s *aliDnsAccountLister) List(selector labels.Selector) (ret []*v1.AliDnsAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AliDnsAccount))
	})
	return ret, err
}

// Get retrieves the AliDnsAccount from the index for a given name.
func (s *aliDnsAccountLister) Get(name string) (*v1.AliDnsAccount, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("alidnsaccount"), name)
	}
	return obj.(*v1.AliDnsAccount), nil
}
