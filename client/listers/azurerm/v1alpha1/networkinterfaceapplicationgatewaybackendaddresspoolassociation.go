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

// NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationLister helps list NetworkInterfaceApplicationGatewayBackendAddressPoolAssociations.
type NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationLister interface {
	// List lists all NetworkInterfaceApplicationGatewayBackendAddressPoolAssociations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation, err error)
	// NetworkInterfaceApplicationGatewayBackendAddressPoolAssociations returns an object that can list and get NetworkInterfaceApplicationGatewayBackendAddressPoolAssociations.
	NetworkInterfaceApplicationGatewayBackendAddressPoolAssociations(namespace string) NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationNamespaceLister
	NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationListerExpansion
}

// networkInterfaceApplicationGatewayBackendAddressPoolAssociationLister implements the NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationLister interface.
type networkInterfaceApplicationGatewayBackendAddressPoolAssociationLister struct {
	indexer cache.Indexer
}

// NewNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationLister returns a new NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationLister.
func NewNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationLister(indexer cache.Indexer) NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationLister {
	return &networkInterfaceApplicationGatewayBackendAddressPoolAssociationLister{indexer: indexer}
}

// List lists all NetworkInterfaceApplicationGatewayBackendAddressPoolAssociations in the indexer.
func (s *networkInterfaceApplicationGatewayBackendAddressPoolAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation))
	})
	return ret, err
}

// NetworkInterfaceApplicationGatewayBackendAddressPoolAssociations returns an object that can list and get NetworkInterfaceApplicationGatewayBackendAddressPoolAssociations.
func (s *networkInterfaceApplicationGatewayBackendAddressPoolAssociationLister) NetworkInterfaceApplicationGatewayBackendAddressPoolAssociations(namespace string) NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationNamespaceLister {
	return networkInterfaceApplicationGatewayBackendAddressPoolAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationNamespaceLister helps list and get NetworkInterfaceApplicationGatewayBackendAddressPoolAssociations.
type NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationNamespaceLister interface {
	// List lists all NetworkInterfaceApplicationGatewayBackendAddressPoolAssociations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation, err error)
	// Get retrieves the NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation, error)
	NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationNamespaceListerExpansion
}

// networkInterfaceApplicationGatewayBackendAddressPoolAssociationNamespaceLister implements the NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationNamespaceLister
// interface.
type networkInterfaceApplicationGatewayBackendAddressPoolAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkInterfaceApplicationGatewayBackendAddressPoolAssociations in the indexer for a given namespace.
func (s networkInterfaceApplicationGatewayBackendAddressPoolAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation))
	})
	return ret, err
}

// Get retrieves the NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation from the indexer for a given namespace and name.
func (s networkInterfaceApplicationGatewayBackendAddressPoolAssociationNamespaceLister) Get(name string) (*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networkinterfaceapplicationgatewaybackendaddresspoolassociation"), name)
	}
	return obj.(*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation), nil
}