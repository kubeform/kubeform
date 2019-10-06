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

// VirtualNetworkLister helps list VirtualNetworks.
type VirtualNetworkLister interface {
	// List lists all VirtualNetworks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualNetwork, err error)
	// VirtualNetworks returns an object that can list and get VirtualNetworks.
	VirtualNetworks(namespace string) VirtualNetworkNamespaceLister
	VirtualNetworkListerExpansion
}

// virtualNetworkLister implements the VirtualNetworkLister interface.
type virtualNetworkLister struct {
	indexer cache.Indexer
}

// NewVirtualNetworkLister returns a new VirtualNetworkLister.
func NewVirtualNetworkLister(indexer cache.Indexer) VirtualNetworkLister {
	return &virtualNetworkLister{indexer: indexer}
}

// List lists all VirtualNetworks in the indexer.
func (s *virtualNetworkLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualNetwork, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualNetwork))
	})
	return ret, err
}

// VirtualNetworks returns an object that can list and get VirtualNetworks.
func (s *virtualNetworkLister) VirtualNetworks(namespace string) VirtualNetworkNamespaceLister {
	return virtualNetworkNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualNetworkNamespaceLister helps list and get VirtualNetworks.
type VirtualNetworkNamespaceLister interface {
	// List lists all VirtualNetworks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualNetwork, err error)
	// Get retrieves the VirtualNetwork from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VirtualNetwork, error)
	VirtualNetworkNamespaceListerExpansion
}

// virtualNetworkNamespaceLister implements the VirtualNetworkNamespaceLister
// interface.
type virtualNetworkNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualNetworks in the indexer for a given namespace.
func (s virtualNetworkNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualNetwork, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualNetwork))
	})
	return ret, err
}

// Get retrieves the VirtualNetwork from the indexer for a given namespace and name.
func (s virtualNetworkNamespaceLister) Get(name string) (*v1alpha1.VirtualNetwork, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("virtualnetwork"), name)
	}
	return obj.(*v1alpha1.VirtualNetwork), nil
}
