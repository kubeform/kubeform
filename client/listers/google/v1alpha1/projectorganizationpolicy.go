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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProjectOrganizationPolicyLister helps list ProjectOrganizationPolicies.
type ProjectOrganizationPolicyLister interface {
	// List lists all ProjectOrganizationPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ProjectOrganizationPolicy, err error)
	// ProjectOrganizationPolicies returns an object that can list and get ProjectOrganizationPolicies.
	ProjectOrganizationPolicies(namespace string) ProjectOrganizationPolicyNamespaceLister
	ProjectOrganizationPolicyListerExpansion
}

// projectOrganizationPolicyLister implements the ProjectOrganizationPolicyLister interface.
type projectOrganizationPolicyLister struct {
	indexer cache.Indexer
}

// NewProjectOrganizationPolicyLister returns a new ProjectOrganizationPolicyLister.
func NewProjectOrganizationPolicyLister(indexer cache.Indexer) ProjectOrganizationPolicyLister {
	return &projectOrganizationPolicyLister{indexer: indexer}
}

// List lists all ProjectOrganizationPolicies in the indexer.
func (s *projectOrganizationPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.ProjectOrganizationPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProjectOrganizationPolicy))
	})
	return ret, err
}

// ProjectOrganizationPolicies returns an object that can list and get ProjectOrganizationPolicies.
func (s *projectOrganizationPolicyLister) ProjectOrganizationPolicies(namespace string) ProjectOrganizationPolicyNamespaceLister {
	return projectOrganizationPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProjectOrganizationPolicyNamespaceLister helps list and get ProjectOrganizationPolicies.
type ProjectOrganizationPolicyNamespaceLister interface {
	// List lists all ProjectOrganizationPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ProjectOrganizationPolicy, err error)
	// Get retrieves the ProjectOrganizationPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ProjectOrganizationPolicy, error)
	ProjectOrganizationPolicyNamespaceListerExpansion
}

// projectOrganizationPolicyNamespaceLister implements the ProjectOrganizationPolicyNamespaceLister
// interface.
type projectOrganizationPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProjectOrganizationPolicies in the indexer for a given namespace.
func (s projectOrganizationPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProjectOrganizationPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProjectOrganizationPolicy))
	})
	return ret, err
}

// Get retrieves the ProjectOrganizationPolicy from the indexer for a given namespace and name.
func (s projectOrganizationPolicyNamespaceLister) Get(name string) (*v1alpha1.ProjectOrganizationPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("projectorganizationpolicy"), name)
	}
	return obj.(*v1alpha1.ProjectOrganizationPolicy), nil
}
