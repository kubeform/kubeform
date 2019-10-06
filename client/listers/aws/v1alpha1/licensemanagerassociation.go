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

// LicensemanagerAssociationLister helps list LicensemanagerAssociations.
type LicensemanagerAssociationLister interface {
	// List lists all LicensemanagerAssociations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LicensemanagerAssociation, err error)
	// LicensemanagerAssociations returns an object that can list and get LicensemanagerAssociations.
	LicensemanagerAssociations(namespace string) LicensemanagerAssociationNamespaceLister
	LicensemanagerAssociationListerExpansion
}

// licensemanagerAssociationLister implements the LicensemanagerAssociationLister interface.
type licensemanagerAssociationLister struct {
	indexer cache.Indexer
}

// NewLicensemanagerAssociationLister returns a new LicensemanagerAssociationLister.
func NewLicensemanagerAssociationLister(indexer cache.Indexer) LicensemanagerAssociationLister {
	return &licensemanagerAssociationLister{indexer: indexer}
}

// List lists all LicensemanagerAssociations in the indexer.
func (s *licensemanagerAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.LicensemanagerAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LicensemanagerAssociation))
	})
	return ret, err
}

// LicensemanagerAssociations returns an object that can list and get LicensemanagerAssociations.
func (s *licensemanagerAssociationLister) LicensemanagerAssociations(namespace string) LicensemanagerAssociationNamespaceLister {
	return licensemanagerAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LicensemanagerAssociationNamespaceLister helps list and get LicensemanagerAssociations.
type LicensemanagerAssociationNamespaceLister interface {
	// List lists all LicensemanagerAssociations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LicensemanagerAssociation, err error)
	// Get retrieves the LicensemanagerAssociation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LicensemanagerAssociation, error)
	LicensemanagerAssociationNamespaceListerExpansion
}

// licensemanagerAssociationNamespaceLister implements the LicensemanagerAssociationNamespaceLister
// interface.
type licensemanagerAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LicensemanagerAssociations in the indexer for a given namespace.
func (s licensemanagerAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LicensemanagerAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LicensemanagerAssociation))
	})
	return ret, err
}

// Get retrieves the LicensemanagerAssociation from the indexer for a given namespace and name.
func (s licensemanagerAssociationNamespaceLister) Get(name string) (*v1alpha1.LicensemanagerAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("licensemanagerassociation"), name)
	}
	return obj.(*v1alpha1.LicensemanagerAssociation), nil
}
