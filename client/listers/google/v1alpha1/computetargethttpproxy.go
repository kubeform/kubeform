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

// ComputeTargetHTTPProxyLister helps list ComputeTargetHTTPProxies.
type ComputeTargetHTTPProxyLister interface {
	// List lists all ComputeTargetHTTPProxies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeTargetHTTPProxy, err error)
	// ComputeTargetHTTPProxies returns an object that can list and get ComputeTargetHTTPProxies.
	ComputeTargetHTTPProxies(namespace string) ComputeTargetHTTPProxyNamespaceLister
	ComputeTargetHTTPProxyListerExpansion
}

// computeTargetHTTPProxyLister implements the ComputeTargetHTTPProxyLister interface.
type computeTargetHTTPProxyLister struct {
	indexer cache.Indexer
}

// NewComputeTargetHTTPProxyLister returns a new ComputeTargetHTTPProxyLister.
func NewComputeTargetHTTPProxyLister(indexer cache.Indexer) ComputeTargetHTTPProxyLister {
	return &computeTargetHTTPProxyLister{indexer: indexer}
}

// List lists all ComputeTargetHTTPProxies in the indexer.
func (s *computeTargetHTTPProxyLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeTargetHTTPProxy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeTargetHTTPProxy))
	})
	return ret, err
}

// ComputeTargetHTTPProxies returns an object that can list and get ComputeTargetHTTPProxies.
func (s *computeTargetHTTPProxyLister) ComputeTargetHTTPProxies(namespace string) ComputeTargetHTTPProxyNamespaceLister {
	return computeTargetHTTPProxyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeTargetHTTPProxyNamespaceLister helps list and get ComputeTargetHTTPProxies.
type ComputeTargetHTTPProxyNamespaceLister interface {
	// List lists all ComputeTargetHTTPProxies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeTargetHTTPProxy, err error)
	// Get retrieves the ComputeTargetHTTPProxy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeTargetHTTPProxy, error)
	ComputeTargetHTTPProxyNamespaceListerExpansion
}

// computeTargetHTTPProxyNamespaceLister implements the ComputeTargetHTTPProxyNamespaceLister
// interface.
type computeTargetHTTPProxyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeTargetHTTPProxies in the indexer for a given namespace.
func (s computeTargetHTTPProxyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeTargetHTTPProxy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeTargetHTTPProxy))
	})
	return ret, err
}

// Get retrieves the ComputeTargetHTTPProxy from the indexer for a given namespace and name.
func (s computeTargetHTTPProxyNamespaceLister) Get(name string) (*v1alpha1.ComputeTargetHTTPProxy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computetargethttpproxy"), name)
	}
	return obj.(*v1alpha1.ComputeTargetHTTPProxy), nil
}
