---
title: Changelog | Kubeform
description: Changelog
menu:
  docs_{{.version}}:
    identifier: changelog-kubeform-v2020.10.13
    name: Changelog-v2020.10.13
    parent: welcome
    weight: 20201013
product_name: kubeform
menu_name: docs_{{.version}}
section_menu_id: welcome
url: /docs/{{.version}}/welcome/changelog-v2020.10.13/
aliases:
  - /docs/{{.version}}/CHANGELOG-v2020.10.13/
---

# Kubeform v2020.10.13 (2020-10-13)


## [kubeform/installer](https://github.com/kubeform/installer)

### [v0.2.0](https://github.com/kubeform/installer/releases/tag/v0.2.0)

- [f026037](https://github.com/kubeform/installer/commit/f026037) Prepare for release v0.2.0 (#26)
- [976cb54](https://github.com/kubeform/installer/commit/976cb54) Update repository config (#25)
- [2099bb8](https://github.com/kubeform/installer/commit/2099bb8) Update repository config (#24)
- [d827a76](https://github.com/kubeform/installer/commit/d827a76) Update repository config (#23)
- [e4151e1](https://github.com/kubeform/installer/commit/e4151e1) Update Kubernetes v1.18.9 dependencies (#22)
- [cd74243](https://github.com/kubeform/installer/commit/cd74243) Update repository config (#21)
- [cd8544b](https://github.com/kubeform/installer/commit/cd8544b) Update repository config (#20)
- [08cd53c](https://github.com/kubeform/installer/commit/08cd53c) Make chart registry configurable (#19)
- [503318c](https://github.com/kubeform/installer/commit/503318c) Publish to testing dir for alpha/beta releases
- [e5d5848](https://github.com/kubeform/installer/commit/e5d5848) Update ci.yml
- [c5c1816](https://github.com/kubeform/installer/commit/c5c1816) Enable proxy support for kfc deployment (#14)
- [80caf4f](https://github.com/kubeform/installer/commit/80caf4f) Tag chart and app version as string for yq
- [ebbcd8b](https://github.com/kubeform/installer/commit/ebbcd8b) Update update-release-tracker.sh
- [8e2ab75](https://github.com/kubeform/installer/commit/8e2ab75) Update release.yml
- [7e39627](https://github.com/kubeform/installer/commit/7e39627) Add script to update release tracker on pr merge (#18)
- [95b5165](https://github.com/kubeform/installer/commit/95b5165) Update ci.yml
- [3001fea](https://github.com/kubeform/installer/commit/3001fea) Allow passing namespace for chart testing (#17)
- [6fee2fa](https://github.com/kubeform/installer/commit/6fee2fa) Add commands to update chart (#15)
- [abe2160](https://github.com/kubeform/installer/commit/abe2160) Fix chart release process (#16)
- [da7b5a7](https://github.com/kubeform/installer/commit/da7b5a7) Create .kodiak.toml
- [04d11a2](https://github.com/kubeform/installer/commit/04d11a2) Use recommended kubernetes app labels
- [9b5abd7](https://github.com/kubeform/installer/commit/9b5abd7) Auto generate chart readme file
- [87b3281](https://github.com/kubeform/installer/commit/87b3281) Trigger the workflow on push or pull request
- [2c7b2d1](https://github.com/kubeform/installer/commit/2c7b2d1) Add openapi v3 schema for chart (#11)
- [a7321c1](https://github.com/kubeform/installer/commit/a7321c1) Apply fixes to kubeform chart (#10)
- [22de125](https://github.com/kubeform/installer/commit/22de125) Use kind 0.7.0 (#9)
- [9330fc8](https://github.com/kubeform/installer/commit/9330fc8) Update README.md
- [618a0a4](https://github.com/kubeform/installer/commit/618a0a4) Annotate pods not only deployment



## [kubeform/kfc](https://github.com/kubeform/kfc)

### [v0.2.0](https://github.com/kubeform/kfc/releases/tag/v0.2.0)

- [febd34bd](https://github.com/kubeform/kfc/commit/febd34bd) Prepare for release v0.2.0 (#46)
- [e57ee90d](https://github.com/kubeform/kfc/commit/e57ee90d) Fix observedGeneration calculation (#44)
- [55927499](https://github.com/kubeform/kfc/commit/55927499) Update repository config (#43)
- [1415ec0a](https://github.com/kubeform/kfc/commit/1415ec0a) Update repository config (#42)
- [81fce46e](https://github.com/kubeform/kfc/commit/81fce46e) Update Kubernetes v1.18.9 dependencies (#41)
- [c606c4ce](https://github.com/kubeform/kfc/commit/c606c4ce) Publish docker images to ghcr.io (#40)
- [3817a72a](https://github.com/kubeform/kfc/commit/3817a72a) Update repository config (#39)
- [9ab8983b](https://github.com/kubeform/kfc/commit/9ab8983b) Update Kubernetes v1.18.9 dependencies (#38)
- [68bdbd81](https://github.com/kubeform/kfc/commit/68bdbd81) Update repository config (#37)
- [8a024e6d](https://github.com/kubeform/kfc/commit/8a024e6d) Update repository config (#36)
- [3098c240](https://github.com/kubeform/kfc/commit/3098c240) Update repository config (#35)
- [c8a6e57f](https://github.com/kubeform/kfc/commit/c8a6e57f) Trigger e2e tests on /ok-to-test command (#34)
- [1d2fbe64](https://github.com/kubeform/kfc/commit/1d2fbe64) Update to Kubernetes v1.18.3 (#33)
- [cd0b0f27](https://github.com/kubeform/kfc/commit/cd0b0f27) Update ci.yml
- [4d9cfa70](https://github.com/kubeform/kfc/commit/4d9cfa70) Update update-release-tracker.sh
- [52e185ae](https://github.com/kubeform/kfc/commit/52e185ae) Add script to update release tracker on pr merge (#32)
- [1547ae07](https://github.com/kubeform/kfc/commit/1547ae07) Create .kodiak.toml
- [634a2061](https://github.com/kubeform/kfc/commit/634a2061) Update crazy-max/ghaction-docker-buildx flag
- [3d54ac20](https://github.com/kubeform/kfc/commit/3d54ac20) Trigger the workflow on push or pull request
- [127f3aa6](https://github.com/kubeform/kfc/commit/127f3aa6) Build once in CI pipeline (#31)



## [kubeform/kubeform](https://github.com/kubeform/kubeform)

### [v0.2.0](https://github.com/kubeform/kubeform/releases/tag/v0.2.0)

- [b84896b9](https://github.com/kubeform/kubeform/commit/b84896b91) Update repository config (#48)
- [409988e6](https://github.com/kubeform/kubeform/commit/409988e6b) Update repository config (#47)
- [eeabd422](https://github.com/kubeform/kubeform/commit/eeabd4223) Remove gitignore rule breaking azure sdk vendoring (#45)
- [17067c44](https://github.com/kubeform/kubeform/commit/17067c444) Format CI files
- [69849941](https://github.com/kubeform/kubeform/commit/698499417) Use code-generator v1.18.3 for generating clientset (#44)
- [0e579a89](https://github.com/kubeform/kubeform/commit/0e579a89b) Update to Kubernetes v1.18.3 (#43)
- [18fa0573](https://github.com/kubeform/kubeform/commit/18fa0573b) Update update-release-tracker.sh
- [d6984afc](https://github.com/kubeform/kubeform/commit/d6984afc6) Add script to update release tracker on pr merge (#41)
- [6ffe5f38](https://github.com/kubeform/kubeform/commit/6ffe5f388) Create .kodiak.toml
- [1d71959c](https://github.com/kubeform/kubeform/commit/1d71959ca) Use recommended kubernetes app labels
- [25f23dc2](https://github.com/kubeform/kubeform/commit/25f23dc23) Trigger the workflow on push or pull request
- [d6f90526](https://github.com/kubeform/kubeform/commit/d6f90526c) Update README.md
- [4b59d709](https://github.com/kubeform/kubeform/commit/4b59d7092) Use controller-tools@v0.2.2 to generate structural schema (#40)




