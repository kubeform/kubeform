/*
Copyright 2019 The Kubeform Authors.

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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// NetworkSecurityRuleLister helps list NetworkSecurityRules.
type NetworkSecurityRuleLister interface {
	// List lists all NetworkSecurityRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkSecurityRule, err error)
	// NetworkSecurityRules returns an object that can list and get NetworkSecurityRules.
	NetworkSecurityRules(namespace string) NetworkSecurityRuleNamespaceLister
	NetworkSecurityRuleListerExpansion
}

// networkSecurityRuleLister implements the NetworkSecurityRuleLister interface.
type networkSecurityRuleLister struct {
	indexer cache.Indexer
}

// NewNetworkSecurityRuleLister returns a new NetworkSecurityRuleLister.
func NewNetworkSecurityRuleLister(indexer cache.Indexer) NetworkSecurityRuleLister {
	return &networkSecurityRuleLister{indexer: indexer}
}

// List lists all NetworkSecurityRules in the indexer.
func (s *networkSecurityRuleLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkSecurityRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkSecurityRule))
	})
	return ret, err
}

// NetworkSecurityRules returns an object that can list and get NetworkSecurityRules.
func (s *networkSecurityRuleLister) NetworkSecurityRules(namespace string) NetworkSecurityRuleNamespaceLister {
	return networkSecurityRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkSecurityRuleNamespaceLister helps list and get NetworkSecurityRules.
type NetworkSecurityRuleNamespaceLister interface {
	// List lists all NetworkSecurityRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkSecurityRule, err error)
	// Get retrieves the NetworkSecurityRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NetworkSecurityRule, error)
	NetworkSecurityRuleNamespaceListerExpansion
}

// networkSecurityRuleNamespaceLister implements the NetworkSecurityRuleNamespaceLister
// interface.
type networkSecurityRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkSecurityRules in the indexer for a given namespace.
func (s networkSecurityRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkSecurityRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkSecurityRule))
	})
	return ret, err
}

// Get retrieves the NetworkSecurityRule from the indexer for a given namespace and name.
func (s networkSecurityRuleNamespaceLister) Get(name string) (*v1alpha1.NetworkSecurityRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networksecurityrule"), name)
	}
	return obj.(*v1alpha1.NetworkSecurityRule), nil
}
