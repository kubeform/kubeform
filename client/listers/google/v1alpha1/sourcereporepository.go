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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SourcerepoRepositoryLister helps list SourcerepoRepositories.
type SourcerepoRepositoryLister interface {
	// List lists all SourcerepoRepositories in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SourcerepoRepository, err error)
	// SourcerepoRepositories returns an object that can list and get SourcerepoRepositories.
	SourcerepoRepositories(namespace string) SourcerepoRepositoryNamespaceLister
	SourcerepoRepositoryListerExpansion
}

// sourcerepoRepositoryLister implements the SourcerepoRepositoryLister interface.
type sourcerepoRepositoryLister struct {
	indexer cache.Indexer
}

// NewSourcerepoRepositoryLister returns a new SourcerepoRepositoryLister.
func NewSourcerepoRepositoryLister(indexer cache.Indexer) SourcerepoRepositoryLister {
	return &sourcerepoRepositoryLister{indexer: indexer}
}

// List lists all SourcerepoRepositories in the indexer.
func (s *sourcerepoRepositoryLister) List(selector labels.Selector) (ret []*v1alpha1.SourcerepoRepository, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SourcerepoRepository))
	})
	return ret, err
}

// SourcerepoRepositories returns an object that can list and get SourcerepoRepositories.
func (s *sourcerepoRepositoryLister) SourcerepoRepositories(namespace string) SourcerepoRepositoryNamespaceLister {
	return sourcerepoRepositoryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SourcerepoRepositoryNamespaceLister helps list and get SourcerepoRepositories.
type SourcerepoRepositoryNamespaceLister interface {
	// List lists all SourcerepoRepositories in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SourcerepoRepository, err error)
	// Get retrieves the SourcerepoRepository from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SourcerepoRepository, error)
	SourcerepoRepositoryNamespaceListerExpansion
}

// sourcerepoRepositoryNamespaceLister implements the SourcerepoRepositoryNamespaceLister
// interface.
type sourcerepoRepositoryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SourcerepoRepositories in the indexer for a given namespace.
func (s sourcerepoRepositoryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SourcerepoRepository, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SourcerepoRepository))
	})
	return ret, err
}

// Get retrieves the SourcerepoRepository from the indexer for a given namespace and name.
func (s sourcerepoRepositoryNamespaceLister) Get(name string) (*v1alpha1.SourcerepoRepository, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sourcereporepository"), name)
	}
	return obj.(*v1alpha1.SourcerepoRepository), nil
}
