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

// ComputeInstanceLister helps list ComputeInstances.
type ComputeInstanceLister interface {
	// List lists all ComputeInstances in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeInstance, err error)
	// ComputeInstances returns an object that can list and get ComputeInstances.
	ComputeInstances(namespace string) ComputeInstanceNamespaceLister
	ComputeInstanceListerExpansion
}

// computeInstanceLister implements the ComputeInstanceLister interface.
type computeInstanceLister struct {
	indexer cache.Indexer
}

// NewComputeInstanceLister returns a new ComputeInstanceLister.
func NewComputeInstanceLister(indexer cache.Indexer) ComputeInstanceLister {
	return &computeInstanceLister{indexer: indexer}
}

// List lists all ComputeInstances in the indexer.
func (s *computeInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeInstance))
	})
	return ret, err
}

// ComputeInstances returns an object that can list and get ComputeInstances.
func (s *computeInstanceLister) ComputeInstances(namespace string) ComputeInstanceNamespaceLister {
	return computeInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeInstanceNamespaceLister helps list and get ComputeInstances.
type ComputeInstanceNamespaceLister interface {
	// List lists all ComputeInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeInstance, err error)
	// Get retrieves the ComputeInstance from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeInstance, error)
	ComputeInstanceNamespaceListerExpansion
}

// computeInstanceNamespaceLister implements the ComputeInstanceNamespaceLister
// interface.
type computeInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeInstances in the indexer for a given namespace.
func (s computeInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeInstance))
	})
	return ret, err
}

// Get retrieves the ComputeInstance from the indexer for a given namespace and name.
func (s computeInstanceNamespaceLister) Get(name string) (*v1alpha1.ComputeInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computeinstance"), name)
	}
	return obj.(*v1alpha1.ComputeInstance), nil
}
