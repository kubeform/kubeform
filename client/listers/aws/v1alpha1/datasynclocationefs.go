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

// DatasyncLocationEfsLister helps list DatasyncLocationEfses.
type DatasyncLocationEfsLister interface {
	// List lists all DatasyncLocationEfses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DatasyncLocationEfs, err error)
	// DatasyncLocationEfses returns an object that can list and get DatasyncLocationEfses.
	DatasyncLocationEfses(namespace string) DatasyncLocationEfsNamespaceLister
	DatasyncLocationEfsListerExpansion
}

// datasyncLocationEfsLister implements the DatasyncLocationEfsLister interface.
type datasyncLocationEfsLister struct {
	indexer cache.Indexer
}

// NewDatasyncLocationEfsLister returns a new DatasyncLocationEfsLister.
func NewDatasyncLocationEfsLister(indexer cache.Indexer) DatasyncLocationEfsLister {
	return &datasyncLocationEfsLister{indexer: indexer}
}

// List lists all DatasyncLocationEfses in the indexer.
func (s *datasyncLocationEfsLister) List(selector labels.Selector) (ret []*v1alpha1.DatasyncLocationEfs, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DatasyncLocationEfs))
	})
	return ret, err
}

// DatasyncLocationEfses returns an object that can list and get DatasyncLocationEfses.
func (s *datasyncLocationEfsLister) DatasyncLocationEfses(namespace string) DatasyncLocationEfsNamespaceLister {
	return datasyncLocationEfsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DatasyncLocationEfsNamespaceLister helps list and get DatasyncLocationEfses.
type DatasyncLocationEfsNamespaceLister interface {
	// List lists all DatasyncLocationEfses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DatasyncLocationEfs, err error)
	// Get retrieves the DatasyncLocationEfs from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DatasyncLocationEfs, error)
	DatasyncLocationEfsNamespaceListerExpansion
}

// datasyncLocationEfsNamespaceLister implements the DatasyncLocationEfsNamespaceLister
// interface.
type datasyncLocationEfsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DatasyncLocationEfses in the indexer for a given namespace.
func (s datasyncLocationEfsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DatasyncLocationEfs, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DatasyncLocationEfs))
	})
	return ret, err
}

// Get retrieves the DatasyncLocationEfs from the indexer for a given namespace and name.
func (s datasyncLocationEfsNamespaceLister) Get(name string) (*v1alpha1.DatasyncLocationEfs, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("datasynclocationefs"), name)
	}
	return obj.(*v1alpha1.DatasyncLocationEfs), nil
}
