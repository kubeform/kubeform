---
title: Uninstall
description: Kubeform Uninstall
menu:
  product_kubeform_0.0.1:
    identifier: uninstall-Kubeform
    name: Uninstall
    parent: setup
    weight: 20
product_name: Kubeform
menu_name: product_kubeform_0.0.1
section_menu_id: setup
---
# Uninstall Kubeform

To uninstall kubeform installed using helm, use the following command.

```console
$ helm delete kfc
```

If you want to remove the release from the store and make its name free for later use, Then provide the `--purge` flag:

```console
$ helm delete kfc --purge
```