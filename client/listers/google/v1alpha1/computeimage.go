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

// ComputeImageLister helps list ComputeImages.
type ComputeImageLister interface {
	// List lists all ComputeImages in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeImage, err error)
	// ComputeImages returns an object that can list and get ComputeImages.
	ComputeImages(namespace string) ComputeImageNamespaceLister
	ComputeImageListerExpansion
}

// computeImageLister implements the ComputeImageLister interface.
type computeImageLister struct {
	indexer cache.Indexer
}

// NewComputeImageLister returns a new ComputeImageLister.
func NewComputeImageLister(indexer cache.Indexer) ComputeImageLister {
	return &computeImageLister{indexer: indexer}
}

// List lists all ComputeImages in the indexer.
func (s *computeImageLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeImage, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeImage))
	})
	return ret, err
}

// ComputeImages returns an object that can list and get ComputeImages.
func (s *computeImageLister) ComputeImages(namespace string) ComputeImageNamespaceLister {
	return computeImageNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeImageNamespaceLister helps list and get ComputeImages.
type ComputeImageNamespaceLister interface {
	// List lists all ComputeImages in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeImage, err error)
	// Get retrieves the ComputeImage from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeImage, error)
	ComputeImageNamespaceListerExpansion
}

// computeImageNamespaceLister implements the ComputeImageNamespaceLister
// interface.
type computeImageNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeImages in the indexer for a given namespace.
func (s computeImageNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeImage, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeImage))
	})
	return ret, err
}

// Get retrieves the ComputeImage from the indexer for a given namespace and name.
func (s computeImageNamespaceLister) Get(name string) (*v1alpha1.ComputeImage, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computeimage"), name)
	}
	return obj.(*v1alpha1.ComputeImage), nil
}
