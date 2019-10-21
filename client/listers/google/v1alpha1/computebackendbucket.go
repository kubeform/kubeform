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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// ComputeBackendBucketLister helps list ComputeBackendBuckets.
type ComputeBackendBucketLister interface {
	// List lists all ComputeBackendBuckets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeBackendBucket, err error)
	// ComputeBackendBuckets returns an object that can list and get ComputeBackendBuckets.
	ComputeBackendBuckets(namespace string) ComputeBackendBucketNamespaceLister
	ComputeBackendBucketListerExpansion
}

// computeBackendBucketLister implements the ComputeBackendBucketLister interface.
type computeBackendBucketLister struct {
	indexer cache.Indexer
}

// NewComputeBackendBucketLister returns a new ComputeBackendBucketLister.
func NewComputeBackendBucketLister(indexer cache.Indexer) ComputeBackendBucketLister {
	return &computeBackendBucketLister{indexer: indexer}
}

// List lists all ComputeBackendBuckets in the indexer.
func (s *computeBackendBucketLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeBackendBucket, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeBackendBucket))
	})
	return ret, err
}

// ComputeBackendBuckets returns an object that can list and get ComputeBackendBuckets.
func (s *computeBackendBucketLister) ComputeBackendBuckets(namespace string) ComputeBackendBucketNamespaceLister {
	return computeBackendBucketNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeBackendBucketNamespaceLister helps list and get ComputeBackendBuckets.
type ComputeBackendBucketNamespaceLister interface {
	// List lists all ComputeBackendBuckets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeBackendBucket, err error)
	// Get retrieves the ComputeBackendBucket from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeBackendBucket, error)
	ComputeBackendBucketNamespaceListerExpansion
}

// computeBackendBucketNamespaceLister implements the ComputeBackendBucketNamespaceLister
// interface.
type computeBackendBucketNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeBackendBuckets in the indexer for a given namespace.
func (s computeBackendBucketNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeBackendBucket, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeBackendBucket))
	})
	return ret, err
}

// Get retrieves the ComputeBackendBucket from the indexer for a given namespace and name.
func (s computeBackendBucketNamespaceLister) Get(name string) (*v1alpha1.ComputeBackendBucket, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computebackendbucket"), name)
	}
	return obj.(*v1alpha1.ComputeBackendBucket), nil
}