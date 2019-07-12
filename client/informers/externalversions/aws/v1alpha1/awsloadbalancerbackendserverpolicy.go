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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	awsv1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	versioned "kubeform.dev/kubeform/client/clientset/versioned"
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/kubeform/client/listers/aws/v1alpha1"
)

// AwsLoadBalancerBackendServerPolicyInformer provides access to a shared informer and lister for
// AwsLoadBalancerBackendServerPolicies.
type AwsLoadBalancerBackendServerPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AwsLoadBalancerBackendServerPolicyLister
}

type awsLoadBalancerBackendServerPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAwsLoadBalancerBackendServerPolicyInformer constructs a new informer for AwsLoadBalancerBackendServerPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAwsLoadBalancerBackendServerPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAwsLoadBalancerBackendServerPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAwsLoadBalancerBackendServerPolicyInformer constructs a new informer for AwsLoadBalancerBackendServerPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAwsLoadBalancerBackendServerPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().AwsLoadBalancerBackendServerPolicies().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().AwsLoadBalancerBackendServerPolicies().Watch(options)
			},
		},
		&awsv1alpha1.AwsLoadBalancerBackendServerPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *awsLoadBalancerBackendServerPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAwsLoadBalancerBackendServerPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *awsLoadBalancerBackendServerPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&awsv1alpha1.AwsLoadBalancerBackendServerPolicy{}, f.defaultInformer)
}

func (f *awsLoadBalancerBackendServerPolicyInformer) Lister() v1alpha1.AwsLoadBalancerBackendServerPolicyLister {
	return v1alpha1.NewAwsLoadBalancerBackendServerPolicyLister(f.Informer().GetIndexer())
}
