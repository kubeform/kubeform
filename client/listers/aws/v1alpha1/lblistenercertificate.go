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

// LbListenerCertificateLister helps list LbListenerCertificates.
type LbListenerCertificateLister interface {
	// List lists all LbListenerCertificates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LbListenerCertificate, err error)
	// LbListenerCertificates returns an object that can list and get LbListenerCertificates.
	LbListenerCertificates(namespace string) LbListenerCertificateNamespaceLister
	LbListenerCertificateListerExpansion
}

// lbListenerCertificateLister implements the LbListenerCertificateLister interface.
type lbListenerCertificateLister struct {
	indexer cache.Indexer
}

// NewLbListenerCertificateLister returns a new LbListenerCertificateLister.
func NewLbListenerCertificateLister(indexer cache.Indexer) LbListenerCertificateLister {
	return &lbListenerCertificateLister{indexer: indexer}
}

// List lists all LbListenerCertificates in the indexer.
func (s *lbListenerCertificateLister) List(selector labels.Selector) (ret []*v1alpha1.LbListenerCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LbListenerCertificate))
	})
	return ret, err
}

// LbListenerCertificates returns an object that can list and get LbListenerCertificates.
func (s *lbListenerCertificateLister) LbListenerCertificates(namespace string) LbListenerCertificateNamespaceLister {
	return lbListenerCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LbListenerCertificateNamespaceLister helps list and get LbListenerCertificates.
type LbListenerCertificateNamespaceLister interface {
	// List lists all LbListenerCertificates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LbListenerCertificate, err error)
	// Get retrieves the LbListenerCertificate from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LbListenerCertificate, error)
	LbListenerCertificateNamespaceListerExpansion
}

// lbListenerCertificateNamespaceLister implements the LbListenerCertificateNamespaceLister
// interface.
type lbListenerCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LbListenerCertificates in the indexer for a given namespace.
func (s lbListenerCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LbListenerCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LbListenerCertificate))
	})
	return ret, err
}

// Get retrieves the LbListenerCertificate from the indexer for a given namespace and name.
func (s lbListenerCertificateNamespaceLister) Get(name string) (*v1alpha1.LbListenerCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("lblistenercertificate"), name)
	}
	return obj.(*v1alpha1.LbListenerCertificate), nil
}
