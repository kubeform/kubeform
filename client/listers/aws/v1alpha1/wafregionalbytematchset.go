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

// WafregionalByteMatchSetLister helps list WafregionalByteMatchSets.
type WafregionalByteMatchSetLister interface {
	// List lists all WafregionalByteMatchSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.WafregionalByteMatchSet, err error)
	// WafregionalByteMatchSets returns an object that can list and get WafregionalByteMatchSets.
	WafregionalByteMatchSets(namespace string) WafregionalByteMatchSetNamespaceLister
	WafregionalByteMatchSetListerExpansion
}

// wafregionalByteMatchSetLister implements the WafregionalByteMatchSetLister interface.
type wafregionalByteMatchSetLister struct {
	indexer cache.Indexer
}

// NewWafregionalByteMatchSetLister returns a new WafregionalByteMatchSetLister.
func NewWafregionalByteMatchSetLister(indexer cache.Indexer) WafregionalByteMatchSetLister {
	return &wafregionalByteMatchSetLister{indexer: indexer}
}

// List lists all WafregionalByteMatchSets in the indexer.
func (s *wafregionalByteMatchSetLister) List(selector labels.Selector) (ret []*v1alpha1.WafregionalByteMatchSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafregionalByteMatchSet))
	})
	return ret, err
}

// WafregionalByteMatchSets returns an object that can list and get WafregionalByteMatchSets.
func (s *wafregionalByteMatchSetLister) WafregionalByteMatchSets(namespace string) WafregionalByteMatchSetNamespaceLister {
	return wafregionalByteMatchSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WafregionalByteMatchSetNamespaceLister helps list and get WafregionalByteMatchSets.
type WafregionalByteMatchSetNamespaceLister interface {
	// List lists all WafregionalByteMatchSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.WafregionalByteMatchSet, err error)
	// Get retrieves the WafregionalByteMatchSet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.WafregionalByteMatchSet, error)
	WafregionalByteMatchSetNamespaceListerExpansion
}

// wafregionalByteMatchSetNamespaceLister implements the WafregionalByteMatchSetNamespaceLister
// interface.
type wafregionalByteMatchSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WafregionalByteMatchSets in the indexer for a given namespace.
func (s wafregionalByteMatchSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WafregionalByteMatchSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafregionalByteMatchSet))
	})
	return ret, err
}

// Get retrieves the WafregionalByteMatchSet from the indexer for a given namespace and name.
func (s wafregionalByteMatchSetNamespaceLister) Get(name string) (*v1alpha1.WafregionalByteMatchSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("wafregionalbytematchset"), name)
	}
	return obj.(*v1alpha1.WafregionalByteMatchSet), nil
}
