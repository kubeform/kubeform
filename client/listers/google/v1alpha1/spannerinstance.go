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

// SpannerInstanceLister helps list SpannerInstances.
type SpannerInstanceLister interface {
	// List lists all SpannerInstances in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SpannerInstance, err error)
	// SpannerInstances returns an object that can list and get SpannerInstances.
	SpannerInstances(namespace string) SpannerInstanceNamespaceLister
	SpannerInstanceListerExpansion
}

// spannerInstanceLister implements the SpannerInstanceLister interface.
type spannerInstanceLister struct {
	indexer cache.Indexer
}

// NewSpannerInstanceLister returns a new SpannerInstanceLister.
func NewSpannerInstanceLister(indexer cache.Indexer) SpannerInstanceLister {
	return &spannerInstanceLister{indexer: indexer}
}

// List lists all SpannerInstances in the indexer.
func (s *spannerInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.SpannerInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SpannerInstance))
	})
	return ret, err
}

// SpannerInstances returns an object that can list and get SpannerInstances.
func (s *spannerInstanceLister) SpannerInstances(namespace string) SpannerInstanceNamespaceLister {
	return spannerInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SpannerInstanceNamespaceLister helps list and get SpannerInstances.
type SpannerInstanceNamespaceLister interface {
	// List lists all SpannerInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SpannerInstance, err error)
	// Get retrieves the SpannerInstance from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SpannerInstance, error)
	SpannerInstanceNamespaceListerExpansion
}

// spannerInstanceNamespaceLister implements the SpannerInstanceNamespaceLister
// interface.
type spannerInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SpannerInstances in the indexer for a given namespace.
func (s spannerInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SpannerInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SpannerInstance))
	})
	return ret, err
}

// Get retrieves the SpannerInstance from the indexer for a given namespace and name.
func (s spannerInstanceNamespaceLister) Get(name string) (*v1alpha1.SpannerInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("spannerinstance"), name)
	}
	return obj.(*v1alpha1.SpannerInstance), nil
}
