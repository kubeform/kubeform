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

// OrganizationsOrganizationLister helps list OrganizationsOrganizations.
type OrganizationsOrganizationLister interface {
	// List lists all OrganizationsOrganizations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.OrganizationsOrganization, err error)
	// OrganizationsOrganizations returns an object that can list and get OrganizationsOrganizations.
	OrganizationsOrganizations(namespace string) OrganizationsOrganizationNamespaceLister
	OrganizationsOrganizationListerExpansion
}

// organizationsOrganizationLister implements the OrganizationsOrganizationLister interface.
type organizationsOrganizationLister struct {
	indexer cache.Indexer
}

// NewOrganizationsOrganizationLister returns a new OrganizationsOrganizationLister.
func NewOrganizationsOrganizationLister(indexer cache.Indexer) OrganizationsOrganizationLister {
	return &organizationsOrganizationLister{indexer: indexer}
}

// List lists all OrganizationsOrganizations in the indexer.
func (s *organizationsOrganizationLister) List(selector labels.Selector) (ret []*v1alpha1.OrganizationsOrganization, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OrganizationsOrganization))
	})
	return ret, err
}

// OrganizationsOrganizations returns an object that can list and get OrganizationsOrganizations.
func (s *organizationsOrganizationLister) OrganizationsOrganizations(namespace string) OrganizationsOrganizationNamespaceLister {
	return organizationsOrganizationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OrganizationsOrganizationNamespaceLister helps list and get OrganizationsOrganizations.
type OrganizationsOrganizationNamespaceLister interface {
	// List lists all OrganizationsOrganizations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.OrganizationsOrganization, err error)
	// Get retrieves the OrganizationsOrganization from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.OrganizationsOrganization, error)
	OrganizationsOrganizationNamespaceListerExpansion
}

// organizationsOrganizationNamespaceLister implements the OrganizationsOrganizationNamespaceLister
// interface.
type organizationsOrganizationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OrganizationsOrganizations in the indexer for a given namespace.
func (s organizationsOrganizationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OrganizationsOrganization, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OrganizationsOrganization))
	})
	return ret, err
}

// Get retrieves the OrganizationsOrganization from the indexer for a given namespace and name.
func (s organizationsOrganizationNamespaceLister) Get(name string) (*v1alpha1.OrganizationsOrganization, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("organizationsorganization"), name)
	}
	return obj.(*v1alpha1.OrganizationsOrganization), nil
}
