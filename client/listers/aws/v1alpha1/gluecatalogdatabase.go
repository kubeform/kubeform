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

// GlueCatalogDatabaseLister helps list GlueCatalogDatabases.
type GlueCatalogDatabaseLister interface {
	// List lists all GlueCatalogDatabases in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GlueCatalogDatabase, err error)
	// GlueCatalogDatabases returns an object that can list and get GlueCatalogDatabases.
	GlueCatalogDatabases(namespace string) GlueCatalogDatabaseNamespaceLister
	GlueCatalogDatabaseListerExpansion
}

// glueCatalogDatabaseLister implements the GlueCatalogDatabaseLister interface.
type glueCatalogDatabaseLister struct {
	indexer cache.Indexer
}

// NewGlueCatalogDatabaseLister returns a new GlueCatalogDatabaseLister.
func NewGlueCatalogDatabaseLister(indexer cache.Indexer) GlueCatalogDatabaseLister {
	return &glueCatalogDatabaseLister{indexer: indexer}
}

// List lists all GlueCatalogDatabases in the indexer.
func (s *glueCatalogDatabaseLister) List(selector labels.Selector) (ret []*v1alpha1.GlueCatalogDatabase, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlueCatalogDatabase))
	})
	return ret, err
}

// GlueCatalogDatabases returns an object that can list and get GlueCatalogDatabases.
func (s *glueCatalogDatabaseLister) GlueCatalogDatabases(namespace string) GlueCatalogDatabaseNamespaceLister {
	return glueCatalogDatabaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GlueCatalogDatabaseNamespaceLister helps list and get GlueCatalogDatabases.
type GlueCatalogDatabaseNamespaceLister interface {
	// List lists all GlueCatalogDatabases in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.GlueCatalogDatabase, err error)
	// Get retrieves the GlueCatalogDatabase from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.GlueCatalogDatabase, error)
	GlueCatalogDatabaseNamespaceListerExpansion
}

// glueCatalogDatabaseNamespaceLister implements the GlueCatalogDatabaseNamespaceLister
// interface.
type glueCatalogDatabaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GlueCatalogDatabases in the indexer for a given namespace.
func (s glueCatalogDatabaseNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GlueCatalogDatabase, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlueCatalogDatabase))
	})
	return ret, err
}

// Get retrieves the GlueCatalogDatabase from the indexer for a given namespace and name.
func (s glueCatalogDatabaseNamespaceLister) Get(name string) (*v1alpha1.GlueCatalogDatabase, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("gluecatalogdatabase"), name)
	}
	return obj.(*v1alpha1.GlueCatalogDatabase), nil
}
