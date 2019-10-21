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

// SecurityGroupLister helps list SecurityGroups.
type SecurityGroupLister interface {
	// List lists all SecurityGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SecurityGroup, err error)
	// SecurityGroups returns an object that can list and get SecurityGroups.
	SecurityGroups(namespace string) SecurityGroupNamespaceLister
	SecurityGroupListerExpansion
}

// securityGroupLister implements the SecurityGroupLister interface.
type securityGroupLister struct {
	indexer cache.Indexer
}

// NewSecurityGroupLister returns a new SecurityGroupLister.
func NewSecurityGroupLister(indexer cache.Indexer) SecurityGroupLister {
	return &securityGroupLister{indexer: indexer}
}

// List lists all SecurityGroups in the indexer.
func (s *securityGroupLister) List(selector labels.Selector) (ret []*v1alpha1.SecurityGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecurityGroup))
	})
	return ret, err
}

// SecurityGroups returns an object that can list and get SecurityGroups.
func (s *securityGroupLister) SecurityGroups(namespace string) SecurityGroupNamespaceLister {
	return securityGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SecurityGroupNamespaceLister helps list and get SecurityGroups.
type SecurityGroupNamespaceLister interface {
	// List lists all SecurityGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SecurityGroup, err error)
	// Get retrieves the SecurityGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SecurityGroup, error)
	SecurityGroupNamespaceListerExpansion
}

// securityGroupNamespaceLister implements the SecurityGroupNamespaceLister
// interface.
type securityGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SecurityGroups in the indexer for a given namespace.
func (s securityGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SecurityGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecurityGroup))
	})
	return ret, err
}

// Get retrieves the SecurityGroup from the indexer for a given namespace and name.
func (s securityGroupNamespaceLister) Get(name string) (*v1alpha1.SecurityGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("securitygroup"), name)
	}
	return obj.(*v1alpha1.SecurityGroup), nil
}