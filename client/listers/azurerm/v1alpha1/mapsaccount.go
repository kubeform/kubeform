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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MapsAccountLister helps list MapsAccounts.
type MapsAccountLister interface {
	// List lists all MapsAccounts in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MapsAccount, err error)
	// MapsAccounts returns an object that can list and get MapsAccounts.
	MapsAccounts(namespace string) MapsAccountNamespaceLister
	MapsAccountListerExpansion
}

// mapsAccountLister implements the MapsAccountLister interface.
type mapsAccountLister struct {
	indexer cache.Indexer
}

// NewMapsAccountLister returns a new MapsAccountLister.
func NewMapsAccountLister(indexer cache.Indexer) MapsAccountLister {
	return &mapsAccountLister{indexer: indexer}
}

// List lists all MapsAccounts in the indexer.
func (s *mapsAccountLister) List(selector labels.Selector) (ret []*v1alpha1.MapsAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MapsAccount))
	})
	return ret, err
}

// MapsAccounts returns an object that can list and get MapsAccounts.
func (s *mapsAccountLister) MapsAccounts(namespace string) MapsAccountNamespaceLister {
	return mapsAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MapsAccountNamespaceLister helps list and get MapsAccounts.
type MapsAccountNamespaceLister interface {
	// List lists all MapsAccounts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MapsAccount, err error)
	// Get retrieves the MapsAccount from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MapsAccount, error)
	MapsAccountNamespaceListerExpansion
}

// mapsAccountNamespaceLister implements the MapsAccountNamespaceLister
// interface.
type mapsAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MapsAccounts in the indexer for a given namespace.
func (s mapsAccountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MapsAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MapsAccount))
	})
	return ret, err
}

// Get retrieves the MapsAccount from the indexer for a given namespace and name.
func (s mapsAccountNamespaceLister) Get(name string) (*v1alpha1.MapsAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mapsaccount"), name)
	}
	return obj.(*v1alpha1.MapsAccount), nil
}
