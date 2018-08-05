package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerDnsaction64() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_dnsaction64,
		Read:          read_dnsaction64,
		Update:        update_dnsaction64,
		Delete:        delete_dnsaction64,
		Schema: map[string]*schema.Schema{
			"actionname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"excluderule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"mappedrule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_dnsaction64(d *schema.ResourceData) string {
	return d.Get("actionname").(string)
}

func get_dnsaction64(d *schema.ResourceData) nitro.Dnsaction64 {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Dnsaction64{
		Actionname:  d.Get("actionname").(string),
		Excluderule: d.Get("excluderule").(string),
		Mappedrule:  d.Get("mappedrule").(string),
		Prefix:      d.Get("prefix").(string),
	}

	return resource
}

func set_dnsaction64(d *schema.ResourceData, resource *nitro.Dnsaction64) {
	var _ = strconv.Itoa

	d.Set("actionname", resource.Actionname)
	d.Set("excluderule", resource.Excluderule)
	d.Set("mappedrule", resource.Mappedrule)
	d.Set("prefix", resource.Prefix)

	var key []string

	key = append(key, resource.Actionname)
	d.SetId(strings.Join(key, "-"))
}

func create_dnsaction64(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_dnsaction64")

	client := meta.(*nitro.NitroClient)

	key := key_dnsaction64(d)

	exists, err := client.ExistsDnsaction64(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetDnsaction64(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_dnsaction64(d, resource)
	} else {
		err := client.AddDnsaction64(get_dnsaction64(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetDnsaction64(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_dnsaction64(d, resource)
	}

	return nil
}

func read_dnsaction64(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_dnsaction64")

	client := meta.(*nitro.NitroClient)

	key := key_dnsaction64(d)

	exists, err := client.ExistsDnsaction64(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetDnsaction64(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_dnsaction64(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_dnsaction64(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_dnsaction64")

	return nil
}

func delete_dnsaction64(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_dnsaction64")

	client := meta.(*nitro.NitroClient)

	key := key_dnsaction64(d)

	exists, err := client.ExistsDnsaction64(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteDnsaction64(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
