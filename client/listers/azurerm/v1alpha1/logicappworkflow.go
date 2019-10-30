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

// LogicAppWorkflowLister helps list LogicAppWorkflows.
type LogicAppWorkflowLister interface {
	// List lists all LogicAppWorkflows in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LogicAppWorkflow, err error)
	// LogicAppWorkflows returns an object that can list and get LogicAppWorkflows.
	LogicAppWorkflows(namespace string) LogicAppWorkflowNamespaceLister
	LogicAppWorkflowListerExpansion
}

// logicAppWorkflowLister implements the LogicAppWorkflowLister interface.
type logicAppWorkflowLister struct {
	indexer cache.Indexer
}

// NewLogicAppWorkflowLister returns a new LogicAppWorkflowLister.
func NewLogicAppWorkflowLister(indexer cache.Indexer) LogicAppWorkflowLister {
	return &logicAppWorkflowLister{indexer: indexer}
}

// List lists all LogicAppWorkflows in the indexer.
func (s *logicAppWorkflowLister) List(selector labels.Selector) (ret []*v1alpha1.LogicAppWorkflow, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LogicAppWorkflow))
	})
	return ret, err
}

// LogicAppWorkflows returns an object that can list and get LogicAppWorkflows.
func (s *logicAppWorkflowLister) LogicAppWorkflows(namespace string) LogicAppWorkflowNamespaceLister {
	return logicAppWorkflowNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LogicAppWorkflowNamespaceLister helps list and get LogicAppWorkflows.
type LogicAppWorkflowNamespaceLister interface {
	// List lists all LogicAppWorkflows in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LogicAppWorkflow, err error)
	// Get retrieves the LogicAppWorkflow from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LogicAppWorkflow, error)
	LogicAppWorkflowNamespaceListerExpansion
}

// logicAppWorkflowNamespaceLister implements the LogicAppWorkflowNamespaceLister
// interface.
type logicAppWorkflowNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LogicAppWorkflows in the indexer for a given namespace.
func (s logicAppWorkflowNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LogicAppWorkflow, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LogicAppWorkflow))
	})
	return ret, err
}

// Get retrieves the LogicAppWorkflow from the indexer for a given namespace and name.
func (s logicAppWorkflowNamespaceLister) Get(name string) (*v1alpha1.LogicAppWorkflow, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("logicappworkflow"), name)
	}
	return obj.(*v1alpha1.LogicAppWorkflow), nil
}
