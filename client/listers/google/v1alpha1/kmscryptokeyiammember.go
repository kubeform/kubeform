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

// KmsCryptoKeyIamMemberLister helps list KmsCryptoKeyIamMembers.
type KmsCryptoKeyIamMemberLister interface {
	// List lists all KmsCryptoKeyIamMembers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KmsCryptoKeyIamMember, err error)
	// KmsCryptoKeyIamMembers returns an object that can list and get KmsCryptoKeyIamMembers.
	KmsCryptoKeyIamMembers(namespace string) KmsCryptoKeyIamMemberNamespaceLister
	KmsCryptoKeyIamMemberListerExpansion
}

// kmsCryptoKeyIamMemberLister implements the KmsCryptoKeyIamMemberLister interface.
type kmsCryptoKeyIamMemberLister struct {
	indexer cache.Indexer
}

// NewKmsCryptoKeyIamMemberLister returns a new KmsCryptoKeyIamMemberLister.
func NewKmsCryptoKeyIamMemberLister(indexer cache.Indexer) KmsCryptoKeyIamMemberLister {
	return &kmsCryptoKeyIamMemberLister{indexer: indexer}
}

// List lists all KmsCryptoKeyIamMembers in the indexer.
func (s *kmsCryptoKeyIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.KmsCryptoKeyIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KmsCryptoKeyIamMember))
	})
	return ret, err
}

// KmsCryptoKeyIamMembers returns an object that can list and get KmsCryptoKeyIamMembers.
func (s *kmsCryptoKeyIamMemberLister) KmsCryptoKeyIamMembers(namespace string) KmsCryptoKeyIamMemberNamespaceLister {
	return kmsCryptoKeyIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KmsCryptoKeyIamMemberNamespaceLister helps list and get KmsCryptoKeyIamMembers.
type KmsCryptoKeyIamMemberNamespaceLister interface {
	// List lists all KmsCryptoKeyIamMembers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.KmsCryptoKeyIamMember, err error)
	// Get retrieves the KmsCryptoKeyIamMember from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.KmsCryptoKeyIamMember, error)
	KmsCryptoKeyIamMemberNamespaceListerExpansion
}

// kmsCryptoKeyIamMemberNamespaceLister implements the KmsCryptoKeyIamMemberNamespaceLister
// interface.
type kmsCryptoKeyIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KmsCryptoKeyIamMembers in the indexer for a given namespace.
func (s kmsCryptoKeyIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KmsCryptoKeyIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KmsCryptoKeyIamMember))
	})
	return ret, err
}

// Get retrieves the KmsCryptoKeyIamMember from the indexer for a given namespace and name.
func (s kmsCryptoKeyIamMemberNamespaceLister) Get(name string) (*v1alpha1.KmsCryptoKeyIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kmscryptokeyiammember"), name)
	}
	return obj.(*v1alpha1.KmsCryptoKeyIamMember), nil
}