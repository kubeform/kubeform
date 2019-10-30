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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AutomationVariableBoolLister helps list AutomationVariableBools.
type AutomationVariableBoolLister interface {
	// List lists all AutomationVariableBools in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AutomationVariableBool, err error)
	// AutomationVariableBools returns an object that can list and get AutomationVariableBools.
	AutomationVariableBools(namespace string) AutomationVariableBoolNamespaceLister
	AutomationVariableBoolListerExpansion
}

// automationVariableBoolLister implements the AutomationVariableBoolLister interface.
type automationVariableBoolLister struct {
	indexer cache.Indexer
}

// NewAutomationVariableBoolLister returns a new AutomationVariableBoolLister.
func NewAutomationVariableBoolLister(indexer cache.Indexer) AutomationVariableBoolLister {
	return &automationVariableBoolLister{indexer: indexer}
}

// List lists all AutomationVariableBools in the indexer.
func (s *automationVariableBoolLister) List(selector labels.Selector) (ret []*v1alpha1.AutomationVariableBool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutomationVariableBool))
	})
	return ret, err
}

// AutomationVariableBools returns an object that can list and get AutomationVariableBools.
func (s *automationVariableBoolLister) AutomationVariableBools(namespace string) AutomationVariableBoolNamespaceLister {
	return automationVariableBoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AutomationVariableBoolNamespaceLister helps list and get AutomationVariableBools.
type AutomationVariableBoolNamespaceLister interface {
	// List lists all AutomationVariableBools in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AutomationVariableBool, err error)
	// Get retrieves the AutomationVariableBool from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AutomationVariableBool, error)
	AutomationVariableBoolNamespaceListerExpansion
}

// automationVariableBoolNamespaceLister implements the AutomationVariableBoolNamespaceLister
// interface.
type automationVariableBoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AutomationVariableBools in the indexer for a given namespace.
func (s automationVariableBoolNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AutomationVariableBool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutomationVariableBool))
	})
	return ret, err
}

// Get retrieves the AutomationVariableBool from the indexer for a given namespace and name.
func (s automationVariableBoolNamespaceLister) Get(name string) (*v1alpha1.AutomationVariableBool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("automationvariablebool"), name)
	}
	return obj.(*v1alpha1.AutomationVariableBool), nil
}
