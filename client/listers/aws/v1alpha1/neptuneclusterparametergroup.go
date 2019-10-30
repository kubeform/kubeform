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

// NeptuneClusterParameterGroupLister helps list NeptuneClusterParameterGroups.
type NeptuneClusterParameterGroupLister interface {
	// List lists all NeptuneClusterParameterGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NeptuneClusterParameterGroup, err error)
	// NeptuneClusterParameterGroups returns an object that can list and get NeptuneClusterParameterGroups.
	NeptuneClusterParameterGroups(namespace string) NeptuneClusterParameterGroupNamespaceLister
	NeptuneClusterParameterGroupListerExpansion
}

// neptuneClusterParameterGroupLister implements the NeptuneClusterParameterGroupLister interface.
type neptuneClusterParameterGroupLister struct {
	indexer cache.Indexer
}

// NewNeptuneClusterParameterGroupLister returns a new NeptuneClusterParameterGroupLister.
func NewNeptuneClusterParameterGroupLister(indexer cache.Indexer) NeptuneClusterParameterGroupLister {
	return &neptuneClusterParameterGroupLister{indexer: indexer}
}

// List lists all NeptuneClusterParameterGroups in the indexer.
func (s *neptuneClusterParameterGroupLister) List(selector labels.Selector) (ret []*v1alpha1.NeptuneClusterParameterGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NeptuneClusterParameterGroup))
	})
	return ret, err
}

// NeptuneClusterParameterGroups returns an object that can list and get NeptuneClusterParameterGroups.
func (s *neptuneClusterParameterGroupLister) NeptuneClusterParameterGroups(namespace string) NeptuneClusterParameterGroupNamespaceLister {
	return neptuneClusterParameterGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NeptuneClusterParameterGroupNamespaceLister helps list and get NeptuneClusterParameterGroups.
type NeptuneClusterParameterGroupNamespaceLister interface {
	// List lists all NeptuneClusterParameterGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NeptuneClusterParameterGroup, err error)
	// Get retrieves the NeptuneClusterParameterGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NeptuneClusterParameterGroup, error)
	NeptuneClusterParameterGroupNamespaceListerExpansion
}

// neptuneClusterParameterGroupNamespaceLister implements the NeptuneClusterParameterGroupNamespaceLister
// interface.
type neptuneClusterParameterGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NeptuneClusterParameterGroups in the indexer for a given namespace.
func (s neptuneClusterParameterGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NeptuneClusterParameterGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NeptuneClusterParameterGroup))
	})
	return ret, err
}

// Get retrieves the NeptuneClusterParameterGroup from the indexer for a given namespace and name.
func (s neptuneClusterParameterGroupNamespaceLister) Get(name string) (*v1alpha1.NeptuneClusterParameterGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("neptuneclusterparametergroup"), name)
	}
	return obj.(*v1alpha1.NeptuneClusterParameterGroup), nil
}
