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

// AlbListenerRuleLister helps list AlbListenerRules.
type AlbListenerRuleLister interface {
	// List lists all AlbListenerRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AlbListenerRule, err error)
	// AlbListenerRules returns an object that can list and get AlbListenerRules.
	AlbListenerRules(namespace string) AlbListenerRuleNamespaceLister
	AlbListenerRuleListerExpansion
}

// albListenerRuleLister implements the AlbListenerRuleLister interface.
type albListenerRuleLister struct {
	indexer cache.Indexer
}

// NewAlbListenerRuleLister returns a new AlbListenerRuleLister.
func NewAlbListenerRuleLister(indexer cache.Indexer) AlbListenerRuleLister {
	return &albListenerRuleLister{indexer: indexer}
}

// List lists all AlbListenerRules in the indexer.
func (s *albListenerRuleLister) List(selector labels.Selector) (ret []*v1alpha1.AlbListenerRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AlbListenerRule))
	})
	return ret, err
}

// AlbListenerRules returns an object that can list and get AlbListenerRules.
func (s *albListenerRuleLister) AlbListenerRules(namespace string) AlbListenerRuleNamespaceLister {
	return albListenerRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AlbListenerRuleNamespaceLister helps list and get AlbListenerRules.
type AlbListenerRuleNamespaceLister interface {
	// List lists all AlbListenerRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AlbListenerRule, err error)
	// Get retrieves the AlbListenerRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AlbListenerRule, error)
	AlbListenerRuleNamespaceListerExpansion
}

// albListenerRuleNamespaceLister implements the AlbListenerRuleNamespaceLister
// interface.
type albListenerRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AlbListenerRules in the indexer for a given namespace.
func (s albListenerRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AlbListenerRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AlbListenerRule))
	})
	return ret, err
}

// Get retrieves the AlbListenerRule from the indexer for a given namespace and name.
func (s albListenerRuleNamespaceLister) Get(name string) (*v1alpha1.AlbListenerRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("alblistenerrule"), name)
	}
	return obj.(*v1alpha1.AlbListenerRule), nil
}
