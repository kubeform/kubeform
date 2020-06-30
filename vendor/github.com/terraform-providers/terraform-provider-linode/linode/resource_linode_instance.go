package linode

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/linode/linodego"
)

const (
	LinodeInstanceCreateTimeout = 10 * time.Minute
	LinodeInstanceUpdateTimeout = 20 * time.Minute
	LinodeInstanceDeleteTimeout = 10 * time.Minute
)

var linodeInstanceDownsizeFailedMessage = `
Did you try to resize a linode with implicit, default disks to a smaller type? The provider does
not automatically scale the boot disk to fit an updated instance type. You may need to switch to
an explicit disk configuration.

Take a look at the example here:
https://www.terraform.io/docs/providers/linode/r/instance.html#linode-instance-with-explicit-configs-and-disks`

func resourceLinodeInstanceDeviceDisk() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disk_label": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The `label` of the `disk` to map to this `device` slot.",
			},
			"disk_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The Disk ID to map to this disk slot",
			},
			"volume_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The Block Storage volume ID to map to this disk slot",
			},
		},
	}
}

func resourceLinodeInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceLinodeInstanceCreate,
		Read:   resourceLinodeInstanceRead,
		Update: resourceLinodeInstanceUpdate,
		Delete: resourceLinodeInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(LinodeInstanceCreateTimeout),
			Update: schema.DefaultTimeout(LinodeInstanceUpdateTimeout),
			Delete: schema.DefaultTimeout(LinodeInstanceDeleteTimeout),
		},
		Schema: map[string]*schema.Schema{
			"image": {
				Type:          schema.TypeString,
				Description:   "An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with private/. See /images for more information on the Images available for you to use.",
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"disk", "config", "backup_id"},
			},
			"backup_id": {
				Type:          schema.TypeInt,
				Description:   "A Backup ID from another Linode's available backups. Your User must have read_write access to that Linode, the Backup must have a status of successful, and the Linode must be deployed to the same region as the Backup. See /linode/instances/{linodeId}/backups for a Linode's available backups. This field and the image field are mutually exclusive.",
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"image", "disk", "config"},
			},
			"stackscript_id": {
				Type:          schema.TypeInt,
				Description:   "The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an Image that is compatible with this StackScript.",
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"disk", "config"},
			},
			"stackscript_data": {
				Type:          schema.TypeMap,
				Description:   "An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode. Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.",
				Optional:      true,
				ForceNew:      true,
				Sensitive:     true,
				ConflictsWith: []string{"disk", "config"},
			},
			"label": {
				Type:         schema.TypeString,
				Description:  "The Linode's label is for display purposes only. If no label is provided for a Linode, a default will be assigned",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(3, 50),
			},
			"group": {
				Type:        schema.TypeString,
				Description: "The display group of the Linode instance.",
				Optional:    true,
			},
			"tags": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "An array of tags applied to this object. Tags are for organizational purposes only.",
			},
			"boot_config_label": {
				Type:        schema.TypeString,
				Description: "The Label of the Instance Config that should be used to boot the Linode instance.",
				Optional:    true,
				Computed:    true,
			},
			"region": {
				Type:         schema.TypeString,
				Description:  "This is the location where the Linode was deployed. This cannot be changed without opening a support ticket.",
				Required:     true,
				ForceNew:     true,
				InputDefault: "us-east",
			},
			"type": {
				Type:        schema.TypeString,
				Description: "The type of instance to be deployed, determining the price and size.",
				Optional:    true,
				Default:     "g6-standard-1",
			},
			"status": {
				Type:        schema.TypeString,
				Description: "The status of the instance, indicating the current readiness state.",
				Computed:    true,
			},
			"ip_address": {
				Type:        schema.TypeString,
				Description: "This Linode's Public IPv4 Address. If there are multiple public IPv4 addresses on this Instance, an arbitrary address will be used for this field.",
				Computed:    true,
			},
			"ipv6": {
				Type:        schema.TypeString,
				Description: "This Linode's IPv6 SLAAC addresses. This address is specific to a Linode, and may not be shared.",
				Computed:    true,
			},

			"ipv4": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "This Linode's IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a single private IPv4 address if needed. You may need to open a support ticket to get additional IPv4 addresses.",
				Computed:    true,
			},

			"private_ip": {
				Type:        schema.TypeBool,
				Description: "If true, the created Linode will have private networking enabled, allowing use of the 192.168.128.0/17 network within the Linode's region.",
				Optional:    true,
			},
			"private_ip_address": {
				Type:        schema.TypeString,
				Description: "This Linode's Private IPv4 Address.  The regional private IP address range is 192.168.128/17 address shared by all Linode Instances in a region.",
				Computed:    true,
			},
			"authorized_keys": {
				Type:          schema.TypeList,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Description:   "A list of SSH public keys to deploy for the root user on the newly created Linode. Only accepted if 'image' is provided.",
				Optional:      true,
				ForceNew:      true,
				StateFunc:     sshKeyState,
				ConflictsWith: []string{"disk", "config"},
			},
			"authorized_users": {
				Type:          schema.TypeList,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Description:   "A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the `root` user's `~/.ssh/authorized_keys` file automatically. Only accepted if 'image' is provided.",
				Optional:      true,
				ForceNew:      true,
				StateFunc:     sshKeyState,
				ConflictsWith: []string{"disk", "config"},
			},
			"root_pass": {
				Type:          schema.TypeString,
				Description:   "The password that will be initialially assigned to the 'root' user account.",
				Sensitive:     true,
				Optional:      true,
				ForceNew:      true,
				StateFunc:     rootPasswordState,
				ConflictsWith: []string{"disk", "config"},
			},
			"swap_size": {
				Type:          schema.TypeInt,
				Description:   "When deploying from an Image, this field is optional with a Linode API default of 512mb, otherwise it is ignored. This is used to set the swap disk size for the newly-created Linode.",
				Optional:      true,
				Computed:      true,
				Default:       nil,
				ConflictsWith: []string{"disk", "config"},
			},
			"backups_enabled": {
				Type:        schema.TypeBool,
				Description: "If this field is set to true, the created Linode will automatically be enrolled in the Linode Backup service. This will incur an additional charge. The cost for the Backup service is dependent on the Type of Linode deployed.",
				Optional:    true,
				Computed:    true,
				Default:     nil,
			},
			"watchdog_enabled": {
				Type:        schema.TypeBool,
				Description: "The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie will give up if there have been more than 5 boot jobs issued within 15 minutes.",
				Optional:    true,
				Default:     true,
			},
			"specs": {
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disk": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The amount of storage space, in GB. this Linode has access to. A typical Linode will divide this space between a primary disk with an image deployed to it, and a swap disk, usually 512 MB. This is the default configuration created when deploying a Linode with an image without specifying disks.",
						},
						"memory": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The amount of RAM, in MB, this Linode has access to. Typically a Linode will choose to boot with all of its available RAM, but this can be configured in a Config profile.",
						},
						"vcpus": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of vcpus this Linode has access to. Typically a Linode will choose to boot with all of its available vcpus, but this can be configured in a Config Profile.",
						},
						"transfer": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The amount of network transfer this Linode is allotted each month.",
						},
					},
				},
			},

			"alerts": {
				Computed: true,
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The percentage of CPU usage required to trigger an alert. If the average CPU usage over two hours exceeds this value, we'll send you an alert. If this is set to 0, the alert is disabled.",
						},
						"network_in": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The amount of incoming traffic, in Mbit/s, required to trigger an alert. If the average incoming traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.",
						},
						"network_out": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The amount of outbound traffic, in Mbit/s, required to trigger an alert. If the average outbound traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.",
						},
						"transfer_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The percentage of network transfer that may be used before an alert is triggered. When this value is exceeded, we'll alert you. If this is set to 0 (zero), the alert is disabled.",
						},
						"io": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The amount of disk IO operation per second required to trigger an alert. If the average disk IO over two hours exceeds this value, we'll send you an alert. If set to 0, this alert is disabled.",
						},
					},
				},
			},
			"backups": {
				Type:        schema.TypeList,
				Description: "Information about this Linode's backups status.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If this Linode has the Backup service enabled.",
						},
						"schedule": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								// TODO(displague) these fields are updatable via PUT to instance
								Schema: map[string]*schema.Schema{
									"day": {
										Type:        schema.TypeString,
										Description: "The day ('Sunday'-'Saturday') of the week that your Linode's weekly Backup is taken. If not set manually, a day will be chosen for you. Backups are taken every day, but backups taken on this day are preferred when selecting backups to retain for a longer period.  If not set manually, then when backups are initially enabled, this may come back as 'Scheduling' until the day is automatically selected.",
										Computed:    true,
									},
									"window": {
										Type:        schema.TypeString,
										Description: "The window ('W0'-'W22') in which your backups will be taken, in UTC. A backups window is a two-hour span of time in which the backup may occur. For example, 'W10' indicates that your backups should be taken between 10:00 and 12:00. If you do not choose a backup window, one will be selected for you automatically.  If not set manually, when backups are initially enabled this may come back as Scheduling until the window is automatically selected.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"config": {
				Optional:      true,
				Description:   "Configuration profiles define the VM settings and boot behavior of the Linode Instance.",
				Type:          schema.TypeList,
				ConflictsWith: []string{"image", "root_pass", "authorized_keys", "authorized_users", "swap_size", "backup_id", "stackscript_id"},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					_, hasImage := d.GetOk("image")
					return hasImage
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": {
							Type:         schema.TypeString,
							Description:  "The Config's label for display purposes.  Also used by `boot_config_label`.",
							Required:     true,
							ValidateFunc: validation.StringLenBetween(1, 48),
						},
						"helpers": {
							Type:        schema.TypeList,
							Description: "Helpers enabled when booting to this Linode Config.",
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"updatedb_disabled": {
										Type:        schema.TypeBool,
										Description: "Disables updatedb cron job to avoid disk thrashing.",
										Optional:    true,
										Default:     true,
									},
									"distro": {
										Type:        schema.TypeBool,
										Description: "Controls the behavior of the Linode Config's Distribution Helper setting.",
										Optional:    true,
										Default:     true,
									},
									"modules_dep": {
										Type:        schema.TypeBool,
										Description: "Creates a modules dependency file for the Kernel you run.",
										Optional:    true,
										Default:     true,
									},
									"network": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Controls the behavior of the Linode Config's Network Helper setting, used to automatically configure additional IP addresses assigned to this instance.",
										Default:     true,
									},
									"devtmpfs_automount": {
										Type:        schema.TypeBool,
										Description: "Populates the /dev directory early during boot without udev. Defaults to false.",
										Optional:    true,
										Default:     false,
									},
								},
							},
						},
						"devices": {
							Type:        schema.TypeList,
							Description: "Device sda-sdh can be either a Disk or Volume identified by disk_label or volume_id. Only one type per slot allowed.",
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sda": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Computed: true,
										Optional: true,
										Elem:     resourceLinodeInstanceDeviceDisk(),
									},
									"sdb": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Computed: true,
										Elem:     resourceLinodeInstanceDeviceDisk(),
									},
									"sdc": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Computed: true,
										Elem:     resourceLinodeInstanceDeviceDisk(),
									},
									"sdd": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Computed: true,
										Elem:     resourceLinodeInstanceDeviceDisk(),
									},
									"sde": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Computed: true,
										Elem:     resourceLinodeInstanceDeviceDisk(),
									},
									"sdf": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Computed: true,
										Elem:     resourceLinodeInstanceDeviceDisk(),
									},
									"sdg": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Computed: true,
										Elem:     resourceLinodeInstanceDeviceDisk(),
									},
									"sdh": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Computed: true,
										Elem:     resourceLinodeInstanceDeviceDisk(),
									},
								},
							},
						},
						"kernel": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A Kernel ID to boot a Linode with. Default is based on image choice. (examples: linode/latest-64bit, linode/grub2, linode/direct-disk)",
						},
						"run_level": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "Defines the state of your Linode after booting. Defaults to default.",
							Default:      "default",
							ValidateFunc: validation.StringInSlice([]string{"default", "single", "binbash"}, false),
						},
						"virt_mode": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "Controls the virtualization mode. Defaults to paravirt.",
							Default:      "paravirt",
							ValidateFunc: validation.StringInSlice([]string{"paravirt", "fullvirt"}, false),
						},
						"root_device": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The root device to boot. The corresponding disk must be attached.",
						},
						"comments": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Optional field for arbitrary User comments on this Config.",
						},

						"memory_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Defaults to the total RAM of the Linode",
						},
					},
				},
			},
			"disk": {
				Optional:      true,
				ConflictsWith: []string{"image", "root_pass", "authorized_keys", "authorized_users", "swap_size", "backup_id", "stackscript_id"},
				Type:          schema.TypeList,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					_, hasImage := d.GetOk("image")
					return hasImage
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": {
							Type:         schema.TypeString,
							Description:  "The disks label, which acts as an identifier in Terraform.",
							Required:     true,
							ValidateFunc: validation.StringLenBetween(1, 48),
						},
						"size": {
							Type:        schema.TypeInt,
							Description: "The size of the Disk in MB.",
							Required:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "The ID of the Disk (for use in Linode Image resources and Linode Instance Config Devices)",
							Computed:    true,
						},
						"filesystem": {
							Type:         schema.TypeString,
							Description:  "The Disk filesystem can be one of: raw, swap, ext3, ext4, initrd (max 32mb)",
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
							ValidateFunc: validation.StringInSlice([]string{"raw", "swap", "ext3", "ext4", "initrd"}, false),
						},
						"read_only": {
							Type:        schema.TypeBool,
							Description: "If true, this Disk is read-only.",
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"image": {
							Type:        schema.TypeString,
							Description: "An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with private/.",
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								// the API does not return this field for existing disks, so must be ignored for diffs/updates
								return !d.HasChange("label")
							},
						},
						"authorized_keys": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "A list of SSH public keys to deploy for the root user on the newly created Linode. Only accepted if 'image' is provided.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								// the API does not return this field for existing disks, so must be ignored for diffs/updates
								return !d.HasChange("label")
							},
							Optional:  true,
							ForceNew:  true,
							StateFunc: sshKeyState,
						},
						"authorized_users": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the `root` user's `~/.ssh/authorized_keys` file automatically. Only accepted if 'image' is provided.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								// the API does not return this field for existing disks, so must be ignored for diffs/updates
								return !d.HasChange("label")
							},
							Optional:  true,
							ForceNew:  true,
							StateFunc: sshKeyState,
						},
						"stackscript_id": {
							Type:        schema.TypeInt,
							Description: "The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an Image that is compatible with this StackScript.",
							Computed:    true,
							Optional:    true,
							ForceNew:    true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								// the API does not return this field for existing disks, so must be ignored for diffs/updates
								return !d.HasChange("label")
							},
							Default: nil,
						},
						"stackscript_data": {
							Type:        schema.TypeMap,
							Description: "An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode. Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.",
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Sensitive:   true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								// the API does not return this field for existing disks, so must be ignored for diffs/updates
								return !d.HasChange("label")
							},
							Default: nil,
						},
						"root_pass": {
							Type:        schema.TypeString,
							Description: "The password that will be initialially assigned to the 'root' user account.",
							Sensitive:   true,
							Optional:    true,
							ForceNew:    true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								// the API does not return this field for existing disks, so must be ignored for diffs/updates
								return !d.HasChange("label")
							},
							ValidateFunc: validation.StringLenBetween(6, 128),
							StateFunc:    rootPasswordState,
						},
					},
				},
			},
		},
	}
}

func resourceLinodeInstanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode instance ID %s as int: %s", d.Id(), err)
	}

	instance, err := client.GetInstance(context.Background(), int(id))

	if err != nil {
		if lerr, ok := err.(*linodego.Error); ok && lerr.Code == 404 {
			log.Printf("[WARN] removing Linode Instance ID %q from state because it no longer exists", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error finding the specified Linode instance: %s", err)
	}

	instanceNetwork, err := client.GetInstanceIPAddresses(context.Background(), int(id))

	if err != nil {
		return fmt.Errorf("Error getting the IPs for Linode instance %s: %s", d.Id(), err)
	}

	var ips []string
	for _, ip := range instance.IPv4 {
		ips = append(ips, ip.String())
	}
	d.Set("ipv4", ips)
	d.Set("ipv6", instance.IPv6)
	public, private := instanceNetwork.IPv4.Public, instanceNetwork.IPv4.Private

	if len(public) > 0 {
		d.Set("ip_address", public[0].Address)

		d.SetConnInfo(map[string]string{
			"type": "ssh",
			"host": public[0].Address,
		})
		// TODO(displague) to determine 'user', need to check disk.image
		// "linode/containerlinux" is "core", else "root"
		// might be better to make this a resource field and avoid lookups
	}

	if len(private) > 0 {
		d.Set("private_ip", true)
		d.Set("private_ip_address", private[0].Address)
	} else {
		d.Set("private_ip", false)
	}

	d.Set("label", instance.Label)
	d.Set("status", instance.Status)
	d.Set("type", instance.Type)
	d.Set("region", instance.Region)
	d.Set("watchdog_enabled", instance.WatchdogEnabled)
	d.Set("group", instance.Group)
	d.Set("tags", instance.Tags)

	flatSpecs := flattenInstanceSpecs(*instance)
	flatAlerts := flattenInstanceAlerts(*instance)
	flatBackups := flattenInstanceBackups(*instance)

	d.Set("backups", flatBackups)
	d.Set("specs", flatSpecs)
	d.Set("alerts", flatAlerts)

	instanceDisks, err := client.ListInstanceDisks(context.Background(), int(id), nil)
	if err != nil {
		return fmt.Errorf("Error getting the disks for the Linode instance %d: %s", id, err)
	}

	disks, swapSize := flattenInstanceDisks(instanceDisks)
	d.Set("disk", disks)
	d.Set("swap_size", swapSize)

	instanceConfigs, err := client.ListInstanceConfigs(context.Background(), int(id), nil)

	if err != nil {
		return fmt.Errorf("Error getting the config for Linode instance %d (%s): %s", instance.ID, instance.Label, err)
	}
	diskLabelIDMap := make(map[int]string, len(instanceDisks))
	for _, disk := range instanceDisks {
		diskLabelIDMap[disk.ID] = disk.Label

	}

	configs := flattenInstanceConfigs(instanceConfigs, diskLabelIDMap)

	d.Set("config", configs)
	if len(instanceConfigs) == 1 {
		d.Set("boot_config_label", instanceConfigs[0].Label)
	}

	return nil
}

func resourceLinodeInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	client, ok := meta.(linodego.Client)
	if !ok {
		return fmt.Errorf("Invalid Client when creating Linode Instance")
	}

	bootConfig := 0
	createOpts := linodego.InstanceCreateOptions{
		Region:         d.Get("region").(string),
		Type:           d.Get("type").(string),
		Label:          d.Get("label").(string),
		Group:          d.Get("group").(string),
		BackupsEnabled: d.Get("backups_enabled").(bool),
		PrivateIP:      d.Get("private_ip").(bool),
	}

	if tagsRaw, tagsOk := d.GetOk("tags"); tagsOk {
		for _, tag := range tagsRaw.(*schema.Set).List() {
			createOpts.Tags = append(createOpts.Tags, tag.(string))
		}
	}

	_, disksOk := d.GetOk("disk")
	_, configsOk := d.GetOk("config")

	// If we don't have disks and we don't have configs, use the single API call approach
	if !disksOk && !configsOk {
		for _, key := range d.Get("authorized_keys").([]interface{}) {
			createOpts.AuthorizedKeys = append(createOpts.AuthorizedKeys, key.(string))
		}
		for _, key := range d.Get("authorized_users").([]interface{}) {
			createOpts.AuthorizedUsers = append(createOpts.AuthorizedUsers, key.(string))
		}
		createOpts.RootPass = d.Get("root_pass").(string)
		if createOpts.RootPass == "" {
			var err error
			createOpts.RootPass, err = createRandomRootPassword()
			if err != nil {
				return err
			}
		}
		createOpts.Image = d.Get("image").(string)
		createOpts.Booted = &boolTrue
		createOpts.BackupID = d.Get("backup_id").(int)
		if swapSize := d.Get("swap_size").(int); swapSize > 0 {
			createOpts.SwapSize = &swapSize
		}

		createOpts.StackScriptID = d.Get("stackscript_id").(int)

		if stackscriptDataRaw, ok := d.GetOk("stackscript_data"); ok {
			stackscriptData, ok := stackscriptDataRaw.(map[string]interface{})
			if !ok {
				return fmt.Errorf("Error parsing stackscript_data: expected map[string]interface{}")
			}
			createOpts.StackScriptData = make(map[string]string, len(stackscriptData))
			for name, value := range stackscriptData {
				createOpts.StackScriptData[name] = value.(string)
			}
		}
	} else {
		createOpts.Booted = &boolFalse // necessary to prepare disks and configs
	}

	instance, err := client.CreateInstance(context.Background(), createOpts)
	if err != nil {
		return fmt.Errorf("Error creating a Linode Instance: %s", err)
	}

	d.SetId(fmt.Sprintf("%d", instance.ID))

	var ips []string
	for _, ip := range instance.IPv4 {
		ips = append(ips, ip.String())
	}

	d.Set("ipv4", ips)
	d.Set("ipv6", instance.IPv6)

	for _, address := range instance.IPv4 {
		if private := privateIP(*address); private {
			d.Set("private_ip_address", address.String())
		} else {
			d.Set("ip_address", address.String())
		}
	}

	updateOpts := linodego.InstanceUpdateOptions{}
	doUpdate := false

	if _, watchdogEnabledOk := d.GetOk("watchdog_enabled"); watchdogEnabledOk {
		doUpdate = true
		watchdogEnabled := d.Get("watchdog_enabled").(bool)
		updateOpts.WatchdogEnabled = &watchdogEnabled
	}

	if _, alertsOk := d.GetOk("alerts.0"); alertsOk {
		doUpdate = true
		updateOpts.Alerts = &linodego.InstanceAlert{}

		// TODO(displague) only set specified alerts
		updateOpts.Alerts.CPU = d.Get("alerts.0.cpu").(int)
		updateOpts.Alerts.IO = d.Get("alerts.0.io").(int)
		updateOpts.Alerts.NetworkIn = d.Get("alerts.0.network_in").(int)
		updateOpts.Alerts.NetworkOut = d.Get("alerts.0.network_out").(int)
		updateOpts.Alerts.TransferQuota = d.Get("alerts.0.transfer_quota").(int)
	}

	if doUpdate {
		instance, err = client.UpdateInstance(context.Background(), instance.ID, updateOpts)
		if err != nil {
			return err
		}
	}

	// Look up tables for any disks and configs we create
	// - so configs and initrd can reference disks by label
	// - so configs can be referenced as a boot_config_label param
	var diskIDLabelMap map[string]int
	var configIDLabelMap map[string]int

	if disksOk {
		_, err = client.WaitForEventFinished(context.Background(), instance.ID, linodego.EntityLinode, linodego.ActionLinodeCreate, *instance.Created, int(d.Timeout(schema.TimeoutCreate).Seconds()))
		if err != nil {
			return fmt.Errorf("Error waiting for Instance to finish creating: %s", err)
		}

		diskSpecs := d.Get("disk").([]interface{})
		diskIDLabelMap = make(map[string]int, len(diskSpecs))

		for _, diskSpec := range diskSpecs {
			diskSpec := diskSpec.(map[string]interface{})

			instanceDisk, err := createInstanceDisk(client, *instance, diskSpec, d)
			if err != nil {
				return err
			}
			diskIDLabelMap[instanceDisk.Label] = instanceDisk.ID
		}
	}

	if configsOk {
		configSpecs := d.Get("config").([]interface{})
		detacher := makeVolumeDetacher(client, d)

		configIDMap, err := createInstanceConfigsFromSet(client, instance.ID, configSpecs, diskIDLabelMap, detacher)
		if err != nil {
			return err
		}

		configIDLabelMap = make(map[string]int, len(configIDMap))
		for k, v := range configIDMap {
			if len(configIDLabelMap) == 1 {
				bootConfig = k
			}

			configIDLabelMap[v.Label] = k
		}
	}

	if createOpts.Booted == nil || !*createOpts.Booted {
		if disksOk && configsOk {
			if err = client.BootInstance(context.Background(), instance.ID, bootConfig); err != nil {
				return fmt.Errorf("Error booting Linode instance %d: %s", instance.ID, err)
			}

			if _, err = client.WaitForEventFinished(context.Background(), instance.ID, linodego.EntityLinode, linodego.ActionLinodeBoot, *instance.Created, int(d.Timeout(schema.TimeoutCreate).Seconds())); err != nil {
				return fmt.Errorf("Error booting Linode instance %d: %s", instance.ID, err)
			}

			if _, err = client.WaitForInstanceStatus(context.Background(), instance.ID, linodego.InstanceRunning, int(d.Timeout(schema.TimeoutCreate).Seconds())); err != nil {
				return fmt.Errorf("Timed-out waiting for Linode instance %d to boot: %s", instance.ID, err)
			}
		} else {
			if _, err = client.WaitForInstanceStatus(context.Background(), instance.ID, linodego.InstanceOffline, int(d.Timeout(schema.TimeoutCreate).Seconds())); err != nil {
				return fmt.Errorf("Timed-out waiting for Linode instance %d to be created: %s", instance.ID, err)
			}
		}
	} else {
		if _, err = client.WaitForInstanceStatus(context.Background(), instance.ID, linodego.InstanceRunning, int(d.Timeout(schema.TimeoutCreate).Seconds())); err != nil {
			return fmt.Errorf("Timed-out waiting for Linode instance %d to boot: %s", instance.ID, err)
		}
	}

	return resourceLinodeInstanceRead(d, meta)
}

