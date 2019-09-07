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

// AlbListenerLister helps list AlbListeners.
type AlbListenerLister interface {
	// List lists all AlbListeners in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AlbListener, err error)
	// AlbListeners returns an object that can list and get AlbListeners.
	AlbListeners(namespace string) AlbListenerNamespaceLister
	AlbListenerListerExpansion
}

// albListenerLister implements the AlbListenerLister interface.
type albListenerLister struct {
	indexer cache.Indexer
}

// NewAlbListenerLister returns a new AlbListenerLister.
func NewAlbListenerLister(indexer cache.Indexer) AlbListenerLister {
	return &albListenerLister{indexer: indexer}
}

// List lists all AlbListeners in the indexer.
func (s *albListenerLister) List(selector labels.Selector) (ret []*v1alpha1.AlbListener, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AlbListener))
	})
	return ret, err
}

// AlbListeners returns an object that can list and get AlbListeners.
func (s *albListenerLister) AlbListeners(namespace string) AlbListenerNamespaceLister {
	return albListenerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AlbListenerNamespaceLister helps list and get AlbListeners.
type AlbListenerNamespaceLister interface {
	// List lists all AlbListeners in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AlbListener, err error)
	// Get retrieves the AlbListener from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AlbListener, error)
	AlbListenerNamespaceListerExpansion
}

// albListenerNamespaceLister implements the AlbListenerNamespaceLister
// interface.
type albListenerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AlbListeners in the indexer for a given namespace.
func (s albListenerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AlbListener, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AlbListener))
	})
	return ret, err
}

// Get retrieves the AlbListener from the indexer for a given namespace and name.
func (s albListenerNamespaceLister) Get(name string) (*v1alpha1.AlbListener, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("alblistener"), name)
	}
	return obj.(*v1alpha1.AlbListener), nil
}