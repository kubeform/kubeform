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

// DmsReplicationTaskLister helps list DmsReplicationTasks.
type DmsReplicationTaskLister interface {
	// List lists all DmsReplicationTasks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DmsReplicationTask, err error)
	// DmsReplicationTasks returns an object that can list and get DmsReplicationTasks.
	DmsReplicationTasks(namespace string) DmsReplicationTaskNamespaceLister
	DmsReplicationTaskListerExpansion
}

// dmsReplicationTaskLister implements the DmsReplicationTaskLister interface.
type dmsReplicationTaskLister struct {
	indexer cache.Indexer
}

// NewDmsReplicationTaskLister returns a new DmsReplicationTaskLister.
func NewDmsReplicationTaskLister(indexer cache.Indexer) DmsReplicationTaskLister {
	return &dmsReplicationTaskLister{indexer: indexer}
}

// List lists all DmsReplicationTasks in the indexer.
func (s *dmsReplicationTaskLister) List(selector labels.Selector) (ret []*v1alpha1.DmsReplicationTask, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DmsReplicationTask))
	})
	return ret, err
}

// DmsReplicationTasks returns an object that can list and get DmsReplicationTasks.
func (s *dmsReplicationTaskLister) DmsReplicationTasks(namespace string) DmsReplicationTaskNamespaceLister {
	return dmsReplicationTaskNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DmsReplicationTaskNamespaceLister helps list and get DmsReplicationTasks.
type DmsReplicationTaskNamespaceLister interface {
	// List lists all DmsReplicationTasks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DmsReplicationTask, err error)
	// Get retrieves the DmsReplicationTask from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DmsReplicationTask, error)
	DmsReplicationTaskNamespaceListerExpansion
}

// dmsReplicationTaskNamespaceLister implements the DmsReplicationTaskNamespaceLister
// interface.
type dmsReplicationTaskNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DmsReplicationTasks in the indexer for a given namespace.
func (s dmsReplicationTaskNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DmsReplicationTask, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DmsReplicationTask))
	})
	return ret, err
}

// Get retrieves the DmsReplicationTask from the indexer for a given namespace and name.
func (s dmsReplicationTaskNamespaceLister) Get(name string) (*v1alpha1.DmsReplicationTask, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dmsreplicationtask"), name)
	}
	return obj.(*v1alpha1.DmsReplicationTask), nil
}
