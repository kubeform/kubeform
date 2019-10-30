/*
Copyright The Kubeform Authors.

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

// ComputeSSLCertificateLister helps list ComputeSSLCertificates.
type ComputeSSLCertificateLister interface {
	// List lists all ComputeSSLCertificates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeSSLCertificate, err error)
	// ComputeSSLCertificates returns an object that can list and get ComputeSSLCertificates.
	ComputeSSLCertificates(namespace string) ComputeSSLCertificateNamespaceLister
	ComputeSSLCertificateListerExpansion
}

// computeSSLCertificateLister implements the ComputeSSLCertificateLister interface.
type computeSSLCertificateLister struct {
	indexer cache.Indexer
}

// NewComputeSSLCertificateLister returns a new ComputeSSLCertificateLister.
func NewComputeSSLCertificateLister(indexer cache.Indexer) ComputeSSLCertificateLister {
	return &computeSSLCertificateLister{indexer: indexer}
}

// List lists all ComputeSSLCertificates in the indexer.
func (s *computeSSLCertificateLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeSSLCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeSSLCertificate))
	})
	return ret, err
}

// ComputeSSLCertificates returns an object that can list and get ComputeSSLCertificates.
func (s *computeSSLCertificateLister) ComputeSSLCertificates(namespace string) ComputeSSLCertificateNamespaceLister {
	return computeSSLCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeSSLCertificateNamespaceLister helps list and get ComputeSSLCertificates.
type ComputeSSLCertificateNamespaceLister interface {
	// List lists all ComputeSSLCertificates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeSSLCertificate, err error)
	// Get retrieves the ComputeSSLCertificate from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeSSLCertificate, error)
	ComputeSSLCertificateNamespaceListerExpansion
}

// computeSSLCertificateNamespaceLister implements the ComputeSSLCertificateNamespaceLister
// interface.
type computeSSLCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeSSLCertificates in the indexer for a given namespace.
func (s computeSSLCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeSSLCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeSSLCertificate))
	})
	return ret, err
}

// Get retrieves the ComputeSSLCertificate from the indexer for a given namespace and name.
func (s computeSSLCertificateNamespaceLister) Get(name string) (*v1alpha1.ComputeSSLCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computesslcertificate"), name)
	}
	return obj.(*v1alpha1.ComputeSSLCertificate), nil
}
