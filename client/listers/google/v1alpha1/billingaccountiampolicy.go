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

// BillingAccountIamPolicyLister helps list BillingAccountIamPolicies.
type BillingAccountIamPolicyLister interface {
	// List lists all BillingAccountIamPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.BillingAccountIamPolicy, err error)
	// BillingAccountIamPolicies returns an object that can list and get BillingAccountIamPolicies.
	BillingAccountIamPolicies(namespace string) BillingAccountIamPolicyNamespaceLister
	BillingAccountIamPolicyListerExpansion
}

// billingAccountIamPolicyLister implements the BillingAccountIamPolicyLister interface.
type billingAccountIamPolicyLister struct {
	indexer cache.Indexer
}

// NewBillingAccountIamPolicyLister returns a new BillingAccountIamPolicyLister.
func NewBillingAccountIamPolicyLister(indexer cache.Indexer) BillingAccountIamPolicyLister {
	return &billingAccountIamPolicyLister{indexer: indexer}
}

// List lists all BillingAccountIamPolicies in the indexer.
func (s *billingAccountIamPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.BillingAccountIamPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BillingAccountIamPolicy))
	})
	return ret, err
}

// BillingAccountIamPolicies returns an object that can list and get BillingAccountIamPolicies.
func (s *billingAccountIamPolicyLister) BillingAccountIamPolicies(namespace string) BillingAccountIamPolicyNamespaceLister {
	return billingAccountIamPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BillingAccountIamPolicyNamespaceLister helps list and get BillingAccountIamPolicies.
type BillingAccountIamPolicyNamespaceLister interface {
	// List lists all BillingAccountIamPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.BillingAccountIamPolicy, err error)
	// Get retrieves the BillingAccountIamPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.BillingAccountIamPolicy, error)
	BillingAccountIamPolicyNamespaceListerExpansion
}

// billingAccountIamPolicyNamespaceLister implements the BillingAccountIamPolicyNamespaceLister
// interface.
type billingAccountIamPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BillingAccountIamPolicies in the indexer for a given namespace.
func (s billingAccountIamPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BillingAccountIamPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BillingAccountIamPolicy))
	})
	return ret, err
}

// Get retrieves the BillingAccountIamPolicy from the indexer for a given namespace and name.
func (s billingAccountIamPolicyNamespaceLister) Get(name string) (*v1alpha1.BillingAccountIamPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("billingaccountiampolicy"), name)
	}
	return obj.(*v1alpha1.BillingAccountIamPolicy), nil
}
