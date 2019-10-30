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

// LbCookieStickinessPolicyLister helps list LbCookieStickinessPolicies.
type LbCookieStickinessPolicyLister interface {
	// List lists all LbCookieStickinessPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LbCookieStickinessPolicy, err error)
	// LbCookieStickinessPolicies returns an object that can list and get LbCookieStickinessPolicies.
	LbCookieStickinessPolicies(namespace string) LbCookieStickinessPolicyNamespaceLister
	LbCookieStickinessPolicyListerExpansion
}

// lbCookieStickinessPolicyLister implements the LbCookieStickinessPolicyLister interface.
type lbCookieStickinessPolicyLister struct {
	indexer cache.Indexer
}

// NewLbCookieStickinessPolicyLister returns a new LbCookieStickinessPolicyLister.
func NewLbCookieStickinessPolicyLister(indexer cache.Indexer) LbCookieStickinessPolicyLister {
	return &lbCookieStickinessPolicyLister{indexer: indexer}
}

// List lists all LbCookieStickinessPolicies in the indexer.
func (s *lbCookieStickinessPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.LbCookieStickinessPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LbCookieStickinessPolicy))
	})
	return ret, err
}

// LbCookieStickinessPolicies returns an object that can list and get LbCookieStickinessPolicies.
func (s *lbCookieStickinessPolicyLister) LbCookieStickinessPolicies(namespace string) LbCookieStickinessPolicyNamespaceLister {
	return lbCookieStickinessPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LbCookieStickinessPolicyNamespaceLister helps list and get LbCookieStickinessPolicies.
type LbCookieStickinessPolicyNamespaceLister interface {
	// List lists all LbCookieStickinessPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LbCookieStickinessPolicy, err error)
	// Get retrieves the LbCookieStickinessPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LbCookieStickinessPolicy, error)
	LbCookieStickinessPolicyNamespaceListerExpansion
}

// lbCookieStickinessPolicyNamespaceLister implements the LbCookieStickinessPolicyNamespaceLister
// interface.
type lbCookieStickinessPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LbCookieStickinessPolicies in the indexer for a given namespace.
func (s lbCookieStickinessPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LbCookieStickinessPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LbCookieStickinessPolicy))
	})
	return ret, err
}

// Get retrieves the LbCookieStickinessPolicy from the indexer for a given namespace and name.
func (s lbCookieStickinessPolicyNamespaceLister) Get(name string) (*v1alpha1.LbCookieStickinessPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("lbcookiestickinesspolicy"), name)
	}
	return obj.(*v1alpha1.LbCookieStickinessPolicy), nil
}
