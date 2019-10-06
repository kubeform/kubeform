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

// SecurityCenterWorkspaceLister helps list SecurityCenterWorkspaces.
type SecurityCenterWorkspaceLister interface {
	// List lists all SecurityCenterWorkspaces in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SecurityCenterWorkspace, err error)
	// SecurityCenterWorkspaces returns an object that can list and get SecurityCenterWorkspaces.
	SecurityCenterWorkspaces(namespace string) SecurityCenterWorkspaceNamespaceLister
	SecurityCenterWorkspaceListerExpansion
}

// securityCenterWorkspaceLister implements the SecurityCenterWorkspaceLister interface.
type securityCenterWorkspaceLister struct {
	indexer cache.Indexer
}

// NewSecurityCenterWorkspaceLister returns a new SecurityCenterWorkspaceLister.
func NewSecurityCenterWorkspaceLister(indexer cache.Indexer) SecurityCenterWorkspaceLister {
	return &securityCenterWorkspaceLister{indexer: indexer}
}

// List lists all SecurityCenterWorkspaces in the indexer.
func (s *securityCenterWorkspaceLister) List(selector labels.Selector) (ret []*v1alpha1.SecurityCenterWorkspace, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecurityCenterWorkspace))
	})
	return ret, err
}

// SecurityCenterWorkspaces returns an object that can list and get SecurityCenterWorkspaces.
func (s *securityCenterWorkspaceLister) SecurityCenterWorkspaces(namespace string) SecurityCenterWorkspaceNamespaceLister {
	return securityCenterWorkspaceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SecurityCenterWorkspaceNamespaceLister helps list and get SecurityCenterWorkspaces.
type SecurityCenterWorkspaceNamespaceLister interface {
	// List lists all SecurityCenterWorkspaces in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SecurityCenterWorkspace, err error)
	// Get retrieves the SecurityCenterWorkspace from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SecurityCenterWorkspace, error)
	SecurityCenterWorkspaceNamespaceListerExpansion
}

// securityCenterWorkspaceNamespaceLister implements the SecurityCenterWorkspaceNamespaceLister
// interface.
type securityCenterWorkspaceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SecurityCenterWorkspaces in the indexer for a given namespace.
func (s securityCenterWorkspaceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SecurityCenterWorkspace, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecurityCenterWorkspace))
	})
	return ret, err
}

// Get retrieves the SecurityCenterWorkspace from the indexer for a given namespace and name.
func (s securityCenterWorkspaceNamespaceLister) Get(name string) (*v1alpha1.SecurityCenterWorkspace, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("securitycenterworkspace"), name)
	}
	return obj.(*v1alpha1.SecurityCenterWorkspace), nil
}
