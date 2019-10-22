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

// DxConnectionLister helps list DxConnections.
type DxConnectionLister interface {
	// List lists all DxConnections in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DxConnection, err error)
	// DxConnections returns an object that can list and get DxConnections.
	DxConnections(namespace string) DxConnectionNamespaceLister
	DxConnectionListerExpansion
}

// dxConnectionLister implements the DxConnectionLister interface.
type dxConnectionLister struct {
	indexer cache.Indexer
}

// NewDxConnectionLister returns a new DxConnectionLister.
func NewDxConnectionLister(indexer cache.Indexer) DxConnectionLister {
	return &dxConnectionLister{indexer: indexer}
}

// List lists all DxConnections in the indexer.
func (s *dxConnectionLister) List(selector labels.Selector) (ret []*v1alpha1.DxConnection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DxConnection))
	})
	return ret, err
}

// DxConnections returns an object that can list and get DxConnections.
func (s *dxConnectionLister) DxConnections(namespace string) DxConnectionNamespaceLister {
	return dxConnectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DxConnectionNamespaceLister helps list and get DxConnections.
type DxConnectionNamespaceLister interface {
	// List lists all DxConnections in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DxConnection, err error)
	// Get retrieves the DxConnection from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DxConnection, error)
	DxConnectionNamespaceListerExpansion
}

// dxConnectionNamespaceLister implements the DxConnectionNamespaceLister
// interface.
type dxConnectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DxConnections in the indexer for a given namespace.
func (s dxConnectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DxConnection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DxConnection))
	})
	return ret, err
}

// Get retrieves the DxConnection from the indexer for a given namespace and name.
func (s dxConnectionNamespaceLister) Get(name string) (*v1alpha1.DxConnection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dxconnection"), name)
	}
	return obj.(*v1alpha1.DxConnection), nil
}
