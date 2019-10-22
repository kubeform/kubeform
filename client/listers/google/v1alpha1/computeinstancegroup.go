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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ComputeInstanceGroupLister helps list ComputeInstanceGroups.
type ComputeInstanceGroupLister interface {
	// List lists all ComputeInstanceGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeInstanceGroup, err error)
	// ComputeInstanceGroups returns an object that can list and get ComputeInstanceGroups.
	ComputeInstanceGroups(namespace string) ComputeInstanceGroupNamespaceLister
	ComputeInstanceGroupListerExpansion
}

// computeInstanceGroupLister implements the ComputeInstanceGroupLister interface.
type computeInstanceGroupLister struct {
	indexer cache.Indexer
}

// NewComputeInstanceGroupLister returns a new ComputeInstanceGroupLister.
func NewComputeInstanceGroupLister(indexer cache.Indexer) ComputeInstanceGroupLister {
	return &computeInstanceGroupLister{indexer: indexer}
}

// List lists all ComputeInstanceGroups in the indexer.
func (s *computeInstanceGroupLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeInstanceGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeInstanceGroup))
	})
	return ret, err
}

// ComputeInstanceGroups returns an object that can list and get ComputeInstanceGroups.
func (s *computeInstanceGroupLister) ComputeInstanceGroups(namespace string) ComputeInstanceGroupNamespaceLister {
	return computeInstanceGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeInstanceGroupNamespaceLister helps list and get ComputeInstanceGroups.
type ComputeInstanceGroupNamespaceLister interface {
	// List lists all ComputeInstanceGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeInstanceGroup, err error)
	// Get retrieves the ComputeInstanceGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeInstanceGroup, error)
	ComputeInstanceGroupNamespaceListerExpansion
}

// computeInstanceGroupNamespaceLister implements the ComputeInstanceGroupNamespaceLister
// interface.
type computeInstanceGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeInstanceGroups in the indexer for a given namespace.
func (s computeInstanceGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeInstanceGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeInstanceGroup))
	})
	return ret, err
}

// Get retrieves the ComputeInstanceGroup from the indexer for a given namespace and name.
func (s computeInstanceGroupNamespaceLister) Get(name string) (*v1alpha1.ComputeInstanceGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computeinstancegroup"), name)
	}
	return obj.(*v1alpha1.ComputeInstanceGroup), nil
}
