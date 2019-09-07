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

// RamResourceAssociationLister helps list RamResourceAssociations.
type RamResourceAssociationLister interface {
	// List lists all RamResourceAssociations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RamResourceAssociation, err error)
	// RamResourceAssociations returns an object that can list and get RamResourceAssociations.
	RamResourceAssociations(namespace string) RamResourceAssociationNamespaceLister
	RamResourceAssociationListerExpansion
}

// ramResourceAssociationLister implements the RamResourceAssociationLister interface.
type ramResourceAssociationLister struct {
	indexer cache.Indexer
}

// NewRamResourceAssociationLister returns a new RamResourceAssociationLister.
func NewRamResourceAssociationLister(indexer cache.Indexer) RamResourceAssociationLister {
	return &ramResourceAssociationLister{indexer: indexer}
}

// List lists all RamResourceAssociations in the indexer.
func (s *ramResourceAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.RamResourceAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RamResourceAssociation))
	})
	return ret, err
}

// RamResourceAssociations returns an object that can list and get RamResourceAssociations.
func (s *ramResourceAssociationLister) RamResourceAssociations(namespace string) RamResourceAssociationNamespaceLister {
	return ramResourceAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RamResourceAssociationNamespaceLister helps list and get RamResourceAssociations.
type RamResourceAssociationNamespaceLister interface {
	// List lists all RamResourceAssociations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RamResourceAssociation, err error)
	// Get retrieves the RamResourceAssociation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RamResourceAssociation, error)
	RamResourceAssociationNamespaceListerExpansion
}

// ramResourceAssociationNamespaceLister implements the RamResourceAssociationNamespaceLister
// interface.
type ramResourceAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RamResourceAssociations in the indexer for a given namespace.
func (s ramResourceAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RamResourceAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RamResourceAssociation))
	})
	return ret, err
}

// Get retrieves the RamResourceAssociation from the indexer for a given namespace and name.
func (s ramResourceAssociationNamespaceLister) Get(name string) (*v1alpha1.RamResourceAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ramresourceassociation"), name)
	}
	return obj.(*v1alpha1.RamResourceAssociation), nil
}