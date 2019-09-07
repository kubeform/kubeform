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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// IothubSharedAccessPolicyLister helps list IothubSharedAccessPolicies.
type IothubSharedAccessPolicyLister interface {
	// List lists all IothubSharedAccessPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IothubSharedAccessPolicy, err error)
	// IothubSharedAccessPolicies returns an object that can list and get IothubSharedAccessPolicies.
	IothubSharedAccessPolicies(namespace string) IothubSharedAccessPolicyNamespaceLister
	IothubSharedAccessPolicyListerExpansion
}

// iothubSharedAccessPolicyLister implements the IothubSharedAccessPolicyLister interface.
type iothubSharedAccessPolicyLister struct {
	indexer cache.Indexer
}

// NewIothubSharedAccessPolicyLister returns a new IothubSharedAccessPolicyLister.
func NewIothubSharedAccessPolicyLister(indexer cache.Indexer) IothubSharedAccessPolicyLister {
	return &iothubSharedAccessPolicyLister{indexer: indexer}
}

// List lists all IothubSharedAccessPolicies in the indexer.
func (s *iothubSharedAccessPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.IothubSharedAccessPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IothubSharedAccessPolicy))
	})
	return ret, err
}

// IothubSharedAccessPolicies returns an object that can list and get IothubSharedAccessPolicies.
func (s *iothubSharedAccessPolicyLister) IothubSharedAccessPolicies(namespace string) IothubSharedAccessPolicyNamespaceLister {
	return iothubSharedAccessPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IothubSharedAccessPolicyNamespaceLister helps list and get IothubSharedAccessPolicies.
type IothubSharedAccessPolicyNamespaceLister interface {
	// List lists all IothubSharedAccessPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IothubSharedAccessPolicy, err error)
	// Get retrieves the IothubSharedAccessPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IothubSharedAccessPolicy, error)
	IothubSharedAccessPolicyNamespaceListerExpansion
}

// iothubSharedAccessPolicyNamespaceLister implements the IothubSharedAccessPolicyNamespaceLister
// interface.
type iothubSharedAccessPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IothubSharedAccessPolicies in the indexer for a given namespace.
func (s iothubSharedAccessPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IothubSharedAccessPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IothubSharedAccessPolicy))
	})
	return ret, err
}

// Get retrieves the IothubSharedAccessPolicy from the indexer for a given namespace and name.
func (s iothubSharedAccessPolicyNamespaceLister) Get(name string) (*v1alpha1.IothubSharedAccessPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iothubsharedaccesspolicy"), name)
	}
	return obj.(*v1alpha1.IothubSharedAccessPolicy), nil
}