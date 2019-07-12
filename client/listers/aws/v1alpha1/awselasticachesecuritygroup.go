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

// AwsElasticacheSecurityGroupLister helps list AwsElasticacheSecurityGroups.
type AwsElasticacheSecurityGroupLister interface {
	// List lists all AwsElasticacheSecurityGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsElasticacheSecurityGroup, err error)
	// Get retrieves the AwsElasticacheSecurityGroup from the index for a given name.
	Get(name string) (*v1alpha1.AwsElasticacheSecurityGroup, error)
	AwsElasticacheSecurityGroupListerExpansion
}

// awsElasticacheSecurityGroupLister implements the AwsElasticacheSecurityGroupLister interface.
type awsElasticacheSecurityGroupLister struct {
	indexer cache.Indexer
}

// NewAwsElasticacheSecurityGroupLister returns a new AwsElasticacheSecurityGroupLister.
func NewAwsElasticacheSecurityGroupLister(indexer cache.Indexer) AwsElasticacheSecurityGroupLister {
	return &awsElasticacheSecurityGroupLister{indexer: indexer}
}

// List lists all AwsElasticacheSecurityGroups in the indexer.
func (s *awsElasticacheSecurityGroupLister) List(selector labels.Selector) (ret []*v1alpha1.AwsElasticacheSecurityGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsElasticacheSecurityGroup))
	})
	return ret, err
}

// Get retrieves the AwsElasticacheSecurityGroup from the index for a given name.
func (s *awsElasticacheSecurityGroupLister) Get(name string) (*v1alpha1.AwsElasticacheSecurityGroup, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awselasticachesecuritygroup"), name)
	}
	return obj.(*v1alpha1.AwsElasticacheSecurityGroup), nil
}
