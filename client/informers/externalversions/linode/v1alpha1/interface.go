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
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// LinodeDomains returns a LinodeDomainInformer.
	LinodeDomains() LinodeDomainInformer
	// LinodeDomainRecords returns a LinodeDomainRecordInformer.
	LinodeDomainRecords() LinodeDomainRecordInformer
	// LinodeImages returns a LinodeImageInformer.
	LinodeImages() LinodeImageInformer
	// LinodeInstances returns a LinodeInstanceInformer.
	LinodeInstances() LinodeInstanceInformer
	// LinodeNodebalancers returns a LinodeNodebalancerInformer.
	LinodeNodebalancers() LinodeNodebalancerInformer
	// LinodeNodebalancerConfigs returns a LinodeNodebalancerConfigInformer.
	LinodeNodebalancerConfigs() LinodeNodebalancerConfigInformer
	// LinodeNodebalancerNodes returns a LinodeNodebalancerNodeInformer.
	LinodeNodebalancerNodes() LinodeNodebalancerNodeInformer
	// LinodeRdnses returns a LinodeRdnsInformer.
	LinodeRdnses() LinodeRdnsInformer
	// LinodeSshkeys returns a LinodeSshkeyInformer.
	LinodeSshkeys() LinodeSshkeyInformer
	// LinodeStackscripts returns a LinodeStackscriptInformer.
	LinodeStackscripts() LinodeStackscriptInformer
	// LinodeTokens returns a LinodeTokenInformer.
	LinodeTokens() LinodeTokenInformer
	// LinodeVolumes returns a LinodeVolumeInformer.
	LinodeVolumes() LinodeVolumeInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// LinodeDomains returns a LinodeDomainInformer.
func (v *version) LinodeDomains() LinodeDomainInformer {
	return &linodeDomainInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LinodeDomainRecords returns a LinodeDomainRecordInformer.
func (v *version) LinodeDomainRecords() LinodeDomainRecordInformer {
	return &linodeDomainRecordInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LinodeImages returns a LinodeImageInformer.
func (v *version) LinodeImages() LinodeImageInformer {
	return &linodeImageInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LinodeInstances returns a LinodeInstanceInformer.
func (v *version) LinodeInstances() LinodeInstanceInformer {
	return &linodeInstanceInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LinodeNodebalancers returns a LinodeNodebalancerInformer.
func (v *version) LinodeNodebalancers() LinodeNodebalancerInformer {
	return &linodeNodebalancerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LinodeNodebalancerConfigs returns a LinodeNodebalancerConfigInformer.
func (v *version) LinodeNodebalancerConfigs() LinodeNodebalancerConfigInformer {
	return &linodeNodebalancerConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LinodeNodebalancerNodes returns a LinodeNodebalancerNodeInformer.
func (v *version) LinodeNodebalancerNodes() LinodeNodebalancerNodeInformer {
	return &linodeNodebalancerNodeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LinodeRdnses returns a LinodeRdnsInformer.
func (v *version) LinodeRdnses() LinodeRdnsInformer {
	return &linodeRdnsInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LinodeSshkeys returns a LinodeSshkeyInformer.
func (v *version) LinodeSshkeys() LinodeSshkeyInformer {
	return &linodeSshkeyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LinodeStackscripts returns a LinodeStackscriptInformer.
func (v *version) LinodeStackscripts() LinodeStackscriptInformer {
	return &linodeStackscriptInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LinodeTokens returns a LinodeTokenInformer.
func (v *version) LinodeTokens() LinodeTokenInformer {
	return &linodeTokenInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LinodeVolumes returns a LinodeVolumeInformer.
func (v *version) LinodeVolumes() LinodeVolumeInformer {
	return &linodeVolumeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
