package linode

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linode/linodego"
	"golang.org/x/crypto/sha3"
)

var (
	boolFalse = false
	boolTrue  = true
)

type flattenedAccountCreditCard map[string]string

type flattenedProfileReferrals map[string]interface{}

func flattenAccountCreditCard(card linodego.CreditCard) []flattenedAccountCreditCard {
	return []flattenedAccountCreditCard{{
		"expiry":    card.Expiry,
		"last_four": card.LastFour,
	}}
}

func flattenProfileReferrals(referrals linodego.ProfileReferrals) []flattenedProfileReferrals {
	return []flattenedProfileReferrals{{
		"code":      referrals.Code,
		"url":       referrals.URL,
		"total":     referrals.Total,
		"completed": referrals.Completed,
		"pending":   referrals.Pending,
		"credit":    referrals.Credit,
	}}
}

func flattenInstanceSpecs(instance linodego.Instance) []map[string]int {
	return []map[string]int{{
		"vcpus":    instance.Specs.VCPUs,
		"disk":     instance.Specs.Disk,
		"memory":   instance.Specs.Memory,
		"transfer": instance.Specs.Transfer,
	}}
}

func flattenInstanceAlerts(instance linodego.Instance) []map[string]int {
	return []map[string]int{{
		"cpu":            instance.Alerts.CPU,
		"io":             instance.Alerts.IO,
		"network_in":     instance.Alerts.NetworkIn,
		"network_out":    instance.Alerts.NetworkOut,
		"transfer_quota": instance.Alerts.TransferQuota,
	}}
}

type flattenedInstanceBackupSchedule [1]struct {
	day, window string
}

type flattenedInstanceBackup [1]struct {
	enabled  bool
	schedule flattenedInstanceBackupSchedule
}

func flattenInstanceBackups(instance linodego.Instance) flattenedInstanceBackup {
	return flattenedInstanceBackup{{
		instance.Backups.Enabled,
		flattenedInstanceBackupSchedule{{
			instance.Backups.Schedule.Day,
			instance.Backups.Schedule.Window,
		}},
	}}
}

func flattenInstanceDisks(instanceDisks []linodego.InstanceDisk) (disks []map[string]interface{}, swapSize int) {
	for _, disk := range instanceDisks {
		// Determine if swap exists and the size.  If it does not exist, swap_size=0
		if disk.Filesystem == "swap" {
			swapSize += disk.Size
		}
		disks = append(disks, map[string]interface{}{
			"id":         disk.ID,
			"size":       disk.Size,
			"label":      disk.Label,
			"filesystem": string(disk.Filesystem),
		})
	}
	return
}

func flattenInstanceConfigs(instanceConfigs []linodego.InstanceConfig, diskLabelIDMap map[int]string) (configs []map[string]interface{}) {
	for _, config := range instanceConfigs {

		devices := []map[string]interface{}{{
			"sda": flattenInstanceConfigDevice(config.Devices.SDA, diskLabelIDMap),
			"sdb": flattenInstanceConfigDevice(config.Devices.SDB, diskLabelIDMap),
			"sdc": flattenInstanceConfigDevice(config.Devices.SDC, diskLabelIDMap),
			"sdd": flattenInstanceConfigDevice(config.Devices.SDD, diskLabelIDMap),
			"sde": flattenInstanceConfigDevice(config.Devices.SDE, diskLabelIDMap),
			"sdf": flattenInstanceConfigDevice(config.Devices.SDF, diskLabelIDMap),
			"sdg": flattenInstanceConfigDevice(config.Devices.SDG, diskLabelIDMap),
			"sdh": flattenInstanceConfigDevice(config.Devices.SDH, diskLabelIDMap),
		}}

		// Determine if swap exists and the size.  If it does not exist, swap_size=0
		c := map[string]interface{}{
			"root_device":  config.RootDevice,
			"kernel":       config.Kernel,
			"run_level":    string(config.RunLevel),
			"virt_mode":    string(config.VirtMode),
			"comments":     config.Comments,
			"memory_limit": config.MemoryLimit,
			"label":        config.Label,
			"helpers": []map[string]bool{{
				"updatedb_disabled":  config.Helpers.UpdateDBDisabled,
				"distro":             config.Helpers.Distro,
				"modules_dep":        config.Helpers.ModulesDep,
				"network":            config.Helpers.Network,
				"devtmpfs_automount": config.Helpers.DevTmpFsAutomount,
			}},
			"devices": devices,
		}

		configs = append(configs, c)
	}
	return
}

