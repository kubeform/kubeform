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

// CognitiveAccountLister helps list CognitiveAccounts.
type CognitiveAccountLister interface {
	// List lists all CognitiveAccounts in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CognitiveAccount, err error)
	// CognitiveAccounts returns an object that can list and get CognitiveAccounts.
	CognitiveAccounts(namespace string) CognitiveAccountNamespaceLister
	CognitiveAccountListerExpansion
}

// cognitiveAccountLister implements the CognitiveAccountLister interface.
type cognitiveAccountLister struct {
	indexer cache.Indexer
}

// NewCognitiveAccountLister returns a new CognitiveAccountLister.
func NewCognitiveAccountLister(indexer cache.Indexer) CognitiveAccountLister {
	return &cognitiveAccountLister{indexer: indexer}
}

// List lists all CognitiveAccounts in the indexer.
func (s *cognitiveAccountLister) List(selector labels.Selector) (ret []*v1alpha1.CognitiveAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CognitiveAccount))
	})
	return ret, err
}

// CognitiveAccounts returns an object that can list and get CognitiveAccounts.
func (s *cognitiveAccountLister) CognitiveAccounts(namespace string) CognitiveAccountNamespaceLister {
	return cognitiveAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CognitiveAccountNamespaceLister helps list and get CognitiveAccounts.
type CognitiveAccountNamespaceLister interface {
	// List lists all CognitiveAccounts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CognitiveAccount, err error)
	// Get retrieves the CognitiveAccount from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CognitiveAccount, error)
	CognitiveAccountNamespaceListerExpansion
}

// cognitiveAccountNamespaceLister implements the CognitiveAccountNamespaceLister
// interface.
type cognitiveAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CognitiveAccounts in the indexer for a given namespace.
func (s cognitiveAccountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CognitiveAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CognitiveAccount))
	})
	return ret, err
}

// Get retrieves the CognitiveAccount from the indexer for a given namespace and name.
func (s cognitiveAccountNamespaceLister) Get(name string) (*v1alpha1.CognitiveAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cognitiveaccount"), name)
	}
	return obj.(*v1alpha1.CognitiveAccount), nil
}