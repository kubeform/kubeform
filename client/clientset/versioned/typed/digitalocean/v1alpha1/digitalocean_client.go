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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
	"kubeform.dev/kubeform/client/clientset/versioned/scheme"

	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type DigitaloceanV1alpha1Interface interface {
	RESTClient() rest.Interface
	CdnsGetter
	CertificatesGetter
	DatabaseClustersGetter
	DomainsGetter
	DropletsGetter
	DropletSnapshotsGetter
	FirewallsGetter
	FloatingIPsGetter
	FloatingIPAssignmentsGetter
	KubernetesClustersGetter
	KubernetesNodePoolsGetter
	LoadbalancersGetter
	ProjectsGetter
	RecordsGetter
	SpacesBucketsGetter
	SshKeysGetter
	TagsGetter
	VolumesGetter
	VolumeAttachmentsGetter
	VolumeSnapshotsGetter
}

// DigitaloceanV1alpha1Client is used to interact with features provided by the digitalocean.kubeform.com group.
type DigitaloceanV1alpha1Client struct {
	restClient rest.Interface
}

func (c *DigitaloceanV1alpha1Client) Cdns(namespace string) CdnInterface {
	return newCdns(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) Certificates(namespace string) CertificateInterface {
	return newCertificates(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) DatabaseClusters(namespace string) DatabaseClusterInterface {
	return newDatabaseClusters(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) Domains(namespace string) DomainInterface {
	return newDomains(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) Droplets(namespace string) DropletInterface {
	return newDroplets(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) DropletSnapshots(namespace string) DropletSnapshotInterface {
	return newDropletSnapshots(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) Firewalls(namespace string) FirewallInterface {
	return newFirewalls(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) FloatingIPs(namespace string) FloatingIPInterface {
	return newFloatingIPs(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) FloatingIPAssignments(namespace string) FloatingIPAssignmentInterface {
	return newFloatingIPAssignments(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) KubernetesClusters(namespace string) KubernetesClusterInterface {
	return newKubernetesClusters(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) KubernetesNodePools(namespace string) KubernetesNodePoolInterface {
	return newKubernetesNodePools(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) Loadbalancers(namespace string) LoadbalancerInterface {
	return newLoadbalancers(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) Projects(namespace string) ProjectInterface {
	return newProjects(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) Records(namespace string) RecordInterface {
	return newRecords(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) SpacesBuckets(namespace string) SpacesBucketInterface {
	return newSpacesBuckets(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) SshKeys(namespace string) SshKeyInterface {
	return newSshKeys(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) Tags(namespace string) TagInterface {
	return newTags(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) Volumes(namespace string) VolumeInterface {
	return newVolumes(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) VolumeAttachments(namespace string) VolumeAttachmentInterface {
	return newVolumeAttachments(c, namespace)
}

func (c *DigitaloceanV1alpha1Client) VolumeSnapshots(namespace string) VolumeSnapshotInterface {
	return newVolumeSnapshots(c, namespace)
}

// NewForConfig creates a new DigitaloceanV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*DigitaloceanV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &DigitaloceanV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new DigitaloceanV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *DigitaloceanV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new DigitaloceanV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *DigitaloceanV1alpha1Client {
	return &DigitaloceanV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *DigitaloceanV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
