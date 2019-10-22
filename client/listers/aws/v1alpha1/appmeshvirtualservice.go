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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AppmeshVirtualServiceLister helps list AppmeshVirtualServices.
type AppmeshVirtualServiceLister interface {
	// List lists all AppmeshVirtualServices in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AppmeshVirtualService, err error)
	// AppmeshVirtualServices returns an object that can list and get AppmeshVirtualServices.
	AppmeshVirtualServices(namespace string) AppmeshVirtualServiceNamespaceLister
	AppmeshVirtualServiceListerExpansion
}

// appmeshVirtualServiceLister implements the AppmeshVirtualServiceLister interface.
type appmeshVirtualServiceLister struct {
	indexer cache.Indexer
}

// NewAppmeshVirtualServiceLister returns a new AppmeshVirtualServiceLister.
func NewAppmeshVirtualServiceLister(indexer cache.Indexer) AppmeshVirtualServiceLister {
	return &appmeshVirtualServiceLister{indexer: indexer}
}

// List lists all AppmeshVirtualServices in the indexer.
func (s *appmeshVirtualServiceLister) List(selector labels.Selector) (ret []*v1alpha1.AppmeshVirtualService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppmeshVirtualService))
	})
	return ret, err
}

// AppmeshVirtualServices returns an object that can list and get AppmeshVirtualServices.
func (s *appmeshVirtualServiceLister) AppmeshVirtualServices(namespace string) AppmeshVirtualServiceNamespaceLister {
	return appmeshVirtualServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppmeshVirtualServiceNamespaceLister helps list and get AppmeshVirtualServices.
type AppmeshVirtualServiceNamespaceLister interface {
	// List lists all AppmeshVirtualServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AppmeshVirtualService, err error)
	// Get retrieves the AppmeshVirtualService from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AppmeshVirtualService, error)
	AppmeshVirtualServiceNamespaceListerExpansion
}

// appmeshVirtualServiceNamespaceLister implements the AppmeshVirtualServiceNamespaceLister
// interface.
type appmeshVirtualServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppmeshVirtualServices in the indexer for a given namespace.
func (s appmeshVirtualServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AppmeshVirtualService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppmeshVirtualService))
	})
	return ret, err
}

// Get retrieves the AppmeshVirtualService from the indexer for a given namespace and name.
func (s appmeshVirtualServiceNamespaceLister) Get(name string) (*v1alpha1.AppmeshVirtualService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("appmeshvirtualservice"), name)
	}
	return obj.(*v1alpha1.AppmeshVirtualService), nil
}
