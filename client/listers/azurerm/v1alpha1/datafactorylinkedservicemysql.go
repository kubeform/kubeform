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

// DataFactoryLinkedServiceMysqlLister helps list DataFactoryLinkedServiceMysqls.
type DataFactoryLinkedServiceMysqlLister interface {
	// List lists all DataFactoryLinkedServiceMysqls in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DataFactoryLinkedServiceMysql, err error)
	// DataFactoryLinkedServiceMysqls returns an object that can list and get DataFactoryLinkedServiceMysqls.
	DataFactoryLinkedServiceMysqls(namespace string) DataFactoryLinkedServiceMysqlNamespaceLister
	DataFactoryLinkedServiceMysqlListerExpansion
}

// dataFactoryLinkedServiceMysqlLister implements the DataFactoryLinkedServiceMysqlLister interface.
type dataFactoryLinkedServiceMysqlLister struct {
	indexer cache.Indexer
}

// NewDataFactoryLinkedServiceMysqlLister returns a new DataFactoryLinkedServiceMysqlLister.
func NewDataFactoryLinkedServiceMysqlLister(indexer cache.Indexer) DataFactoryLinkedServiceMysqlLister {
	return &dataFactoryLinkedServiceMysqlLister{indexer: indexer}
}

// List lists all DataFactoryLinkedServiceMysqls in the indexer.
func (s *dataFactoryLinkedServiceMysqlLister) List(selector labels.Selector) (ret []*v1alpha1.DataFactoryLinkedServiceMysql, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataFactoryLinkedServiceMysql))
	})
	return ret, err
}

// DataFactoryLinkedServiceMysqls returns an object that can list and get DataFactoryLinkedServiceMysqls.
func (s *dataFactoryLinkedServiceMysqlLister) DataFactoryLinkedServiceMysqls(namespace string) DataFactoryLinkedServiceMysqlNamespaceLister {
	return dataFactoryLinkedServiceMysqlNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataFactoryLinkedServiceMysqlNamespaceLister helps list and get DataFactoryLinkedServiceMysqls.
type DataFactoryLinkedServiceMysqlNamespaceLister interface {
	// List lists all DataFactoryLinkedServiceMysqls in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DataFactoryLinkedServiceMysql, err error)
	// Get retrieves the DataFactoryLinkedServiceMysql from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DataFactoryLinkedServiceMysql, error)
	DataFactoryLinkedServiceMysqlNamespaceListerExpansion
}

// dataFactoryLinkedServiceMysqlNamespaceLister implements the DataFactoryLinkedServiceMysqlNamespaceLister
// interface.
type dataFactoryLinkedServiceMysqlNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataFactoryLinkedServiceMysqls in the indexer for a given namespace.
func (s dataFactoryLinkedServiceMysqlNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DataFactoryLinkedServiceMysql, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataFactoryLinkedServiceMysql))
	})
	return ret, err
}

// Get retrieves the DataFactoryLinkedServiceMysql from the indexer for a given namespace and name.
func (s dataFactoryLinkedServiceMysqlNamespaceLister) Get(name string) (*v1alpha1.DataFactoryLinkedServiceMysql, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("datafactorylinkedservicemysql"), name)
	}
	return obj.(*v1alpha1.DataFactoryLinkedServiceMysql), nil
}
