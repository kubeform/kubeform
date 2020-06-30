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

// MariadbVirtualNetworkRuleLister helps list MariadbVirtualNetworkRules.
type MariadbVirtualNetworkRuleLister interface {
	// List lists all MariadbVirtualNetworkRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MariadbVirtualNetworkRule, err error)
	// MariadbVirtualNetworkRules returns an object that can list and get MariadbVirtualNetworkRules.
	MariadbVirtualNetworkRules(namespace string) MariadbVirtualNetworkRuleNamespaceLister
	MariadbVirtualNetworkRuleListerExpansion
}

// mariadbVirtualNetworkRuleLister implements the MariadbVirtualNetworkRuleLister interface.
type mariadbVirtualNetworkRuleLister struct {
	indexer cache.Indexer
}

// NewMariadbVirtualNetworkRuleLister returns a new MariadbVirtualNetworkRuleLister.
func NewMariadbVirtualNetworkRuleLister(indexer cache.Indexer) MariadbVirtualNetworkRuleLister {
	return &mariadbVirtualNetworkRuleLister{indexer: indexer}
}

// List lists all MariadbVirtualNetworkRules in the indexer.
func (s *mariadbVirtualNetworkRuleLister) List(selector labels.Selector) (ret []*v1alpha1.MariadbVirtualNetworkRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MariadbVirtualNetworkRule))
	})
	return ret, err
}

// MariadbVirtualNetworkRules returns an object that can list and get MariadbVirtualNetworkRules.
func (s *mariadbVirtualNetworkRuleLister) MariadbVirtualNetworkRules(namespace string) MariadbVirtualNetworkRuleNamespaceLister {
	return mariadbVirtualNetworkRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MariadbVirtualNetworkRuleNamespaceLister helps list and get MariadbVirtualNetworkRules.
type MariadbVirtualNetworkRuleNamespaceLister interface {
	// List lists all MariadbVirtualNetworkRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MariadbVirtualNetworkRule, err error)
	// Get retrieves the MariadbVirtualNetworkRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MariadbVirtualNetworkRule, error)
	MariadbVirtualNetworkRuleNamespaceListerExpansion
}

// mariadbVirtualNetworkRuleNamespaceLister implements the MariadbVirtualNetworkRuleNamespaceLister
// interface.
type mariadbVirtualNetworkRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MariadbVirtualNetworkRules in the indexer for a given namespace.
func (s mariadbVirtualNetworkRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MariadbVirtualNetworkRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MariadbVirtualNetworkRule))
	})
	return ret, err
}

// Get retrieves the MariadbVirtualNetworkRule from the indexer for a given namespace and name.
func (s mariadbVirtualNetworkRuleNamespaceLister) Get(name string) (*v1alpha1.MariadbVirtualNetworkRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mariadbvirtualnetworkrule"), name)
	}
	return obj.(*v1alpha1.MariadbVirtualNetworkRule), nil
}
