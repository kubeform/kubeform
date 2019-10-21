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

// StoragegatewayCachedIscsiVolumeLister helps list StoragegatewayCachedIscsiVolumes.
type StoragegatewayCachedIscsiVolumeLister interface {
	// List lists all StoragegatewayCachedIscsiVolumes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StoragegatewayCachedIscsiVolume, err error)
	// StoragegatewayCachedIscsiVolumes returns an object that can list and get StoragegatewayCachedIscsiVolumes.
	StoragegatewayCachedIscsiVolumes(namespace string) StoragegatewayCachedIscsiVolumeNamespaceLister
	StoragegatewayCachedIscsiVolumeListerExpansion
}

// storagegatewayCachedIscsiVolumeLister implements the StoragegatewayCachedIscsiVolumeLister interface.
type storagegatewayCachedIscsiVolumeLister struct {
	indexer cache.Indexer
}

// NewStoragegatewayCachedIscsiVolumeLister returns a new StoragegatewayCachedIscsiVolumeLister.
func NewStoragegatewayCachedIscsiVolumeLister(indexer cache.Indexer) StoragegatewayCachedIscsiVolumeLister {
	return &storagegatewayCachedIscsiVolumeLister{indexer: indexer}
}

// List lists all StoragegatewayCachedIscsiVolumes in the indexer.
func (s *storagegatewayCachedIscsiVolumeLister) List(selector labels.Selector) (ret []*v1alpha1.StoragegatewayCachedIscsiVolume, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StoragegatewayCachedIscsiVolume))
	})
	return ret, err
}

// StoragegatewayCachedIscsiVolumes returns an object that can list and get StoragegatewayCachedIscsiVolumes.
func (s *storagegatewayCachedIscsiVolumeLister) StoragegatewayCachedIscsiVolumes(namespace string) StoragegatewayCachedIscsiVolumeNamespaceLister {
	return storagegatewayCachedIscsiVolumeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StoragegatewayCachedIscsiVolumeNamespaceLister helps list and get StoragegatewayCachedIscsiVolumes.
type StoragegatewayCachedIscsiVolumeNamespaceLister interface {
	// List lists all StoragegatewayCachedIscsiVolumes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StoragegatewayCachedIscsiVolume, err error)
	// Get retrieves the StoragegatewayCachedIscsiVolume from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StoragegatewayCachedIscsiVolume, error)
	StoragegatewayCachedIscsiVolumeNamespaceListerExpansion
}

// storagegatewayCachedIscsiVolumeNamespaceLister implements the StoragegatewayCachedIscsiVolumeNamespaceLister
// interface.
type storagegatewayCachedIscsiVolumeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StoragegatewayCachedIscsiVolumes in the indexer for a given namespace.
func (s storagegatewayCachedIscsiVolumeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StoragegatewayCachedIscsiVolume, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StoragegatewayCachedIscsiVolume))
	})
	return ret, err
}

// Get retrieves the StoragegatewayCachedIscsiVolume from the indexer for a given namespace and name.
func (s storagegatewayCachedIscsiVolumeNamespaceLister) Get(name string) (*v1alpha1.StoragegatewayCachedIscsiVolume, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storagegatewaycachediscsivolume"), name)
	}
	return obj.(*v1alpha1.StoragegatewayCachedIscsiVolume), nil
}