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

// NetappPoolLister helps list NetappPools.
type NetappPoolLister interface {
	// List lists all NetappPools in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NetappPool, err error)
	// NetappPools returns an object that can list and get NetappPools.
	NetappPools(namespace string) NetappPoolNamespaceLister
	NetappPoolListerExpansion
}

// netappPoolLister implements the NetappPoolLister interface.
type netappPoolLister struct {
	indexer cache.Indexer
}

// NewNetappPoolLister returns a new NetappPoolLister.
func NewNetappPoolLister(indexer cache.Indexer) NetappPoolLister {
	return &netappPoolLister{indexer: indexer}
}

// List lists all NetappPools in the indexer.
func (s *netappPoolLister) List(selector labels.Selector) (ret []*v1alpha1.NetappPool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetappPool))
	})
	return ret, err
}

// NetappPools returns an object that can list and get NetappPools.
func (s *netappPoolLister) NetappPools(namespace string) NetappPoolNamespaceLister {
	return netappPoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetappPoolNamespaceLister helps list and get NetappPools.
type NetappPoolNamespaceLister interface {
	// List lists all NetappPools in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NetappPool, err error)
	// Get retrieves the NetappPool from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NetappPool, error)
	NetappPoolNamespaceListerExpansion
}

// netappPoolNamespaceLister implements the NetappPoolNamespaceLister
// interface.
type netappPoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetappPools in the indexer for a given namespace.
func (s netappPoolNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetappPool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetappPool))
	})
	return ret, err
}

// Get retrieves the NetappPool from the indexer for a given namespace and name.
func (s netappPoolNamespaceLister) Get(name string) (*v1alpha1.NetappPool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("netapppool"), name)
	}
	return obj.(*v1alpha1.NetappPool), nil
}