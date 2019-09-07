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

// StorageBucketIamMemberLister helps list StorageBucketIamMembers.
type StorageBucketIamMemberLister interface {
	// List lists all StorageBucketIamMembers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StorageBucketIamMember, err error)
	// StorageBucketIamMembers returns an object that can list and get StorageBucketIamMembers.
	StorageBucketIamMembers(namespace string) StorageBucketIamMemberNamespaceLister
	StorageBucketIamMemberListerExpansion
}

// storageBucketIamMemberLister implements the StorageBucketIamMemberLister interface.
type storageBucketIamMemberLister struct {
	indexer cache.Indexer
}

// NewStorageBucketIamMemberLister returns a new StorageBucketIamMemberLister.
func NewStorageBucketIamMemberLister(indexer cache.Indexer) StorageBucketIamMemberLister {
	return &storageBucketIamMemberLister{indexer: indexer}
}

// List lists all StorageBucketIamMembers in the indexer.
func (s *storageBucketIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.StorageBucketIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageBucketIamMember))
	})
	return ret, err
}

// StorageBucketIamMembers returns an object that can list and get StorageBucketIamMembers.
func (s *storageBucketIamMemberLister) StorageBucketIamMembers(namespace string) StorageBucketIamMemberNamespaceLister {
	return storageBucketIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StorageBucketIamMemberNamespaceLister helps list and get StorageBucketIamMembers.
type StorageBucketIamMemberNamespaceLister interface {
	// List lists all StorageBucketIamMembers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StorageBucketIamMember, err error)
	// Get retrieves the StorageBucketIamMember from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StorageBucketIamMember, error)
	StorageBucketIamMemberNamespaceListerExpansion
}

// storageBucketIamMemberNamespaceLister implements the StorageBucketIamMemberNamespaceLister
// interface.
type storageBucketIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StorageBucketIamMembers in the indexer for a given namespace.
func (s storageBucketIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StorageBucketIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageBucketIamMember))
	})
	return ret, err
}

// Get retrieves the StorageBucketIamMember from the indexer for a given namespace and name.
func (s storageBucketIamMemberNamespaceLister) Get(name string) (*v1alpha1.StorageBucketIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storagebucketiammember"), name)
	}
	return obj.(*v1alpha1.StorageBucketIamMember), nil
}