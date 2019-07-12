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

// AwsOpsworksInstanceLister helps list AwsOpsworksInstances.
type AwsOpsworksInstanceLister interface {
	// List lists all AwsOpsworksInstances in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsOpsworksInstance, err error)
	// Get retrieves the AwsOpsworksInstance from the index for a given name.
	Get(name string) (*v1alpha1.AwsOpsworksInstance, error)
	AwsOpsworksInstanceListerExpansion
}

// awsOpsworksInstanceLister implements the AwsOpsworksInstanceLister interface.
type awsOpsworksInstanceLister struct {
	indexer cache.Indexer
}

// NewAwsOpsworksInstanceLister returns a new AwsOpsworksInstanceLister.
func NewAwsOpsworksInstanceLister(indexer cache.Indexer) AwsOpsworksInstanceLister {
	return &awsOpsworksInstanceLister{indexer: indexer}
}

// List lists all AwsOpsworksInstances in the indexer.
func (s *awsOpsworksInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsOpsworksInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsOpsworksInstance))
	})
	return ret, err
}

// Get retrieves the AwsOpsworksInstance from the index for a given name.
func (s *awsOpsworksInstanceLister) Get(name string) (*v1alpha1.AwsOpsworksInstance, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsopsworksinstance"), name)
	}
	return obj.(*v1alpha1.AwsOpsworksInstance), nil
}
