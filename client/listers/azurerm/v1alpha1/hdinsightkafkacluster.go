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

// HdinsightKafkaClusterLister helps list HdinsightKafkaClusters.
type HdinsightKafkaClusterLister interface {
	// List lists all HdinsightKafkaClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.HdinsightKafkaCluster, err error)
	// HdinsightKafkaClusters returns an object that can list and get HdinsightKafkaClusters.
	HdinsightKafkaClusters(namespace string) HdinsightKafkaClusterNamespaceLister
	HdinsightKafkaClusterListerExpansion
}

// hdinsightKafkaClusterLister implements the HdinsightKafkaClusterLister interface.
type hdinsightKafkaClusterLister struct {
	indexer cache.Indexer
}

// NewHdinsightKafkaClusterLister returns a new HdinsightKafkaClusterLister.
func NewHdinsightKafkaClusterLister(indexer cache.Indexer) HdinsightKafkaClusterLister {
	return &hdinsightKafkaClusterLister{indexer: indexer}
}

// List lists all HdinsightKafkaClusters in the indexer.
func (s *hdinsightKafkaClusterLister) List(selector labels.Selector) (ret []*v1alpha1.HdinsightKafkaCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HdinsightKafkaCluster))
	})
	return ret, err
}

// HdinsightKafkaClusters returns an object that can list and get HdinsightKafkaClusters.
func (s *hdinsightKafkaClusterLister) HdinsightKafkaClusters(namespace string) HdinsightKafkaClusterNamespaceLister {
	return hdinsightKafkaClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HdinsightKafkaClusterNamespaceLister helps list and get HdinsightKafkaClusters.
type HdinsightKafkaClusterNamespaceLister interface {
	// List lists all HdinsightKafkaClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.HdinsightKafkaCluster, err error)
	// Get retrieves the HdinsightKafkaCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.HdinsightKafkaCluster, error)
	HdinsightKafkaClusterNamespaceListerExpansion
}

// hdinsightKafkaClusterNamespaceLister implements the HdinsightKafkaClusterNamespaceLister
// interface.
type hdinsightKafkaClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HdinsightKafkaClusters in the indexer for a given namespace.
func (s hdinsightKafkaClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HdinsightKafkaCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HdinsightKafkaCluster))
	})
	return ret, err
}

// Get retrieves the HdinsightKafkaCluster from the indexer for a given namespace and name.
func (s hdinsightKafkaClusterNamespaceLister) Get(name string) (*v1alpha1.HdinsightKafkaCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hdinsightkafkacluster"), name)
	}
	return obj.(*v1alpha1.HdinsightKafkaCluster), nil
}