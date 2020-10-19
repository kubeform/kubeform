---
title: Upgrade | Kubeform
description: Kubeform Upgrade
menu:
  docs_{{ .version }}:
    identifier: upgrade-kubeform
    name: Upgrade
    parent: setup
    weight: 30
product_name: kubeform
menu_name: docs_{{ .version }}
section_menu_id: setup
---

# Upgrading Kubeform

This guide will show you how to upgrade Kubeform operator. Here, we are going to show how to update the license and how to upgrade between two Kubeform versions.

## Updating License

Kubeform support updating license without requiring any re-installation or restart. Kubeform creates a Secret named `<helm release name>-license` with the license file. You just need to update the Secret. The changes will propagate automatically to the operator and it will use the updated license going forward.

Follow the below instructions to update the license:

- Get a new license and save it into a file.
- Then, run the following upgrade command based on your installation.

<ul class="nav nav-tabs" id="installerTab" role="tablist">
  <li class="nav-item">
    <a class="nav-link active" id="helm3-tab" data-toggle="tab" href="#helm3" role="tab" aria-controls="helm3" aria-selected="true">Helm 3</a>
  </li>
  <li class="nav-item">
    <a class="nav-link" id="helm2-tab" data-toggle="tab" href="#helm2" role="tab" aria-controls="helm2" aria-selected="false">Helm 2</a>
  </li>
  <li class="nav-item">
    <a class="nav-link" id="script-tab" data-toggle="tab" href="#script" role="tab" aria-controls="script" aria-selected="false">YAML</a>
  </li>
</ul>
<div class="tab-content" id="installerTabContent">
  <div class="tab-pane fade show active" id="helm3" role="tabpanel" aria-labelledby="helm3-tab">

## Using Helm 3

```bash
$ helm upgrade kubeform -n kube-system appscode/kubeform  \
    --reuse-values                                        \
    --set-file license=/path/to/new/license.txt
```

</div>
<div class="tab-pane fade" id="helm2" role="tabpanel" aria-labelledby="helm2-tab">

## Using Helm 2

```bash
$ helm upgrade kubeform appscode/kubeform                 \
    --reuse-values                                        \
    --set-file license=/path/to/new/license.txt
```

</div>
<div class="tab-pane fade" id="script" role="tabpanel" aria-labelledby="script-tab">

## Using YAML (with helm 3)

```bash
$ helm template kubeform appscode/kubeform                \
    --set-file license=/path/to/new/license.txt           \
    --show-only templates/license.yaml                    \
    --no-hooks | kubectl apply -f -
```

</div>
</div>

## Upgrading between patch versions

If you are upgrading Kubeform to a patch release, please reapply the [installation instructions](/docs/setup/README.md). That will upgrade the operator pod to the new version and fix any RBAC issues.
