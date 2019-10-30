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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DmsReplicationInstanceLister helps list DmsReplicationInstances.
type DmsReplicationInstanceLister interface {
	// List lists all DmsReplicationInstances in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DmsReplicationInstance, err error)
	// DmsReplicationInstances returns an object that can list and get DmsReplicationInstances.
	DmsReplicationInstances(namespace string) DmsReplicationInstanceNamespaceLister
	DmsReplicationInstanceListerExpansion
}

// dmsReplicationInstanceLister implements the DmsReplicationInstanceLister interface.
type dmsReplicationInstanceLister struct {
	indexer cache.Indexer
}

// NewDmsReplicationInstanceLister returns a new DmsReplicationInstanceLister.
func NewDmsReplicationInstanceLister(indexer cache.Indexer) DmsReplicationInstanceLister {
	return &dmsReplicationInstanceLister{indexer: indexer}
}

// List lists all DmsReplicationInstances in the indexer.
func (s *dmsReplicationInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.DmsReplicationInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DmsReplicationInstance))
	})
	return ret, err
}

// DmsReplicationInstances returns an object that can list and get DmsReplicationInstances.
func (s *dmsReplicationInstanceLister) DmsReplicationInstances(namespace string) DmsReplicationInstanceNamespaceLister {
	return dmsReplicationInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DmsReplicationInstanceNamespaceLister helps list and get DmsReplicationInstances.
type DmsReplicationInstanceNamespaceLister interface {
	// List lists all DmsReplicationInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DmsReplicationInstance, err error)
	// Get retrieves the DmsReplicationInstance from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DmsReplicationInstance, error)
	DmsReplicationInstanceNamespaceListerExpansion
}

// dmsReplicationInstanceNamespaceLister implements the DmsReplicationInstanceNamespaceLister
// interface.
type dmsReplicationInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DmsReplicationInstances in the indexer for a given namespace.
func (s dmsReplicationInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DmsReplicationInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DmsReplicationInstance))
	})
	return ret, err
}

// Get retrieves the DmsReplicationInstance from the indexer for a given namespace and name.
func (s dmsReplicationInstanceNamespaceLister) Get(name string) (*v1alpha1.DmsReplicationInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dmsreplicationinstance"), name)
	}
	return obj.(*v1alpha1.DmsReplicationInstance), nil
}
