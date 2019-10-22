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
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RdnsLister helps list Rdnses.
type RdnsLister interface {
	// List lists all Rdnses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Rdns, err error)
	// Rdnses returns an object that can list and get Rdnses.
	Rdnses(namespace string) RdnsNamespaceLister
	RdnsListerExpansion
}

// rdnsLister implements the RdnsLister interface.
type rdnsLister struct {
	indexer cache.Indexer
}

// NewRdnsLister returns a new RdnsLister.
func NewRdnsLister(indexer cache.Indexer) RdnsLister {
	return &rdnsLister{indexer: indexer}
}

// List lists all Rdnses in the indexer.
func (s *rdnsLister) List(selector labels.Selector) (ret []*v1alpha1.Rdns, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Rdns))
	})
	return ret, err
}

// Rdnses returns an object that can list and get Rdnses.
func (s *rdnsLister) Rdnses(namespace string) RdnsNamespaceLister {
	return rdnsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RdnsNamespaceLister helps list and get Rdnses.
type RdnsNamespaceLister interface {
	// List lists all Rdnses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Rdns, err error)
	// Get retrieves the Rdns from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Rdns, error)
	RdnsNamespaceListerExpansion
}

// rdnsNamespaceLister implements the RdnsNamespaceLister
// interface.
type rdnsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Rdnses in the indexer for a given namespace.
func (s rdnsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Rdns, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Rdns))
	})
	return ret, err
}

// Get retrieves the Rdns from the indexer for a given namespace and name.
func (s rdnsNamespaceLister) Get(name string) (*v1alpha1.Rdns, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("rdns"), name)
	}
	return obj.(*v1alpha1.Rdns), nil
}
