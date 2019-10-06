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

// IamRolePolicyLister helps list IamRolePolicies.
type IamRolePolicyLister interface {
	// List lists all IamRolePolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IamRolePolicy, err error)
	// IamRolePolicies returns an object that can list and get IamRolePolicies.
	IamRolePolicies(namespace string) IamRolePolicyNamespaceLister
	IamRolePolicyListerExpansion
}

// iamRolePolicyLister implements the IamRolePolicyLister interface.
type iamRolePolicyLister struct {
	indexer cache.Indexer
}

// NewIamRolePolicyLister returns a new IamRolePolicyLister.
func NewIamRolePolicyLister(indexer cache.Indexer) IamRolePolicyLister {
	return &iamRolePolicyLister{indexer: indexer}
}

// List lists all IamRolePolicies in the indexer.
func (s *iamRolePolicyLister) List(selector labels.Selector) (ret []*v1alpha1.IamRolePolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamRolePolicy))
	})
	return ret, err
}

// IamRolePolicies returns an object that can list and get IamRolePolicies.
func (s *iamRolePolicyLister) IamRolePolicies(namespace string) IamRolePolicyNamespaceLister {
	return iamRolePolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IamRolePolicyNamespaceLister helps list and get IamRolePolicies.
type IamRolePolicyNamespaceLister interface {
	// List lists all IamRolePolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IamRolePolicy, err error)
	// Get retrieves the IamRolePolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IamRolePolicy, error)
	IamRolePolicyNamespaceListerExpansion
}

// iamRolePolicyNamespaceLister implements the IamRolePolicyNamespaceLister
// interface.
type iamRolePolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IamRolePolicies in the indexer for a given namespace.
func (s iamRolePolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IamRolePolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamRolePolicy))
	})
	return ret, err
}

// Get retrieves the IamRolePolicy from the indexer for a given namespace and name.
func (s iamRolePolicyNamespaceLister) Get(name string) (*v1alpha1.IamRolePolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iamrolepolicy"), name)
	}
	return obj.(*v1alpha1.IamRolePolicy), nil
}
