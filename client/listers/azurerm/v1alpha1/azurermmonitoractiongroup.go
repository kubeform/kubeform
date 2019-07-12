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

// AzurermMonitorActionGroupLister helps list AzurermMonitorActionGroups.
type AzurermMonitorActionGroupLister interface {
	// List lists all AzurermMonitorActionGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AzurermMonitorActionGroup, err error)
	// Get retrieves the AzurermMonitorActionGroup from the index for a given name.
	Get(name string) (*v1alpha1.AzurermMonitorActionGroup, error)
	AzurermMonitorActionGroupListerExpansion
}

// azurermMonitorActionGroupLister implements the AzurermMonitorActionGroupLister interface.
type azurermMonitorActionGroupLister struct {
	indexer cache.Indexer
}

// NewAzurermMonitorActionGroupLister returns a new AzurermMonitorActionGroupLister.
func NewAzurermMonitorActionGroupLister(indexer cache.Indexer) AzurermMonitorActionGroupLister {
	return &azurermMonitorActionGroupLister{indexer: indexer}
}

// List lists all AzurermMonitorActionGroups in the indexer.
func (s *azurermMonitorActionGroupLister) List(selector labels.Selector) (ret []*v1alpha1.AzurermMonitorActionGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AzurermMonitorActionGroup))
	})
	return ret, err
}

// Get retrieves the AzurermMonitorActionGroup from the index for a given name.
func (s *azurermMonitorActionGroupLister) Get(name string) (*v1alpha1.AzurermMonitorActionGroup, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("azurermmonitoractiongroup"), name)
	}
	return obj.(*v1alpha1.AzurermMonitorActionGroup), nil
}
