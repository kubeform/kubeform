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

// HdinsightHbaseClusterLister helps list HdinsightHbaseClusters.
type HdinsightHbaseClusterLister interface {
	// List lists all HdinsightHbaseClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.HdinsightHbaseCluster, err error)
	// HdinsightHbaseClusters returns an object that can list and get HdinsightHbaseClusters.
	HdinsightHbaseClusters(namespace string) HdinsightHbaseClusterNamespaceLister
	HdinsightHbaseClusterListerExpansion
}

// hdinsightHbaseClusterLister implements the HdinsightHbaseClusterLister interface.
type hdinsightHbaseClusterLister struct {
	indexer cache.Indexer
}

// NewHdinsightHbaseClusterLister returns a new HdinsightHbaseClusterLister.
func NewHdinsightHbaseClusterLister(indexer cache.Indexer) HdinsightHbaseClusterLister {
	return &hdinsightHbaseClusterLister{indexer: indexer}
}

// List lists all HdinsightHbaseClusters in the indexer.
func (s *hdinsightHbaseClusterLister) List(selector labels.Selector) (ret []*v1alpha1.HdinsightHbaseCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HdinsightHbaseCluster))
	})
	return ret, err
}

// HdinsightHbaseClusters returns an object that can list and get HdinsightHbaseClusters.
func (s *hdinsightHbaseClusterLister) HdinsightHbaseClusters(namespace string) HdinsightHbaseClusterNamespaceLister {
	return hdinsightHbaseClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HdinsightHbaseClusterNamespaceLister helps list and get HdinsightHbaseClusters.
type HdinsightHbaseClusterNamespaceLister interface {
	// List lists all HdinsightHbaseClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.HdinsightHbaseCluster, err error)
	// Get retrieves the HdinsightHbaseCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.HdinsightHbaseCluster, error)
	HdinsightHbaseClusterNamespaceListerExpansion
}

// hdinsightHbaseClusterNamespaceLister implements the HdinsightHbaseClusterNamespaceLister
// interface.
type hdinsightHbaseClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HdinsightHbaseClusters in the indexer for a given namespace.
func (s hdinsightHbaseClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HdinsightHbaseCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HdinsightHbaseCluster))
	})
	return ret, err
}

// Get retrieves the HdinsightHbaseCluster from the indexer for a given namespace and name.
func (s hdinsightHbaseClusterNamespaceLister) Get(name string) (*v1alpha1.HdinsightHbaseCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hdinsighthbasecluster"), name)
	}
	return obj.(*v1alpha1.HdinsightHbaseCluster), nil
}