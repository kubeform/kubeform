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

// AwsDefaultVpcDhcpOptionsLister helps list AwsDefaultVpcDhcpOptionses.
type AwsDefaultVpcDhcpOptionsLister interface {
	// List lists all AwsDefaultVpcDhcpOptionses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsDefaultVpcDhcpOptions, err error)
	// Get retrieves the AwsDefaultVpcDhcpOptions from the index for a given name.
	Get(name string) (*v1alpha1.AwsDefaultVpcDhcpOptions, error)
	AwsDefaultVpcDhcpOptionsListerExpansion
}

// awsDefaultVpcDhcpOptionsLister implements the AwsDefaultVpcDhcpOptionsLister interface.
type awsDefaultVpcDhcpOptionsLister struct {
	indexer cache.Indexer
}

// NewAwsDefaultVpcDhcpOptionsLister returns a new AwsDefaultVpcDhcpOptionsLister.
func NewAwsDefaultVpcDhcpOptionsLister(indexer cache.Indexer) AwsDefaultVpcDhcpOptionsLister {
	return &awsDefaultVpcDhcpOptionsLister{indexer: indexer}
}

// List lists all AwsDefaultVpcDhcpOptionses in the indexer.
func (s *awsDefaultVpcDhcpOptionsLister) List(selector labels.Selector) (ret []*v1alpha1.AwsDefaultVpcDhcpOptions, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsDefaultVpcDhcpOptions))
	})
	return ret, err
}

// Get retrieves the AwsDefaultVpcDhcpOptions from the index for a given name.
func (s *awsDefaultVpcDhcpOptionsLister) Get(name string) (*v1alpha1.AwsDefaultVpcDhcpOptions, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsdefaultvpcdhcpoptions"), name)
	}
	return obj.(*v1alpha1.AwsDefaultVpcDhcpOptions), nil
}
