---
title: Google
menu:
  docs_v0.0.1:
    identifier: readme-google
    name: Google
    parent: google-guides
    weight: 10
menu_name: docs_v0.0.1
section_menu_id: guides
url: /docs/v0.0.1/guides/google/
aliases:
  - /docs/v0.0.1/guides/google/README/
---

# Google

This guide will show you how to provision a Google Storage Bucket using Kubeform.

> Examples used in this guide can be found [here](/docs/examples/google).

Look at the `Terraform` configuration below:

```
provider "google" {
    credentials = "json Credential"

    project = "Project Name"

    region = "us-central1"
}⏎

resource "google_storage_bucket" "test1" {
    location = "EU"

    name = "bucket_test1"

    website = {
      main_page_suffix = "index.html"
      not_found_page = "404.html"
    }
}
```

This config creates a Google Storage Bucket. We'll create the exact configuration using kubeform. The steps are given below:

## 1. Create CRD:

At first, create the CRD of Google Storage Bucket using the following yaml:

```yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: storagebuckets.google.kubeform.com
spec:
  group: google.kubeform.com
  version: v1alpha1
  names:
    kind: StorageBucket
    plural: storagebuckets
  scope: Namespaced
```

Save it in a file (eg. `crd.yaml`) then apply it using kubectl.

```bash
$ kubectl apply -f crd.yaml
```

## 2. Create Google Provider Secret

Then create the secret which is necessary for provisioning the Storage Bucket in Google.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: google
type: "kubeform.com/google"
data:
  credentials: <Base64 encoded value of google json service account>
  project: <Base64 encoded value of project name>
  region: dXMtY2VudHJhbDEK # Base64 encoded value of us-central1
```

Here we can see that, the data of the secret is same as the field of the provider part in the terraform config file. Save it in a file (eg. `secret.yaml`) then apply it using kubectl.

```bash
$ kubectl apply -f secret.yaml
```
> **Note:** here, data key (eg. `project`, `region` etc.) must be in snake case format (same as the tf configuration file)

## 3. Create Google Storage Bucket

Now, we'll create the Google Storage Bucket CRD. The yaml is given below:

```yaml
apiVersion: google.kubeform.com/v1alpha1
kind: StorageBucket
metadata:
  name: test1
spec:
  name: bucket_test1
  location: EU
  website:
    - mainPageSuffix: index.html
      notFoundPage: 404.html
  providerRef:
    name: google
```

Here, we can see that the provider secret is referenced using a field called `providerRef`.

Save it in a file (eg. `bucket.yaml`) then apply it using kubectl.

```bash
$ kubectl apply -f bucket.yaml
```

After that, an Google Storage Bucket will be created!

## Delete Google Storage Bucket

To delete the Google Storage Bucket just run:

```bash
kubectl delete -f bucket.yaml
```
