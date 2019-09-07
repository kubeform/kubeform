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

// CognitoResourceServerLister helps list CognitoResourceServers.
type CognitoResourceServerLister interface {
	// List lists all CognitoResourceServers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CognitoResourceServer, err error)
	// CognitoResourceServers returns an object that can list and get CognitoResourceServers.
	CognitoResourceServers(namespace string) CognitoResourceServerNamespaceLister
	CognitoResourceServerListerExpansion
}

// cognitoResourceServerLister implements the CognitoResourceServerLister interface.
type cognitoResourceServerLister struct {
	indexer cache.Indexer
}

// NewCognitoResourceServerLister returns a new CognitoResourceServerLister.
func NewCognitoResourceServerLister(indexer cache.Indexer) CognitoResourceServerLister {
	return &cognitoResourceServerLister{indexer: indexer}
}

// List lists all CognitoResourceServers in the indexer.
func (s *cognitoResourceServerLister) List(selector labels.Selector) (ret []*v1alpha1.CognitoResourceServer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CognitoResourceServer))
	})
	return ret, err
}

// CognitoResourceServers returns an object that can list and get CognitoResourceServers.
func (s *cognitoResourceServerLister) CognitoResourceServers(namespace string) CognitoResourceServerNamespaceLister {
	return cognitoResourceServerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CognitoResourceServerNamespaceLister helps list and get CognitoResourceServers.
type CognitoResourceServerNamespaceLister interface {
	// List lists all CognitoResourceServers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CognitoResourceServer, err error)
	// Get retrieves the CognitoResourceServer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CognitoResourceServer, error)
	CognitoResourceServerNamespaceListerExpansion
}

// cognitoResourceServerNamespaceLister implements the CognitoResourceServerNamespaceLister
// interface.
type cognitoResourceServerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CognitoResourceServers in the indexer for a given namespace.
func (s cognitoResourceServerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CognitoResourceServer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CognitoResourceServer))
	})
	return ret, err
}

// Get retrieves the CognitoResourceServer from the indexer for a given namespace and name.
func (s cognitoResourceServerNamespaceLister) Get(name string) (*v1alpha1.CognitoResourceServer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cognitoresourceserver"), name)
	}
	return obj.(*v1alpha1.CognitoResourceServer), nil
}