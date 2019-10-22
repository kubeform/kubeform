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

// DxHostedPrivateVirtualInterfaceAccepterLister helps list DxHostedPrivateVirtualInterfaceAccepters.
type DxHostedPrivateVirtualInterfaceAccepterLister interface {
	// List lists all DxHostedPrivateVirtualInterfaceAccepters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DxHostedPrivateVirtualInterfaceAccepter, err error)
	// DxHostedPrivateVirtualInterfaceAccepters returns an object that can list and get DxHostedPrivateVirtualInterfaceAccepters.
	DxHostedPrivateVirtualInterfaceAccepters(namespace string) DxHostedPrivateVirtualInterfaceAccepterNamespaceLister
	DxHostedPrivateVirtualInterfaceAccepterListerExpansion
}

// dxHostedPrivateVirtualInterfaceAccepterLister implements the DxHostedPrivateVirtualInterfaceAccepterLister interface.
type dxHostedPrivateVirtualInterfaceAccepterLister struct {
	indexer cache.Indexer
}

// NewDxHostedPrivateVirtualInterfaceAccepterLister returns a new DxHostedPrivateVirtualInterfaceAccepterLister.
func NewDxHostedPrivateVirtualInterfaceAccepterLister(indexer cache.Indexer) DxHostedPrivateVirtualInterfaceAccepterLister {
	return &dxHostedPrivateVirtualInterfaceAccepterLister{indexer: indexer}
}

// List lists all DxHostedPrivateVirtualInterfaceAccepters in the indexer.
func (s *dxHostedPrivateVirtualInterfaceAccepterLister) List(selector labels.Selector) (ret []*v1alpha1.DxHostedPrivateVirtualInterfaceAccepter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DxHostedPrivateVirtualInterfaceAccepter))
	})
	return ret, err
}

// DxHostedPrivateVirtualInterfaceAccepters returns an object that can list and get DxHostedPrivateVirtualInterfaceAccepters.
func (s *dxHostedPrivateVirtualInterfaceAccepterLister) DxHostedPrivateVirtualInterfaceAccepters(namespace string) DxHostedPrivateVirtualInterfaceAccepterNamespaceLister {
	return dxHostedPrivateVirtualInterfaceAccepterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DxHostedPrivateVirtualInterfaceAccepterNamespaceLister helps list and get DxHostedPrivateVirtualInterfaceAccepters.
type DxHostedPrivateVirtualInterfaceAccepterNamespaceLister interface {
	// List lists all DxHostedPrivateVirtualInterfaceAccepters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DxHostedPrivateVirtualInterfaceAccepter, err error)
	// Get retrieves the DxHostedPrivateVirtualInterfaceAccepter from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DxHostedPrivateVirtualInterfaceAccepter, error)
	DxHostedPrivateVirtualInterfaceAccepterNamespaceListerExpansion
}

// dxHostedPrivateVirtualInterfaceAccepterNamespaceLister implements the DxHostedPrivateVirtualInterfaceAccepterNamespaceLister
// interface.
type dxHostedPrivateVirtualInterfaceAccepterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DxHostedPrivateVirtualInterfaceAccepters in the indexer for a given namespace.
func (s dxHostedPrivateVirtualInterfaceAccepterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DxHostedPrivateVirtualInterfaceAccepter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DxHostedPrivateVirtualInterfaceAccepter))
	})
	return ret, err
}

// Get retrieves the DxHostedPrivateVirtualInterfaceAccepter from the indexer for a given namespace and name.
func (s dxHostedPrivateVirtualInterfaceAccepterNamespaceLister) Get(name string) (*v1alpha1.DxHostedPrivateVirtualInterfaceAccepter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dxhostedprivatevirtualinterfaceaccepter"), name)
	}
	return obj.(*v1alpha1.DxHostedPrivateVirtualInterfaceAccepter), nil
}
