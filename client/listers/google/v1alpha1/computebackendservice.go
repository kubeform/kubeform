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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ComputeBackendServiceLister helps list ComputeBackendServices.
type ComputeBackendServiceLister interface {
	// List lists all ComputeBackendServices in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeBackendService, err error)
	// ComputeBackendServices returns an object that can list and get ComputeBackendServices.
	ComputeBackendServices(namespace string) ComputeBackendServiceNamespaceLister
	ComputeBackendServiceListerExpansion
}

// computeBackendServiceLister implements the ComputeBackendServiceLister interface.
type computeBackendServiceLister struct {
	indexer cache.Indexer
}

// NewComputeBackendServiceLister returns a new ComputeBackendServiceLister.
func NewComputeBackendServiceLister(indexer cache.Indexer) ComputeBackendServiceLister {
	return &computeBackendServiceLister{indexer: indexer}
}

// List lists all ComputeBackendServices in the indexer.
func (s *computeBackendServiceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeBackendService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeBackendService))
	})
	return ret, err
}

// ComputeBackendServices returns an object that can list and get ComputeBackendServices.
func (s *computeBackendServiceLister) ComputeBackendServices(namespace string) ComputeBackendServiceNamespaceLister {
	return computeBackendServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeBackendServiceNamespaceLister helps list and get ComputeBackendServices.
type ComputeBackendServiceNamespaceLister interface {
	// List lists all ComputeBackendServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeBackendService, err error)
	// Get retrieves the ComputeBackendService from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeBackendService, error)
	ComputeBackendServiceNamespaceListerExpansion
}

// computeBackendServiceNamespaceLister implements the ComputeBackendServiceNamespaceLister
// interface.
type computeBackendServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeBackendServices in the indexer for a given namespace.
func (s computeBackendServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeBackendService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeBackendService))
	})
	return ret, err
}

// Get retrieves the ComputeBackendService from the indexer for a given namespace and name.
func (s computeBackendServiceNamespaceLister) Get(name string) (*v1alpha1.ComputeBackendService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computebackendservice"), name)
	}
	return obj.(*v1alpha1.ComputeBackendService), nil
}
