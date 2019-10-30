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

// SpannerInstanceIamBindingLister helps list SpannerInstanceIamBindings.
type SpannerInstanceIamBindingLister interface {
	// List lists all SpannerInstanceIamBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SpannerInstanceIamBinding, err error)
	// SpannerInstanceIamBindings returns an object that can list and get SpannerInstanceIamBindings.
	SpannerInstanceIamBindings(namespace string) SpannerInstanceIamBindingNamespaceLister
	SpannerInstanceIamBindingListerExpansion
}

// spannerInstanceIamBindingLister implements the SpannerInstanceIamBindingLister interface.
type spannerInstanceIamBindingLister struct {
	indexer cache.Indexer
}

// NewSpannerInstanceIamBindingLister returns a new SpannerInstanceIamBindingLister.
func NewSpannerInstanceIamBindingLister(indexer cache.Indexer) SpannerInstanceIamBindingLister {
	return &spannerInstanceIamBindingLister{indexer: indexer}
}

// List lists all SpannerInstanceIamBindings in the indexer.
func (s *spannerInstanceIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.SpannerInstanceIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SpannerInstanceIamBinding))
	})
	return ret, err
}

// SpannerInstanceIamBindings returns an object that can list and get SpannerInstanceIamBindings.
func (s *spannerInstanceIamBindingLister) SpannerInstanceIamBindings(namespace string) SpannerInstanceIamBindingNamespaceLister {
	return spannerInstanceIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SpannerInstanceIamBindingNamespaceLister helps list and get SpannerInstanceIamBindings.
type SpannerInstanceIamBindingNamespaceLister interface {
	// List lists all SpannerInstanceIamBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SpannerInstanceIamBinding, err error)
	// Get retrieves the SpannerInstanceIamBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SpannerInstanceIamBinding, error)
	SpannerInstanceIamBindingNamespaceListerExpansion
}

// spannerInstanceIamBindingNamespaceLister implements the SpannerInstanceIamBindingNamespaceLister
// interface.
type spannerInstanceIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SpannerInstanceIamBindings in the indexer for a given namespace.
func (s spannerInstanceIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SpannerInstanceIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SpannerInstanceIamBinding))
	})
	return ret, err
}

// Get retrieves the SpannerInstanceIamBinding from the indexer for a given namespace and name.
func (s spannerInstanceIamBindingNamespaceLister) Get(name string) (*v1alpha1.SpannerInstanceIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("spannerinstanceiambinding"), name)
	}
	return obj.(*v1alpha1.SpannerInstanceIamBinding), nil
}
