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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RoleAssignmentLister helps list RoleAssignments.
type RoleAssignmentLister interface {
	// List lists all RoleAssignments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RoleAssignment, err error)
	// RoleAssignments returns an object that can list and get RoleAssignments.
	RoleAssignments(namespace string) RoleAssignmentNamespaceLister
	RoleAssignmentListerExpansion
}

// roleAssignmentLister implements the RoleAssignmentLister interface.
type roleAssignmentLister struct {
	indexer cache.Indexer
}

// NewRoleAssignmentLister returns a new RoleAssignmentLister.
func NewRoleAssignmentLister(indexer cache.Indexer) RoleAssignmentLister {
	return &roleAssignmentLister{indexer: indexer}
}

// List lists all RoleAssignments in the indexer.
func (s *roleAssignmentLister) List(selector labels.Selector) (ret []*v1alpha1.RoleAssignment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RoleAssignment))
	})
	return ret, err
}

// RoleAssignments returns an object that can list and get RoleAssignments.
func (s *roleAssignmentLister) RoleAssignments(namespace string) RoleAssignmentNamespaceLister {
	return roleAssignmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RoleAssignmentNamespaceLister helps list and get RoleAssignments.
type RoleAssignmentNamespaceLister interface {
	// List lists all RoleAssignments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RoleAssignment, err error)
	// Get retrieves the RoleAssignment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RoleAssignment, error)
	RoleAssignmentNamespaceListerExpansion
}

// roleAssignmentNamespaceLister implements the RoleAssignmentNamespaceLister
// interface.
type roleAssignmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RoleAssignments in the indexer for a given namespace.
func (s roleAssignmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RoleAssignment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RoleAssignment))
	})
	return ret, err
}

// Get retrieves the RoleAssignment from the indexer for a given namespace and name.
func (s roleAssignmentNamespaceLister) Get(name string) (*v1alpha1.RoleAssignment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("roleassignment"), name)
	}
	return obj.(*v1alpha1.RoleAssignment), nil
}
