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

// FirewallNATRuleCollectionLister helps list FirewallNATRuleCollections.
type FirewallNATRuleCollectionLister interface {
	// List lists all FirewallNATRuleCollections in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FirewallNATRuleCollection, err error)
	// FirewallNATRuleCollections returns an object that can list and get FirewallNATRuleCollections.
	FirewallNATRuleCollections(namespace string) FirewallNATRuleCollectionNamespaceLister
	FirewallNATRuleCollectionListerExpansion
}

// firewallNATRuleCollectionLister implements the FirewallNATRuleCollectionLister interface.
type firewallNATRuleCollectionLister struct {
	indexer cache.Indexer
}

// NewFirewallNATRuleCollectionLister returns a new FirewallNATRuleCollectionLister.
func NewFirewallNATRuleCollectionLister(indexer cache.Indexer) FirewallNATRuleCollectionLister {
	return &firewallNATRuleCollectionLister{indexer: indexer}
}

// List lists all FirewallNATRuleCollections in the indexer.
func (s *firewallNATRuleCollectionLister) List(selector labels.Selector) (ret []*v1alpha1.FirewallNATRuleCollection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FirewallNATRuleCollection))
	})
	return ret, err
}

// FirewallNATRuleCollections returns an object that can list and get FirewallNATRuleCollections.
func (s *firewallNATRuleCollectionLister) FirewallNATRuleCollections(namespace string) FirewallNATRuleCollectionNamespaceLister {
	return firewallNATRuleCollectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FirewallNATRuleCollectionNamespaceLister helps list and get FirewallNATRuleCollections.
type FirewallNATRuleCollectionNamespaceLister interface {
	// List lists all FirewallNATRuleCollections in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FirewallNATRuleCollection, err error)
	// Get retrieves the FirewallNATRuleCollection from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FirewallNATRuleCollection, error)
	FirewallNATRuleCollectionNamespaceListerExpansion
}

// firewallNATRuleCollectionNamespaceLister implements the FirewallNATRuleCollectionNamespaceLister
// interface.
type firewallNATRuleCollectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FirewallNATRuleCollections in the indexer for a given namespace.
func (s firewallNATRuleCollectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FirewallNATRuleCollection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FirewallNATRuleCollection))
	})
	return ret, err
}

// Get retrieves the FirewallNATRuleCollection from the indexer for a given namespace and name.
func (s firewallNATRuleCollectionNamespaceLister) Get(name string) (*v1alpha1.FirewallNATRuleCollection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("firewallnatrulecollection"), name)
	}
	return obj.(*v1alpha1.FirewallNATRuleCollection), nil
}
