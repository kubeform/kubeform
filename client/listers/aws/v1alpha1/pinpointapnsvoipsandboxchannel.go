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

// PinpointApnsVoipSandboxChannelLister helps list PinpointApnsVoipSandboxChannels.
type PinpointApnsVoipSandboxChannelLister interface {
	// List lists all PinpointApnsVoipSandboxChannels in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PinpointApnsVoipSandboxChannel, err error)
	// PinpointApnsVoipSandboxChannels returns an object that can list and get PinpointApnsVoipSandboxChannels.
	PinpointApnsVoipSandboxChannels(namespace string) PinpointApnsVoipSandboxChannelNamespaceLister
	PinpointApnsVoipSandboxChannelListerExpansion
}

// pinpointApnsVoipSandboxChannelLister implements the PinpointApnsVoipSandboxChannelLister interface.
type pinpointApnsVoipSandboxChannelLister struct {
	indexer cache.Indexer
}

// NewPinpointApnsVoipSandboxChannelLister returns a new PinpointApnsVoipSandboxChannelLister.
func NewPinpointApnsVoipSandboxChannelLister(indexer cache.Indexer) PinpointApnsVoipSandboxChannelLister {
	return &pinpointApnsVoipSandboxChannelLister{indexer: indexer}
}

// List lists all PinpointApnsVoipSandboxChannels in the indexer.
func (s *pinpointApnsVoipSandboxChannelLister) List(selector labels.Selector) (ret []*v1alpha1.PinpointApnsVoipSandboxChannel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PinpointApnsVoipSandboxChannel))
	})
	return ret, err
}

// PinpointApnsVoipSandboxChannels returns an object that can list and get PinpointApnsVoipSandboxChannels.
func (s *pinpointApnsVoipSandboxChannelLister) PinpointApnsVoipSandboxChannels(namespace string) PinpointApnsVoipSandboxChannelNamespaceLister {
	return pinpointApnsVoipSandboxChannelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PinpointApnsVoipSandboxChannelNamespaceLister helps list and get PinpointApnsVoipSandboxChannels.
type PinpointApnsVoipSandboxChannelNamespaceLister interface {
	// List lists all PinpointApnsVoipSandboxChannels in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PinpointApnsVoipSandboxChannel, err error)
	// Get retrieves the PinpointApnsVoipSandboxChannel from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PinpointApnsVoipSandboxChannel, error)
	PinpointApnsVoipSandboxChannelNamespaceListerExpansion
}

// pinpointApnsVoipSandboxChannelNamespaceLister implements the PinpointApnsVoipSandboxChannelNamespaceLister
// interface.
type pinpointApnsVoipSandboxChannelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PinpointApnsVoipSandboxChannels in the indexer for a given namespace.
func (s pinpointApnsVoipSandboxChannelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PinpointApnsVoipSandboxChannel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PinpointApnsVoipSandboxChannel))
	})
	return ret, err
}

// Get retrieves the PinpointApnsVoipSandboxChannel from the indexer for a given namespace and name.
func (s pinpointApnsVoipSandboxChannelNamespaceLister) Get(name string) (*v1alpha1.PinpointApnsVoipSandboxChannel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pinpointapnsvoipsandboxchannel"), name)
	}
	return obj.(*v1alpha1.PinpointApnsVoipSandboxChannel), nil
}