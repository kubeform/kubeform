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

// SubnetNetworkSecurityGroupAssociationLister helps list SubnetNetworkSecurityGroupAssociations.
type SubnetNetworkSecurityGroupAssociationLister interface {
	// List lists all SubnetNetworkSecurityGroupAssociations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SubnetNetworkSecurityGroupAssociation, err error)
	// SubnetNetworkSecurityGroupAssociations returns an object that can list and get SubnetNetworkSecurityGroupAssociations.
	SubnetNetworkSecurityGroupAssociations(namespace string) SubnetNetworkSecurityGroupAssociationNamespaceLister
	SubnetNetworkSecurityGroupAssociationListerExpansion
}

// subnetNetworkSecurityGroupAssociationLister implements the SubnetNetworkSecurityGroupAssociationLister interface.
type subnetNetworkSecurityGroupAssociationLister struct {
	indexer cache.Indexer
}

// NewSubnetNetworkSecurityGroupAssociationLister returns a new SubnetNetworkSecurityGroupAssociationLister.
func NewSubnetNetworkSecurityGroupAssociationLister(indexer cache.Indexer) SubnetNetworkSecurityGroupAssociationLister {
	return &subnetNetworkSecurityGroupAssociationLister{indexer: indexer}
}

// List lists all SubnetNetworkSecurityGroupAssociations in the indexer.
func (s *subnetNetworkSecurityGroupAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.SubnetNetworkSecurityGroupAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SubnetNetworkSecurityGroupAssociation))
	})
	return ret, err
}

// SubnetNetworkSecurityGroupAssociations returns an object that can list and get SubnetNetworkSecurityGroupAssociations.
func (s *subnetNetworkSecurityGroupAssociationLister) SubnetNetworkSecurityGroupAssociations(namespace string) SubnetNetworkSecurityGroupAssociationNamespaceLister {
	return subnetNetworkSecurityGroupAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SubnetNetworkSecurityGroupAssociationNamespaceLister helps list and get SubnetNetworkSecurityGroupAssociations.
type SubnetNetworkSecurityGroupAssociationNamespaceLister interface {
	// List lists all SubnetNetworkSecurityGroupAssociations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SubnetNetworkSecurityGroupAssociation, err error)
	// Get retrieves the SubnetNetworkSecurityGroupAssociation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SubnetNetworkSecurityGroupAssociation, error)
	SubnetNetworkSecurityGroupAssociationNamespaceListerExpansion
}

// subnetNetworkSecurityGroupAssociationNamespaceLister implements the SubnetNetworkSecurityGroupAssociationNamespaceLister
// interface.
type subnetNetworkSecurityGroupAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SubnetNetworkSecurityGroupAssociations in the indexer for a given namespace.
func (s subnetNetworkSecurityGroupAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SubnetNetworkSecurityGroupAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SubnetNetworkSecurityGroupAssociation))
	})
	return ret, err
}

// Get retrieves the SubnetNetworkSecurityGroupAssociation from the indexer for a given namespace and name.
func (s subnetNetworkSecurityGroupAssociationNamespaceLister) Get(name string) (*v1alpha1.SubnetNetworkSecurityGroupAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("subnetnetworksecuritygroupassociation"), name)
	}
	return obj.(*v1alpha1.SubnetNetworkSecurityGroupAssociation), nil
}