func createInstanceConfigsFromSet(client linodego.Client, instanceID int, cset []interface{}, diskIDLabelMap map[string]int, detacher volumeDetacher) (map[int]linodego.InstanceConfig, error) {
	configIDMap := make(map[int]linodego.InstanceConfig, len(cset))

	for _, v := range cset {
		config, ok := v.(map[string]interface{})

		if !ok {
			return configIDMap, fmt.Errorf("Error parsing configs")
		}

		configOpts := linodego.InstanceConfigCreateOptions{}

		configOpts.Kernel = config["kernel"].(string)
		configOpts.Label = config["label"].(string)
		configOpts.Comments = config["comments"].(string)

		if helpers, helpersOk := config["helpers"].([]interface{}); helpersOk {
			for _, helper := range helpers {
				if helperMap, helperMapOk := helper.(map[string]interface{}); helperMapOk {
					configOpts.Helpers = &linodego.InstanceConfigHelpers{}
					if updateDBDisabled, found := helperMap["updatedb_disabled"]; found {
						if value, updateDBDisabledOk := updateDBDisabled.(bool); updateDBDisabledOk {
							configOpts.Helpers.UpdateDBDisabled = value
						}
					}
					if distro, found := helperMap["distro"]; found {
						if value, distroOk := distro.(bool); distroOk {
							configOpts.Helpers.Distro = value
						}
					}
					if modulesDep, found := helperMap["modules_dep"]; found {
						if value, modulesDepOk := modulesDep.(bool); modulesDepOk {
							configOpts.Helpers.ModulesDep = value
						}
					}
					if network, found := helperMap["network"]; found {
						if value, networkOk := network.(bool); networkOk {
							configOpts.Helpers.Network = value
						}
					}
					if devTmpFsAutomount, found := helperMap["devtmpfs_automount"]; found {
						if value, devTmpFsAutomountOk := devTmpFsAutomount.(bool); devTmpFsAutomountOk {
							configOpts.Helpers.DevTmpFsAutomount = value
						}
					}
				}
			}
		}

		rootDevice := config["root_device"].(string)
		if rootDevice != "" {
			configOpts.RootDevice = &rootDevice
		}
		// configOpts.InitRD = config["initrd"].(string)
		// TODO(displague) need a disk_label to initrd lookup?
		devices, ok := config["devices"].([]interface{})
		if !ok {
			return configIDMap, fmt.Errorf("Error converting config devices")
		}
		// TODO(displague) ok needed? check it
		for _, device := range devices {
			deviceMap, ok := device.(map[string]interface{})
			if !ok {
				return configIDMap, fmt.Errorf("Error converting config device %#v", device)
			}
			confDevices, err := expandInstanceConfigDeviceMap(deviceMap, diskIDLabelMap)
			if err != nil {
				return configIDMap, err
			}
			if confDevices != nil {
				configOpts.Devices = *confDevices
			}
		}

		if err := detachConfigVolumes(configOpts.Devices, detacher); err != nil {
			return configIDMap, err
		}

		instanceConfig, err := client.CreateInstanceConfig(context.Background(), instanceID, configOpts)
		if err != nil {
			return configIDMap, fmt.Errorf("Error creating Instance Config: %s", err)
		}
		configIDMap[instanceConfig.ID] = *instanceConfig
	}
	return configIDMap, nil

}

