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

// AthenaDatabaseLister helps list AthenaDatabases.
type AthenaDatabaseLister interface {
	// List lists all AthenaDatabases in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AthenaDatabase, err error)
	// AthenaDatabases returns an object that can list and get AthenaDatabases.
	AthenaDatabases(namespace string) AthenaDatabaseNamespaceLister
	AthenaDatabaseListerExpansion
}

// athenaDatabaseLister implements the AthenaDatabaseLister interface.
type athenaDatabaseLister struct {
	indexer cache.Indexer
}

// NewAthenaDatabaseLister returns a new AthenaDatabaseLister.
func NewAthenaDatabaseLister(indexer cache.Indexer) AthenaDatabaseLister {
	return &athenaDatabaseLister{indexer: indexer}
}

// List lists all AthenaDatabases in the indexer.
func (s *athenaDatabaseLister) List(selector labels.Selector) (ret []*v1alpha1.AthenaDatabase, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AthenaDatabase))
	})
	return ret, err
}

// AthenaDatabases returns an object that can list and get AthenaDatabases.
func (s *athenaDatabaseLister) AthenaDatabases(namespace string) AthenaDatabaseNamespaceLister {
	return athenaDatabaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AthenaDatabaseNamespaceLister helps list and get AthenaDatabases.
type AthenaDatabaseNamespaceLister interface {
	// List lists all AthenaDatabases in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AthenaDatabase, err error)
	// Get retrieves the AthenaDatabase from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AthenaDatabase, error)
	AthenaDatabaseNamespaceListerExpansion
}

// athenaDatabaseNamespaceLister implements the AthenaDatabaseNamespaceLister
// interface.
type athenaDatabaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AthenaDatabases in the indexer for a given namespace.
func (s athenaDatabaseNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AthenaDatabase, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AthenaDatabase))
	})
	return ret, err
}

// Get retrieves the AthenaDatabase from the indexer for a given namespace and name.
func (s athenaDatabaseNamespaceLister) Get(name string) (*v1alpha1.AthenaDatabase, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("athenadatabase"), name)
	}
	return obj.(*v1alpha1.AthenaDatabase), nil
}
