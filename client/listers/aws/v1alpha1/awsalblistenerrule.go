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

// AwsAlbListenerRuleLister helps list AwsAlbListenerRules.
type AwsAlbListenerRuleLister interface {
	// List lists all AwsAlbListenerRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsAlbListenerRule, err error)
	// Get retrieves the AwsAlbListenerRule from the index for a given name.
	Get(name string) (*v1alpha1.AwsAlbListenerRule, error)
	AwsAlbListenerRuleListerExpansion
}

// awsAlbListenerRuleLister implements the AwsAlbListenerRuleLister interface.
type awsAlbListenerRuleLister struct {
	indexer cache.Indexer
}

// NewAwsAlbListenerRuleLister returns a new AwsAlbListenerRuleLister.
func NewAwsAlbListenerRuleLister(indexer cache.Indexer) AwsAlbListenerRuleLister {
	return &awsAlbListenerRuleLister{indexer: indexer}
}

// List lists all AwsAlbListenerRules in the indexer.
func (s *awsAlbListenerRuleLister) List(selector labels.Selector) (ret []*v1alpha1.AwsAlbListenerRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsAlbListenerRule))
	})
	return ret, err
}

// Get retrieves the AwsAlbListenerRule from the index for a given name.
func (s *awsAlbListenerRuleLister) Get(name string) (*v1alpha1.AwsAlbListenerRule, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsalblistenerrule"), name)
	}
	return obj.(*v1alpha1.AwsAlbListenerRule), nil
}
