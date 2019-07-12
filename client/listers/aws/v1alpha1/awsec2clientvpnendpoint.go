/*
Copyright 2019 The KubeDB Authors.

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

// AwsEc2ClientVpnEndpointLister helps list AwsEc2ClientVpnEndpoints.
type AwsEc2ClientVpnEndpointLister interface {
	// List lists all AwsEc2ClientVpnEndpoints in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsEc2ClientVpnEndpoint, err error)
	// Get retrieves the AwsEc2ClientVpnEndpoint from the index for a given name.
	Get(name string) (*v1alpha1.AwsEc2ClientVpnEndpoint, error)
	AwsEc2ClientVpnEndpointListerExpansion
}

// awsEc2ClientVpnEndpointLister implements the AwsEc2ClientVpnEndpointLister interface.
type awsEc2ClientVpnEndpointLister struct {
	indexer cache.Indexer
}

// NewAwsEc2ClientVpnEndpointLister returns a new AwsEc2ClientVpnEndpointLister.
func NewAwsEc2ClientVpnEndpointLister(indexer cache.Indexer) AwsEc2ClientVpnEndpointLister {
	return &awsEc2ClientVpnEndpointLister{indexer: indexer}
}

// List lists all AwsEc2ClientVpnEndpoints in the indexer.
func (s *awsEc2ClientVpnEndpointLister) List(selector labels.Selector) (ret []*v1alpha1.AwsEc2ClientVpnEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsEc2ClientVpnEndpoint))
	})
	return ret, err
}

// Get retrieves the AwsEc2ClientVpnEndpoint from the index for a given name.
func (s *awsEc2ClientVpnEndpointLister) Get(name string) (*v1alpha1.AwsEc2ClientVpnEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsec2clientvpnendpoint"), name)
	}
	return obj.(*v1alpha1.AwsEc2ClientVpnEndpoint), nil
}
