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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// DataflowJobLister helps list DataflowJobs.
type DataflowJobLister interface {
	// List lists all DataflowJobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DataflowJob, err error)
	// DataflowJobs returns an object that can list and get DataflowJobs.
	DataflowJobs(namespace string) DataflowJobNamespaceLister
	DataflowJobListerExpansion
}

// dataflowJobLister implements the DataflowJobLister interface.
type dataflowJobLister struct {
	indexer cache.Indexer
}

// NewDataflowJobLister returns a new DataflowJobLister.
func NewDataflowJobLister(indexer cache.Indexer) DataflowJobLister {
	return &dataflowJobLister{indexer: indexer}
}

// List lists all DataflowJobs in the indexer.
func (s *dataflowJobLister) List(selector labels.Selector) (ret []*v1alpha1.DataflowJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataflowJob))
	})
	return ret, err
}

// DataflowJobs returns an object that can list and get DataflowJobs.
func (s *dataflowJobLister) DataflowJobs(namespace string) DataflowJobNamespaceLister {
	return dataflowJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataflowJobNamespaceLister helps list and get DataflowJobs.
type DataflowJobNamespaceLister interface {
	// List lists all DataflowJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DataflowJob, err error)
	// Get retrieves the DataflowJob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DataflowJob, error)
	DataflowJobNamespaceListerExpansion
}

// dataflowJobNamespaceLister implements the DataflowJobNamespaceLister
// interface.
type dataflowJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataflowJobs in the indexer for a given namespace.
func (s dataflowJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DataflowJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataflowJob))
	})
	return ret, err
}

// Get retrieves the DataflowJob from the indexer for a given namespace and name.
func (s dataflowJobNamespaceLister) Get(name string) (*v1alpha1.DataflowJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dataflowjob"), name)
	}
	return obj.(*v1alpha1.DataflowJob), nil
}