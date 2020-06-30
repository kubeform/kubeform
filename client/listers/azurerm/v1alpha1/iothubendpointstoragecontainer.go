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

// IothubEndpointStorageContainerLister helps list IothubEndpointStorageContainers.
type IothubEndpointStorageContainerLister interface {
	// List lists all IothubEndpointStorageContainers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IothubEndpointStorageContainer, err error)
	// IothubEndpointStorageContainers returns an object that can list and get IothubEndpointStorageContainers.
	IothubEndpointStorageContainers(namespace string) IothubEndpointStorageContainerNamespaceLister
	IothubEndpointStorageContainerListerExpansion
}

// iothubEndpointStorageContainerLister implements the IothubEndpointStorageContainerLister interface.
type iothubEndpointStorageContainerLister struct {
	indexer cache.Indexer
}

// NewIothubEndpointStorageContainerLister returns a new IothubEndpointStorageContainerLister.
func NewIothubEndpointStorageContainerLister(indexer cache.Indexer) IothubEndpointStorageContainerLister {
	return &iothubEndpointStorageContainerLister{indexer: indexer}
}

// List lists all IothubEndpointStorageContainers in the indexer.
func (s *iothubEndpointStorageContainerLister) List(selector labels.Selector) (ret []*v1alpha1.IothubEndpointStorageContainer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IothubEndpointStorageContainer))
	})
	return ret, err
}

// IothubEndpointStorageContainers returns an object that can list and get IothubEndpointStorageContainers.
func (s *iothubEndpointStorageContainerLister) IothubEndpointStorageContainers(namespace string) IothubEndpointStorageContainerNamespaceLister {
	return iothubEndpointStorageContainerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IothubEndpointStorageContainerNamespaceLister helps list and get IothubEndpointStorageContainers.
type IothubEndpointStorageContainerNamespaceLister interface {
	// List lists all IothubEndpointStorageContainers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IothubEndpointStorageContainer, err error)
	// Get retrieves the IothubEndpointStorageContainer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IothubEndpointStorageContainer, error)
	IothubEndpointStorageContainerNamespaceListerExpansion
}

// iothubEndpointStorageContainerNamespaceLister implements the IothubEndpointStorageContainerNamespaceLister
// interface.
type iothubEndpointStorageContainerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IothubEndpointStorageContainers in the indexer for a given namespace.
func (s iothubEndpointStorageContainerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IothubEndpointStorageContainer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IothubEndpointStorageContainer))
	})
	return ret, err
}

// Get retrieves the IothubEndpointStorageContainer from the indexer for a given namespace and name.
func (s iothubEndpointStorageContainerNamespaceLister) Get(name string) (*v1alpha1.IothubEndpointStorageContainer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iothubendpointstoragecontainer"), name)
	}
	return obj.(*v1alpha1.IothubEndpointStorageContainer), nil
}
