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

// AcmCertificateValidationLister helps list AcmCertificateValidations.
type AcmCertificateValidationLister interface {
	// List lists all AcmCertificateValidations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AcmCertificateValidation, err error)
	// AcmCertificateValidations returns an object that can list and get AcmCertificateValidations.
	AcmCertificateValidations(namespace string) AcmCertificateValidationNamespaceLister
	AcmCertificateValidationListerExpansion
}

// acmCertificateValidationLister implements the AcmCertificateValidationLister interface.
type acmCertificateValidationLister struct {
	indexer cache.Indexer
}

// NewAcmCertificateValidationLister returns a new AcmCertificateValidationLister.
func NewAcmCertificateValidationLister(indexer cache.Indexer) AcmCertificateValidationLister {
	return &acmCertificateValidationLister{indexer: indexer}
}

// List lists all AcmCertificateValidations in the indexer.
func (s *acmCertificateValidationLister) List(selector labels.Selector) (ret []*v1alpha1.AcmCertificateValidation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AcmCertificateValidation))
	})
	return ret, err
}

// AcmCertificateValidations returns an object that can list and get AcmCertificateValidations.
func (s *acmCertificateValidationLister) AcmCertificateValidations(namespace string) AcmCertificateValidationNamespaceLister {
	return acmCertificateValidationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AcmCertificateValidationNamespaceLister helps list and get AcmCertificateValidations.
type AcmCertificateValidationNamespaceLister interface {
	// List lists all AcmCertificateValidations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AcmCertificateValidation, err error)
	// Get retrieves the AcmCertificateValidation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AcmCertificateValidation, error)
	AcmCertificateValidationNamespaceListerExpansion
}

// acmCertificateValidationNamespaceLister implements the AcmCertificateValidationNamespaceLister
// interface.
type acmCertificateValidationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AcmCertificateValidations in the indexer for a given namespace.
func (s acmCertificateValidationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AcmCertificateValidation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AcmCertificateValidation))
	})
	return ret, err
}

// Get retrieves the AcmCertificateValidation from the indexer for a given namespace and name.
func (s acmCertificateValidationNamespaceLister) Get(name string) (*v1alpha1.AcmCertificateValidation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("acmcertificatevalidation"), name)
	}
	return obj.(*v1alpha1.AcmCertificateValidation), nil
}
