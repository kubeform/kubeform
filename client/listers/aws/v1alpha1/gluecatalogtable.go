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

// GlueCatalogTableLister helps list GlueCatalogTables.
type GlueCatalogTableLister interface {
	// List lists all GlueCatalogTables in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GlueCatalogTable, err error)
	// GlueCatalogTables returns an object that can list and get GlueCatalogTables.
	GlueCatalogTables(namespace string) GlueCatalogTableNamespaceLister
	GlueCatalogTableListerExpansion
}

// glueCatalogTableLister implements the GlueCatalogTableLister interface.
type glueCatalogTableLister struct {
	indexer cache.Indexer
}

// NewGlueCatalogTableLister returns a new GlueCatalogTableLister.
func NewGlueCatalogTableLister(indexer cache.Indexer) GlueCatalogTableLister {
	return &glueCatalogTableLister{indexer: indexer}
}

// List lists all GlueCatalogTables in the indexer.
func (s *glueCatalogTableLister) List(selector labels.Selector) (ret []*v1alpha1.GlueCatalogTable, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlueCatalogTable))
	})
	return ret, err
}

// GlueCatalogTables returns an object that can list and get GlueCatalogTables.
func (s *glueCatalogTableLister) GlueCatalogTables(namespace string) GlueCatalogTableNamespaceLister {
	return glueCatalogTableNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GlueCatalogTableNamespaceLister helps list and get GlueCatalogTables.
type GlueCatalogTableNamespaceLister interface {
	// List lists all GlueCatalogTables in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.GlueCatalogTable, err error)
	// Get retrieves the GlueCatalogTable from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.GlueCatalogTable, error)
	GlueCatalogTableNamespaceListerExpansion
}

// glueCatalogTableNamespaceLister implements the GlueCatalogTableNamespaceLister
// interface.
type glueCatalogTableNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GlueCatalogTables in the indexer for a given namespace.
func (s glueCatalogTableNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GlueCatalogTable, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlueCatalogTable))
	})
	return ret, err
}

// Get retrieves the GlueCatalogTable from the indexer for a given namespace and name.
func (s glueCatalogTableNamespaceLister) Get(name string) (*v1alpha1.GlueCatalogTable, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("gluecatalogtable"), name)
	}
	return obj.(*v1alpha1.GlueCatalogTable), nil
}
