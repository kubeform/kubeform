---
title: AWS
menu:
  docs_{{ .version }}:
    identifier: readme-aws
    name: Overview
    parent: aws-guides
    weight: 10
menu_name: docs_{{ .version }}
section_menu_id: guides
url: /docs/{{ .version }}/guides/aws/
aliases:
  - /docs/{{ .version }}/guides/aws/README/
---

# AWS

This guide will show you how to provision (Create, Update, Delete) an AWS S3 Bucket using Kubeform.

> Examples used in this guide can be found [here](https://github.com/kubeform/kubeform/tree/{{< param "info.version" >}}/docs/examples/aws).

At first, let's look at the `Terraform` configuration for an AWS S3 Bucket below:

```
provider "aws" {
    access_key = "ACCESS_KEY"

    region = "REGION_NAME"

    secret_key = "SECRET_KEY"
}

resource "aws_s3_bucket" "test1" {
    bucket = "s3-bucket-test1"
}
```

Now, if we apply `terraform apply` this config will create an AWS S3 Bucket. We'll create the exact configuration using `kubeform`. The steps are given below:

## 1. Create CRD:

At first, create the CRD of AWS S3 Bucket using the following kubectl command:

```console
$ kubectl apply -f https://raw.githubusercontent.com/kubeform/provider-aws-api/master/crds/s3.aws.kubeform.com_buckets.yaml
```

## 2. Create AWS Provider Secret

Then create the secret which is necessary for provisioning the S3 Bucket in AWS.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: aws-provider-secret
stringData:
  provider: |
    {
        "region": "<REGION_NAME>",
        "access_key": "<ACCESS_KEY>",
        "secret_key": "<SECRET_KEY>"
    }
```

Here we can see that, the `provider` field of the `stringData` of the secret is same as the field of the provider part in the terraform config file. The provider secret needs to be provided in json format, under the `provider` key. Save it in a file (eg. `provider-secret.yaml`) then apply it using kubectl.

```console
$ kubectl apply -f provider-secret.yaml
```

> **Note:** Here, key of the provider field of the stringData (eg. `"access_key"`, `"secret_key"` ) must be in snake case format (same as the tf configuration file)

## 3. Create S3 Bucket

Now, we'll create the S3 Bucket CRD. The yaml is given below:

```yaml
apiVersion: s3.aws.kubeform.com/v1alpha1
kind: Bucket
metadata:
  name: test1
spec:
  resource:
    bucket: s3-bucket-test1
  providerRef:
    name: aws-provider-secret
  terminationPolicy: DoNotTerminate
```

Here, the `resource` field contains the AWS S3 Bucket resource spec. Also, we can see that the provider secret is referenced using a field called `providerRef`.

> We can see a field named `terminationPolicy`, this is a feature of kubeform. This field can have two values, `Delete` or `DoNotTerminate`. When the value of this field is set to `DoNotTerminate` then the resource won't get deleted even though we apply `kubectl delete` operation, this field needs to be set to `Delete` to delete the resource. It helps to avoid accidental deletion of the resource. We will see the use of this field in `Delete S3 Bucket` part later on this guide. 

Save it in a file (eg. `aws-s3-bucket.yaml`) then apply it using kubectl.

```console
$ kubectl apply -f aws-s3-bucket.yaml
```

After applying this command, the resource will be in `InProgress` phase until the cloud creates the resource. Once the cloud resource get created, the resource will be in `Current` phase which means we have successfully created the resource.

## 4. Update S3 Bucket

Now, we'll update the S3 Bucket CRD. For updating the Bucket, we will modify the existing `aws-s3-bucket.yaml`, we will use a different `bucket` (`s3-bucket-test1-update`) field. The modified yaml is given below:

```yaml
apiVersion: s3.aws.kubeform.com/v1alpha1
kind: Bucket
metadata:
    name: test1
spec:
    resource:
        bucket: s3-bucket-test1-update
    providerRef:
        name: aws-provider-secret
    terminationPolicy: DoNotTerminate
```

Now, apply it using kubectl command.

```console
$ kubectl apply -f aws-s3-bucket.yaml
```

After that, existing AWS S3 Bucket will be updated!

> **Note:** Here, we have changed the `bucket` field which is Immutable, means if we change an immutable field then the resource will first get deleted and then created. But, there are some fields which are mutable, means changing those fields, resource will be only updated/changed. So, be careful!

## 5. Delete S3 Bucket

To delete the S3 Bucket just run:

```console
$ kubectl delete -f aws-s3-bucket.yaml
```

After applying this command we will get below error message, as we have set `terminationPolicy: DoNotTerminate`:

```text
Error from server (bucket "default/test1" can't be terminated. To delete, change spec.terminationPolicy to Delete): error when deleting "aws-s3-bucket.yaml": admission webhook "bucket.s3.aws.kubeform.com" denied the request: bucket "default/test1" can't be terminated. To delete, change spec.terminationPolicy to Delete
```

Let's change the `terminationPolicy` to `Delete` by using kubectl patch command.

```console
$ kubectl patch -n default bucket test1 -p '{"spec":{"terminationPolicy":"Delete"}}' --type="merge"
```

Now, we can delete the Bucket.

```console
$ kubectl delete -f aws-s3-bucket.yaml
```

After applying this command the resource will be in `Terminating` phase until the cloud resource get destroyed. Once the cloud resource get destroyed, the resource will get deleted successfully. 



