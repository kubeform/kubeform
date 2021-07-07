---
title: Linode
menu:
  docs_{{ .version }}:
    identifier: readme-linode
    name: Overview
    parent: linode-guides
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: guides
url: /docs/{{ .version }}/guides/linode/
aliases:
  - /docs/{{ .version }}/guides/linode/README/
---

# Linode

This guide will show you how to provision (Create, Update, Delete) a Linode Instance using Kubeform.

> Examples used in this guide can be found [here](https://github.com/kubeform/kubeform/tree/{{< param "info.version" >}}/docs/examples/linode).

At first, let's look at the `Terraform` configuration for a Linode Instance below:

```
provider "linode" {
    token = "LINODE_API_TOKEN"
}

resource "linode_instance" "test1" {
    image = "linode/ubuntu18.04"

    label = "instance-test1"

    region = "us-east"

    root_pass = "TestPassWord@123"
}
```

Now, if we apply `terraform apply` this config will create a Linode Instance. We'll create the exact configuration using Kubeform. The steps are given below:

## 1. Create CRD:

At first, create the CRD of Linode Instance using the following kubectl command:

```console
$ kubectl apply -f https://raw.githubusercontent.com/kubeform/provider-linode-api/master/crds/instance.linode.kubeform.com_instances.yaml
```

## 2. Create Linode Provider Secret

Then create the secret which is necessary for provisioning the Instance in Linode.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: linode-provider-secret
stringData:
  provider: |
    {
        "token": "LINODE_API_TOKEN"
    }

```

Here we can see that, the `provider` field of the `stringData` of the secret is same as the field of the provider part in the terraform config file. Save it in a file (eg. `provider-secret.yaml`) then apply it using kubectl command.

```console
$ kubectl apply -f provider-secret.yaml
```

> **Note:** Here, key of the provider field of the stringData (eg. `"token"`) must be in snake case format (same as the tf configuration file)

## 3. Create Secrets for Sensitive Data

If we look at the terraform config, we can see that there is a field called `root_pass`. This is a sensitive field. So, we should not use these kinds of sensitive values directly in the yaml. We'll create a secret to store the sensitive values like this:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: linode-sensitive-secret
stringData:
  input: |
    {
        "root_pass": "TestPassWord@123"
    }
```

we'll reference this secret from the Instance CRD. Save it in a file (eg. `sensitive-secret.yaml`) then apply it using kubectl command.

```console
$ kubectl apply -f sensitive-secret.yaml
```

> **Note:** Here, key of input field of the stringData (eg. `"root_pass"`) must be in snake case format (same as the tf configuration file)

## 4. Create Instance

Now, we'll create the Instance CRD. The yaml is given below:

```yaml
apiVersion: instance.linode.kubeform.com/v1alpha1
kind: Instance
metadata:
  name: test1
spec:
  resource:
    region: us-east
    image: linode/ubuntu18.04
    label: instance-test1
  providerRef:
    name: linode-provider-secret
  secretRef:
    name: linode-sensitive-secret
  terminationPolicy: DoNotTerminate
```

Here, the `resource` field contains the Linode Instance spec. Also, we can see that the provider secret is referenced using a field called `providerRef` and the sensitive value secret is referenced using a field called `secretRef`. 

> We can see a field named `terminationPolicy`, this is a feature of kubeform. This field can have two types of values, `Delete` or `DoNotTerminate`. When the value of this field is set to `DoNotTerminate` then the Instance won't get deleted even though we apply `kubectl delete` operation, this field needs to be set to `Delete` to delete the Instance. It helps to avoid accidental deletion of the resource. We will see the use of this field in `Delete Instance` part later on this page. 

Save it in a file (eg. `linode-instance.yaml`) then apply it using kubectl.

```console
$ kubectl apply -f linode-instance.yaml
```

After that, a Linode Instance will be created!

## 5. Update Instance

Now, we'll update the Instance CRD. For updating the Instance, we will modify the existing `linode-instnace.yaml`, we will use a different `label` (`instance-test1-update`). The modified yaml is given below:

```yaml
apiVersion: instance.linode.kubeform.com/v1alpha1
kind: Instance
metadata:
  name: test1
spec:
  resource:
    region: us-east
    image: linode/ubuntu18.04
    label: instance-test1-update
  providerRef:
    name: linode-provider-secret
  secretRef:
    name: linode-sensitive-secret
  terminationPolicy: DoNotTerminate
```

Now, apply it using kubectl command.

```console
$ kubectl apply -f linode-instance.yaml
```

After that, existing Linode Instance will be updated!

> **Note:** Here, we have changed the `label` field which is mutable, means if we change a mutable field it will get updated/changed only. But, there are some fields which are immutable, means changing those fields, resource will first get deleted and then again recreated. So, be careful!


## 6. Delete Instance

To delete the Instance just run:

```console
$ kubectl delete -f linode-instance.yaml
```

After applying this command we will get below error message, as we have set `terminationPolicy: DoNotTerminate`:

```text
Error from server (instance "default/test1" can't be terminated. To delete, change spec.terminationPolicy to Delete): error when deleting "linode-instance.yaml": admission webhook "instance.instance.linode.kubeform.com" denied the request: instance "default/test1" can't be terminated. To delete, change spec.terminationPolicy to Delete
```

Let's change the `terminationPolicy` to `Delete` by using kubectl patch command.

```console
$ kubectl patch -n default instance test1 -p '{"spec":{"terminationPolicy":"Delete"}}' --type="merge"
```

Now, we can delete the Instance.

```console
$ kubectl delete -f linode-instance.yaml
```

After applying this command we can see that the Instance has successfully got deleted!
