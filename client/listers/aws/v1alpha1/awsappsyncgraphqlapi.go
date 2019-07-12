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

// AwsAppsyncGraphqlApiLister helps list AwsAppsyncGraphqlApis.
type AwsAppsyncGraphqlApiLister interface {
	// List lists all AwsAppsyncGraphqlApis in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsAppsyncGraphqlApi, err error)
	// Get retrieves the AwsAppsyncGraphqlApi from the index for a given name.
	Get(name string) (*v1alpha1.AwsAppsyncGraphqlApi, error)
	AwsAppsyncGraphqlApiListerExpansion
}

// awsAppsyncGraphqlApiLister implements the AwsAppsyncGraphqlApiLister interface.
type awsAppsyncGraphqlApiLister struct {
	indexer cache.Indexer
}

// NewAwsAppsyncGraphqlApiLister returns a new AwsAppsyncGraphqlApiLister.
func NewAwsAppsyncGraphqlApiLister(indexer cache.Indexer) AwsAppsyncGraphqlApiLister {
	return &awsAppsyncGraphqlApiLister{indexer: indexer}
}

// List lists all AwsAppsyncGraphqlApis in the indexer.
func (s *awsAppsyncGraphqlApiLister) List(selector labels.Selector) (ret []*v1alpha1.AwsAppsyncGraphqlApi, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsAppsyncGraphqlApi))
	})
	return ret, err
}

// Get retrieves the AwsAppsyncGraphqlApi from the index for a given name.
func (s *awsAppsyncGraphqlApiLister) Get(name string) (*v1alpha1.AwsAppsyncGraphqlApi, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsappsyncgraphqlapi"), name)
	}
	return obj.(*v1alpha1.AwsAppsyncGraphqlApi), nil
}
