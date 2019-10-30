/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DataLakeStoreFirewallRuleLister helps list DataLakeStoreFirewallRules.
type DataLakeStoreFirewallRuleLister interface {
	// List lists all DataLakeStoreFirewallRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DataLakeStoreFirewallRule, err error)
	// DataLakeStoreFirewallRules returns an object that can list and get DataLakeStoreFirewallRules.
	DataLakeStoreFirewallRules(namespace string) DataLakeStoreFirewallRuleNamespaceLister
	DataLakeStoreFirewallRuleListerExpansion
}

// dataLakeStoreFirewallRuleLister implements the DataLakeStoreFirewallRuleLister interface.
type dataLakeStoreFirewallRuleLister struct {
	indexer cache.Indexer
}

// NewDataLakeStoreFirewallRuleLister returns a new DataLakeStoreFirewallRuleLister.
func NewDataLakeStoreFirewallRuleLister(indexer cache.Indexer) DataLakeStoreFirewallRuleLister {
	return &dataLakeStoreFirewallRuleLister{indexer: indexer}
}

// List lists all DataLakeStoreFirewallRules in the indexer.
func (s *dataLakeStoreFirewallRuleLister) List(selector labels.Selector) (ret []*v1alpha1.DataLakeStoreFirewallRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataLakeStoreFirewallRule))
	})
	return ret, err
}

// DataLakeStoreFirewallRules returns an object that can list and get DataLakeStoreFirewallRules.
func (s *dataLakeStoreFirewallRuleLister) DataLakeStoreFirewallRules(namespace string) DataLakeStoreFirewallRuleNamespaceLister {
	return dataLakeStoreFirewallRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataLakeStoreFirewallRuleNamespaceLister helps list and get DataLakeStoreFirewallRules.
type DataLakeStoreFirewallRuleNamespaceLister interface {
	// List lists all DataLakeStoreFirewallRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DataLakeStoreFirewallRule, err error)
	// Get retrieves the DataLakeStoreFirewallRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DataLakeStoreFirewallRule, error)
	DataLakeStoreFirewallRuleNamespaceListerExpansion
}

// dataLakeStoreFirewallRuleNamespaceLister implements the DataLakeStoreFirewallRuleNamespaceLister
// interface.
type dataLakeStoreFirewallRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataLakeStoreFirewallRules in the indexer for a given namespace.
func (s dataLakeStoreFirewallRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DataLakeStoreFirewallRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataLakeStoreFirewallRule))
	})
	return ret, err
}

// Get retrieves the DataLakeStoreFirewallRule from the indexer for a given namespace and name.
func (s dataLakeStoreFirewallRuleNamespaceLister) Get(name string) (*v1alpha1.DataLakeStoreFirewallRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("datalakestorefirewallrule"), name)
	}
	return obj.(*v1alpha1.DataLakeStoreFirewallRule), nil
}
