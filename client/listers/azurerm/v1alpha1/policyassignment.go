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

// PolicyAssignmentLister helps list PolicyAssignments.
type PolicyAssignmentLister interface {
	// List lists all PolicyAssignments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PolicyAssignment, err error)
	// PolicyAssignments returns an object that can list and get PolicyAssignments.
	PolicyAssignments(namespace string) PolicyAssignmentNamespaceLister
	PolicyAssignmentListerExpansion
}

// policyAssignmentLister implements the PolicyAssignmentLister interface.
type policyAssignmentLister struct {
	indexer cache.Indexer
}

// NewPolicyAssignmentLister returns a new PolicyAssignmentLister.
func NewPolicyAssignmentLister(indexer cache.Indexer) PolicyAssignmentLister {
	return &policyAssignmentLister{indexer: indexer}
}

// List lists all PolicyAssignments in the indexer.
func (s *policyAssignmentLister) List(selector labels.Selector) (ret []*v1alpha1.PolicyAssignment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PolicyAssignment))
	})
	return ret, err
}

// PolicyAssignments returns an object that can list and get PolicyAssignments.
func (s *policyAssignmentLister) PolicyAssignments(namespace string) PolicyAssignmentNamespaceLister {
	return policyAssignmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PolicyAssignmentNamespaceLister helps list and get PolicyAssignments.
type PolicyAssignmentNamespaceLister interface {
	// List lists all PolicyAssignments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PolicyAssignment, err error)
	// Get retrieves the PolicyAssignment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PolicyAssignment, error)
	PolicyAssignmentNamespaceListerExpansion
}

// policyAssignmentNamespaceLister implements the PolicyAssignmentNamespaceLister
// interface.
type policyAssignmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PolicyAssignments in the indexer for a given namespace.
func (s policyAssignmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PolicyAssignment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PolicyAssignment))
	})
	return ret, err
}

// Get retrieves the PolicyAssignment from the indexer for a given namespace and name.
func (s policyAssignmentNamespaceLister) Get(name string) (*v1alpha1.PolicyAssignment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("policyassignment"), name)
	}
	return obj.(*v1alpha1.PolicyAssignment), nil
}
