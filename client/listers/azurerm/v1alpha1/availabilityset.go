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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// AvailabilitySetLister helps list AvailabilitySets.
type AvailabilitySetLister interface {
	// List lists all AvailabilitySets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AvailabilitySet, err error)
	// AvailabilitySets returns an object that can list and get AvailabilitySets.
	AvailabilitySets(namespace string) AvailabilitySetNamespaceLister
	AvailabilitySetListerExpansion
}

// availabilitySetLister implements the AvailabilitySetLister interface.
type availabilitySetLister struct {
	indexer cache.Indexer
}

// NewAvailabilitySetLister returns a new AvailabilitySetLister.
func NewAvailabilitySetLister(indexer cache.Indexer) AvailabilitySetLister {
	return &availabilitySetLister{indexer: indexer}
}

// List lists all AvailabilitySets in the indexer.
func (s *availabilitySetLister) List(selector labels.Selector) (ret []*v1alpha1.AvailabilitySet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AvailabilitySet))
	})
	return ret, err
}

// AvailabilitySets returns an object that can list and get AvailabilitySets.
func (s *availabilitySetLister) AvailabilitySets(namespace string) AvailabilitySetNamespaceLister {
	return availabilitySetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AvailabilitySetNamespaceLister helps list and get AvailabilitySets.
type AvailabilitySetNamespaceLister interface {
	// List lists all AvailabilitySets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AvailabilitySet, err error)
	// Get retrieves the AvailabilitySet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AvailabilitySet, error)
	AvailabilitySetNamespaceListerExpansion
}

// availabilitySetNamespaceLister implements the AvailabilitySetNamespaceLister
// interface.
type availabilitySetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AvailabilitySets in the indexer for a given namespace.
func (s availabilitySetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AvailabilitySet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AvailabilitySet))
	})
	return ret, err
}

// Get retrieves the AvailabilitySet from the indexer for a given namespace and name.
func (s availabilitySetNamespaceLister) Get(name string) (*v1alpha1.AvailabilitySet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("availabilityset"), name)
	}
	return obj.(*v1alpha1.AvailabilitySet), nil
}