func updateInstanceConfigs(client linodego.Client, d *schema.ResourceData, instance linodego.Instance, tfConfigsOld, tfConfigsNew interface{}, diskIDLabelMap map[string]int) (bool, map[string]int, []*linodego.InstanceConfig, error) {
	var updatedConfigMap map[string]int
	var rebootInstance bool
	var updatedConfigs []*linodego.InstanceConfig

	configs, err := client.ListInstanceConfigs(context.Background(), int(instance.ID), nil)
	if err != nil {
		return rebootInstance, updatedConfigMap, updatedConfigs, fmt.Errorf("Error fetching the config for Instance %d: %s", instance.ID, err)
	}

	configMap := make(map[string]linodego.InstanceConfig, len(configs))
	for _, config := range configs {
		if _, duplicate := configMap[config.Label]; duplicate {
			return rebootInstance, updatedConfigMap, updatedConfigs, fmt.Errorf("Error indexing Instance %d Configs: Label '%s' is assigned to multiple configs", instance.ID, config.Label)
		}
		configMap[config.Label] = config
	}

	oldConfigLabels := make([]string, len(tfConfigsOld.([]interface{})))

	for _, tfConfigOld := range tfConfigsOld.([]interface{}) {
		if oldConfig, ok := tfConfigOld.(map[string]interface{}); ok {
			oldConfigLabels = append(oldConfigLabels, oldConfig["label"].(string))
		}
	}
	tfConfigs := tfConfigsNew.([]interface{})
	updatedConfigs = make([]*linodego.InstanceConfig, len(tfConfigs))
	updatedConfigMap = make(map[string]int, len(tfConfigs))
	for _, tfConfig := range tfConfigs {
		tfc, _ := tfConfig.(map[string]interface{})
		label, _ := tfc["label"].(string)
		rootDevice, _ := tfc["root_device"].(string)
		if existingConfig, existing := configMap[label]; existing {
			configUpdateOpts := existingConfig.GetUpdateOptions()
			configUpdateOpts.Kernel = tfc["kernel"].(string)
			configUpdateOpts.RunLevel = tfc["run_level"].(string)
			configUpdateOpts.VirtMode = tfc["virt_mode"].(string)
			configUpdateOpts.RootDevice = rootDevice
			configUpdateOpts.Comments = tfc["comments"].(string)
			configUpdateOpts.MemoryLimit = tfc["memory_limit"].(int)

			tfcHelpersRaw, helpersFound := tfc["helpers"]
			if tfcHelpers, ok := tfcHelpersRaw.([]interface{}); helpersFound && ok {
				helpersMap := tfcHelpers[0].(map[string]interface{})
				configUpdateOpts.Helpers = &linodego.InstanceConfigHelpers{
					UpdateDBDisabled:  helpersMap["updatedb_disabled"].(bool),
					Distro:            helpersMap["distro"].(bool),
					ModulesDep:        helpersMap["modules_dep"].(bool),
					Network:           helpersMap["network"].(bool),
					DevTmpFsAutomount: helpersMap["devtmpfs_automount"].(bool),
				}

			}

			tfcDevicesRaw, devicesFound := tfc["devices"]
			if tfcDevices, ok := tfcDevicesRaw.([]interface{}); devicesFound && ok {
				devices := tfcDevices[0].(map[string]interface{})

				configUpdateOpts.Devices, err = expandInstanceConfigDeviceMap(devices, diskIDLabelMap)

				if err != nil {
					return rebootInstance, updatedConfigMap, updatedConfigs, err
				}
				if configUpdateOpts.Devices != nil && emptyConfigDeviceMap(*configUpdateOpts.Devices) {
					configUpdateOpts.Devices = nil
				}
			} else {
				configUpdateOpts.Devices = nil
			}

			if configUpdateOpts.Devices != nil {
				detacher := makeVolumeDetacher(client, d)

				if detachErr := detachConfigVolumes(*configUpdateOpts.Devices, detacher); detachErr != nil {
					return rebootInstance, updatedConfigMap, updatedConfigs, detachErr
				}
			}

			updatedConfig, err := client.UpdateInstanceConfig(context.Background(), instance.ID, existingConfig.ID, configUpdateOpts)
			if err != nil {
				return rebootInstance, updatedConfigMap, updatedConfigs, fmt.Errorf("Error updating Instance %d Config %d: %s", instance.ID, existingConfig.ID, err)
			}

			updatedConfigMap[updatedConfig.Label] = updatedConfig.ID
		} else {
			detacher := makeVolumeDetacher(client, d)

			configIDMap, err := createInstanceConfigsFromSet(client, instance.ID, []interface{}{tfc}, diskIDLabelMap, detacher)
			if err != nil {
				return rebootInstance, updatedConfigMap, updatedConfigs, err
			}
			for _, config := range configIDMap {
				updatedConfigMap[config.Label] = config.ID
				updatedConfigs = append(updatedConfigs, &config)
			}
		}
	}

	updatedConfigMap, err = deleteInstanceConfigs(client, instance.ID, oldConfigLabels, updatedConfigMap, configMap)
	if err != nil {
		return rebootInstance, updatedConfigMap, updatedConfigs, err
	}

	return rebootInstance, updatedConfigMap, updatedConfigs, nil
}

func deleteInstanceConfigs(client linodego.Client, instanceID int, oldConfigLabels []string, newConfigLabels map[string]int, configMap map[string]linodego.InstanceConfig) (map[string]int, error) {
	for _, oldLabel := range oldConfigLabels {
		if _, found := newConfigLabels[oldLabel]; !found {
			if listedConfig, found := configMap[oldLabel]; found {
				if err := client.DeleteInstanceConfig(context.Background(), instanceID, listedConfig.ID); err != nil {
					return newConfigLabels, err
				}
				delete(newConfigLabels, oldLabel)
			}
		}
	}
	return newConfigLabels, nil
}

