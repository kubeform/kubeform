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

// DnsSrvRecordLister helps list DnsSrvRecords.
type DnsSrvRecordLister interface {
	// List lists all DnsSrvRecords in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DnsSrvRecord, err error)
	// DnsSrvRecords returns an object that can list and get DnsSrvRecords.
	DnsSrvRecords(namespace string) DnsSrvRecordNamespaceLister
	DnsSrvRecordListerExpansion
}

// dnsSrvRecordLister implements the DnsSrvRecordLister interface.
type dnsSrvRecordLister struct {
	indexer cache.Indexer
}

// NewDnsSrvRecordLister returns a new DnsSrvRecordLister.
func NewDnsSrvRecordLister(indexer cache.Indexer) DnsSrvRecordLister {
	return &dnsSrvRecordLister{indexer: indexer}
}

// List lists all DnsSrvRecords in the indexer.
func (s *dnsSrvRecordLister) List(selector labels.Selector) (ret []*v1alpha1.DnsSrvRecord, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsSrvRecord))
	})
	return ret, err
}

// DnsSrvRecords returns an object that can list and get DnsSrvRecords.
func (s *dnsSrvRecordLister) DnsSrvRecords(namespace string) DnsSrvRecordNamespaceLister {
	return dnsSrvRecordNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DnsSrvRecordNamespaceLister helps list and get DnsSrvRecords.
type DnsSrvRecordNamespaceLister interface {
	// List lists all DnsSrvRecords in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DnsSrvRecord, err error)
	// Get retrieves the DnsSrvRecord from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DnsSrvRecord, error)
	DnsSrvRecordNamespaceListerExpansion
}

// dnsSrvRecordNamespaceLister implements the DnsSrvRecordNamespaceLister
// interface.
type dnsSrvRecordNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DnsSrvRecords in the indexer for a given namespace.
func (s dnsSrvRecordNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DnsSrvRecord, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsSrvRecord))
	})
	return ret, err
}

// Get retrieves the DnsSrvRecord from the indexer for a given namespace and name.
func (s dnsSrvRecordNamespaceLister) Get(name string) (*v1alpha1.DnsSrvRecord, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dnssrvrecord"), name)
	}
	return obj.(*v1alpha1.DnsSrvRecord), nil
}