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

// AwsSecurityhubStandardsSubscriptionInformer provides access to a shared informer and lister for
// AwsSecurityhubStandardsSubscriptions.
type AwsSecurityhubStandardsSubscriptionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AwsSecurityhubStandardsSubscriptionLister
}

type awsSecurityhubStandardsSubscriptionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAwsSecurityhubStandardsSubscriptionInformer constructs a new informer for AwsSecurityhubStandardsSubscription type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAwsSecurityhubStandardsSubscriptionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAwsSecurityhubStandardsSubscriptionInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAwsSecurityhubStandardsSubscriptionInformer constructs a new informer for AwsSecurityhubStandardsSubscription type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAwsSecurityhubStandardsSubscriptionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().AwsSecurityhubStandardsSubscriptions().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().AwsSecurityhubStandardsSubscriptions().Watch(options)
			},
		},
		&awsv1alpha1.AwsSecurityhubStandardsSubscription{},
		resyncPeriod,
		indexers,
	)
}

func (f *awsSecurityhubStandardsSubscriptionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAwsSecurityhubStandardsSubscriptionInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *awsSecurityhubStandardsSubscriptionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&awsv1alpha1.AwsSecurityhubStandardsSubscription{}, f.defaultInformer)
}

func (f *awsSecurityhubStandardsSubscriptionInformer) Lister() v1alpha1.AwsSecurityhubStandardsSubscriptionLister {
	return v1alpha1.NewAwsSecurityhubStandardsSubscriptionLister(f.Informer().GetIndexer())
}
