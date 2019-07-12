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

// AwsXraySamplingRuleLister helps list AwsXraySamplingRules.
type AwsXraySamplingRuleLister interface {
	// List lists all AwsXraySamplingRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsXraySamplingRule, err error)
	// Get retrieves the AwsXraySamplingRule from the index for a given name.
	Get(name string) (*v1alpha1.AwsXraySamplingRule, error)
	AwsXraySamplingRuleListerExpansion
}

// awsXraySamplingRuleLister implements the AwsXraySamplingRuleLister interface.
type awsXraySamplingRuleLister struct {
	indexer cache.Indexer
}

// NewAwsXraySamplingRuleLister returns a new AwsXraySamplingRuleLister.
func NewAwsXraySamplingRuleLister(indexer cache.Indexer) AwsXraySamplingRuleLister {
	return &awsXraySamplingRuleLister{indexer: indexer}
}

// List lists all AwsXraySamplingRules in the indexer.
func (s *awsXraySamplingRuleLister) List(selector labels.Selector) (ret []*v1alpha1.AwsXraySamplingRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsXraySamplingRule))
	})
	return ret, err
}

// Get retrieves the AwsXraySamplingRule from the index for a given name.
func (s *awsXraySamplingRuleLister) Get(name string) (*v1alpha1.AwsXraySamplingRule, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsxraysamplingrule"), name)
	}
	return obj.(*v1alpha1.AwsXraySamplingRule), nil
}
