/*
Copyright The Kubeform Authors.

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
	"context"
	time "time"

	awsv1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	versioned "kubeform.dev/kubeform/client/clientset/versioned"
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/kubeform/client/listers/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KinesisFirehoseDeliveryStreamInformer provides access to a shared informer and lister for
// KinesisFirehoseDeliveryStreams.
type KinesisFirehoseDeliveryStreamInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KinesisFirehoseDeliveryStreamLister
}

type kinesisFirehoseDeliveryStreamInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKinesisFirehoseDeliveryStreamInformer constructs a new informer for KinesisFirehoseDeliveryStream type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKinesisFirehoseDeliveryStreamInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKinesisFirehoseDeliveryStreamInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKinesisFirehoseDeliveryStreamInformer constructs a new informer for KinesisFirehoseDeliveryStream type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKinesisFirehoseDeliveryStreamInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().KinesisFirehoseDeliveryStreams(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().KinesisFirehoseDeliveryStreams(namespace).Watch(context.TODO(), options)
			},
		},
		&awsv1alpha1.KinesisFirehoseDeliveryStream{},
		resyncPeriod,
		indexers,
	)
}

func (f *kinesisFirehoseDeliveryStreamInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKinesisFirehoseDeliveryStreamInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kinesisFirehoseDeliveryStreamInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&awsv1alpha1.KinesisFirehoseDeliveryStream{}, f.defaultInformer)
}

func (f *kinesisFirehoseDeliveryStreamInformer) Lister() v1alpha1.KinesisFirehoseDeliveryStreamLister {
	return v1alpha1.NewKinesisFirehoseDeliveryStreamLister(f.Informer().GetIndexer())
}
