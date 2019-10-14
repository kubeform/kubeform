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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// MediaStoreContainerLister helps list MediaStoreContainers.
type MediaStoreContainerLister interface {
	// List lists all MediaStoreContainers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MediaStoreContainer, err error)
	// MediaStoreContainers returns an object that can list and get MediaStoreContainers.
	MediaStoreContainers(namespace string) MediaStoreContainerNamespaceLister
	MediaStoreContainerListerExpansion
}

// mediaStoreContainerLister implements the MediaStoreContainerLister interface.
type mediaStoreContainerLister struct {
	indexer cache.Indexer
}

// NewMediaStoreContainerLister returns a new MediaStoreContainerLister.
func NewMediaStoreContainerLister(indexer cache.Indexer) MediaStoreContainerLister {
	return &mediaStoreContainerLister{indexer: indexer}
}

// List lists all MediaStoreContainers in the indexer.
func (s *mediaStoreContainerLister) List(selector labels.Selector) (ret []*v1alpha1.MediaStoreContainer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MediaStoreContainer))
	})
	return ret, err
}

// MediaStoreContainers returns an object that can list and get MediaStoreContainers.
func (s *mediaStoreContainerLister) MediaStoreContainers(namespace string) MediaStoreContainerNamespaceLister {
	return mediaStoreContainerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MediaStoreContainerNamespaceLister helps list and get MediaStoreContainers.
type MediaStoreContainerNamespaceLister interface {
	// List lists all MediaStoreContainers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MediaStoreContainer, err error)
	// Get retrieves the MediaStoreContainer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MediaStoreContainer, error)
	MediaStoreContainerNamespaceListerExpansion
}

// mediaStoreContainerNamespaceLister implements the MediaStoreContainerNamespaceLister
// interface.
type mediaStoreContainerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MediaStoreContainers in the indexer for a given namespace.
func (s mediaStoreContainerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MediaStoreContainer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MediaStoreContainer))
	})
	return ret, err
}

// Get retrieves the MediaStoreContainer from the indexer for a given namespace and name.
func (s mediaStoreContainerNamespaceLister) Get(name string) (*v1alpha1.MediaStoreContainer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mediastorecontainer"), name)
	}
	return obj.(*v1alpha1.MediaStoreContainer), nil
}
