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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DbSnapshotLister helps list DbSnapshots.
type DbSnapshotLister interface {
	// List lists all DbSnapshots in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DbSnapshot, err error)
	// DbSnapshots returns an object that can list and get DbSnapshots.
	DbSnapshots(namespace string) DbSnapshotNamespaceLister
	DbSnapshotListerExpansion
}

// dbSnapshotLister implements the DbSnapshotLister interface.
type dbSnapshotLister struct {
	indexer cache.Indexer
}

// NewDbSnapshotLister returns a new DbSnapshotLister.
func NewDbSnapshotLister(indexer cache.Indexer) DbSnapshotLister {
	return &dbSnapshotLister{indexer: indexer}
}

// List lists all DbSnapshots in the indexer.
func (s *dbSnapshotLister) List(selector labels.Selector) (ret []*v1alpha1.DbSnapshot, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DbSnapshot))
	})
	return ret, err
}

// DbSnapshots returns an object that can list and get DbSnapshots.
func (s *dbSnapshotLister) DbSnapshots(namespace string) DbSnapshotNamespaceLister {
	return dbSnapshotNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DbSnapshotNamespaceLister helps list and get DbSnapshots.
type DbSnapshotNamespaceLister interface {
	// List lists all DbSnapshots in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DbSnapshot, err error)
	// Get retrieves the DbSnapshot from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DbSnapshot, error)
	DbSnapshotNamespaceListerExpansion
}

// dbSnapshotNamespaceLister implements the DbSnapshotNamespaceLister
// interface.
type dbSnapshotNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DbSnapshots in the indexer for a given namespace.
func (s dbSnapshotNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DbSnapshot, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DbSnapshot))
	})
	return ret, err
}

// Get retrieves the DbSnapshot from the indexer for a given namespace and name.
func (s dbSnapshotNamespaceLister) Get(name string) (*v1alpha1.DbSnapshot, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dbsnapshot"), name)
	}
	return obj.(*v1alpha1.DbSnapshot), nil
}
