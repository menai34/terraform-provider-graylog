package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/menai34/terraform-provider-graylog/pkg/graylog"
	"github.com/menai34/terraform-provider-graylog/pkg/types"
)

func resourceGraylogSidecarCollector() *schema.Resource {
	return &schema.Resource{
		Create: resourceGraylogSidecarCollectorCreate,
		Read:   resourceGraylogSidecarCollectorRead,
		Update: resourceGraylogSidecarCollectorUpdate,
		Delete: resourceGraylogSidecarCollectorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"node_operating_system": {
				Type:     schema.TypeString,
				Required: true,
			},
			"executable_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"execute_parameters": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"validation_parameters": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_template": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceGraylogSidecarCollectorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)

	request, err := resourceGraylogSidecarCollectorGenerateCreateRequest(d)
	if err != nil {
		return err
	}

	response := new(types.SidecarCollectorCreateResponse)
	err = client.Post("/api/sidecar/collectors", request, response)
	if err != nil {
		return err
	}

	d.SetId(response.ID)

	return resourceGraylogSidecarCollectorRead(d, meta)
}

func resourceGraylogSidecarCollectorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)

	input := new(types.SidecarCollectorCreateResponse)
	url := fmt.Sprintf("/api/sidecar/collectors/%s", d.Id())
	err := client.Get(url, input)
	_ = err // TODO: Gracefully handle 404s

	d.Set("id", input.ID)
	d.Set("name", input.Name)
	d.Set("service_type", input.ServiceType)
	d.Set("node_operating_system", input.NodeOperatingSystem)
	d.Set("executable_path", input.ExecutablePath)
	d.Set("execute_parameters", input.ExecuteParameters)
	d.Set("validation_parameters", input.ValidationCommand)
	d.Set("default_template", input.DefaultTemplate)

	return nil
}

func resourceGraylogSidecarCollectorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)

	request, err := resourceGraylogSidecarCollectorGenerateCreateRequest(d)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("/api/sidecar/collectors/%s", d.Id())
	err = client.Put(url, request, nil)
	if err != nil {
		return err
	}

	return resourceGraylogSidecarCollectorRead(d, meta)
}

func resourceGraylogSidecarCollectorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)
	url := fmt.Sprintf("/api/sidecar/collectors/%s", d.Id())
	return client.Delete(url)
}

func resourceGraylogSidecarCollectorGenerateCreateRequest(d *schema.ResourceData) (*types.SidecarCollectorCreateRequest, error) {
	request := &types.SidecarCollectorCreateRequest{
		Name:                d.Get("name").(string),
		ServiceType:         d.Get("service_type").(string),
		NodeOperatingSystem: d.Get("node_operating_system").(string),
		ExecutablePath:      d.Get("executable_path").(string),
		ExecuteParameters:   d.Get("execute_parameters").(string),
		ValidationCommand:   d.Get("validation_parameters").(string),
		DefaultTemplate:     d.Get("default_template").(string),
	}

	return request, nil
}
