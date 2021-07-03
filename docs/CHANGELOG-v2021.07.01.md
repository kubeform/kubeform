---
title: Changelog | Kubeform
description: Changelog
menu:
  docs_{{.version}}:
    identifier: changelog-kubeform-v2021.07.01
    name: Changelog-v2021.07.01
    parent: welcome
    weight: 20210701
product_name: kubeform
menu_name: docs_{{.version}}
section_menu_id: welcome
url: /docs/{{.version}}/welcome/changelog-v2021.07.01/
aliases:
  - /docs/{{.version}}/CHANGELOG-v2021.07.01/
---

# Kubeform v2021.07.01 (2021-07-03)


## [kubeform/installer](https://github.com/kubeform/installer)

### [v2021.07.01](https://github.com/kubeform/installer/releases/tag/v2021.07.01)

- [17817522](https://github.com/kubeform/installer/commit/17817522) Prepare for release v2021.07.01 (#72)
- [96b09af7](https://github.com/kubeform/installer/commit/96b09af7) Add prepare-release script (#71)
- [74c1f661](https://github.com/kubeform/installer/commit/74c1f661) Update azurerm installer for provider@v1.44.1-0.20210702030130-ff3fb4ec388d gen@8ea2428 (#70)
- [4dbac78f](https://github.com/kubeform/installer/commit/4dbac78f) Update google installer for provider@v1.20.1-0.20210628164239-4d5f8365936e gen@50bb776 (#68)
- [3e751e0a](https://github.com/kubeform/installer/commit/3e751e0a) Update aws installer for provider@v1.60.1-0.20210624223533-7ed5f82d440d gen@a0bbaa1 (#69)
- [0d2b6836](https://github.com/kubeform/installer/commit/0d2b6836) Update controller deployment args (#61)
- [f581b4b3](https://github.com/kubeform/installer/commit/f581b4b3) Update azurerm installer for provider@v1.44.1-0.20210702030130-ff3fb4ec388d gen@b764ec6 (#60)
- [22583694](https://github.com/kubeform/installer/commit/22583694) Update aws installer for provider@v1.60.1-0.20210624223533-7ed5f82d440d gen@715606a (#59)
- [0255f21d](https://github.com/kubeform/installer/commit/0255f21d) Update google installer for provider@v1.20.1-0.20210628164239-4d5f8365936e gen@c066b52 (#58)
- [69d426b8](https://github.com/kubeform/installer/commit/69d426b8) Update digitalocean installer for provider@v1.23.1-0.20210629201646-266cb440f82b gen@ecdba3f (#57)
- [aecce8d2](https://github.com/kubeform/installer/commit/aecce8d2) Update linode installer for provider@v1.19.1 gen@44ef72c (#56)
- [fe41e17b](https://github.com/kubeform/installer/commit/fe41e17b) Fix writing crds
- [8bfb272f](https://github.com/kubeform/installer/commit/8bfb272f) Update Makefile
- [8d78644a](https://github.com/kubeform/installer/commit/8d78644a) Add correct license file
- [f56f3855](https://github.com/kubeform/installer/commit/f56f3855) Generate catalog.json
- [8835aeec](https://github.com/kubeform/installer/commit/8835aeec) Update generated crd yamls
- [8f1457d6](https://github.com/kubeform/installer/commit/8f1457d6) Generate schema for chart values files
- [33d07bb7](https://github.com/kubeform/installer/commit/33d07bb7) Update generated code
- [a25a27ce](https://github.com/kubeform/installer/commit/a25a27ce) Remove description for crd schema
- [28aa5e9b](https://github.com/kubeform/installer/commit/28aa5e9b) Generate bundled crd yamls
- [5fd44fd3](https://github.com/kubeform/installer/commit/5fd44fd3) Enable per provider installation (#52)
- [82cacfcd](https://github.com/kubeform/installer/commit/82cacfcd) Use Kubernetes v1.21.0 toolchain (#51)
- [7f158f79](https://github.com/kubeform/installer/commit/7f158f79) Use Kubernetes v1.21.0 toolchain (#50)
- [9508d7ab](https://github.com/kubeform/installer/commit/9508d7ab) Use Kubernetes v1.21.0 toolchain (#49)
- [e9f0cdab](https://github.com/kubeform/installer/commit/e9f0cdab) Update Kubernetes toolchain to v1.21.0 (#48)
- [993de1f6](https://github.com/kubeform/installer/commit/993de1f6) Update README.md
- [b9a55c6b](https://github.com/kubeform/installer/commit/b9a55c6b) Update repository config (#45)
- [9ef0a2aa](https://github.com/kubeform/installer/commit/9ef0a2aa) Update repository config (#44)
- [ced08d5b](https://github.com/kubeform/installer/commit/ced08d5b) Update Kubernetes v1.18.9 dependencies (#43)
- [0380d98e](https://github.com/kubeform/installer/commit/0380d98e) Update Kubernetes v1.18.9 dependencies (#42)
- [26e12b5d](https://github.com/kubeform/installer/commit/26e12b5d) Update repository config (#41)
- [17758984](https://github.com/kubeform/installer/commit/17758984) Update repository config (#40)
- [542438ba](https://github.com/kubeform/installer/commit/542438ba) Remove generated openapi schema
- [bd2d875e](https://github.com/kubeform/installer/commit/bd2d875e) Speed up schema generation process (#39)
- [02a4d5b2](https://github.com/kubeform/installer/commit/02a4d5b2) Update repository config (#38)
- [a3cd80c1](https://github.com/kubeform/installer/commit/a3cd80c1) Update repository config (#37)
- [be8ccaf9](https://github.com/kubeform/installer/commit/be8ccaf9) Update Kubernetes v1.18.9 dependencies (#36)
- [987a20b6](https://github.com/kubeform/installer/commit/987a20b6) Update repository config (#35)
- [7ccfdec6](https://github.com/kubeform/installer/commit/7ccfdec6) Update Kubernetes v1.18.9 dependencies (#34)
- [9ff986f1](https://github.com/kubeform/installer/commit/9ff986f1) Update Kubernetes v1.18.9 dependencies (#33)
- [a89c8233](https://github.com/kubeform/installer/commit/a89c8233) Update Kubernetes v1.18.9 dependencies (#32)



## [kubeform/provider-aws-api](https://github.com/kubeform/provider-aws-api)

### [v0.1.0](https://github.com/kubeform/provider-aws-api/releases/tag/v0.1.0)

- [7434d30](https://github.com/kubeform/provider-aws-api/commit/7434d30) Cleanup release tracker script
- [f0759e2](https://github.com/kubeform/provider-aws-api/commit/f0759e2) Update Makefile
- [6156c5c](https://github.com/kubeform/provider-aws-api/commit/6156c5c) Generate code for provider@ gen@9d9130d (#2)
- [1da61ef](https://github.com/kubeform/provider-aws-api/commit/1da61ef) Don't run openapi target
- [16353a8](https://github.com/kubeform/provider-aws-api/commit/16353a8) Generate code for provider: gen:9b63997 (#1)



## [kubeform/provider-aws-controller](https://github.com/kubeform/provider-aws-controller)

### [v0.1.0](https://github.com/kubeform/provider-aws-controller/releases/tag/v0.1.0)

- [56bd3b3](https://github.com/kubeform/provider-aws-controller/commit/56bd3b3) Prepare for release v0.1.0 (#6)
- [f7f8acf](https://github.com/kubeform/provider-aws-controller/commit/f7f8acf) Generate code for provider@v1.60.1-0.20210624223533-7ed5f82d440d gen@a0bbaa1 (#5)
- [076a1c3](https://github.com/kubeform/provider-aws-controller/commit/076a1c3) Add docker files
- [126b900](https://github.com/kubeform/provider-aws-controller/commit/126b900) Generate code for provider@v1.60.1-0.20210624223533-7ed5f82d440d gen@715606a (#2)
- [437d84c](https://github.com/kubeform/provider-aws-controller/commit/437d84c) Update dependencies
- [956cffa](https://github.com/kubeform/provider-aws-controller/commit/956cffa) Update Makefile
- [b136c68](https://github.com/kubeform/provider-aws-controller/commit/b136c68) Enable GH actions



## [kubeform/provider-azurerm-api](https://github.com/kubeform/provider-azurerm-api)

### [v0.1.0](https://github.com/kubeform/provider-azurerm-api/releases/tag/v0.1.0)

- [366b7a7](https://github.com/kubeform/provider-azurerm-api/commit/366b7a7) Generate code for provider@v1.44.1-0.20210702030130-ff3fb4ec388d gen@b764ec6 (#2)
- [e7e5850](https://github.com/kubeform/provider-azurerm-api/commit/e7e5850) Update Makefile
- [b0f38a6](https://github.com/kubeform/provider-azurerm-api/commit/b0f38a6) Generate code for provider@ gen@70f5b23 (#1)
- [9b3e3d9](https://github.com/kubeform/provider-azurerm-api/commit/9b3e3d9) Cleanup release tracker script
- [6f0e3a0](https://github.com/kubeform/provider-azurerm-api/commit/6f0e3a0) Don't run openapi target



## [kubeform/provider-azurerm-controller](https://github.com/kubeform/provider-azurerm-controller)

### [v0.1.0](https://github.com/kubeform/provider-azurerm-controller/releases/tag/v0.1.0)

- [57f13b82](https://github.com/kubeform/provider-azurerm-controller/commit/57f13b82) Prepare for release v0.1.0 (#4)
- [09a4f37a](https://github.com/kubeform/provider-azurerm-controller/commit/09a4f37a) Generate code for provider@v1.44.1-0.20210702030130-ff3fb4ec388d gen@8ea2428 (#3)
- [a03ce88b](https://github.com/kubeform/provider-azurerm-controller/commit/a03ce88b) Increase linter timeout to 30 m
- [76535d27](https://github.com/kubeform/provider-azurerm-controller/commit/76535d27) Generate controller
- [8f28c9e4](https://github.com/kubeform/provider-azurerm-controller/commit/8f28c9e4) Add docker files
- [a0eb5ada](https://github.com/kubeform/provider-azurerm-controller/commit/a0eb5ada) Update Makefile
- [7cf10220](https://github.com/kubeform/provider-azurerm-controller/commit/7cf10220) Enable GH actions



## [kubeform/provider-digitalocean-api](https://github.com/kubeform/provider-digitalocean-api)

### [v0.1.0](https://github.com/kubeform/provider-digitalocean-api/releases/tag/v0.1.0)

- [95140a6](https://github.com/kubeform/provider-digitalocean-api/commit/95140a6) Generate clientset
- [69b173e](https://github.com/kubeform/provider-digitalocean-api/commit/69b173e) Generate code for provider@v1.23.1-0.20210629201646-266cb440f82b gen@ecdba3f (#2)
- [f585e6a](https://github.com/kubeform/provider-digitalocean-api/commit/f585e6a) Fix Makefile
- [235a4b8](https://github.com/kubeform/provider-digitalocean-api/commit/235a4b8) Cleanup release tracker script
- [8548241](https://github.com/kubeform/provider-digitalocean-api/commit/8548241) Don't run openapi target



## [kubeform/provider-digitalocean-controller](https://github.com/kubeform/provider-digitalocean-controller)

### [v0.1.0](https://github.com/kubeform/provider-digitalocean-controller/releases/tag/v0.1.0)

- [3670ce8](https://github.com/kubeform/provider-digitalocean-controller/commit/3670ce8) Prepare for release v0.1.0 (#6)
- [84b89a8](https://github.com/kubeform/provider-digitalocean-controller/commit/84b89a8) Generate code for provider@v1.23.1-0.20210629201646-266cb440f82b gen@56e8876 (#5)
- [6b2248b](https://github.com/kubeform/provider-digitalocean-controller/commit/6b2248b) Add docker files
- [138241e](https://github.com/kubeform/provider-digitalocean-controller/commit/138241e) Generate code for provider@v1.23.1-0.20210629201646-266cb440f82b gen@ecdba3f (#2)
- [2a490ad](https://github.com/kubeform/provider-digitalocean-controller/commit/2a490ad) Update Makefile
- [4cedee0](https://github.com/kubeform/provider-digitalocean-controller/commit/4cedee0) Enable GH actions



## [kubeform/provider-google-api](https://github.com/kubeform/provider-google-api)

### [v0.1.0](https://github.com/kubeform/provider-google-api/releases/tag/v0.1.0)

- [465f62b](https://github.com/kubeform/provider-google-api/commit/465f62b) Update Makefile
- [1be2b6e](https://github.com/kubeform/provider-google-api/commit/1be2b6e) Generate code for provider@ gen@0786da0 (#1)
- [9070da0](https://github.com/kubeform/provider-google-api/commit/9070da0) Cleanup release tracker script
- [dfa1c84](https://github.com/kubeform/provider-google-api/commit/dfa1c84) Don't run openapi target



## [kubeform/provider-google-controller](https://github.com/kubeform/provider-google-controller)

### [v0.1.0](https://github.com/kubeform/provider-google-controller/releases/tag/v0.1.0)

- [6bdf081](https://github.com/kubeform/provider-google-controller/commit/6bdf081) Prepare for release v0.1.0 (#6)
- [55f2f5c](https://github.com/kubeform/provider-google-controller/commit/55f2f5c) Generate code for provider@v1.20.1-0.20210628164239-4d5f8365936e gen@50bb776 (#5)
- [9cae6cb](https://github.com/kubeform/provider-google-controller/commit/9cae6cb) Generate code for provider@v1.20.1-0.20210628164239-4d5f8365936e gen@c066b52 (#2)
- [165374b](https://github.com/kubeform/provider-google-controller/commit/165374b) Update Makefile
- [fc8a0a3](https://github.com/kubeform/provider-google-controller/commit/fc8a0a3) Enable GH actions



## [kubeform/provider-linode-api](https://github.com/kubeform/provider-linode-api)

### [v0.1.0](https://github.com/kubeform/provider-linode-api/releases/tag/v0.1.0)

- [4a62563](https://github.com/kubeform/provider-linode-api/commit/4a62563) Update Makefile
- [06f1308](https://github.com/kubeform/provider-linode-api/commit/06f1308) Cleanup release tracker script
- [6d2280d](https://github.com/kubeform/provider-linode-api/commit/6d2280d) Generate code for provider@v1.19.1 gen@0a31599 (#3)
- [e558670](https://github.com/kubeform/provider-linode-api/commit/e558670) Don't run openapi target



## [kubeform/provider-linode-controller](https://github.com/kubeform/provider-linode-controller)

### [v0.1.0](https://github.com/kubeform/provider-linode-controller/releases/tag/v0.1.0)

- [5e27eca](https://github.com/kubeform/provider-linode-controller/commit/5e27eca) Prepare for release v0.1.0 (#10)
- [adcfe6a](https://github.com/kubeform/provider-linode-controller/commit/adcfe6a) Generate code for provider@v1.19.1 gen@7f18b4d (#9)
- [9d617cb](https://github.com/kubeform/provider-linode-controller/commit/9d617cb) Use kuebform namespace
- [7d5faec](https://github.com/kubeform/provider-linode-controller/commit/7d5faec) Update image labels
- [1769a52](https://github.com/kubeform/provider-linode-controller/commit/1769a52) Generate code for provider@v1.19.1 gen@dad33a7 (#7)
- [5388383](https://github.com/kubeform/provider-linode-controller/commit/5388383) Update Makefile
- [f0ef47e](https://github.com/kubeform/provider-linode-controller/commit/f0ef47e) Enable GH actions
- [9b91dc8](https://github.com/kubeform/provider-linode-controller/commit/9b91dc8) Generate code for provider@v1.19.1 gen@587f937 (#5)
- [1895e0e](https://github.com/kubeform/provider-linode-controller/commit/1895e0e) Add Dockerfiles (#3)
- [0c7466b](https://github.com/kubeform/provider-linode-controller/commit/0c7466b) Update Makefile (#2)
- [715b0ac](https://github.com/kubeform/provider-linode-controller/commit/715b0ac) Add latest controllers (#1)