func flattenInstanceConfigDevice(dev *linodego.InstanceConfigDevice, diskLabelIDMap map[int]string) []map[string]interface{} {
	if dev == nil || emptyInstanceConfigDevice(*dev) {
		return nil
	}

	if dev.DiskID > 0 {
		ret := map[string]interface{}{
			"disk_id": dev.DiskID,
		}
		if label, found := diskLabelIDMap[dev.DiskID]; found {
			ret["disk_label"] = label
		}
		return []map[string]interface{}{ret}
	}
	return []map[string]interface{}{{
		"volume_id": dev.VolumeID,
	}}

}

// expandInstanceConfigDeviceMap converts a terraform linode_instance config.*.devices map to a InstanceConfigDeviceMap for the Linode API
func expandInstanceConfigDeviceMap(m map[string]interface{}, diskIDLabelMap map[string]int) (deviceMap *linodego.InstanceConfigDeviceMap, err error) {
	if len(m) == 0 {
		return nil, nil
	}
	deviceMap = &linodego.InstanceConfigDeviceMap{}
	for k, rdev := range m {
		devSlots := rdev.([]interface{})
		for _, rrdev := range devSlots {
			dev := rrdev.(map[string]interface{})
			tDevice := new(linodego.InstanceConfigDevice)
			if err := assignConfigDevice(tDevice, dev, diskIDLabelMap); err != nil {
				return nil, err
			}

			*deviceMap = changeInstanceConfigDevice(*deviceMap, k, tDevice)
		}
	}
	return deviceMap, nil
}

// changeInstanceConfigDevice returns a copy of a config device map with the specified disk slot changed to the provided device
func changeInstanceConfigDevice(deviceMap linodego.InstanceConfigDeviceMap, namedSlot string, device *linodego.InstanceConfigDevice) linodego.InstanceConfigDeviceMap {
	tDevice := device
	if tDevice != nil && emptyInstanceConfigDevice(*tDevice) {
		tDevice = nil
	}
	switch namedSlot {
	case "sda":
		deviceMap.SDA = tDevice
	case "sdb":
		deviceMap.SDB = tDevice
	case "sdc":
		deviceMap.SDC = tDevice
	case "sdd":
		deviceMap.SDD = tDevice
	case "sde":
		deviceMap.SDE = tDevice
	case "sdf":
		deviceMap.SDF = tDevice
	case "sdg":
		deviceMap.SDG = tDevice
	case "sdh":
		deviceMap.SDH = tDevice
	}

	return deviceMap
}

// emptyInstanceConfigDevice returns true only when neither the disk or volume have been assigned to a config device
func emptyInstanceConfigDevice(dev linodego.InstanceConfigDevice) bool {
	return (dev.DiskID == 0 && dev.VolumeID == 0)
}

// emptyConfigDeviceMap returns true only when none of the disks in a config device map have been assigned
func emptyConfigDeviceMap(dmap linodego.InstanceConfigDeviceMap) bool {
	drives := []*linodego.InstanceConfigDevice{
		dmap.SDA, dmap.SDB, dmap.SDC, dmap.SDD, dmap.SDE, dmap.SDF, dmap.SDG, dmap.SDH,
	}
	empty := true
	for _, drive := range drives {
		if drive != nil && !emptyInstanceConfigDevice(*drive) {
			empty = false
			break
		}
	}
	return empty
}

type volumeDetacher func(context.Context, int, string) error

func makeVolumeDetacher(client linodego.Client, d *schema.ResourceData) volumeDetacher {
	return func(ctx context.Context, volumeID int, reason string) error {
		log.Printf("[INFO] Detaching Linode Volume %d %s", volumeID, reason)
		if err := client.DetachVolume(ctx, volumeID); err != nil {
			return err
		}

		log.Printf("[INFO] Waiting for Linode Volume %d to detach ...", volumeID)
		if _, err := client.WaitForVolumeLinodeID(ctx, volumeID, nil, int(d.Timeout(schema.TimeoutUpdate).Seconds())); err != nil {
			return err
		}
		return nil
	}
}

func expandInstanceConfigDevice(m map[string]interface{}) *linodego.InstanceConfigDevice {
	var dev *linodego.InstanceConfigDevice
	// be careful of `disk_label string` in m
	if diskID, ok := m["disk_id"]; ok && diskID.(int) > 0 {
		dev = &linodego.InstanceConfigDevice{
			DiskID: diskID.(int),
		}
	} else if volumeID, ok := m["volume_id"]; ok && volumeID.(int) > 0 {
		dev = &linodego.InstanceConfigDevice{
			VolumeID: m["volume_id"].(int),
		}
	}
	return dev
}

