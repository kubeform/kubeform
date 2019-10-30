/*
Copyright The Kubeform Authors.

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

// CodecommitRepositoryLister helps list CodecommitRepositories.
type CodecommitRepositoryLister interface {
	// List lists all CodecommitRepositories in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CodecommitRepository, err error)
	// CodecommitRepositories returns an object that can list and get CodecommitRepositories.
	CodecommitRepositories(namespace string) CodecommitRepositoryNamespaceLister
	CodecommitRepositoryListerExpansion
}

// codecommitRepositoryLister implements the CodecommitRepositoryLister interface.
type codecommitRepositoryLister struct {
	indexer cache.Indexer
}

// NewCodecommitRepositoryLister returns a new CodecommitRepositoryLister.
func NewCodecommitRepositoryLister(indexer cache.Indexer) CodecommitRepositoryLister {
	return &codecommitRepositoryLister{indexer: indexer}
}

// List lists all CodecommitRepositories in the indexer.
func (s *codecommitRepositoryLister) List(selector labels.Selector) (ret []*v1alpha1.CodecommitRepository, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CodecommitRepository))
	})
	return ret, err
}

// CodecommitRepositories returns an object that can list and get CodecommitRepositories.
func (s *codecommitRepositoryLister) CodecommitRepositories(namespace string) CodecommitRepositoryNamespaceLister {
	return codecommitRepositoryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CodecommitRepositoryNamespaceLister helps list and get CodecommitRepositories.
type CodecommitRepositoryNamespaceLister interface {
	// List lists all CodecommitRepositories in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CodecommitRepository, err error)
	// Get retrieves the CodecommitRepository from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CodecommitRepository, error)
	CodecommitRepositoryNamespaceListerExpansion
}

// codecommitRepositoryNamespaceLister implements the CodecommitRepositoryNamespaceLister
// interface.
type codecommitRepositoryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CodecommitRepositories in the indexer for a given namespace.
func (s codecommitRepositoryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CodecommitRepository, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CodecommitRepository))
	})
	return ret, err
}

// Get retrieves the CodecommitRepository from the indexer for a given namespace and name.
func (s codecommitRepositoryNamespaceLister) Get(name string) (*v1alpha1.CodecommitRepository, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("codecommitrepository"), name)
	}
	return obj.(*v1alpha1.CodecommitRepository), nil
}
