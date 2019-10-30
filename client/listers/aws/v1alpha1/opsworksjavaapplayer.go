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

// OpsworksJavaAppLayerLister helps list OpsworksJavaAppLayers.
type OpsworksJavaAppLayerLister interface {
	// List lists all OpsworksJavaAppLayers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.OpsworksJavaAppLayer, err error)
	// OpsworksJavaAppLayers returns an object that can list and get OpsworksJavaAppLayers.
	OpsworksJavaAppLayers(namespace string) OpsworksJavaAppLayerNamespaceLister
	OpsworksJavaAppLayerListerExpansion
}

// opsworksJavaAppLayerLister implements the OpsworksJavaAppLayerLister interface.
type opsworksJavaAppLayerLister struct {
	indexer cache.Indexer
}

// NewOpsworksJavaAppLayerLister returns a new OpsworksJavaAppLayerLister.
func NewOpsworksJavaAppLayerLister(indexer cache.Indexer) OpsworksJavaAppLayerLister {
	return &opsworksJavaAppLayerLister{indexer: indexer}
}

// List lists all OpsworksJavaAppLayers in the indexer.
func (s *opsworksJavaAppLayerLister) List(selector labels.Selector) (ret []*v1alpha1.OpsworksJavaAppLayer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpsworksJavaAppLayer))
	})
	return ret, err
}

// OpsworksJavaAppLayers returns an object that can list and get OpsworksJavaAppLayers.
func (s *opsworksJavaAppLayerLister) OpsworksJavaAppLayers(namespace string) OpsworksJavaAppLayerNamespaceLister {
	return opsworksJavaAppLayerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OpsworksJavaAppLayerNamespaceLister helps list and get OpsworksJavaAppLayers.
type OpsworksJavaAppLayerNamespaceLister interface {
	// List lists all OpsworksJavaAppLayers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.OpsworksJavaAppLayer, err error)
	// Get retrieves the OpsworksJavaAppLayer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.OpsworksJavaAppLayer, error)
	OpsworksJavaAppLayerNamespaceListerExpansion
}

// opsworksJavaAppLayerNamespaceLister implements the OpsworksJavaAppLayerNamespaceLister
// interface.
type opsworksJavaAppLayerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OpsworksJavaAppLayers in the indexer for a given namespace.
func (s opsworksJavaAppLayerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OpsworksJavaAppLayer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpsworksJavaAppLayer))
	})
	return ret, err
}

// Get retrieves the OpsworksJavaAppLayer from the indexer for a given namespace and name.
func (s opsworksJavaAppLayerNamespaceLister) Get(name string) (*v1alpha1.OpsworksJavaAppLayer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("opsworksjavaapplayer"), name)
	}
	return obj.(*v1alpha1.OpsworksJavaAppLayer), nil
}
