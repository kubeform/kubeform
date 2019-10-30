---
title: Instance
menu:
  docs_{{ .version }}:
    identifier: instance-linode.kubeform.com
    name: Instance
    parent: linode.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Instance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `linode.kubeform.com/v1alpha1` |
|    `kind` | string | `Instance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[InstanceSpec](#instancespec)***||
| `status` | ***[InstanceStatus](#instancestatus)***||
## InstanceSpec

Appears on:[Instance](#instance), [InstanceStatus](#instancestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `alerts` | ***[[]InstanceSpecAlerts](#instancespecalerts)***| ***(Optional)*** |
| `authorizedKeys` | ***[]string***| ***(Optional)*** A list of SSH public keys to deploy for the root user on the newly created Linode. Only accepted if 'image' is provided.|
| `authorizedUsers` | ***[]string***| ***(Optional)*** A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the `root` user's `~/.ssh/authorized_keys` file automatically. Only accepted if 'image' is provided.|
| `backupID` | ***int***| ***(Optional)*** A Backup ID from another Linode's available backups. Your User must have read_write access to that Linode, the Backup must have a status of successful, and the Linode must be deployed to the same region as the Backup. See /linode/instances/{linodeId}/backups for a Linode's available backups. This field and the image field are mutually exclusive.|
| `backups` | ***[[]InstanceSpecBackups](#instancespecbackups)***| ***(Optional)*** Information about this Linode's backups status.|
| `backupsEnabled` | ***bool***| ***(Optional)*** If this field is set to true, the created Linode will automatically be enrolled in the Linode Backup service. This will incur an additional charge. The cost for the Backup service is dependent on the Type of Linode deployed.|
| `bootConfigLabel` | ***string***| ***(Optional)*** The Label of the Instance Config that should be used to boot the Linode instance.|
| `config` | ***[[]InstanceSpecConfig](#instancespecconfig)***| ***(Optional)*** Configuration profiles define the VM settings and boot behavior of the Linode Instance.|
| `disk` | ***[[]InstanceSpecDisk](#instancespecdisk)***| ***(Optional)*** |
| `group` | ***string***| ***(Optional)*** The display group of the Linode instance.|
| `image` | ***string***| ***(Optional)*** An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with private/. See /images for more information on the Images available for you to use.|
| `ipAddress` | ***string***| ***(Optional)*** This Linode's Public IPv4 Address. If there are multiple public IPv4 addresses on this Instance, an arbitrary address will be used for this field.|
| `ipv4` | ***[]string***| ***(Optional)*** This Linode's IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a single private IPv4 address if needed. You may need to open a support ticket to get additional IPv4 addresses.|
| `ipv6` | ***string***| ***(Optional)*** This Linode's IPv6 SLAAC addresses. This address is specific to a Linode, and may not be shared.|
| `label` | ***string***| ***(Optional)*** The Linode's label is for display purposes only. If no label is provided for a Linode, a default will be assigned|
| `privateIP` | ***bool***| ***(Optional)*** If true, the created Linode will have private networking enabled, allowing use of the 192.168.128.0/17 network within the Linode's region.|
| `privateIPAddress` | ***string***| ***(Optional)*** This Linode's Private IPv4 Address.  The regional private IP address range is 192.168.128/17 address shared by all Linode Instances in a region.|
| `region` | ***string***|This is the location where the Linode was deployed. This cannot be changed without opening a support ticket.|
| `specs` | ***[[]InstanceSpecSpecs](#instancespecspecs)***| ***(Optional)*** |
| `stackscriptID` | ***int***| ***(Optional)*** The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an Image that is compatible with this StackScript.|
| `status` | ***string***| ***(Optional)*** The status of the instance, indicating the current readiness state.|
| `swapSize` | ***int***| ***(Optional)*** When deploying from an Image, this field is optional with a Linode API default of 512mb, otherwise it is ignored. This is used to set the swap disk size for the newly-created Linode.|
| `tags` | ***[]string***| ***(Optional)*** An array of tags applied to this object. Tags are for organizational purposes only.|
| `type` | ***string***| ***(Optional)*** The type of instance to be deployed, determining the price and size.|
| `watchdogEnabled` | ***bool***| ***(Optional)*** The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie will give up if there have been more than 5 boot jobs issued within 15 minutes.|
## InstanceSpecAlerts

Appears on:[InstanceSpec](#instancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cpu` | ***int***| ***(Optional)*** The percentage of CPU usage required to trigger an alert. If the average CPU usage over two hours exceeds this value, we'll send you an alert. If this is set to 0, the alert is disabled.|
| `io` | ***int***| ***(Optional)*** The amount of disk IO operation per second required to trigger an alert. If the average disk IO over two hours exceeds this value, we'll send you an alert. If set to 0, this alert is disabled.|
| `networkIn` | ***int***| ***(Optional)*** The amount of incoming traffic, in Mbit/s, required to trigger an alert. If the average incoming traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.|
| `networkOut` | ***int***| ***(Optional)*** The amount of outbound traffic, in Mbit/s, required to trigger an alert. If the average outbound traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.|
| `transferQuota` | ***int***| ***(Optional)*** The percentage of network transfer that may be used before an alert is triggered. When this value is exceeded, we'll alert you. If this is set to 0 (zero), the alert is disabled.|
## InstanceSpecBackups

