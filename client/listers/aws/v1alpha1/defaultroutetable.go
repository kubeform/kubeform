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

// DefaultRouteTableLister helps list DefaultRouteTables.
type DefaultRouteTableLister interface {
	// List lists all DefaultRouteTables in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DefaultRouteTable, err error)
	// DefaultRouteTables returns an object that can list and get DefaultRouteTables.
	DefaultRouteTables(namespace string) DefaultRouteTableNamespaceLister
	DefaultRouteTableListerExpansion
}

// defaultRouteTableLister implements the DefaultRouteTableLister interface.
type defaultRouteTableLister struct {
	indexer cache.Indexer
}

// NewDefaultRouteTableLister returns a new DefaultRouteTableLister.
func NewDefaultRouteTableLister(indexer cache.Indexer) DefaultRouteTableLister {
	return &defaultRouteTableLister{indexer: indexer}
}

// List lists all DefaultRouteTables in the indexer.
func (s *defaultRouteTableLister) List(selector labels.Selector) (ret []*v1alpha1.DefaultRouteTable, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DefaultRouteTable))
	})
	return ret, err
}

// DefaultRouteTables returns an object that can list and get DefaultRouteTables.
func (s *defaultRouteTableLister) DefaultRouteTables(namespace string) DefaultRouteTableNamespaceLister {
	return defaultRouteTableNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DefaultRouteTableNamespaceLister helps list and get DefaultRouteTables.
type DefaultRouteTableNamespaceLister interface {
	// List lists all DefaultRouteTables in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DefaultRouteTable, err error)
	// Get retrieves the DefaultRouteTable from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DefaultRouteTable, error)
	DefaultRouteTableNamespaceListerExpansion
}

// defaultRouteTableNamespaceLister implements the DefaultRouteTableNamespaceLister
// interface.
type defaultRouteTableNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DefaultRouteTables in the indexer for a given namespace.
func (s defaultRouteTableNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DefaultRouteTable, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DefaultRouteTable))
	})
	return ret, err
}

// Get retrieves the DefaultRouteTable from the indexer for a given namespace and name.
func (s defaultRouteTableNamespaceLister) Get(name string) (*v1alpha1.DefaultRouteTable, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("defaultroutetable"), name)
	}
	return obj.(*v1alpha1.DefaultRouteTable), nil
}
