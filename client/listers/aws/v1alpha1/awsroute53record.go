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

// AwsRoute53RecordLister helps list AwsRoute53Records.
type AwsRoute53RecordLister interface {
	// List lists all AwsRoute53Records in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsRoute53Record, err error)
	// Get retrieves the AwsRoute53Record from the index for a given name.
	Get(name string) (*v1alpha1.AwsRoute53Record, error)
	AwsRoute53RecordListerExpansion
}

// awsRoute53RecordLister implements the AwsRoute53RecordLister interface.
type awsRoute53RecordLister struct {
	indexer cache.Indexer
}

// NewAwsRoute53RecordLister returns a new AwsRoute53RecordLister.
func NewAwsRoute53RecordLister(indexer cache.Indexer) AwsRoute53RecordLister {
	return &awsRoute53RecordLister{indexer: indexer}
}

// List lists all AwsRoute53Records in the indexer.
func (s *awsRoute53RecordLister) List(selector labels.Selector) (ret []*v1alpha1.AwsRoute53Record, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsRoute53Record))
	})
	return ret, err
}

// Get retrieves the AwsRoute53Record from the index for a given name.
func (s *awsRoute53RecordLister) Get(name string) (*v1alpha1.AwsRoute53Record, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsroute53record"), name)
	}
	return obj.(*v1alpha1.AwsRoute53Record), nil
}