Appears on:[InstanceSpec](#instancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** If this Linode has the Backup service enabled.|
| `schedule` | ***[[]InstanceSpecBackupsSchedule](#instancespecbackupsschedule)***| ***(Optional)*** |
## InstanceSpecBackupsSchedule

Appears on:[InstanceSpecBackups](#instancespecbackups)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `day` | ***string***| ***(Optional)*** The day ('Sunday'-'Saturday') of the week that your Linode's weekly Backup is taken. If not set manually, a day will be chosen for you. Backups are taken every day, but backups taken on this day are preferred when selecting backups to retain for a longer period.  If not set manually, then when backups are initially enabled, this may come back as 'Scheduling' until the day is automatically selected.|
| `window` | ***string***| ***(Optional)*** The window ('W0'-'W22') in which your backups will be taken, in UTC. A backups window is a two-hour span of time in which the backup may occur. For example, 'W10' indicates that your backups should be taken between 10:00 and 12:00. If you do not choose a backup window, one will be selected for you automatically.  If not set manually, when backups are initially enabled this may come back as Scheduling until the window is automatically selected.|
## InstanceSpecConfig

Appears on:[InstanceSpec](#instancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `comments` | ***string***| ***(Optional)*** Optional field for arbitrary User comments on this Config.|
| `devices` | ***[[]InstanceSpecConfigDevices](#instancespecconfigdevices)***| ***(Optional)*** Device sda-sdh can be either a Disk or Volume identified by disk_label or volume_id. Only one type per slot allowed.|
| `helpers` | ***[[]InstanceSpecConfigHelpers](#instancespecconfighelpers)***| ***(Optional)*** Helpers enabled when booting to this Linode Config.|
| `kernel` | ***string***| ***(Optional)*** A Kernel ID to boot a Linode with. Default is based on image choice. (examples: linode/latest-64bit, linode/grub2, linode/direct-disk)|
| `label` | ***string***|The Config's label for display purposes.  Also used by `boot_config_label`.|
| `memoryLimit` | ***int***| ***(Optional)*** Defaults to the total RAM of the Linode|
| `rootDevice` | ***string***| ***(Optional)*** The root device to boot. The corresponding disk must be attached.|
| `runLevel` | ***string***| ***(Optional)*** Defines the state of your Linode after booting. Defaults to default.|
| `virtMode` | ***string***| ***(Optional)*** Controls the virtualization mode. Defaults to paravirt.|
## InstanceSpecConfigDevices

Appears on:[InstanceSpecConfig](#instancespecconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `sda` | ***[[]InstanceSpecConfigDevicesSda](#instancespecconfigdevicessda)***| ***(Optional)*** |
| `sdb` | ***[[]InstanceSpecConfigDevicesSdb](#instancespecconfigdevicessdb)***| ***(Optional)*** |
| `sdc` | ***[[]InstanceSpecConfigDevicesSdc](#instancespecconfigdevicessdc)***| ***(Optional)*** |
| `sdd` | ***[[]InstanceSpecConfigDevicesSdd](#instancespecconfigdevicessdd)***| ***(Optional)*** |
| `sde` | ***[[]InstanceSpecConfigDevicesSde](#instancespecconfigdevicessde)***| ***(Optional)*** |
| `sdf` | ***[[]InstanceSpecConfigDevicesSdf](#instancespecconfigdevicessdf)***| ***(Optional)*** |
| `sdg` | ***[[]InstanceSpecConfigDevicesSdg](#instancespecconfigdevicessdg)***| ***(Optional)*** |
| `sdh` | ***[[]InstanceSpecConfigDevicesSdh](#instancespecconfigdevicessdh)***| ***(Optional)*** |
## InstanceSpecConfigDevicesSda

Appears on:[InstanceSpecConfigDevices](#instancespecconfigdevices)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskID` | ***int***| ***(Optional)*** The Disk ID to map to this disk slot|
| `diskLabel` | ***string***| ***(Optional)*** The `label` of the `disk` to map to this `device` slot.|
| `volumeID` | ***int***| ***(Optional)*** The Block Storage volume ID to map to this disk slot|
## InstanceSpecConfigDevicesSdb

Appears on:[InstanceSpecConfigDevices](#instancespecconfigdevices)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskID` | ***int***| ***(Optional)*** |
| `diskLabel` | ***string***| ***(Optional)*** |
| `volumeID` | ***int***| ***(Optional)*** |
## InstanceSpecConfigDevicesSdc

Appears on:[InstanceSpecConfigDevices](#instancespecconfigdevices)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskID` | ***int***| ***(Optional)*** |
| `diskLabel` | ***string***| ***(Optional)*** |
| `volumeID` | ***int***| ***(Optional)*** |
## InstanceSpecConfigDevicesSdd

Appears on:[InstanceSpecConfigDevices](#instancespecconfigdevices)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskID` | ***int***| ***(Optional)*** |
| `diskLabel` | ***string***| ***(Optional)*** |
| `volumeID` | ***int***| ***(Optional)*** |
## InstanceSpecConfigDevicesSde

Appears on:[InstanceSpecConfigDevices](#instancespecconfigdevices)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskID` | ***int***| ***(Optional)*** |
| `diskLabel` | ***string***| ***(Optional)*** |
| `volumeID` | ***int***| ***(Optional)*** |
## InstanceSpecConfigDevicesSdf

Appears on:[InstanceSpecConfigDevices](#instancespecconfigdevices)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskID` | ***int***| ***(Optional)*** |
| `diskLabel` | ***string***| ***(Optional)*** |
| `volumeID` | ***int***| ***(Optional)*** |
## InstanceSpecConfigDevicesSdg

Appears on:[InstanceSpecConfigDevices](#instancespecconfigdevices)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskID` | ***int***| ***(Optional)*** |
| `diskLabel` | ***string***| ***(Optional)*** |
| `volumeID` | ***int***| ***(Optional)*** |
## InstanceSpecConfigDevicesSdh

Appears on:[InstanceSpecConfigDevices](#instancespecconfigdevices)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskID` | ***int***| ***(Optional)*** |
| `diskLabel` | ***string***| ***(Optional)*** |
| `volumeID` | ***int***| ***(Optional)*** |
## InstanceSpecConfigHelpers

Appears on:[InstanceSpecConfig](#instancespecconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `devtmpfsAutomount` | ***bool***| ***(Optional)*** Populates the /dev directory early during boot without udev. Defaults to false.|
| `distro` | ***bool***| ***(Optional)*** Controls the behavior of the Linode Config's Distribution Helper setting.|
| `modulesDep` | ***bool***| ***(Optional)*** Creates a modules dependency file for the Kernel you run.|
| `network` | ***bool***| ***(Optional)*** Controls the behavior of the Linode Config's Network Helper setting, used to automatically configure additional IP addresses assigned to this instance.|
| `updatedbDisabled` | ***bool***| ***(Optional)*** Disables updatedb cron job to avoid disk thrashing.|
## InstanceSpecDisk

Appears on:[InstanceSpec](#instancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authorizedKeys` | ***[]string***| ***(Optional)*** A list of SSH public keys to deploy for the root user on the newly created Linode. Only accepted if 'image' is provided.|
| `authorizedUsers` | ***[]string***| ***(Optional)*** A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the `root` user's `~/.ssh/authorized_keys` file automatically. Only accepted if 'image' is provided.|
| `filesystem` | ***string***| ***(Optional)*** The Disk filesystem can be one of: raw, swap, ext3, ext4, initrd (max 32mb)|
| `ID` | ***int***| ***(Optional)*** The ID of the Disk (for use in Linode Image resources and Linode Instance Config Devices)|
| `image` | ***string***| ***(Optional)*** An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with private/.|
| `label` | ***string***|The disks label, which acts as an identifier in Terraform.|
| `readOnly` | ***bool***| ***(Optional)*** If true, this Disk is read-only.|
| `size` | ***int***|The size of the Disk in MB.|
| `stackscriptID` | ***int***| ***(Optional)*** The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an Image that is compatible with this StackScript.|
## InstanceSpecSpecs

Appears on:[InstanceSpec](#instancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `disk` | ***int***| ***(Optional)*** The amount of storage space, in GB. this Linode has access to. A typical Linode will divide this space between a primary disk with an image deployed to it, and a swap disk, usually 512 MB. This is the default configuration created when deploying a Linode with an image without specifying disks.|
| `memory` | ***int***| ***(Optional)*** The amount of RAM, in MB, this Linode has access to. Typically a Linode will choose to boot with all of its available RAM, but this can be configured in a Config profile.|
| `transfer` | ***int***| ***(Optional)*** The amount of network transfer this Linode is allotted each month.|
| `vcpus` | ***int***| ***(Optional)*** The number of vcpus this Linode has access to. Typically a Linode will choose to boot with all of its available vcpus, but this can be configured in a Config Profile.|
## InstanceStatus

Appears on:[Instance](#instance)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[InstanceSpec](#instancespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[InstanceStatus](#instancestatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `disk.<index>.root_pass` | ***string*** |The password that will be initialially assigned to the 'root' user account.|
| `disk.<index>.stackscript_data` | ***map[string]string*** |An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode. Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.|
| `root_pass` | ***string*** |The password that will be initialially assigned to the 'root' user account.|
| `stackscript_data` | ***map[string]string*** |An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode. Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.|
