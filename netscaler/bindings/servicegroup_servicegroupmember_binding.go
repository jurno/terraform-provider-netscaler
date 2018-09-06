package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerServicegroupServicegroupmemberBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_servicegroup_servicegroupmember_binding,
		Read:          read_servicegroup_servicegroupmember_binding,
		Update:        nil,
		Delete:        delete_servicegroup_servicegroupmember_binding,
		Schema: map[string]*schema.Schema{
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicegroupname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_servicegroup_servicegroupmember_binding(d *schema.ResourceData) nitro.ServicegroupServicegroupmemberBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.ServicegroupServicegroupmemberBinding{
		Port:             d.Get("port").(int),
		Servername:       d.Get("servername").(string),
		Servicegroupname: d.Get("servicegroupname").(string),
		Weight:           d.Get("weight").(int),
	}

	return resource
}

func set_servicegroup_servicegroupmember_binding(d *schema.ResourceData, resource *nitro.ServicegroupServicegroupmemberBinding) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("port", resource.Port)
	d.Set("servername", resource.Servername)
	d.Set("servicegroupname", resource.Servicegroupname)
	d.Set("weight", resource.Weight)

	var key []string

	key = append(key, resource.Servicegroupname)
	key = append(key, resource.Servername)
	key = append(key, strconv.Itoa(resource.Port))
	d.SetId(strings.Join(key, "-"))
}

func get_servicegroup_servicegroupmember_binding_key(d *schema.ResourceData) nitro.ServicegroupServicegroupmemberBindingKey {

	key := nitro.ServicegroupServicegroupmemberBindingKey{
		d.Get("servicegroupname").(string),
		d.Get("servername").(string),
		d.Get("port").(int),
	}
	return key
}

func create_servicegroup_servicegroupmember_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_servicegroup_servicegroupmember_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_servicegroup_servicegroupmember_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsServicegroupServicegroupmemberBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetServicegroupServicegroupmemberBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_servicegroup_servicegroupmember_binding(d, resource)
	} else {
		err := client.AddServicegroupServicegroupmemberBinding(get_servicegroup_servicegroupmember_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetServicegroupServicegroupmemberBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_servicegroup_servicegroupmember_binding(d, resource)
	}

	return nil
}

func read_servicegroup_servicegroupmember_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_servicegroup_servicegroupmember_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_servicegroup_servicegroupmember_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsServicegroupServicegroupmemberBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetServicegroupServicegroupmemberBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_servicegroup_servicegroupmember_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_servicegroup_servicegroupmember_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_servicegroup_servicegroupmember_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_servicegroup_servicegroupmember_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsServicegroupServicegroupmemberBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteServicegroupServicegroupmemberBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
