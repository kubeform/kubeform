---
title: DigitalOcean
menu:
  docs_{{ .version }}:
    identifier: readme-digitalocean
    name: Overview
    parent: digitalocean-guides
    weight: 10
menu_name: docs_{{ .version }}
section_menu_id: guides
url: /docs/{{ .version }}/guides/digitalocean/
aliases:
  - /docs/{{ .version }}/guides/digitalocean/README/
---

# DigitalOcean

This guide will show you how to provision a Digitalocean Database Cluster using Kubeform.

> Examples used in this guide can be found [here](https://github.com/kubeform/docs/tree/{{< param "info.version" >}}/docs/examples/digitalocean).

Look at the `Terraform` configuration below:

```
provider "digitalocean" {
    token = "DigitalOcean Token"
}

resource "digitalocean_database_cluster" "test1" {
    engine = "pg"

    name = "example-cluster"

    node_count" = 1

    region = "nyc1"

    size = "db-s-1vcpu-1gb"

    version = "11"
}⏎
```

This config creates a Digitalocean Database Cluster. We'll create the exact configuration using kubeform. The steps are given below:

## 1. Create CRD:

At first, create the CRD of Digitalocean Database Cluster using the following kubectl command:

```bash
$ kubectl apply -f https://github.com/kubeform/kubeform/raw/master/api/crds/digitalocean.kubeform.com_databaseclusters.yaml
```

## 2. Create DigitalOcean Provider Secret

Then create the secret which is necessary for provisioning the Database Cluster in Digitalocean.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: do
type: "kubeform.com/digitalocean"
data:
  token: '<base64 encoded secret key>'
```

Here we can see that, the data of the secret is same as the field of the provider part in the terraform config file. Save it in a file (eg. `secret.yaml`) then apply it using kubectl.

```bash
$ kubectl apply -f secret.yaml
```
> **Note:** here, data key (eg. `token`) must be in snake case format (same as the tf configuration file)

## 3. Create Digitalocean Database Cluster

Now, we'll create the Digitalocean Database Cluster CRD. The yaml is given below:

```yaml
apiVersion: digitalocean.kubeform.com/v1alpha1
kind: DatabaseCluster
metadata:
  name: test1
spec:
  name: example-cluster
  engine: pg
  version: "11"
  size: db-s-1vcpu-1gb
  region: nyc1
  nodeCount: 1
  providerRef:
    name: do
```

Here, we can see that the provider secret is referenced using a field called `providerRef`.

Save it in a file (eg. `db_cluster.yaml`) then apply it using kubectl.

```bash
$ kubectl apply -f db_cluster.yaml
```

After that, an Digitalocean Database Cluster will be created!

## Delete Digitalocean Database Cluster

To delete the Digitalocean Database Cluster just run:

```bash
kubectl delete -f db_cluster.yaml
```
