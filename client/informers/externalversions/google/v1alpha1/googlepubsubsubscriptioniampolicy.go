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
	googlev1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	versioned "kubeform.dev/kubeform/client/clientset/versioned"
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/kubeform/client/listers/google/v1alpha1"
)

// GooglePubsubSubscriptionIamPolicyInformer provides access to a shared informer and lister for
// GooglePubsubSubscriptionIamPolicies.
type GooglePubsubSubscriptionIamPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.GooglePubsubSubscriptionIamPolicyLister
}

type googlePubsubSubscriptionIamPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewGooglePubsubSubscriptionIamPolicyInformer constructs a new informer for GooglePubsubSubscriptionIamPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGooglePubsubSubscriptionIamPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGooglePubsubSubscriptionIamPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredGooglePubsubSubscriptionIamPolicyInformer constructs a new informer for GooglePubsubSubscriptionIamPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGooglePubsubSubscriptionIamPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GoogleV1alpha1().GooglePubsubSubscriptionIamPolicies().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GoogleV1alpha1().GooglePubsubSubscriptionIamPolicies().Watch(options)
			},
		},
		&googlev1alpha1.GooglePubsubSubscriptionIamPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *googlePubsubSubscriptionIamPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGooglePubsubSubscriptionIamPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *googlePubsubSubscriptionIamPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&googlev1alpha1.GooglePubsubSubscriptionIamPolicy{}, f.defaultInformer)
}

func (f *googlePubsubSubscriptionIamPolicyInformer) Lister() v1alpha1.GooglePubsubSubscriptionIamPolicyLister {
	return v1alpha1.NewGooglePubsubSubscriptionIamPolicyLister(f.Informer().GetIndexer())
}
