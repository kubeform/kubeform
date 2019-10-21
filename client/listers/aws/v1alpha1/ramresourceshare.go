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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// RamResourceShareLister helps list RamResourceShares.
type RamResourceShareLister interface {
	// List lists all RamResourceShares in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RamResourceShare, err error)
	// RamResourceShares returns an object that can list and get RamResourceShares.
	RamResourceShares(namespace string) RamResourceShareNamespaceLister
	RamResourceShareListerExpansion
}

// ramResourceShareLister implements the RamResourceShareLister interface.
type ramResourceShareLister struct {
	indexer cache.Indexer
}

// NewRamResourceShareLister returns a new RamResourceShareLister.
func NewRamResourceShareLister(indexer cache.Indexer) RamResourceShareLister {
	return &ramResourceShareLister{indexer: indexer}
}

// List lists all RamResourceShares in the indexer.
func (s *ramResourceShareLister) List(selector labels.Selector) (ret []*v1alpha1.RamResourceShare, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RamResourceShare))
	})
	return ret, err
}

// RamResourceShares returns an object that can list and get RamResourceShares.
func (s *ramResourceShareLister) RamResourceShares(namespace string) RamResourceShareNamespaceLister {
	return ramResourceShareNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RamResourceShareNamespaceLister helps list and get RamResourceShares.
type RamResourceShareNamespaceLister interface {
	// List lists all RamResourceShares in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RamResourceShare, err error)
	// Get retrieves the RamResourceShare from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RamResourceShare, error)
	RamResourceShareNamespaceListerExpansion
}

// ramResourceShareNamespaceLister implements the RamResourceShareNamespaceLister
// interface.
type ramResourceShareNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RamResourceShares in the indexer for a given namespace.
func (s ramResourceShareNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RamResourceShare, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RamResourceShare))
	})
	return ret, err
}

// Get retrieves the RamResourceShare from the indexer for a given namespace and name.
func (s ramResourceShareNamespaceLister) Get(name string) (*v1alpha1.RamResourceShare, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ramresourceshare"), name)
	}
	return obj.(*v1alpha1.RamResourceShare), nil
}