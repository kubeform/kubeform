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

// BatchComputeEnvironmentLister helps list BatchComputeEnvironments.
type BatchComputeEnvironmentLister interface {
	// List lists all BatchComputeEnvironments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.BatchComputeEnvironment, err error)
	// BatchComputeEnvironments returns an object that can list and get BatchComputeEnvironments.
	BatchComputeEnvironments(namespace string) BatchComputeEnvironmentNamespaceLister
	BatchComputeEnvironmentListerExpansion
}

// batchComputeEnvironmentLister implements the BatchComputeEnvironmentLister interface.
type batchComputeEnvironmentLister struct {
	indexer cache.Indexer
}

// NewBatchComputeEnvironmentLister returns a new BatchComputeEnvironmentLister.
func NewBatchComputeEnvironmentLister(indexer cache.Indexer) BatchComputeEnvironmentLister {
	return &batchComputeEnvironmentLister{indexer: indexer}
}

// List lists all BatchComputeEnvironments in the indexer.
func (s *batchComputeEnvironmentLister) List(selector labels.Selector) (ret []*v1alpha1.BatchComputeEnvironment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BatchComputeEnvironment))
	})
	return ret, err
}

// BatchComputeEnvironments returns an object that can list and get BatchComputeEnvironments.
func (s *batchComputeEnvironmentLister) BatchComputeEnvironments(namespace string) BatchComputeEnvironmentNamespaceLister {
	return batchComputeEnvironmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BatchComputeEnvironmentNamespaceLister helps list and get BatchComputeEnvironments.
type BatchComputeEnvironmentNamespaceLister interface {
	// List lists all BatchComputeEnvironments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.BatchComputeEnvironment, err error)
	// Get retrieves the BatchComputeEnvironment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.BatchComputeEnvironment, error)
	BatchComputeEnvironmentNamespaceListerExpansion
}

// batchComputeEnvironmentNamespaceLister implements the BatchComputeEnvironmentNamespaceLister
// interface.
type batchComputeEnvironmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BatchComputeEnvironments in the indexer for a given namespace.
func (s batchComputeEnvironmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BatchComputeEnvironment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BatchComputeEnvironment))
	})
	return ret, err
}

// Get retrieves the BatchComputeEnvironment from the indexer for a given namespace and name.
func (s batchComputeEnvironmentNamespaceLister) Get(name string) (*v1alpha1.BatchComputeEnvironment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("batchcomputeenvironment"), name)
	}
	return obj.(*v1alpha1.BatchComputeEnvironment), nil
}
