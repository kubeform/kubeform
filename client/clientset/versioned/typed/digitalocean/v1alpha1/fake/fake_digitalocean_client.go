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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/digitalocean/v1alpha1"
)

type FakeDigitaloceanV1alpha1 struct {
	*testing.Fake
}

func (c *FakeDigitaloceanV1alpha1) Cdns(namespace string) v1alpha1.CdnInterface {
	return &FakeCdns{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) Certificates(namespace string) v1alpha1.CertificateInterface {
	return &FakeCertificates{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) DatabaseClusters(namespace string) v1alpha1.DatabaseClusterInterface {
	return &FakeDatabaseClusters{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) Domains(namespace string) v1alpha1.DomainInterface {
	return &FakeDomains{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) Droplets(namespace string) v1alpha1.DropletInterface {
	return &FakeDroplets{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) DropletSnapshots(namespace string) v1alpha1.DropletSnapshotInterface {
	return &FakeDropletSnapshots{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) Firewalls(namespace string) v1alpha1.FirewallInterface {
	return &FakeFirewalls{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) FloatingIPs(namespace string) v1alpha1.FloatingIPInterface {
	return &FakeFloatingIPs{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) FloatingIPAssignments(namespace string) v1alpha1.FloatingIPAssignmentInterface {
	return &FakeFloatingIPAssignments{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) KubernetesClusters(namespace string) v1alpha1.KubernetesClusterInterface {
	return &FakeKubernetesClusters{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) KubernetesNodePools(namespace string) v1alpha1.KubernetesNodePoolInterface {
	return &FakeKubernetesNodePools{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) Loadbalancers(namespace string) v1alpha1.LoadbalancerInterface {
	return &FakeLoadbalancers{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) Projects(namespace string) v1alpha1.ProjectInterface {
	return &FakeProjects{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) Records(namespace string) v1alpha1.RecordInterface {
	return &FakeRecords{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) SpacesBuckets(namespace string) v1alpha1.SpacesBucketInterface {
	return &FakeSpacesBuckets{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) SshKeys(namespace string) v1alpha1.SshKeyInterface {
	return &FakeSshKeys{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) Tags(namespace string) v1alpha1.TagInterface {
	return &FakeTags{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) Volumes(namespace string) v1alpha1.VolumeInterface {
	return &FakeVolumes{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) VolumeAttachments(namespace string) v1alpha1.VolumeAttachmentInterface {
	return &FakeVolumeAttachments{c, namespace}
}

func (c *FakeDigitaloceanV1alpha1) VolumeSnapshots(namespace string) v1alpha1.VolumeSnapshotInterface {
	return &FakeVolumeSnapshots{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDigitaloceanV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}