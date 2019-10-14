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

// DxGatewayAssociationLister helps list DxGatewayAssociations.
type DxGatewayAssociationLister interface {
	// List lists all DxGatewayAssociations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DxGatewayAssociation, err error)
	// DxGatewayAssociations returns an object that can list and get DxGatewayAssociations.
	DxGatewayAssociations(namespace string) DxGatewayAssociationNamespaceLister
	DxGatewayAssociationListerExpansion
}

// dxGatewayAssociationLister implements the DxGatewayAssociationLister interface.
type dxGatewayAssociationLister struct {
	indexer cache.Indexer
}

// NewDxGatewayAssociationLister returns a new DxGatewayAssociationLister.
func NewDxGatewayAssociationLister(indexer cache.Indexer) DxGatewayAssociationLister {
	return &dxGatewayAssociationLister{indexer: indexer}
}

// List lists all DxGatewayAssociations in the indexer.
func (s *dxGatewayAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.DxGatewayAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DxGatewayAssociation))
	})
	return ret, err
}

// DxGatewayAssociations returns an object that can list and get DxGatewayAssociations.
func (s *dxGatewayAssociationLister) DxGatewayAssociations(namespace string) DxGatewayAssociationNamespaceLister {
	return dxGatewayAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DxGatewayAssociationNamespaceLister helps list and get DxGatewayAssociations.
type DxGatewayAssociationNamespaceLister interface {
	// List lists all DxGatewayAssociations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DxGatewayAssociation, err error)
	// Get retrieves the DxGatewayAssociation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DxGatewayAssociation, error)
	DxGatewayAssociationNamespaceListerExpansion
}

// dxGatewayAssociationNamespaceLister implements the DxGatewayAssociationNamespaceLister
// interface.
type dxGatewayAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DxGatewayAssociations in the indexer for a given namespace.
func (s dxGatewayAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DxGatewayAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DxGatewayAssociation))
	})
	return ret, err
}

// Get retrieves the DxGatewayAssociation from the indexer for a given namespace and name.
func (s dxGatewayAssociationNamespaceLister) Get(name string) (*v1alpha1.DxGatewayAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dxgatewayassociation"), name)
	}
	return obj.(*v1alpha1.DxGatewayAssociation), nil
}
