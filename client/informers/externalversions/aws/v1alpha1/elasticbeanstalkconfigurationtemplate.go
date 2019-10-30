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

// ElasticBeanstalkConfigurationTemplateInformer provides access to a shared informer and lister for
// ElasticBeanstalkConfigurationTemplates.
type ElasticBeanstalkConfigurationTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ElasticBeanstalkConfigurationTemplateLister
}

type elasticBeanstalkConfigurationTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewElasticBeanstalkConfigurationTemplateInformer constructs a new informer for ElasticBeanstalkConfigurationTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewElasticBeanstalkConfigurationTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredElasticBeanstalkConfigurationTemplateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredElasticBeanstalkConfigurationTemplateInformer constructs a new informer for ElasticBeanstalkConfigurationTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredElasticBeanstalkConfigurationTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().ElasticBeanstalkConfigurationTemplates(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().ElasticBeanstalkConfigurationTemplates(namespace).Watch(options)
			},
		},
		&awsv1alpha1.ElasticBeanstalkConfigurationTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *elasticBeanstalkConfigurationTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredElasticBeanstalkConfigurationTemplateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *elasticBeanstalkConfigurationTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&awsv1alpha1.ElasticBeanstalkConfigurationTemplate{}, f.defaultInformer)
}

func (f *elasticBeanstalkConfigurationTemplateInformer) Lister() v1alpha1.ElasticBeanstalkConfigurationTemplateLister {
	return v1alpha1.NewElasticBeanstalkConfigurationTemplateLister(f.Informer().GetIndexer())
}
