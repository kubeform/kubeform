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

// NeptuneSubnetGroupLister helps list NeptuneSubnetGroups.
type NeptuneSubnetGroupLister interface {
	// List lists all NeptuneSubnetGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NeptuneSubnetGroup, err error)
	// NeptuneSubnetGroups returns an object that can list and get NeptuneSubnetGroups.
	NeptuneSubnetGroups(namespace string) NeptuneSubnetGroupNamespaceLister
	NeptuneSubnetGroupListerExpansion
}

// neptuneSubnetGroupLister implements the NeptuneSubnetGroupLister interface.
type neptuneSubnetGroupLister struct {
	indexer cache.Indexer
}

// NewNeptuneSubnetGroupLister returns a new NeptuneSubnetGroupLister.
func NewNeptuneSubnetGroupLister(indexer cache.Indexer) NeptuneSubnetGroupLister {
	return &neptuneSubnetGroupLister{indexer: indexer}
}

// List lists all NeptuneSubnetGroups in the indexer.
func (s *neptuneSubnetGroupLister) List(selector labels.Selector) (ret []*v1alpha1.NeptuneSubnetGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NeptuneSubnetGroup))
	})
	return ret, err
}

// NeptuneSubnetGroups returns an object that can list and get NeptuneSubnetGroups.
func (s *neptuneSubnetGroupLister) NeptuneSubnetGroups(namespace string) NeptuneSubnetGroupNamespaceLister {
	return neptuneSubnetGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NeptuneSubnetGroupNamespaceLister helps list and get NeptuneSubnetGroups.
type NeptuneSubnetGroupNamespaceLister interface {
	// List lists all NeptuneSubnetGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NeptuneSubnetGroup, err error)
	// Get retrieves the NeptuneSubnetGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NeptuneSubnetGroup, error)
	NeptuneSubnetGroupNamespaceListerExpansion
}

// neptuneSubnetGroupNamespaceLister implements the NeptuneSubnetGroupNamespaceLister
// interface.
type neptuneSubnetGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NeptuneSubnetGroups in the indexer for a given namespace.
func (s neptuneSubnetGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NeptuneSubnetGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NeptuneSubnetGroup))
	})
	return ret, err
}

// Get retrieves the NeptuneSubnetGroup from the indexer for a given namespace and name.
func (s neptuneSubnetGroupNamespaceLister) Get(name string) (*v1alpha1.NeptuneSubnetGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("neptunesubnetgroup"), name)
	}
	return obj.(*v1alpha1.NeptuneSubnetGroup), nil
}