func createInstanceDisk(client linodego.Client, instance linodego.Instance, v interface{}, d *schema.ResourceData) (*linodego.InstanceDisk, error) {
	disk, ok := v.(map[string]interface{})

	if !ok {
		return nil, fmt.Errorf("Error converting disk from Terraform Set to golang map")
	}

	diskOpts := linodego.InstanceDiskCreateOptions{
		Label:      disk["label"].(string),
		Filesystem: disk["filesystem"].(string),
		Size:       disk["size"].(int),
	}

	if image, ok := disk["image"]; ok {
		diskOpts.Image = image.(string)

		if rootPass, ok := disk["root_pass"]; ok && rootPass != "" {
			diskOpts.RootPass = rootPass.(string)
		} else {
			var err error
			diskOpts.RootPass, err = createRandomRootPassword()
			if err != nil {
				return nil, err
			}
		}

		if authorizedKeys, ok := disk["authorized_keys"]; ok {
			for _, sshKey := range authorizedKeys.([]interface{}) {
				diskOpts.AuthorizedKeys = append(diskOpts.AuthorizedKeys, sshKey.(string))
			}
		}

		if authorizedUsers, ok := disk["authorized_users"]; ok {
			for _, user := range authorizedUsers.([]interface{}) {
				diskOpts.AuthorizedUsers = append(diskOpts.AuthorizedUsers, user.(string))
			}
		}

		if stackscriptID, ok := disk["stackscript_id"]; ok {
			diskOpts.StackscriptID = stackscriptID.(int)
		}

		if stackscriptData, ok := disk["stackscript_data"]; ok {
			for name, value := range stackscriptData.(map[string]interface{}) {
				diskOpts.StackscriptData[name] = value.(string)
			}
		}
	}

	instanceDisk, err := client.CreateInstanceDisk(context.Background(), instance.ID, diskOpts)

	if err != nil {
		return nil, fmt.Errorf("Error creating Linode instance %d disk: %s", instance.ID, err)
	}

	_, err = client.WaitForEventFinished(context.Background(), instance.ID, linodego.EntityLinode, linodego.ActionDiskCreate, instanceDisk.Created, int(d.Timeout(schema.TimeoutCreate).Seconds()))
	if err != nil {
		return nil, fmt.Errorf("Error waiting for Linode instance %d disk: %s", instanceDisk.ID, err)
	}

	return instanceDisk, err
}

func updateInstanceDisks(client linodego.Client, d *schema.ResourceData, instance linodego.Instance, tfDisksOld interface{}, tfDisksNew interface{}) (bool, map[string]int, error) {
	var diskIDLabelMap map[string]int
	var rebootInstance bool

	disks, err := client.ListInstanceDisks(context.Background(), int(instance.ID), nil)
	if err != nil {
		return rebootInstance, diskIDLabelMap, fmt.Errorf("Error fetching the disks for Instance %d: %s", instance.ID, err)
	}

	diskMap := make(map[string]linodego.InstanceDisk, len(disks))
	for _, disk := range disks {
		if _, duplicate := diskMap[disk.Label]; duplicate {
			return rebootInstance, diskIDLabelMap, fmt.Errorf("Error indexing Instance %d Disks: Label '%s' is assigned to multiple disks", instance.ID, disk.Label)
		}
		diskMap[disk.Label] = disk
	}

	oldDiskLabels := make([]string, len(tfDisksOld.([]interface{})))

	for _, tfDiskOld := range tfDisksOld.([]interface{}) {
		if oldDisk, ok := tfDiskOld.(map[string]interface{}); ok {
			oldDiskLabels = append(oldDiskLabels, oldDisk["label"].(string))
		}
	}
	tfDisks := tfDisksNew.([]interface{})

	//updatedDisks := make([]*linodego.InstanceDisk, tfDisks.Len())
	diskIDLabelMap = make(map[string]int, len(tfDisks))

	for _, tfDisk := range tfDisks {
		tfd := tfDisk.(map[string]interface{})

		labelStr, found := tfd["label"]
		if !found {
			return rebootInstance, diskIDLabelMap, fmt.Errorf("Error parsing disk label")
		}

		label, ok := labelStr.(string)
		if !ok {
			return rebootInstance, diskIDLabelMap, fmt.Errorf("Error parsing disk label")
		}

		existingDisk, existing := diskMap[label]

		if existing {
			// The only non-destructive change supported is resize, which requires a reboot
			// Label renames are not supported because this TF provider relies on the label as an identifier
			if tfd["size"].(int) != existingDisk.Size {
				if err := changeInstanceDiskSize(&client, instance, existingDisk, tfd["size"].(int), d); err != nil {
					return rebootInstance, diskIDLabelMap, err
				}
				rebootInstance = true
			}
			if strings.Compare(tfd["filesystem"].(string), string(existingDisk.Filesystem)) != 0 {
				return rebootInstance, diskIDLabelMap, fmt.Errorf("Error updating Instance %d Disk %d: Filesystem changes are not supported ('%s' != '%s')", instance.ID, existingDisk.ID, tfd["filesystem"], existingDisk.Filesystem)
			}
			diskIDLabelMap[existingDisk.Label] = existingDisk.ID

		} else {
			instanceDisk, err := createInstanceDisk(client, instance, tfd, d)
			if err != nil {
				return rebootInstance, diskIDLabelMap, err
			}
			rebootInstance = true
			diskIDLabelMap[instanceDisk.Label] = instanceDisk.ID
		}
	}

	for _, oldLabel := range oldDiskLabels {
		if _, found := diskIDLabelMap[oldLabel]; !found {
			if listedDisk, found := diskMap[oldLabel]; found {
				if err := client.DeleteInstanceDisk(context.Background(), instance.ID, listedDisk.ID); err != nil {
					return rebootInstance, diskIDLabelMap, err
				}
				_, err = client.WaitForEventFinished(context.Background(), instance.ID, linodego.EntityLinode, linodego.ActionDiskDelete, *instance.Created, int(d.Timeout(schema.TimeoutUpdate).Seconds()))
				if err != nil {
					return rebootInstance, diskIDLabelMap, fmt.Errorf("Error waiting for Instance %d Disk %d to finish deleting: %s", instance.ID, listedDisk.ID, err)
				}
				delete(diskIDLabelMap, oldLabel)
			}
		}
	}

	return rebootInstance, diskIDLabelMap, nil
}

