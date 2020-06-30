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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DataprocJobIamBindingLister helps list DataprocJobIamBindings.
type DataprocJobIamBindingLister interface {
	// List lists all DataprocJobIamBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DataprocJobIamBinding, err error)
	// DataprocJobIamBindings returns an object that can list and get DataprocJobIamBindings.
	DataprocJobIamBindings(namespace string) DataprocJobIamBindingNamespaceLister
	DataprocJobIamBindingListerExpansion
}

// dataprocJobIamBindingLister implements the DataprocJobIamBindingLister interface.
type dataprocJobIamBindingLister struct {
	indexer cache.Indexer
}

// NewDataprocJobIamBindingLister returns a new DataprocJobIamBindingLister.
func NewDataprocJobIamBindingLister(indexer cache.Indexer) DataprocJobIamBindingLister {
	return &dataprocJobIamBindingLister{indexer: indexer}
}

// List lists all DataprocJobIamBindings in the indexer.
func (s *dataprocJobIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.DataprocJobIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataprocJobIamBinding))
	})
	return ret, err
}

// DataprocJobIamBindings returns an object that can list and get DataprocJobIamBindings.
func (s *dataprocJobIamBindingLister) DataprocJobIamBindings(namespace string) DataprocJobIamBindingNamespaceLister {
	return dataprocJobIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataprocJobIamBindingNamespaceLister helps list and get DataprocJobIamBindings.
type DataprocJobIamBindingNamespaceLister interface {
	// List lists all DataprocJobIamBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DataprocJobIamBinding, err error)
	// Get retrieves the DataprocJobIamBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DataprocJobIamBinding, error)
	DataprocJobIamBindingNamespaceListerExpansion
}

// dataprocJobIamBindingNamespaceLister implements the DataprocJobIamBindingNamespaceLister
// interface.
type dataprocJobIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataprocJobIamBindings in the indexer for a given namespace.
func (s dataprocJobIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DataprocJobIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataprocJobIamBinding))
	})
	return ret, err
}

// Get retrieves the DataprocJobIamBinding from the indexer for a given namespace and name.
func (s dataprocJobIamBindingNamespaceLister) Get(name string) (*v1alpha1.DataprocJobIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dataprocjobiambinding"), name)
	}
	return obj.(*v1alpha1.DataprocJobIamBinding), nil
}
