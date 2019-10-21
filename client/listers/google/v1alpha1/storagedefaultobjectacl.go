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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// StorageDefaultObjectACLLister helps list StorageDefaultObjectACLs.
type StorageDefaultObjectACLLister interface {
	// List lists all StorageDefaultObjectACLs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StorageDefaultObjectACL, err error)
	// StorageDefaultObjectACLs returns an object that can list and get StorageDefaultObjectACLs.
	StorageDefaultObjectACLs(namespace string) StorageDefaultObjectACLNamespaceLister
	StorageDefaultObjectACLListerExpansion
}

// storageDefaultObjectACLLister implements the StorageDefaultObjectACLLister interface.
type storageDefaultObjectACLLister struct {
	indexer cache.Indexer
}

// NewStorageDefaultObjectACLLister returns a new StorageDefaultObjectACLLister.
func NewStorageDefaultObjectACLLister(indexer cache.Indexer) StorageDefaultObjectACLLister {
	return &storageDefaultObjectACLLister{indexer: indexer}
}

// List lists all StorageDefaultObjectACLs in the indexer.
func (s *storageDefaultObjectACLLister) List(selector labels.Selector) (ret []*v1alpha1.StorageDefaultObjectACL, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageDefaultObjectACL))
	})
	return ret, err
}

// StorageDefaultObjectACLs returns an object that can list and get StorageDefaultObjectACLs.
func (s *storageDefaultObjectACLLister) StorageDefaultObjectACLs(namespace string) StorageDefaultObjectACLNamespaceLister {
	return storageDefaultObjectACLNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StorageDefaultObjectACLNamespaceLister helps list and get StorageDefaultObjectACLs.
type StorageDefaultObjectACLNamespaceLister interface {
	// List lists all StorageDefaultObjectACLs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StorageDefaultObjectACL, err error)
	// Get retrieves the StorageDefaultObjectACL from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StorageDefaultObjectACL, error)
	StorageDefaultObjectACLNamespaceListerExpansion
}

// storageDefaultObjectACLNamespaceLister implements the StorageDefaultObjectACLNamespaceLister
// interface.
type storageDefaultObjectACLNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StorageDefaultObjectACLs in the indexer for a given namespace.
func (s storageDefaultObjectACLNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StorageDefaultObjectACL, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageDefaultObjectACL))
	})
	return ret, err
}

// Get retrieves the StorageDefaultObjectACL from the indexer for a given namespace and name.
func (s storageDefaultObjectACLNamespaceLister) Get(name string) (*v1alpha1.StorageDefaultObjectACL, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storagedefaultobjectacl"), name)
	}
	return obj.(*v1alpha1.StorageDefaultObjectACL), nil
}