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

// AwsEc2TransitGatewayRouteTableLister helps list AwsEc2TransitGatewayRouteTables.
type AwsEc2TransitGatewayRouteTableLister interface {
	// List lists all AwsEc2TransitGatewayRouteTables in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsEc2TransitGatewayRouteTable, err error)
	// Get retrieves the AwsEc2TransitGatewayRouteTable from the index for a given name.
	Get(name string) (*v1alpha1.AwsEc2TransitGatewayRouteTable, error)
	AwsEc2TransitGatewayRouteTableListerExpansion
}

// awsEc2TransitGatewayRouteTableLister implements the AwsEc2TransitGatewayRouteTableLister interface.
type awsEc2TransitGatewayRouteTableLister struct {
	indexer cache.Indexer
}

// NewAwsEc2TransitGatewayRouteTableLister returns a new AwsEc2TransitGatewayRouteTableLister.
func NewAwsEc2TransitGatewayRouteTableLister(indexer cache.Indexer) AwsEc2TransitGatewayRouteTableLister {
	return &awsEc2TransitGatewayRouteTableLister{indexer: indexer}
}

// List lists all AwsEc2TransitGatewayRouteTables in the indexer.
func (s *awsEc2TransitGatewayRouteTableLister) List(selector labels.Selector) (ret []*v1alpha1.AwsEc2TransitGatewayRouteTable, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsEc2TransitGatewayRouteTable))
	})
	return ret, err
}

// Get retrieves the AwsEc2TransitGatewayRouteTable from the index for a given name.
func (s *awsEc2TransitGatewayRouteTableLister) Get(name string) (*v1alpha1.AwsEc2TransitGatewayRouteTable, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsec2transitgatewayroutetable"), name)
	}
	return obj.(*v1alpha1.AwsEc2TransitGatewayRouteTable), nil
}
