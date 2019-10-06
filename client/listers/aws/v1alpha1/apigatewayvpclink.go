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

// ApiGatewayVpcLinkLister helps list ApiGatewayVpcLinks.
type ApiGatewayVpcLinkLister interface {
	// List lists all ApiGatewayVpcLinks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayVpcLink, err error)
	// ApiGatewayVpcLinks returns an object that can list and get ApiGatewayVpcLinks.
	ApiGatewayVpcLinks(namespace string) ApiGatewayVpcLinkNamespaceLister
	ApiGatewayVpcLinkListerExpansion
}

// apiGatewayVpcLinkLister implements the ApiGatewayVpcLinkLister interface.
type apiGatewayVpcLinkLister struct {
	indexer cache.Indexer
}

// NewApiGatewayVpcLinkLister returns a new ApiGatewayVpcLinkLister.
func NewApiGatewayVpcLinkLister(indexer cache.Indexer) ApiGatewayVpcLinkLister {
	return &apiGatewayVpcLinkLister{indexer: indexer}
}

// List lists all ApiGatewayVpcLinks in the indexer.
func (s *apiGatewayVpcLinkLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayVpcLink, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayVpcLink))
	})
	return ret, err
}

// ApiGatewayVpcLinks returns an object that can list and get ApiGatewayVpcLinks.
func (s *apiGatewayVpcLinkLister) ApiGatewayVpcLinks(namespace string) ApiGatewayVpcLinkNamespaceLister {
	return apiGatewayVpcLinkNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiGatewayVpcLinkNamespaceLister helps list and get ApiGatewayVpcLinks.
type ApiGatewayVpcLinkNamespaceLister interface {
	// List lists all ApiGatewayVpcLinks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayVpcLink, err error)
	// Get retrieves the ApiGatewayVpcLink from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiGatewayVpcLink, error)
	ApiGatewayVpcLinkNamespaceListerExpansion
}

// apiGatewayVpcLinkNamespaceLister implements the ApiGatewayVpcLinkNamespaceLister
// interface.
type apiGatewayVpcLinkNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiGatewayVpcLinks in the indexer for a given namespace.
func (s apiGatewayVpcLinkNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayVpcLink, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayVpcLink))
	})
	return ret, err
}

// Get retrieves the ApiGatewayVpcLink from the indexer for a given namespace and name.
func (s apiGatewayVpcLinkNamespaceLister) Get(name string) (*v1alpha1.ApiGatewayVpcLink, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apigatewayvpclink"), name)
	}
	return obj.(*v1alpha1.ApiGatewayVpcLink), nil
}
