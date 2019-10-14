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

// StorageObjectACLLister helps list StorageObjectACLs.
type StorageObjectACLLister interface {
	// List lists all StorageObjectACLs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StorageObjectACL, err error)
	// StorageObjectACLs returns an object that can list and get StorageObjectACLs.
	StorageObjectACLs(namespace string) StorageObjectACLNamespaceLister
	StorageObjectACLListerExpansion
}

// storageObjectACLLister implements the StorageObjectACLLister interface.
type storageObjectACLLister struct {
	indexer cache.Indexer
}

// NewStorageObjectACLLister returns a new StorageObjectACLLister.
func NewStorageObjectACLLister(indexer cache.Indexer) StorageObjectACLLister {
	return &storageObjectACLLister{indexer: indexer}
}

// List lists all StorageObjectACLs in the indexer.
func (s *storageObjectACLLister) List(selector labels.Selector) (ret []*v1alpha1.StorageObjectACL, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageObjectACL))
	})
	return ret, err
}

// StorageObjectACLs returns an object that can list and get StorageObjectACLs.
func (s *storageObjectACLLister) StorageObjectACLs(namespace string) StorageObjectACLNamespaceLister {
	return storageObjectACLNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StorageObjectACLNamespaceLister helps list and get StorageObjectACLs.
type StorageObjectACLNamespaceLister interface {
	// List lists all StorageObjectACLs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StorageObjectACL, err error)
	// Get retrieves the StorageObjectACL from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StorageObjectACL, error)
	StorageObjectACLNamespaceListerExpansion
}

// storageObjectACLNamespaceLister implements the StorageObjectACLNamespaceLister
// interface.
type storageObjectACLNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StorageObjectACLs in the indexer for a given namespace.
func (s storageObjectACLNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StorageObjectACL, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageObjectACL))
	})
	return ret, err
}

// Get retrieves the StorageObjectACL from the indexer for a given namespace and name.
func (s storageObjectACLNamespaceLister) Get(name string) (*v1alpha1.StorageObjectACL, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storageobjectacl"), name)
	}
	return obj.(*v1alpha1.StorageObjectACL), nil
}
