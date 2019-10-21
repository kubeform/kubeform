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

// RdsClusterInstanceLister helps list RdsClusterInstances.
type RdsClusterInstanceLister interface {
	// List lists all RdsClusterInstances in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RdsClusterInstance, err error)
	// RdsClusterInstances returns an object that can list and get RdsClusterInstances.
	RdsClusterInstances(namespace string) RdsClusterInstanceNamespaceLister
	RdsClusterInstanceListerExpansion
}

// rdsClusterInstanceLister implements the RdsClusterInstanceLister interface.
type rdsClusterInstanceLister struct {
	indexer cache.Indexer
}

// NewRdsClusterInstanceLister returns a new RdsClusterInstanceLister.
func NewRdsClusterInstanceLister(indexer cache.Indexer) RdsClusterInstanceLister {
	return &rdsClusterInstanceLister{indexer: indexer}
}

// List lists all RdsClusterInstances in the indexer.
func (s *rdsClusterInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.RdsClusterInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RdsClusterInstance))
	})
	return ret, err
}

// RdsClusterInstances returns an object that can list and get RdsClusterInstances.
func (s *rdsClusterInstanceLister) RdsClusterInstances(namespace string) RdsClusterInstanceNamespaceLister {
	return rdsClusterInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RdsClusterInstanceNamespaceLister helps list and get RdsClusterInstances.
type RdsClusterInstanceNamespaceLister interface {
	// List lists all RdsClusterInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RdsClusterInstance, err error)
	// Get retrieves the RdsClusterInstance from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RdsClusterInstance, error)
	RdsClusterInstanceNamespaceListerExpansion
}

// rdsClusterInstanceNamespaceLister implements the RdsClusterInstanceNamespaceLister
// interface.
type rdsClusterInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RdsClusterInstances in the indexer for a given namespace.
func (s rdsClusterInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RdsClusterInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RdsClusterInstance))
	})
	return ret, err
}

// Get retrieves the RdsClusterInstance from the indexer for a given namespace and name.
func (s rdsClusterInstanceNamespaceLister) Get(name string) (*v1alpha1.RdsClusterInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("rdsclusterinstance"), name)
	}
	return obj.(*v1alpha1.RdsClusterInstance), nil
}