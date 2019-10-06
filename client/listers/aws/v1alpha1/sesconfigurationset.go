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

// SesConfigurationSetLister helps list SesConfigurationSets.
type SesConfigurationSetLister interface {
	// List lists all SesConfigurationSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SesConfigurationSet, err error)
	// SesConfigurationSets returns an object that can list and get SesConfigurationSets.
	SesConfigurationSets(namespace string) SesConfigurationSetNamespaceLister
	SesConfigurationSetListerExpansion
}

// sesConfigurationSetLister implements the SesConfigurationSetLister interface.
type sesConfigurationSetLister struct {
	indexer cache.Indexer
}

// NewSesConfigurationSetLister returns a new SesConfigurationSetLister.
func NewSesConfigurationSetLister(indexer cache.Indexer) SesConfigurationSetLister {
	return &sesConfigurationSetLister{indexer: indexer}
}

// List lists all SesConfigurationSets in the indexer.
func (s *sesConfigurationSetLister) List(selector labels.Selector) (ret []*v1alpha1.SesConfigurationSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SesConfigurationSet))
	})
	return ret, err
}

// SesConfigurationSets returns an object that can list and get SesConfigurationSets.
func (s *sesConfigurationSetLister) SesConfigurationSets(namespace string) SesConfigurationSetNamespaceLister {
	return sesConfigurationSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SesConfigurationSetNamespaceLister helps list and get SesConfigurationSets.
type SesConfigurationSetNamespaceLister interface {
	// List lists all SesConfigurationSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SesConfigurationSet, err error)
	// Get retrieves the SesConfigurationSet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SesConfigurationSet, error)
	SesConfigurationSetNamespaceListerExpansion
}

// sesConfigurationSetNamespaceLister implements the SesConfigurationSetNamespaceLister
// interface.
type sesConfigurationSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SesConfigurationSets in the indexer for a given namespace.
func (s sesConfigurationSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SesConfigurationSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SesConfigurationSet))
	})
	return ret, err
}

// Get retrieves the SesConfigurationSet from the indexer for a given namespace and name.
func (s sesConfigurationSetNamespaceLister) Get(name string) (*v1alpha1.SesConfigurationSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sesconfigurationset"), name)
	}
	return obj.(*v1alpha1.SesConfigurationSet), nil
}
