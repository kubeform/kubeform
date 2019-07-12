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

// AwsCloudwatchLogSubscriptionFilterLister helps list AwsCloudwatchLogSubscriptionFilters.
type AwsCloudwatchLogSubscriptionFilterLister interface {
	// List lists all AwsCloudwatchLogSubscriptionFilters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsCloudwatchLogSubscriptionFilter, err error)
	// Get retrieves the AwsCloudwatchLogSubscriptionFilter from the index for a given name.
	Get(name string) (*v1alpha1.AwsCloudwatchLogSubscriptionFilter, error)
	AwsCloudwatchLogSubscriptionFilterListerExpansion
}

// awsCloudwatchLogSubscriptionFilterLister implements the AwsCloudwatchLogSubscriptionFilterLister interface.
type awsCloudwatchLogSubscriptionFilterLister struct {
	indexer cache.Indexer
}

// NewAwsCloudwatchLogSubscriptionFilterLister returns a new AwsCloudwatchLogSubscriptionFilterLister.
func NewAwsCloudwatchLogSubscriptionFilterLister(indexer cache.Indexer) AwsCloudwatchLogSubscriptionFilterLister {
	return &awsCloudwatchLogSubscriptionFilterLister{indexer: indexer}
}

// List lists all AwsCloudwatchLogSubscriptionFilters in the indexer.
func (s *awsCloudwatchLogSubscriptionFilterLister) List(selector labels.Selector) (ret []*v1alpha1.AwsCloudwatchLogSubscriptionFilter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsCloudwatchLogSubscriptionFilter))
	})
	return ret, err
}

// Get retrieves the AwsCloudwatchLogSubscriptionFilter from the index for a given name.
func (s *awsCloudwatchLogSubscriptionFilterLister) Get(name string) (*v1alpha1.AwsCloudwatchLogSubscriptionFilter, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awscloudwatchlogsubscriptionfilter"), name)
	}
	return obj.(*v1alpha1.AwsCloudwatchLogSubscriptionFilter), nil
}
