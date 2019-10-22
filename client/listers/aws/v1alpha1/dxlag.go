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

// DxLagLister helps list DxLags.
type DxLagLister interface {
	// List lists all DxLags in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DxLag, err error)
	// DxLags returns an object that can list and get DxLags.
	DxLags(namespace string) DxLagNamespaceLister
	DxLagListerExpansion
}

// dxLagLister implements the DxLagLister interface.
type dxLagLister struct {
	indexer cache.Indexer
}

// NewDxLagLister returns a new DxLagLister.
func NewDxLagLister(indexer cache.Indexer) DxLagLister {
	return &dxLagLister{indexer: indexer}
}

// List lists all DxLags in the indexer.
func (s *dxLagLister) List(selector labels.Selector) (ret []*v1alpha1.DxLag, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DxLag))
	})
	return ret, err
}

// DxLags returns an object that can list and get DxLags.
func (s *dxLagLister) DxLags(namespace string) DxLagNamespaceLister {
	return dxLagNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DxLagNamespaceLister helps list and get DxLags.
type DxLagNamespaceLister interface {
	// List lists all DxLags in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DxLag, err error)
	// Get retrieves the DxLag from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DxLag, error)
	DxLagNamespaceListerExpansion
}

// dxLagNamespaceLister implements the DxLagNamespaceLister
// interface.
type dxLagNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DxLags in the indexer for a given namespace.
func (s dxLagNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DxLag, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DxLag))
	})
	return ret, err
}

// Get retrieves the DxLag from the indexer for a given namespace and name.
func (s dxLagNamespaceLister) Get(name string) (*v1alpha1.DxLag, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dxlag"), name)
	}
	return obj.(*v1alpha1.DxLag), nil
}
