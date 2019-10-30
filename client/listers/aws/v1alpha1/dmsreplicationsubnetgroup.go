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

// DmsReplicationSubnetGroupLister helps list DmsReplicationSubnetGroups.
type DmsReplicationSubnetGroupLister interface {
	// List lists all DmsReplicationSubnetGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DmsReplicationSubnetGroup, err error)
	// DmsReplicationSubnetGroups returns an object that can list and get DmsReplicationSubnetGroups.
	DmsReplicationSubnetGroups(namespace string) DmsReplicationSubnetGroupNamespaceLister
	DmsReplicationSubnetGroupListerExpansion
}

// dmsReplicationSubnetGroupLister implements the DmsReplicationSubnetGroupLister interface.
type dmsReplicationSubnetGroupLister struct {
	indexer cache.Indexer
}

// NewDmsReplicationSubnetGroupLister returns a new DmsReplicationSubnetGroupLister.
func NewDmsReplicationSubnetGroupLister(indexer cache.Indexer) DmsReplicationSubnetGroupLister {
	return &dmsReplicationSubnetGroupLister{indexer: indexer}
}

// List lists all DmsReplicationSubnetGroups in the indexer.
func (s *dmsReplicationSubnetGroupLister) List(selector labels.Selector) (ret []*v1alpha1.DmsReplicationSubnetGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DmsReplicationSubnetGroup))
	})
	return ret, err
}

// DmsReplicationSubnetGroups returns an object that can list and get DmsReplicationSubnetGroups.
func (s *dmsReplicationSubnetGroupLister) DmsReplicationSubnetGroups(namespace string) DmsReplicationSubnetGroupNamespaceLister {
	return dmsReplicationSubnetGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DmsReplicationSubnetGroupNamespaceLister helps list and get DmsReplicationSubnetGroups.
type DmsReplicationSubnetGroupNamespaceLister interface {
	// List lists all DmsReplicationSubnetGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DmsReplicationSubnetGroup, err error)
	// Get retrieves the DmsReplicationSubnetGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DmsReplicationSubnetGroup, error)
	DmsReplicationSubnetGroupNamespaceListerExpansion
}

// dmsReplicationSubnetGroupNamespaceLister implements the DmsReplicationSubnetGroupNamespaceLister
// interface.
type dmsReplicationSubnetGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DmsReplicationSubnetGroups in the indexer for a given namespace.
func (s dmsReplicationSubnetGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DmsReplicationSubnetGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DmsReplicationSubnetGroup))
	})
	return ret, err
}

// Get retrieves the DmsReplicationSubnetGroup from the indexer for a given namespace and name.
func (s dmsReplicationSubnetGroupNamespaceLister) Get(name string) (*v1alpha1.DmsReplicationSubnetGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dmsreplicationsubnetgroup"), name)
	}
	return obj.(*v1alpha1.DmsReplicationSubnetGroup), nil
}