// getTotalDiskSize returns the number of disks and their total size.
func getTotalDiskSize(client *linodego.Client, linodeID int) (totalDiskSize int, err error) {
	disks, err := client.ListInstanceDisks(context.Background(), linodeID, nil)
	if err != nil {
		return 0, err
	}

	for _, disk := range disks {
		totalDiskSize += disk.Size
	}

	return totalDiskSize, nil
}

// getBiggestDisk returns the ID and Size of the largest disk attached to the Linode
func getBiggestDisk(client *linodego.Client, linodeID int) (biggestDiskID int, biggestDiskSize int, err error) {
	diskFilter := "{\"+order_by\": \"size\", \"+order\": \"desc\"}"
	disks, err := client.ListInstanceDisks(context.Background(), linodeID, linodego.NewListOptions(1, diskFilter))
	if err != nil {
		return 0, 0, err
	}

	for _, disk := range disks {
		// Find Biggest Disk ID & Size
		if disk.Size > biggestDiskSize {
			biggestDiskID = disk.ID
			biggestDiskSize = disk.Size
		}
	}
	return biggestDiskID, biggestDiskSize, nil
}

// sshKeyState hashes a string passed in as an interface
func sshKeyState(val interface{}) string {
	return hashString(strings.Join(val.([]string), "\n"))
}

// rootPasswordState hashes a string passed in as an interface
func rootPasswordState(val interface{}) string {
	return hashString(val.(string))
}

// hashString hashes a string
func hashString(key string) string {
	hash := sha3.Sum512([]byte(key))
	return base64.StdEncoding.EncodeToString(hash[:])
}

func createRandomRootPassword() (string, error) {
	rawRootPass := make([]byte, 50)
	_, err := rand.Read(rawRootPass)
	if err != nil {
		return "", fmt.Errorf("Failed to generate random password")
	}
	rootPass := base64.StdEncoding.EncodeToString(rawRootPass)
	return rootPass, nil
}

