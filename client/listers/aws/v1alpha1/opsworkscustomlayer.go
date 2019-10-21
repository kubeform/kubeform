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

// OpsworksCustomLayerLister helps list OpsworksCustomLayers.
type OpsworksCustomLayerLister interface {
	// List lists all OpsworksCustomLayers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.OpsworksCustomLayer, err error)
	// OpsworksCustomLayers returns an object that can list and get OpsworksCustomLayers.
	OpsworksCustomLayers(namespace string) OpsworksCustomLayerNamespaceLister
	OpsworksCustomLayerListerExpansion
}

// opsworksCustomLayerLister implements the OpsworksCustomLayerLister interface.
type opsworksCustomLayerLister struct {
	indexer cache.Indexer
}

// NewOpsworksCustomLayerLister returns a new OpsworksCustomLayerLister.
func NewOpsworksCustomLayerLister(indexer cache.Indexer) OpsworksCustomLayerLister {
	return &opsworksCustomLayerLister{indexer: indexer}
}

// List lists all OpsworksCustomLayers in the indexer.
func (s *opsworksCustomLayerLister) List(selector labels.Selector) (ret []*v1alpha1.OpsworksCustomLayer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpsworksCustomLayer))
	})
	return ret, err
}

// OpsworksCustomLayers returns an object that can list and get OpsworksCustomLayers.
func (s *opsworksCustomLayerLister) OpsworksCustomLayers(namespace string) OpsworksCustomLayerNamespaceLister {
	return opsworksCustomLayerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OpsworksCustomLayerNamespaceLister helps list and get OpsworksCustomLayers.
type OpsworksCustomLayerNamespaceLister interface {
	// List lists all OpsworksCustomLayers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.OpsworksCustomLayer, err error)
	// Get retrieves the OpsworksCustomLayer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.OpsworksCustomLayer, error)
	OpsworksCustomLayerNamespaceListerExpansion
}

// opsworksCustomLayerNamespaceLister implements the OpsworksCustomLayerNamespaceLister
// interface.
type opsworksCustomLayerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OpsworksCustomLayers in the indexer for a given namespace.
func (s opsworksCustomLayerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OpsworksCustomLayer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpsworksCustomLayer))
	})
	return ret, err
}

// Get retrieves the OpsworksCustomLayer from the indexer for a given namespace and name.
func (s opsworksCustomLayerNamespaceLister) Get(name string) (*v1alpha1.OpsworksCustomLayer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("opsworkscustomlayer"), name)
	}
	return obj.(*v1alpha1.OpsworksCustomLayer), nil
}