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

// AwsRoute53ResolverEndpointLister helps list AwsRoute53ResolverEndpoints.
type AwsRoute53ResolverEndpointLister interface {
	// List lists all AwsRoute53ResolverEndpoints in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsRoute53ResolverEndpoint, err error)
	// Get retrieves the AwsRoute53ResolverEndpoint from the index for a given name.
	Get(name string) (*v1alpha1.AwsRoute53ResolverEndpoint, error)
	AwsRoute53ResolverEndpointListerExpansion
}

// awsRoute53ResolverEndpointLister implements the AwsRoute53ResolverEndpointLister interface.
type awsRoute53ResolverEndpointLister struct {
	indexer cache.Indexer
}

// NewAwsRoute53ResolverEndpointLister returns a new AwsRoute53ResolverEndpointLister.
func NewAwsRoute53ResolverEndpointLister(indexer cache.Indexer) AwsRoute53ResolverEndpointLister {
	return &awsRoute53ResolverEndpointLister{indexer: indexer}
}

// List lists all AwsRoute53ResolverEndpoints in the indexer.
func (s *awsRoute53ResolverEndpointLister) List(selector labels.Selector) (ret []*v1alpha1.AwsRoute53ResolverEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsRoute53ResolverEndpoint))
	})
	return ret, err
}

// Get retrieves the AwsRoute53ResolverEndpoint from the index for a given name.
func (s *awsRoute53ResolverEndpointLister) Get(name string) (*v1alpha1.AwsRoute53ResolverEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsroute53resolverendpoint"), name)
	}
	return obj.(*v1alpha1.AwsRoute53ResolverEndpoint), nil
}
