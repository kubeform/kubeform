---
title: Upgrade | Kubeform
description: Kubeform Upgrade
menu:
  docs_{{ .version }}:
    identifier: upgrade-kubeform
    name: Upgrade
    parent: setup
    weight: 20
product_name: kubeform
menu_name: docs_{{ .version }}
section_menu_id: setup
---

# Upgrading Kubeform

This guide will show you how to upgrade various Kubeform components. Here, we are going to show how to upgrade from an old Kubeform version to the new version, how to migrate between the enterprise edition and community edition, and how to update the license, etc.

## Migration Between Community Edition and Enterprise Edition

Kubeform `{{< param "info.installer" >}}` supports seamless migration between community edition and enterprise edition. You can run the following commands to migrate between them.

<ul class="nav nav-tabs" id="migrationTab" role="tablist">
  <li class="nav-item">
    <a class="nav-link active" id="mgr-helm3-tab" data-toggle="tab" href="#mgr-helm3" role="tab" aria-controls="mgr-helm3" aria-selected="true">Helm 3</a>
  </li>
  <li class="nav-item">
    <a class="nav-link" id="mgr-yaml-tab" data-toggle="tab" href="#mgr-yaml" role="tab" aria-controls="mgr-yaml" aria-selected="false">YAML</a>
  </li>
</ul>
<div class="tab-content" id="migrationTabContent">
  <div class="tab-pane fade show active" id="mgr-helm3" role="tabpanel" aria-labelledby="mgr-helm3">

#### Using Helm 3

**From Community Edition to Enterprise Edition:**

In order to migrate from Kubeform community edition to Kubeform enterprise edition, please run the following command,

```bash
helm upgrade kubeform-provider-*** -n kubeform appscode/kubeform-provider-*** \
  --reuse-values \
  --set-file kubeform-provider.license=/path/to/kubeform-provider-license.txt
```

**From Enterprise Edition to Community Edition:**

In order to migrate from Kubeform enterprise edition to Kubeform community edition, please run the following command,

```bash
helm upgrade kubeform-provider-*** -n kubeform appscode/kubeform-provider-*** \
  --reuse-values \
  --set-file kubeform-provider.license=/path/to/kubeform-provider-license.txt
```

</div>
<div class="tab-pane fade" id="mgr-yaml" role="tabpanel" aria-labelledby="mgr-yaml">

**Using YAML (with helm 3)**

**From Community Edition to Enterprise Edition:**

In order to migrate from Kubeform community edition to Kubeform enterprise edition, please run the following command,

```bash
# Install Kubeform enterprise edition
helm template kubeform-provider-*** -n kubeform appscode/kubeform-provider-*** \
  --version {{< param "info.version" >}} \
  --set kubeform-provider.cleaner.skip=true \
  --set-file kubeform-provider.license=/path/to/kubeform-provider-license.txt | kubectl apply -f -
```

**From Enterprise Edition to Community Edition:**

In order to migrate from Kubeform enterprise edition to Kubeform community edition, please run the following command,

```bash
# Install Kubeform community edition
helm template kubeform-provider-*** -n kubeform appscode/kubeform-provider-*** \
  --version {{< param "info.version" >}} \
  --set kubeform-provider.cleaner.skip=true \
  --set-file kubeform-provider.license=/path/to/kubeform-provider-license.txt | kubectl apply -f -
```

</div>
</div>

## Updating License

Kubeform support updating license without requiring any re-installation. Kubeform creates a Secret named `<helm release name>-license` with the license file. You just need to update the Secret. The changes will propagate automatically to the operator and it will use the updated license going forward.

Follow the below instructions to update the license:

- Get a new license and save it into a file.
- Then, run the following upgrade command based on your installation.

<ul class="nav nav-tabs" id="luTabs" role="tablist">
  <li class="nav-item">
    <a class="nav-link active" id="lu-helm3-tab" data-toggle="tab" href="#lu-helm3" role="tab" aria-controls="lu-helm3" aria-selected="true">Helm 3</a>
  </li>
  <li class="nav-item">
    <a class="nav-link" id="lu-yaml-tab" data-toggle="tab" href="#lu-yaml" role="tab" aria-controls="lu-yaml" aria-selected="false">YAML</a>
  </li>
</ul>
<div class="tab-content" id="luTabContent">
  <div class="tab-pane fade show active" id="lu-helm3" role="tabpanel" aria-labelledby="lu-helm3">

#### Using Helm 3

```bash
helm upgrade kubeform-provider-*** -n kubeform appscode/kubeform-provider-*** \
  --reuse-values \
  --set-file kubeform-provider.license=/path/to/new/license.txt
```

</div>
<div class="tab-pane fade" id="lu-yaml" role="tabpanel" aria-labelledby="lu-yaml">

#### Using YAML (with helm 3)

```bash
helm template kubeform-provider-*** -n kubeform appscode/kubeform-provider-*** \
  --set kubeform-provider.cleaner.skip=true \
  --show-only appscode/kubeform-provider/templates/license.yaml \
  --set-file kubeform-provider.license=/path/to/new/license.txt | kubectl apply -f -
```

</div>
</div>
