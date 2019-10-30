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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StreamAnalyticsStreamInputEventhubLister helps list StreamAnalyticsStreamInputEventhubs.
type StreamAnalyticsStreamInputEventhubLister interface {
	// List lists all StreamAnalyticsStreamInputEventhubs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsStreamInputEventhub, err error)
	// StreamAnalyticsStreamInputEventhubs returns an object that can list and get StreamAnalyticsStreamInputEventhubs.
	StreamAnalyticsStreamInputEventhubs(namespace string) StreamAnalyticsStreamInputEventhubNamespaceLister
	StreamAnalyticsStreamInputEventhubListerExpansion
}

// streamAnalyticsStreamInputEventhubLister implements the StreamAnalyticsStreamInputEventhubLister interface.
type streamAnalyticsStreamInputEventhubLister struct {
	indexer cache.Indexer
}

// NewStreamAnalyticsStreamInputEventhubLister returns a new StreamAnalyticsStreamInputEventhubLister.
func NewStreamAnalyticsStreamInputEventhubLister(indexer cache.Indexer) StreamAnalyticsStreamInputEventhubLister {
	return &streamAnalyticsStreamInputEventhubLister{indexer: indexer}
}

// List lists all StreamAnalyticsStreamInputEventhubs in the indexer.
func (s *streamAnalyticsStreamInputEventhubLister) List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsStreamInputEventhub, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamAnalyticsStreamInputEventhub))
	})
	return ret, err
}

// StreamAnalyticsStreamInputEventhubs returns an object that can list and get StreamAnalyticsStreamInputEventhubs.
func (s *streamAnalyticsStreamInputEventhubLister) StreamAnalyticsStreamInputEventhubs(namespace string) StreamAnalyticsStreamInputEventhubNamespaceLister {
	return streamAnalyticsStreamInputEventhubNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StreamAnalyticsStreamInputEventhubNamespaceLister helps list and get StreamAnalyticsStreamInputEventhubs.
type StreamAnalyticsStreamInputEventhubNamespaceLister interface {
	// List lists all StreamAnalyticsStreamInputEventhubs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsStreamInputEventhub, err error)
	// Get retrieves the StreamAnalyticsStreamInputEventhub from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StreamAnalyticsStreamInputEventhub, error)
	StreamAnalyticsStreamInputEventhubNamespaceListerExpansion
}

// streamAnalyticsStreamInputEventhubNamespaceLister implements the StreamAnalyticsStreamInputEventhubNamespaceLister
// interface.
type streamAnalyticsStreamInputEventhubNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StreamAnalyticsStreamInputEventhubs in the indexer for a given namespace.
func (s streamAnalyticsStreamInputEventhubNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsStreamInputEventhub, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamAnalyticsStreamInputEventhub))
	})
	return ret, err
}

// Get retrieves the StreamAnalyticsStreamInputEventhub from the indexer for a given namespace and name.
func (s streamAnalyticsStreamInputEventhubNamespaceLister) Get(name string) (*v1alpha1.StreamAnalyticsStreamInputEventhub, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("streamanalyticsstreaminputeventhub"), name)
	}
	return obj.(*v1alpha1.StreamAnalyticsStreamInputEventhub), nil
}
