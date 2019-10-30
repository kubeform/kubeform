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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DnsPtrRecordLister helps list DnsPtrRecords.
type DnsPtrRecordLister interface {
	// List lists all DnsPtrRecords in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DnsPtrRecord, err error)
	// DnsPtrRecords returns an object that can list and get DnsPtrRecords.
	DnsPtrRecords(namespace string) DnsPtrRecordNamespaceLister
	DnsPtrRecordListerExpansion
}

// dnsPtrRecordLister implements the DnsPtrRecordLister interface.
type dnsPtrRecordLister struct {
	indexer cache.Indexer
}

// NewDnsPtrRecordLister returns a new DnsPtrRecordLister.
func NewDnsPtrRecordLister(indexer cache.Indexer) DnsPtrRecordLister {
	return &dnsPtrRecordLister{indexer: indexer}
}

// List lists all DnsPtrRecords in the indexer.
func (s *dnsPtrRecordLister) List(selector labels.Selector) (ret []*v1alpha1.DnsPtrRecord, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsPtrRecord))
	})
	return ret, err
}

// DnsPtrRecords returns an object that can list and get DnsPtrRecords.
func (s *dnsPtrRecordLister) DnsPtrRecords(namespace string) DnsPtrRecordNamespaceLister {
	return dnsPtrRecordNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DnsPtrRecordNamespaceLister helps list and get DnsPtrRecords.
type DnsPtrRecordNamespaceLister interface {
	// List lists all DnsPtrRecords in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DnsPtrRecord, err error)
	// Get retrieves the DnsPtrRecord from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DnsPtrRecord, error)
	DnsPtrRecordNamespaceListerExpansion
}

// dnsPtrRecordNamespaceLister implements the DnsPtrRecordNamespaceLister
// interface.
type dnsPtrRecordNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DnsPtrRecords in the indexer for a given namespace.
func (s dnsPtrRecordNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DnsPtrRecord, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsPtrRecord))
	})
	return ret, err
}

// Get retrieves the DnsPtrRecord from the indexer for a given namespace and name.
func (s dnsPtrRecordNamespaceLister) Get(name string) (*v1alpha1.DnsPtrRecord, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dnsptrrecord"), name)
	}
	return obj.(*v1alpha1.DnsPtrRecord), nil
}
