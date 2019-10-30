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

// BillingAccountIamMemberLister helps list BillingAccountIamMembers.
type BillingAccountIamMemberLister interface {
	// List lists all BillingAccountIamMembers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.BillingAccountIamMember, err error)
	// BillingAccountIamMembers returns an object that can list and get BillingAccountIamMembers.
	BillingAccountIamMembers(namespace string) BillingAccountIamMemberNamespaceLister
	BillingAccountIamMemberListerExpansion
}

// billingAccountIamMemberLister implements the BillingAccountIamMemberLister interface.
type billingAccountIamMemberLister struct {
	indexer cache.Indexer
}

// NewBillingAccountIamMemberLister returns a new BillingAccountIamMemberLister.
func NewBillingAccountIamMemberLister(indexer cache.Indexer) BillingAccountIamMemberLister {
	return &billingAccountIamMemberLister{indexer: indexer}
}

// List lists all BillingAccountIamMembers in the indexer.
func (s *billingAccountIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.BillingAccountIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BillingAccountIamMember))
	})
	return ret, err
}

// BillingAccountIamMembers returns an object that can list and get BillingAccountIamMembers.
func (s *billingAccountIamMemberLister) BillingAccountIamMembers(namespace string) BillingAccountIamMemberNamespaceLister {
	return billingAccountIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BillingAccountIamMemberNamespaceLister helps list and get BillingAccountIamMembers.
type BillingAccountIamMemberNamespaceLister interface {
	// List lists all BillingAccountIamMembers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.BillingAccountIamMember, err error)
	// Get retrieves the BillingAccountIamMember from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.BillingAccountIamMember, error)
	BillingAccountIamMemberNamespaceListerExpansion
}

// billingAccountIamMemberNamespaceLister implements the BillingAccountIamMemberNamespaceLister
// interface.
type billingAccountIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BillingAccountIamMembers in the indexer for a given namespace.
func (s billingAccountIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BillingAccountIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BillingAccountIamMember))
	})
	return ret, err
}

// Get retrieves the BillingAccountIamMember from the indexer for a given namespace and name.
func (s billingAccountIamMemberNamespaceLister) Get(name string) (*v1alpha1.BillingAccountIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("billingaccountiammember"), name)
	}
	return obj.(*v1alpha1.BillingAccountIamMember), nil
}
