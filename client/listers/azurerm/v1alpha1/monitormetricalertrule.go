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

// MonitorMetricAlertruleLister helps list MonitorMetricAlertrules.
type MonitorMetricAlertruleLister interface {
	// List lists all MonitorMetricAlertrules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MonitorMetricAlertrule, err error)
	// MonitorMetricAlertrules returns an object that can list and get MonitorMetricAlertrules.
	MonitorMetricAlertrules(namespace string) MonitorMetricAlertruleNamespaceLister
	MonitorMetricAlertruleListerExpansion
}

// monitorMetricAlertruleLister implements the MonitorMetricAlertruleLister interface.
type monitorMetricAlertruleLister struct {
	indexer cache.Indexer
}

// NewMonitorMetricAlertruleLister returns a new MonitorMetricAlertruleLister.
func NewMonitorMetricAlertruleLister(indexer cache.Indexer) MonitorMetricAlertruleLister {
	return &monitorMetricAlertruleLister{indexer: indexer}
}

// List lists all MonitorMetricAlertrules in the indexer.
func (s *monitorMetricAlertruleLister) List(selector labels.Selector) (ret []*v1alpha1.MonitorMetricAlertrule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MonitorMetricAlertrule))
	})
	return ret, err
}

// MonitorMetricAlertrules returns an object that can list and get MonitorMetricAlertrules.
func (s *monitorMetricAlertruleLister) MonitorMetricAlertrules(namespace string) MonitorMetricAlertruleNamespaceLister {
	return monitorMetricAlertruleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MonitorMetricAlertruleNamespaceLister helps list and get MonitorMetricAlertrules.
type MonitorMetricAlertruleNamespaceLister interface {
	// List lists all MonitorMetricAlertrules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MonitorMetricAlertrule, err error)
	// Get retrieves the MonitorMetricAlertrule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MonitorMetricAlertrule, error)
	MonitorMetricAlertruleNamespaceListerExpansion
}

// monitorMetricAlertruleNamespaceLister implements the MonitorMetricAlertruleNamespaceLister
// interface.
type monitorMetricAlertruleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MonitorMetricAlertrules in the indexer for a given namespace.
func (s monitorMetricAlertruleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MonitorMetricAlertrule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MonitorMetricAlertrule))
	})
	return ret, err
}

// Get retrieves the MonitorMetricAlertrule from the indexer for a given namespace and name.
func (s monitorMetricAlertruleNamespaceLister) Get(name string) (*v1alpha1.MonitorMetricAlertrule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("monitormetricalertrule"), name)
	}
	return obj.(*v1alpha1.MonitorMetricAlertrule), nil
}
