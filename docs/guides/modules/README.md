---
title: Modules
menu:
  docs_{{ .version }}:
    identifier: readme-modules
    name: Overview
    parent: modules-guides
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: guides
url: /docs/{{ .version }}/guides/modules/
aliases:
  - /docs/{{ .version }}/guides/modules/README/
---

# Modules

This guide will show you how to provision an AWS RDS module using Kubeform.

> Examples used in this guide can be found [here](https://github.com/kubeform/docs/tree/{{< param "info.version" >}}/docs/examples/modules).

Look at the `Terraform` configuration below:

```
{
  "module": {
    "test-rds": {
      "source": "terraform-aws-modules/rds/aws",
      "allocated_storage": "5",
      "backup_window": "03:00-04:00",
      "engine": "mysql",
      "engine_version": "5.7.19",
      "family": "mysql5.7",
      "identifier": "demodb",
      "instance_class": "db.t2.large",
      "maintenance_window": "Mon:00:00-Mon:03:00",
      "major_engine_version": "5.7",
      "name": "demodb",
      "options": [
        {
          "option_name": "MARIADB_AUDIT_PLUGIN",
          "option_settings": [
            {
              "name": "SERVER_AUDIT_EVENTS",
              "value": "CONNECT"
            },
            {
              "name": "SERVER_AUDIT_FILE_ROTATIONS",
              "value": "37"
            }
          ]
        }
      ],
      "parameters": [
        {
          "name": "character_set_client",
          "value": "utf8"
        },
        {
          "name": "character_set_server",
          "value": "utf8"
        }
      ],
      "password": "foobar1234",
      "port": "3306",
      "skip_final_snapshot": true,
      "subnet_ids": [
        "<subnet id>"
      ],
      "tags": {
        "Environment": "dev",
        "Owner": "user"
      },
      "username": "user",
      "vpc_security_group_ids": [
        "<vpc id>"
      ]
    }
  }
}
```

This config creates an AWS RDS instance. We'll create the exact configuration using Kubeform. The steps are given below:

## 1. Create CRD:

At first, create the CRD of AWS RDS module using the following kubectl command:

```console
$ kubectl apply -f crd.yaml	$ kubectl apply -f https://github.com/kubeform/kubeform/raw/master/api/crds/modules.kubeform.com_rds.yaml
```

## 2. Create AWS Provider Secret

Then create the secret which is necessary for provisioning the RDS instance in AWS.

```yaml
apiVersion: v1
kind: Secret
metadata:
    name: aws
type: kubeform.com/aws
data:
    region: dXMtZWFzdC0xCg==  # base64 encoded value of `us-east-1`
    access_key: '<base64 encoded access key>'
    secret_key: '<base64 encoded secret key>'
```

Here we can see that, the data of the secret is same as the field of the provider part in the terraform config file. Save it in a file (eg. `secret.yaml`) then apply it using kubectl.

```console
$ kubectl apply -f secret.yaml
```

> **Note:** here, data key (eg. `access_key`, `access_key` ) must be in snake case format (same as the `tf` configuration file)


## 3. Create Secrets for Sensitive Data

If we look at the terraform config, We can see that there is a field called password. This is a sensitive field. So, we should not use this kind of sensitive value directly in the yaml. We'll create a secret to store the sensitive value like this:

```yaml
apiVersion: v1
kind: Secret
metadata:
    name: rds-pass
type: kubeform.com/aws
data:
    password: Zm9vYmFyMTIzNAo=  # base64 encoded value of `foobar1234`
```

We'll reference this secret from the RDS CRD. Save it in a file (eg. `provider.yaml`) then apply it using kubectl.

```console
$ kubectl apply -f provider.yaml
```

> **Note:** here, data key (eg. `password`) must be in snake case format (same as the tf configuration file)

> In module, there is no concept of sensitive module. But if you want to use an input as sensitve, you can just put it in the secret as a key value pair.

## 4. Create RDS

Now, we'll create the RDS CRD. The yaml is given below:

```yaml
apiVersion: modules.kubeform.com/v1alpha1
kind: RDS
metadata:
  name: test-rds
spec:
  providerRef: 
    name: "aws"
  secretRef:
    name: "rds-pass"
  identifier: "demodb"
  engine: "mysql"
  engineVersion: "5.7.19"
  instanceClass: "db.t2.large"
  allocatedStorage: "5"
  storageEncrypted: false
  name: "demodb"
  username: "user"
  port: "3306"
  vpcSecurityGroupIDS: 
    - <vpc security group ID>
  maintenanceWindow: "Mon:00:00-Mon:03:00"
  backupWindow: "03:00-04:00"
  subnetIDS: 
    - <Subnet ID>
  family: "mysql5.7"
  majorEngineVersion: "5.7"
  skipFinalSnapshot: true
  deletionProtection: false
  tags: 
    Owner: "user"
    Environment: "dev"
  parameters:
    - name: "character_set_client"
      value: "utf8"
    - name: "character_set_server"
      value: "utf8"
  options:
     - option_name: "MARIADB_AUDIT_PLUGIN"
       option_settings:
         - name: "SERVER_AUDIT_EVENTS"
           value: "CONNECT"
         - name: "SERVER_AUDIT_FILE_ROTATIONS"
           value: "37"
```

Here, we can see that the provider secret is referenced using a field called `providerRef` and the sensitive value secret is referenced using a field called `secretRef`.

Save it in a file (eg. `rds.yaml`) then apply it using kubectl.

```console
$ kubectl apply -f rds.yaml
```

After that, an AWS RDS will be created!

## Delete RDS

To delete the RDS just run:

```console
kubectl delete -f rds.yaml
```
