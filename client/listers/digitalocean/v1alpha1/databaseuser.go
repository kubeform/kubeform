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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DatabaseUserLister helps list DatabaseUsers.
type DatabaseUserLister interface {
	// List lists all DatabaseUsers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DatabaseUser, err error)
	// DatabaseUsers returns an object that can list and get DatabaseUsers.
	DatabaseUsers(namespace string) DatabaseUserNamespaceLister
	DatabaseUserListerExpansion
}

// databaseUserLister implements the DatabaseUserLister interface.
type databaseUserLister struct {
	indexer cache.Indexer
}

// NewDatabaseUserLister returns a new DatabaseUserLister.
func NewDatabaseUserLister(indexer cache.Indexer) DatabaseUserLister {
	return &databaseUserLister{indexer: indexer}
}

// List lists all DatabaseUsers in the indexer.
func (s *databaseUserLister) List(selector labels.Selector) (ret []*v1alpha1.DatabaseUser, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DatabaseUser))
	})
	return ret, err
}

// DatabaseUsers returns an object that can list and get DatabaseUsers.
func (s *databaseUserLister) DatabaseUsers(namespace string) DatabaseUserNamespaceLister {
	return databaseUserNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DatabaseUserNamespaceLister helps list and get DatabaseUsers.
type DatabaseUserNamespaceLister interface {
	// List lists all DatabaseUsers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DatabaseUser, err error)
	// Get retrieves the DatabaseUser from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DatabaseUser, error)
	DatabaseUserNamespaceListerExpansion
}

// databaseUserNamespaceLister implements the DatabaseUserNamespaceLister
// interface.
type databaseUserNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DatabaseUsers in the indexer for a given namespace.
func (s databaseUserNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DatabaseUser, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DatabaseUser))
	})
	return ret, err
}

// Get retrieves the DatabaseUser from the indexer for a given namespace and name.
func (s databaseUserNamespaceLister) Get(name string) (*v1alpha1.DatabaseUser, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("databaseuser"), name)
	}
	return obj.(*v1alpha1.DatabaseUser), nil
}