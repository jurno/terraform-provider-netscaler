package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerCapolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_capolicy,
		Read:          read_capolicy,
		Update:        update_capolicy,
		Delete:        delete_capolicy,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"logaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"undefaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_capolicy(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_capolicy(d *schema.ResourceData) nitro.Capolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Capolicy{
		Name:        d.Get("name").(string),
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	return resource
}

func set_capolicy(d *schema.ResourceData, resource *nitro.Capolicy) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("action", resource.Action)
	d.Set("comment", resource.Comment)
	d.Set("logaction", resource.Logaction)
	d.Set("rule", resource.Rule)
	d.Set("undefaction", resource.Undefaction)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_capolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_capolicy")

	client := meta.(*nitro.NitroClient)

	key := key_capolicy(d)

	exists, err := client.ExistsCapolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCapolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_capolicy(d, resource)
	} else {
		err := client.AddCapolicy(get_capolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCapolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_capolicy(d, resource)
	}

	return nil
}

func read_capolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_capolicy")

	client := meta.(*nitro.NitroClient)

	key := key_capolicy(d)

	exists, err := client.ExistsCapolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCapolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_capolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_capolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_capolicy")

	return nil
}

func delete_capolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_capolicy")

	client := meta.(*nitro.NitroClient)

	key := key_capolicy(d)

	exists, err := client.ExistsCapolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCapolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