// changeInstanceType resizes the Linode Instance
func changeInstanceType(client *linodego.Client, instance *linodego.Instance, targetType string, d *schema.ResourceData) error {

	diskResize := false
	waitForOnline := true

	resizeOpts := linodego.InstanceResizeOptions{
		AllowAutoDiskResize: &diskResize,
		Type:                targetType,
	}

	// Instance must be either offline or running (with no extra activity) to resize.
	if instance.Status == linodego.InstanceOffline || instance.Status == linodego.InstanceShuttingDown {
		waitForOnline = false
		if _, err := client.WaitForInstanceStatus(context.Background(), instance.ID, linodego.InstanceOffline, int(d.Timeout(schema.TimeoutUpdate).Seconds())); err != nil {
			return fmt.Errorf("Error waiting for Instance %d to go offline: %s", instance.ID, err)
		}
	} else {
		if _, err := client.WaitForInstanceStatus(context.Background(), instance.ID, linodego.InstanceRunning, int(d.Timeout(schema.TimeoutUpdate).Seconds())); err != nil {
			return fmt.Errorf("Error waiting for Instance %d readiness: %s", instance.ID, err)
		}
	}
	// We have to wait through the resize process because if we issue jobs before the complete process is complete, the API
	// Will raise an error.

	// Issue the resize job
	if err := client.ResizeInstance(context.Background(), instance.ID, resizeOpts); err != nil {
		return fmt.Errorf("Error resizing Instance %d: %s", instance.ID, err)
	}

	_, err := client.WaitForEventFinished(context.Background(), instance.ID, linodego.EntityLinode, linodego.ActionLinodeResize, *instance.Created, int(d.Timeout(schema.TimeoutUpdate).Seconds()))
	if err != nil {
		return fmt.Errorf("Error waiting for instance %d to finish resizing: %s", instance.ID, err)
	}

	// Wait for instance status to go offline
	if _, err := client.WaitForInstanceStatus(context.Background(), instance.ID, linodego.InstanceOffline, int(d.Timeout(schema.TimeoutUpdate).Seconds())); err != nil {
		return fmt.Errorf("Error waiting for Instance %d to enter offline state: %s", instance.ID, err)
	}

	// Wait for instance status to go online if necessary
	if waitForOnline == true {
		if _, err := client.WaitForInstanceStatus(context.Background(), instance.ID, linodego.InstanceRunning, int(d.Timeout(schema.TimeoutUpdate).Seconds())); err != nil {
			return fmt.Errorf("Error waiting for Instance %d to enter online state: %s", instance.ID, err)
		}
	}
	return nil
}

// returns the amount of disk space used by the new plan and old plan
func getDiskSizeChange(oldDisk interface{}, newDisk interface{}) (int, int) {

	tfDisksOldInterface := oldDisk.([]interface{})
	tfDisksNewInterface := newDisk.([]interface{})

	oldDiskSize := 0
	newDiskSize := 0

	// Get total amount of disk usage before & after
	for _, disk := range tfDisksOldInterface {
		oldDiskSize += disk.(map[string]interface{})["size"].(int)
	}

	for _, disk := range tfDisksNewInterface {
		newDiskSize += disk.(map[string]interface{})["size"].(int)
	}

	return oldDiskSize, newDiskSize
}

func changeInstanceDiskSize(client *linodego.Client, instance linodego.Instance, disk linodego.InstanceDisk, targetSize int, d *schema.ResourceData) error {
	if instance.Specs.Disk > targetSize {
		switch instance.Status {
		case linodego.InstanceShuttingDown:
			if _, err := client.WaitForInstanceStatus(context.Background(), instance.ID, linodego.InstanceOffline, int(d.Timeout(schema.TimeoutUpdate).Seconds())); err != nil {
				return fmt.Errorf("Error waiting for Instance %d to go offline: %s", instance.ID, err)
			}
		case linodego.InstanceOffline:
		default:
			if err := client.ShutdownInstance(context.Background(), instance.ID); err != nil {
				return err
			}
		}

		// Wait for instance to go offline. Resize the disk once Linode is shut down.
		if _, err := client.WaitForInstanceStatus(context.Background(), instance.ID, linodego.InstanceOffline, int(d.Timeout(schema.TimeoutUpdate).Seconds())); err != nil {
			return fmt.Errorf("Error waiting for Instance %d to go offline: %s", instance.ID, err)
		} else {
			if err := client.ResizeInstanceDisk(context.Background(), instance.ID, disk.ID, targetSize); err != nil {
				return fmt.Errorf("Error resizing disk %d for Instance %d: %s", disk.ID, instance.ID, err)
			}
		}

		// Wait for the disk resize operation to complete, and boot instance.
		_, err := client.WaitForEventFinished(context.Background(), instance.ID, linodego.EntityLinode, linodego.ActionDiskResize, disk.Updated, int(d.Timeout(schema.TimeoutUpdate).Seconds()))
		if err != nil {
			return fmt.Errorf("Error waiting for resize of Instance %d Disk %d: %s", instance.ID, disk.ID, err)
		}
	} else {
		return fmt.Errorf("Error resizing disk %d: size exceeds disk size for Instance %d", disk.ID, instance.ID)
	}
	return nil
}

// privateIP determines if an IP is for private use (RFC1918)
// https://stackoverflow.com/a/41273687
func privateIP(ip net.IP) bool {
	_, private24BitBlock, _ := net.ParseCIDR("10.0.0.0/8")
	_, private20BitBlock, _ := net.ParseCIDR("172.16.0.0/12")
	_, private16BitBlock, _ := net.ParseCIDR("192.168.0.0/16")
	private := private24BitBlock.Contains(ip) || private20BitBlock.Contains(ip) || private16BitBlock.Contains(ip)
	return private
}

