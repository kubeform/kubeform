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

// CodecommitTriggerLister helps list CodecommitTriggers.
type CodecommitTriggerLister interface {
	// List lists all CodecommitTriggers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CodecommitTrigger, err error)
	// CodecommitTriggers returns an object that can list and get CodecommitTriggers.
	CodecommitTriggers(namespace string) CodecommitTriggerNamespaceLister
	CodecommitTriggerListerExpansion
}

// codecommitTriggerLister implements the CodecommitTriggerLister interface.
type codecommitTriggerLister struct {
	indexer cache.Indexer
}

// NewCodecommitTriggerLister returns a new CodecommitTriggerLister.
func NewCodecommitTriggerLister(indexer cache.Indexer) CodecommitTriggerLister {
	return &codecommitTriggerLister{indexer: indexer}
}

// List lists all CodecommitTriggers in the indexer.
func (s *codecommitTriggerLister) List(selector labels.Selector) (ret []*v1alpha1.CodecommitTrigger, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CodecommitTrigger))
	})
	return ret, err
}

// CodecommitTriggers returns an object that can list and get CodecommitTriggers.
func (s *codecommitTriggerLister) CodecommitTriggers(namespace string) CodecommitTriggerNamespaceLister {
	return codecommitTriggerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CodecommitTriggerNamespaceLister helps list and get CodecommitTriggers.
type CodecommitTriggerNamespaceLister interface {
	// List lists all CodecommitTriggers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CodecommitTrigger, err error)
	// Get retrieves the CodecommitTrigger from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CodecommitTrigger, error)
	CodecommitTriggerNamespaceListerExpansion
}

// codecommitTriggerNamespaceLister implements the CodecommitTriggerNamespaceLister
// interface.
type codecommitTriggerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CodecommitTriggers in the indexer for a given namespace.
func (s codecommitTriggerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CodecommitTrigger, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CodecommitTrigger))
	})
	return ret, err
}

// Get retrieves the CodecommitTrigger from the indexer for a given namespace and name.
func (s codecommitTriggerNamespaceLister) Get(name string) (*v1alpha1.CodecommitTrigger, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("codecommittrigger"), name)
	}
	return obj.(*v1alpha1.CodecommitTrigger), nil
}