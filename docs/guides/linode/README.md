# Linode

This guide will show you how to provision a Linode Instance using Kubeform.

> Examples used in this guide can be found [here](/docs/examples/linode).

Look at the `Terraform` configuration below:

```
provider "linode" {
    token = "LINODE_API_TOKEN"
}

resource "linode_instance" "instance-test1" {
    image = "linode/ubuntu18.04"

    label = "instance-test1"

    region = "us-east"

    root_pass = "password"

    stackscript_data = {
      user = "John Doe"
    }

    stackscript_id = 553057
}
```

This config creates a Linode Instance. We'll create the exact configuration using kubeform. The steps are given below:

## 1. Create CRD:

At first, create the CRD of Linode Instance using the following yaml:

```yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: instances.linode.kubeform.com
spec:
  group: linode.kubeform.com
  version: v1alpha1
  names:
    kind: Instance
    plural: instances
  scope: Namespaced
``` 

Save it in a file (eg. `crd.yaml`) then apply it using kubectl.

```bash
$ kubectl apply -f crd.yaml
```

## 2. Create Linode Provider Secret

Then create the secret which is necessary for provisioning the Instance in Linode.

```yaml 
apiVersion: v1
kind: Secret
metadata:
  name: linode
type: "kubeform.com/linode"
data:
  token: '<base64 encoded secret key>'
```

Here we can see that, the data of the secret is same as the field of the provider part in the terraform config file. Save it in a file (eg. `secret.yaml`) then apply it using kubectl.

```bash
$ kubectl apply -f secret.yaml
```
> **Note:** here, data key (eg. `token`) must be in snake case format (same as the tf configuration file)

## 3. Create Secrets for Sensitive Data

If we look at the terraform config, We can see that there is a field called stackscript_data and root_pass. These are sensitive fields. So, we should not use these kind of sensitive values directly in the yaml. We'll create a secret to store the sensitive values like the this:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: data
type: "kubeform.com/linode"
data:
  root_pass: cGFzc3dvcmQK  # base64 encoded value of password
  stackscript_data: ewoidXNlciIgOiAiSm9obiBEb2UiCn0K  # base64 encoded value of "{user = "John Doe"}"
```

we'll reference this secret from the Instance CRD. Save it in a file (eg. `provider.yaml`) then apply it using kubectl.

```bash
$ kubectl apply -f provider.yaml
```
> **Note:** here, data key (eg. `root_pass`, `stackscript_data` etc.) must be in snake case format (same as the tf configuration file)

## 4. Create Instance

Now, we'll create the Instance CRD. The yaml is given below:

```yaml
apiVersion: linode.kubeform.com/v1alpha1
kind: Instance
metadata:
  name: instance-test1
spec:
  region: us-east
  image: linode/ubuntu18.04
  label: instance-test1
  providerRef:
    name: linode
  secretRef:
    name: data
  stackscriptID: 553057
```

Here, we can see that the provider secret is referenced using a field called `providerRef` and the sensitive value secret is referenced using a field called `secretRef`.

Save it in a file (eg. `instance.yaml`) then apply it using kubectl.

```bash
$ kubectl apply -f instance.yaml
```

After that, an Linode Instance will be created!

## Delete Instance

To delete the Instance just run:

```bash
kubectl delete -f instance.yaml
```
