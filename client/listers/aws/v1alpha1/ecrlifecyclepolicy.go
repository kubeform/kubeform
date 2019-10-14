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

// EcrLifecyclePolicyLister helps list EcrLifecyclePolicies.
type EcrLifecyclePolicyLister interface {
	// List lists all EcrLifecyclePolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.EcrLifecyclePolicy, err error)
	// EcrLifecyclePolicies returns an object that can list and get EcrLifecyclePolicies.
	EcrLifecyclePolicies(namespace string) EcrLifecyclePolicyNamespaceLister
	EcrLifecyclePolicyListerExpansion
}

// ecrLifecyclePolicyLister implements the EcrLifecyclePolicyLister interface.
type ecrLifecyclePolicyLister struct {
	indexer cache.Indexer
}

// NewEcrLifecyclePolicyLister returns a new EcrLifecyclePolicyLister.
func NewEcrLifecyclePolicyLister(indexer cache.Indexer) EcrLifecyclePolicyLister {
	return &ecrLifecyclePolicyLister{indexer: indexer}
}

// List lists all EcrLifecyclePolicies in the indexer.
func (s *ecrLifecyclePolicyLister) List(selector labels.Selector) (ret []*v1alpha1.EcrLifecyclePolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EcrLifecyclePolicy))
	})
	return ret, err
}

// EcrLifecyclePolicies returns an object that can list and get EcrLifecyclePolicies.
func (s *ecrLifecyclePolicyLister) EcrLifecyclePolicies(namespace string) EcrLifecyclePolicyNamespaceLister {
	return ecrLifecyclePolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EcrLifecyclePolicyNamespaceLister helps list and get EcrLifecyclePolicies.
type EcrLifecyclePolicyNamespaceLister interface {
	// List lists all EcrLifecyclePolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.EcrLifecyclePolicy, err error)
	// Get retrieves the EcrLifecyclePolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.EcrLifecyclePolicy, error)
	EcrLifecyclePolicyNamespaceListerExpansion
}

// ecrLifecyclePolicyNamespaceLister implements the EcrLifecyclePolicyNamespaceLister
// interface.
type ecrLifecyclePolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EcrLifecyclePolicies in the indexer for a given namespace.
func (s ecrLifecyclePolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EcrLifecyclePolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EcrLifecyclePolicy))
	})
	return ret, err
}

// Get retrieves the EcrLifecyclePolicy from the indexer for a given namespace and name.
func (s ecrLifecyclePolicyNamespaceLister) Get(name string) (*v1alpha1.EcrLifecyclePolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ecrlifecyclepolicy"), name)
	}
	return obj.(*v1alpha1.EcrLifecyclePolicy), nil
}
