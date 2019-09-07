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

// KmsCryptoKeyLister helps list KmsCryptoKeys.
type KmsCryptoKeyLister interface {
	// List lists all KmsCryptoKeys in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KmsCryptoKey, err error)
	// KmsCryptoKeys returns an object that can list and get KmsCryptoKeys.
	KmsCryptoKeys(namespace string) KmsCryptoKeyNamespaceLister
	KmsCryptoKeyListerExpansion
}

// kmsCryptoKeyLister implements the KmsCryptoKeyLister interface.
type kmsCryptoKeyLister struct {
	indexer cache.Indexer
}

// NewKmsCryptoKeyLister returns a new KmsCryptoKeyLister.
func NewKmsCryptoKeyLister(indexer cache.Indexer) KmsCryptoKeyLister {
	return &kmsCryptoKeyLister{indexer: indexer}
}

// List lists all KmsCryptoKeys in the indexer.
func (s *kmsCryptoKeyLister) List(selector labels.Selector) (ret []*v1alpha1.KmsCryptoKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KmsCryptoKey))
	})
	return ret, err
}

// KmsCryptoKeys returns an object that can list and get KmsCryptoKeys.
func (s *kmsCryptoKeyLister) KmsCryptoKeys(namespace string) KmsCryptoKeyNamespaceLister {
	return kmsCryptoKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KmsCryptoKeyNamespaceLister helps list and get KmsCryptoKeys.
type KmsCryptoKeyNamespaceLister interface {
	// List lists all KmsCryptoKeys in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.KmsCryptoKey, err error)
	// Get retrieves the KmsCryptoKey from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.KmsCryptoKey, error)
	KmsCryptoKeyNamespaceListerExpansion
}

// kmsCryptoKeyNamespaceLister implements the KmsCryptoKeyNamespaceLister
// interface.
type kmsCryptoKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KmsCryptoKeys in the indexer for a given namespace.
func (s kmsCryptoKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KmsCryptoKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KmsCryptoKey))
	})
	return ret, err
}

// Get retrieves the KmsCryptoKey from the indexer for a given namespace and name.
func (s kmsCryptoKeyNamespaceLister) Get(name string) (*v1alpha1.KmsCryptoKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kmscryptokey"), name)
	}
	return obj.(*v1alpha1.KmsCryptoKey), nil
}