func diskHashCode(v interface{}) int {
	switch t := v.(type) {
	case linodego.InstanceDisk:
		return schema.HashString(t.Label + ":" + strconv.Itoa(t.Size))
	case map[string]interface{}:
		if _, found := t["size"]; found {
			if size, ok := t["size"].(int); ok {
				if _, found := t["label"]; found {
					if label, ok := t["label"].(string); ok {
						return schema.HashString(label + ":" + strconv.Itoa(size))
					}
				}
			}
		}
		panic(fmt.Sprintf("Error hashing disk for invalid map: %#v", v))
	default:
		panic(fmt.Sprintf("Error hashing config for unknown interface: %#v", v))
	}
}

func labelHashcode(v interface{}) int {
	switch t := v.(type) {
	case string:
		return schema.HashString(v)
	case linodego.InstanceDisk:
		return schema.HashString(t.Label)
	case linodego.InstanceConfig:
		return schema.HashString(t.Label)
	case map[string]interface{}:
		if _, found := t["label"]; found {
			if label, ok := t["label"].(string); ok {
				return schema.HashString(label)
			}
		}
		panic(fmt.Sprintf("Error hashing label for unknown map: %#v", v))
	default:
		panic(fmt.Sprintf("Error hashing label for unknown interface: %#v", v))
	}
}

func configHashcode(v interface{}) int {
	switch t := v.(type) {
	case string:
		return schema.HashString(v)
	case linodego.InstanceConfig:
		return schema.HashString(t.Label)
	case map[string]interface{}:
		if _, found := t["label"]; found {
			if label, ok := t["label"].(string); ok {
				return schema.HashString(label)
			}
		}
		panic(fmt.Sprintf("Error hashing config for unknown map: %#v", v))
	default:
		panic(fmt.Sprintf("Error hashing config for unknown interface: %#v", v))
	}
}

func diskState(v interface{}) string {
	switch t := v.(type) {
	case map[string]interface{}:
		if _, found := t["size"]; found {
			if size, ok := t["size"].(int); ok {
				if _, found := t["label"]; found {
					if label, ok := t["label"].(string); ok {
						return label + ":" + strconv.Itoa(size)
					}
				}
			}
		}
		panic(fmt.Sprintf("Error generating disk state for invalid map: %#v", v))
	default:
		panic(fmt.Sprintf("Error generating disk for unknown interface: %#v", v))
	}
}

func assignConfigDevice(device *linodego.InstanceConfigDevice, dev map[string]interface{}, diskIDLabelMap map[string]int) error {
	if label, ok := dev["disk_label"].(string); ok && len(label) > 0 {
		if dev["disk_id"], ok = diskIDLabelMap[label]; !ok {
			return fmt.Errorf("Error mapping disk label %s to ID", dev["disk_label"])
		}
	}
	expanded := expandInstanceConfigDevice(dev)
	if expanded != nil {
		*device = *expanded
	}
	return nil
}

// detachConfigVolumes detaches any volumes associated with an InstanceConfig.Devices struct
func detachConfigVolumes(dmap linodego.InstanceConfigDeviceMap, detacher volumeDetacher) error {
	// Preallocate our slice of config devices
	drives := []*linodego.InstanceConfigDevice{
		dmap.SDA, dmap.SDB, dmap.SDC, dmap.SDD, dmap.SDE, dmap.SDF, dmap.SDG, dmap.SDH,
	}

	// Make a buffered error channel for our goroutines to send error values back on
	errCh := make(chan error, len(drives))

	// Make a sync.WaitGroup so our devices can signal they're finished
	var wg sync.WaitGroup
	wg.Add(len(drives))

	// For each drive, spawn a goroutine to detach the volume, send an error on the err channel
	// if one exists, and signal the worker process is done
	for _, d := range drives {
		go func(dev *linodego.InstanceConfigDevice) {
			defer wg.Done()

			if dev != nil && dev.VolumeID > 0 {
				err := detacher(context.Background(), dev.VolumeID, "for config attachment")
				if err != nil {
					errCh <- err
				}
			}
		}(d)
	}

	// Wait until all processes are finished and close the error channel so we can range over it
	wg.Wait()
	close(errCh)

	// Build the error from the errors in the channel and return the combined error if any exist
	var errStr string
	for err := range errCh {
		if len(errStr) == 0 {
			errStr += ", "
		}

		errStr += err.Error()
	}

	if len(errStr) > 0 {
		return fmt.Errorf("Error detaching volumes: %s", errStr)
	}

	return nil
}
