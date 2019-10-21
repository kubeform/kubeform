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

// IamSamlProviderLister helps list IamSamlProviders.
type IamSamlProviderLister interface {
	// List lists all IamSamlProviders in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IamSamlProvider, err error)
	// IamSamlProviders returns an object that can list and get IamSamlProviders.
	IamSamlProviders(namespace string) IamSamlProviderNamespaceLister
	IamSamlProviderListerExpansion
}

// iamSamlProviderLister implements the IamSamlProviderLister interface.
type iamSamlProviderLister struct {
	indexer cache.Indexer
}

// NewIamSamlProviderLister returns a new IamSamlProviderLister.
func NewIamSamlProviderLister(indexer cache.Indexer) IamSamlProviderLister {
	return &iamSamlProviderLister{indexer: indexer}
}

// List lists all IamSamlProviders in the indexer.
func (s *iamSamlProviderLister) List(selector labels.Selector) (ret []*v1alpha1.IamSamlProvider, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamSamlProvider))
	})
	return ret, err
}

// IamSamlProviders returns an object that can list and get IamSamlProviders.
func (s *iamSamlProviderLister) IamSamlProviders(namespace string) IamSamlProviderNamespaceLister {
	return iamSamlProviderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IamSamlProviderNamespaceLister helps list and get IamSamlProviders.
type IamSamlProviderNamespaceLister interface {
	// List lists all IamSamlProviders in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IamSamlProvider, err error)
	// Get retrieves the IamSamlProvider from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IamSamlProvider, error)
	IamSamlProviderNamespaceListerExpansion
}

// iamSamlProviderNamespaceLister implements the IamSamlProviderNamespaceLister
// interface.
type iamSamlProviderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IamSamlProviders in the indexer for a given namespace.
func (s iamSamlProviderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IamSamlProvider, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamSamlProvider))
	})
	return ret, err
}

// Get retrieves the IamSamlProvider from the indexer for a given namespace and name.
func (s iamSamlProviderNamespaceLister) Get(name string) (*v1alpha1.IamSamlProvider, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iamsamlprovider"), name)
	}
	return obj.(*v1alpha1.IamSamlProvider), nil
}