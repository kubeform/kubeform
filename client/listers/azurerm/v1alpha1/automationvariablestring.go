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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AutomationVariableStringLister helps list AutomationVariableStrings.
type AutomationVariableStringLister interface {
	// List lists all AutomationVariableStrings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AutomationVariableString, err error)
	// AutomationVariableStrings returns an object that can list and get AutomationVariableStrings.
	AutomationVariableStrings(namespace string) AutomationVariableStringNamespaceLister
	AutomationVariableStringListerExpansion
}

// automationVariableStringLister implements the AutomationVariableStringLister interface.
type automationVariableStringLister struct {
	indexer cache.Indexer
}

// NewAutomationVariableStringLister returns a new AutomationVariableStringLister.
func NewAutomationVariableStringLister(indexer cache.Indexer) AutomationVariableStringLister {
	return &automationVariableStringLister{indexer: indexer}
}

// List lists all AutomationVariableStrings in the indexer.
func (s *automationVariableStringLister) List(selector labels.Selector) (ret []*v1alpha1.AutomationVariableString, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutomationVariableString))
	})
	return ret, err
}

// AutomationVariableStrings returns an object that can list and get AutomationVariableStrings.
func (s *automationVariableStringLister) AutomationVariableStrings(namespace string) AutomationVariableStringNamespaceLister {
	return automationVariableStringNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AutomationVariableStringNamespaceLister helps list and get AutomationVariableStrings.
type AutomationVariableStringNamespaceLister interface {
	// List lists all AutomationVariableStrings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AutomationVariableString, err error)
	// Get retrieves the AutomationVariableString from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AutomationVariableString, error)
	AutomationVariableStringNamespaceListerExpansion
}

// automationVariableStringNamespaceLister implements the AutomationVariableStringNamespaceLister
// interface.
type automationVariableStringNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AutomationVariableStrings in the indexer for a given namespace.
func (s automationVariableStringNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AutomationVariableString, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutomationVariableString))
	})
	return ret, err
}

// Get retrieves the AutomationVariableString from the indexer for a given namespace and name.
func (s automationVariableStringNamespaceLister) Get(name string) (*v1alpha1.AutomationVariableString, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("automationvariablestring"), name)
	}
	return obj.(*v1alpha1.AutomationVariableString), nil
}
