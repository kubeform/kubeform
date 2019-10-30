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

// ComputeAttachedDiskLister helps list ComputeAttachedDisks.
type ComputeAttachedDiskLister interface {
	// List lists all ComputeAttachedDisks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeAttachedDisk, err error)
	// ComputeAttachedDisks returns an object that can list and get ComputeAttachedDisks.
	ComputeAttachedDisks(namespace string) ComputeAttachedDiskNamespaceLister
	ComputeAttachedDiskListerExpansion
}

// computeAttachedDiskLister implements the ComputeAttachedDiskLister interface.
type computeAttachedDiskLister struct {
	indexer cache.Indexer
}

// NewComputeAttachedDiskLister returns a new ComputeAttachedDiskLister.
func NewComputeAttachedDiskLister(indexer cache.Indexer) ComputeAttachedDiskLister {
	return &computeAttachedDiskLister{indexer: indexer}
}

// List lists all ComputeAttachedDisks in the indexer.
func (s *computeAttachedDiskLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeAttachedDisk, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeAttachedDisk))
	})
	return ret, err
}

// ComputeAttachedDisks returns an object that can list and get ComputeAttachedDisks.
func (s *computeAttachedDiskLister) ComputeAttachedDisks(namespace string) ComputeAttachedDiskNamespaceLister {
	return computeAttachedDiskNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeAttachedDiskNamespaceLister helps list and get ComputeAttachedDisks.
type ComputeAttachedDiskNamespaceLister interface {
	// List lists all ComputeAttachedDisks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeAttachedDisk, err error)
	// Get retrieves the ComputeAttachedDisk from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeAttachedDisk, error)
	ComputeAttachedDiskNamespaceListerExpansion
}

// computeAttachedDiskNamespaceLister implements the ComputeAttachedDiskNamespaceLister
// interface.
type computeAttachedDiskNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeAttachedDisks in the indexer for a given namespace.
func (s computeAttachedDiskNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeAttachedDisk, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeAttachedDisk))
	})
	return ret, err
}

// Get retrieves the ComputeAttachedDisk from the indexer for a given namespace and name.
func (s computeAttachedDiskNamespaceLister) Get(name string) (*v1alpha1.ComputeAttachedDisk, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computeattacheddisk"), name)
	}
	return obj.(*v1alpha1.ComputeAttachedDisk), nil
}
