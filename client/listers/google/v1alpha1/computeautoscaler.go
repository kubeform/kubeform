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

// ComputeAutoscalerLister helps list ComputeAutoscalers.
type ComputeAutoscalerLister interface {
	// List lists all ComputeAutoscalers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeAutoscaler, err error)
	// ComputeAutoscalers returns an object that can list and get ComputeAutoscalers.
	ComputeAutoscalers(namespace string) ComputeAutoscalerNamespaceLister
	ComputeAutoscalerListerExpansion
}

// computeAutoscalerLister implements the ComputeAutoscalerLister interface.
type computeAutoscalerLister struct {
	indexer cache.Indexer
}

// NewComputeAutoscalerLister returns a new ComputeAutoscalerLister.
func NewComputeAutoscalerLister(indexer cache.Indexer) ComputeAutoscalerLister {
	return &computeAutoscalerLister{indexer: indexer}
}

// List lists all ComputeAutoscalers in the indexer.
func (s *computeAutoscalerLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeAutoscaler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeAutoscaler))
	})
	return ret, err
}

// ComputeAutoscalers returns an object that can list and get ComputeAutoscalers.
func (s *computeAutoscalerLister) ComputeAutoscalers(namespace string) ComputeAutoscalerNamespaceLister {
	return computeAutoscalerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeAutoscalerNamespaceLister helps list and get ComputeAutoscalers.
type ComputeAutoscalerNamespaceLister interface {
	// List lists all ComputeAutoscalers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeAutoscaler, err error)
	// Get retrieves the ComputeAutoscaler from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeAutoscaler, error)
	ComputeAutoscalerNamespaceListerExpansion
}

// computeAutoscalerNamespaceLister implements the ComputeAutoscalerNamespaceLister
// interface.
type computeAutoscalerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeAutoscalers in the indexer for a given namespace.
func (s computeAutoscalerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeAutoscaler, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeAutoscaler))
	})
	return ret, err
}

// Get retrieves the ComputeAutoscaler from the indexer for a given namespace and name.
func (s computeAutoscalerNamespaceLister) Get(name string) (*v1alpha1.ComputeAutoscaler, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computeautoscaler"), name)
	}
	return obj.(*v1alpha1.ComputeAutoscaler), nil
}