func findDiskByFS(disks []linodego.InstanceDisk, fs linodego.DiskFilesystem) *linodego.InstanceDisk {
	for _, disk := range disks {
		if disk.Filesystem == fs {
			return &disk
		}
	}
	return nil
}

// adjustSwapSizeIfNeeded handles changes to the swap_size attribute if needed. If there is a change, this means resizing
// the underlying main/swap disks on the instance to match the declared swap size allocation.
//
// returns bool describing whether the linode needs to be restarted
func adjustSwapSizeIfNeeded(d *schema.ResourceData, client *linodego.Client, instance *linodego.Instance) (bool, error) {
	if !d.HasChange("swap_size") {
		return false, nil
	}

	// If the swap_size attribute is set, there are two default disks attached to the instance (the main disk of type ext4
	// and a swap disk), as custom disk configuration via "disk" nested attributes conflicts with the swap_size.
	bootDisk, swapDisk, err := getInstanceDefaultDisks(instance.ID, client)
	if err != nil {
		return false, err
	}

	oldSwapVal, newSwapVal := d.GetChange("swap_size")
	oldSwap, newSwap := oldSwapVal.(int), newSwapVal.(int)
	diff := newSwap - oldSwap
	newBootDiskSize := bootDisk.Size - diff

	toResize := []struct {
		size int
		disk *linodego.InstanceDisk
	}{
		{
			size: newBootDiskSize,
			disk: bootDisk,
		},
		{
			size: newSwap,
			disk: swapDisk,
		},
	}

	if bootDisk.Size < newBootDiskSize {
		// swap disk needs to be downsized first to upsize main disk
		toResize[0], toResize[1] = toResize[1], toResize[0]
	}

	for _, resizeOp := range toResize {
		if err := changeInstanceDiskSize(client, *instance, *resizeOp.disk, resizeOp.size, d); err != nil {
			return true, err
		}
	}
	return true, nil
}

func resourceLinodeInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode Instance ID %s as int: %s", d.Id(), err)
	}

	instance, err := client.GetInstance(context.Background(), int(id))
	if err != nil {
		return fmt.Errorf("Error fetching data about the current linode: %s", err)
	}

	updateOpts := linodego.InstanceUpdateOptions{}
	simpleUpdate := false

	if d.HasChange("label") {
		updateOpts.Label = d.Get("label").(string)
		simpleUpdate = true
	}
	if d.HasChange("group") {
		updateOpts.Group = d.Get("group").(string)
		simpleUpdate = true
	}
	if d.HasChange("tags") {
		var tags []string
		for _, tag := range d.Get("tags").(*schema.Set).List() {
			tags = append(tags, tag.(string))
		}
		updateOpts.Tags = &tags
		simpleUpdate = true
	}
	if d.HasChange("watchdog_enabled") {
		watchdogEnabled := d.Get("watchdog_enabled").(bool)
		updateOpts.WatchdogEnabled = &watchdogEnabled
		simpleUpdate = true
	}
	if d.HasChange("alerts") {
		updateOpts.Alerts = &linodego.InstanceAlert{}
		updateOpts.Alerts.CPU = d.Get("alerts.0.cpu").(int)
		updateOpts.Alerts.IO = d.Get("alerts.0.io").(int)
		updateOpts.Alerts.NetworkIn = d.Get("alerts.0.network_in").(int)
		updateOpts.Alerts.NetworkOut = d.Get("alerts.0.network_out").(int)
		updateOpts.Alerts.TransferQuota = d.Get("alerts.0.transfer_quota").(int)
		simpleUpdate = true
	}

	// apply staged simple updates early
	if simpleUpdate {
		if instance, err = client.UpdateInstance(context.Background(), instance.ID, updateOpts); err != nil {
			return fmt.Errorf("Error updating Instance %d: %s", instance.ID, err)
		}
	}

	if d.HasChange("backups_enabled") {
		if d.Get("backups_enabled").(bool) {
			if err = client.EnableInstanceBackups(context.Background(), instance.ID); err != nil {
				return err
			}
		} else {
			if err = client.CancelInstanceBackups(context.Background(), instance.ID); err != nil {
				return err
			}
		}
	}

	rebootInstance := false

	if d.HasChange("private_ip") {
		if _, ok := d.GetOk("private_ip"); !ok {
			return fmt.Errorf("Error removing private IP address for Instance %d: Removing a Private IP address must be handled through a support ticket", instance.ID)
		}

		privateIP, err := client.AddInstanceIPAddress(context.Background(), instance.ID, false)
		if err != nil {
			return fmt.Errorf("Error activating private networking on Instance %d: %s", instance.ID, err)
		}
		d.Set("private_ip_address", privateIP.Address)
		rebootInstance = true
	}

	oldSpec, newSpec, err := getInstanceTypeChange(d, &client)
	if err != nil {
		return fmt.Errorf("Error getting resize info for instance: %s", err)
	}
	upsized := newSpec.Disk > oldSpec.Disk

	if upsized {
		// The linode was upsized; apply before disk changes to allocate more disk
		if instance, err = applyInstanceTypeChange(d, &client, instance, newSpec); err != nil {
			return fmt.Errorf("failed to change instance type: %s", err)
		}
		rebootInstance = true
	}

	if didChange, err := applyInstanceDiskSpec(d, &client, instance, newSpec); err == nil && didChange {
		rebootInstance = true
	} else if err != nil && newSpec.Disk < oldSpec.Disk && !d.HasChange("disk") {
		// Linode was downsized but the pre-existing disk config does not fit new instance spec
		// This might mean the user tried to downsize an instance with an implicit, default
		return fmt.Errorf("failed to apply instance disk spec: %s."+linodeInstanceDownsizeFailedMessage, err)
	} else if err != nil {
		return fmt.Errorf("failed to apply instance disk spec: %s", err)
	}

	if oldSpec.ID != newSpec.ID && !upsized {
		// linode was downsized or changed to a type with the same disk allocation
		if instance, err = applyInstanceTypeChange(d, &client, instance, newSpec); err != nil {
			return fmt.Errorf("failed to change instance type: %s", err)
		}
	}

	if didChange, err := adjustSwapSizeIfNeeded(d, &client, instance); err != nil {
		return err
	} else if didChange {
		rebootInstance = true
	}

	diskIDLabelMap, err := getInstanceDiskLabelIDMap(client, d, instance.ID)
	if err != nil {
		return fmt.Errorf("failed to get disk label to ID mappings")
	}

	tfConfigsOld, tfConfigsNew := d.GetChange("config")
	didChangeConfig, updatedConfigMap, updatedConfigs, err := updateInstanceConfigs(client, d, *instance, tfConfigsOld, tfConfigsNew, diskIDLabelMap)
	if err != nil {
		return err
	}
	rebootInstance = rebootInstance || didChangeConfig

	bootConfig := 0
	bootConfigLabel := d.Get("boot_config_label").(string)
	if bootConfigLabel != "" {
		if foundConfig, found := updatedConfigMap[bootConfigLabel]; found {
			bootConfig = foundConfig
		} else {
			return fmt.Errorf("Error setting boot_config_label: Config label '%s' not found", bootConfigLabel)
		}
	} else if len(updatedConfigs) > 0 {
		bootConfig = updatedConfigs[0].ID
	}

	if rebootInstance && len(diskIDLabelMap) > 0 && len(updatedConfigMap) > 0 && bootConfig > 0 {
		err = client.RebootInstance(context.Background(), instance.ID, bootConfig)

		if err != nil {
			return fmt.Errorf("Error rebooting Instance %d: %s", instance.ID, err)
		}
		_, err = client.WaitForEventFinished(context.Background(), id, linodego.EntityLinode, linodego.ActionLinodeReboot, *instance.Created, int(d.Timeout(schema.TimeoutUpdate).Seconds()))
		if err != nil {
			return fmt.Errorf("Error waiting for Instance %d to finish rebooting: %s", instance.ID, err)
		}
		if _, err = client.WaitForInstanceStatus(context.Background(), instance.ID, linodego.InstanceRunning, int(d.Timeout(schema.TimeoutUpdate).Seconds())); err != nil {
			return fmt.Errorf("Timed-out waiting for Linode instance %d to boot: %s", instance.ID, err)
		}
	}

	return resourceLinodeInstanceRead(d, meta)
}

func resourceLinodeInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode Instance ID %s as int", d.Id())
	}
	minDelete := time.Now().AddDate(0, 0, -1)
	err = client.DeleteInstance(context.Background(), int(id))
	if err != nil {
		return fmt.Errorf("Error deleting Linode instance %d: %s", id, err)
	}
	// Wait for full deletion to assure volumes are detached
	client.WaitForEventFinished(context.Background(), int(id), linodego.EntityLinode, linodego.ActionLinodeDelete, minDelete, int(d.Timeout(schema.TimeoutDelete).Seconds()))

	d.SetId("")
	return nil
}
