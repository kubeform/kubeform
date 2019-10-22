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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StorageShareLister helps list StorageShares.
type StorageShareLister interface {
	// List lists all StorageShares in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StorageShare, err error)
	// StorageShares returns an object that can list and get StorageShares.
	StorageShares(namespace string) StorageShareNamespaceLister
	StorageShareListerExpansion
}

// storageShareLister implements the StorageShareLister interface.
type storageShareLister struct {
	indexer cache.Indexer
}

// NewStorageShareLister returns a new StorageShareLister.
func NewStorageShareLister(indexer cache.Indexer) StorageShareLister {
	return &storageShareLister{indexer: indexer}
}

// List lists all StorageShares in the indexer.
func (s *storageShareLister) List(selector labels.Selector) (ret []*v1alpha1.StorageShare, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageShare))
	})
	return ret, err
}

// StorageShares returns an object that can list and get StorageShares.
func (s *storageShareLister) StorageShares(namespace string) StorageShareNamespaceLister {
	return storageShareNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StorageShareNamespaceLister helps list and get StorageShares.
type StorageShareNamespaceLister interface {
	// List lists all StorageShares in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StorageShare, err error)
	// Get retrieves the StorageShare from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StorageShare, error)
	StorageShareNamespaceListerExpansion
}

// storageShareNamespaceLister implements the StorageShareNamespaceLister
// interface.
type storageShareNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StorageShares in the indexer for a given namespace.
func (s storageShareNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StorageShare, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageShare))
	})
	return ret, err
}

// Get retrieves the StorageShare from the indexer for a given namespace and name.
func (s storageShareNamespaceLister) Get(name string) (*v1alpha1.StorageShare, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storageshare"), name)
	}
	return obj.(*v1alpha1.StorageShare), nil
}
