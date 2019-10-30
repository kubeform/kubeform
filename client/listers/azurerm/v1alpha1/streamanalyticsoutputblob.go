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

// StreamAnalyticsOutputBlobLister helps list StreamAnalyticsOutputBlobs.
type StreamAnalyticsOutputBlobLister interface {
	// List lists all StreamAnalyticsOutputBlobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsOutputBlob, err error)
	// StreamAnalyticsOutputBlobs returns an object that can list and get StreamAnalyticsOutputBlobs.
	StreamAnalyticsOutputBlobs(namespace string) StreamAnalyticsOutputBlobNamespaceLister
	StreamAnalyticsOutputBlobListerExpansion
}

// streamAnalyticsOutputBlobLister implements the StreamAnalyticsOutputBlobLister interface.
type streamAnalyticsOutputBlobLister struct {
	indexer cache.Indexer
}

// NewStreamAnalyticsOutputBlobLister returns a new StreamAnalyticsOutputBlobLister.
func NewStreamAnalyticsOutputBlobLister(indexer cache.Indexer) StreamAnalyticsOutputBlobLister {
	return &streamAnalyticsOutputBlobLister{indexer: indexer}
}

// List lists all StreamAnalyticsOutputBlobs in the indexer.
func (s *streamAnalyticsOutputBlobLister) List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsOutputBlob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamAnalyticsOutputBlob))
	})
	return ret, err
}

// StreamAnalyticsOutputBlobs returns an object that can list and get StreamAnalyticsOutputBlobs.
func (s *streamAnalyticsOutputBlobLister) StreamAnalyticsOutputBlobs(namespace string) StreamAnalyticsOutputBlobNamespaceLister {
	return streamAnalyticsOutputBlobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StreamAnalyticsOutputBlobNamespaceLister helps list and get StreamAnalyticsOutputBlobs.
type StreamAnalyticsOutputBlobNamespaceLister interface {
	// List lists all StreamAnalyticsOutputBlobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsOutputBlob, err error)
	// Get retrieves the StreamAnalyticsOutputBlob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StreamAnalyticsOutputBlob, error)
	StreamAnalyticsOutputBlobNamespaceListerExpansion
}

// streamAnalyticsOutputBlobNamespaceLister implements the StreamAnalyticsOutputBlobNamespaceLister
// interface.
type streamAnalyticsOutputBlobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StreamAnalyticsOutputBlobs in the indexer for a given namespace.
func (s streamAnalyticsOutputBlobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsOutputBlob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamAnalyticsOutputBlob))
	})
	return ret, err
}

// Get retrieves the StreamAnalyticsOutputBlob from the indexer for a given namespace and name.
func (s streamAnalyticsOutputBlobNamespaceLister) Get(name string) (*v1alpha1.StreamAnalyticsOutputBlob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("streamanalyticsoutputblob"), name)
	}
	return obj.(*v1alpha1.StreamAnalyticsOutputBlob), nil
}
