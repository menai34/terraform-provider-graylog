package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/rebuy-de/terraform-provider-graylog/pkg/graylog"
	"github.com/rebuy-de/terraform-provider-graylog/pkg/types"
)

func resourceGraylogSidecarConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceGraylogSidecarConfigurationCreate,
		Read:   resourceGraylogSidecarConfigurationRead,
		Update: resourceGraylogSidecarConfigurationUpdate,
		Delete: resourceGraylogSidecarConfigurationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"collector_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"color": {
				Type:     schema.TypeString,
				Required: true,
			},
			"template": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGraylogSidecarConfigurationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)

	request, err := resourceGraylogSidecarConfigurationGenerateCreateRequest(d)
	if err != nil {
		return err
	}

	response := new(types.SidecarConfigurationCreateResponse)
	err = client.Post("/api/sidecar/configurations", request, response)
	if err != nil {
		return err
	}

	d.SetId(response.ID)

	return resourceGraylogSidecarConfigurationRead(d, meta)
}

func resourceGraylogSidecarConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)

	input := new(types.SidecarConfigurationCreateResponse)
	url := fmt.Sprintf("/api/sidecar/configurations/%s", d.Id())
	err := client.Get(url, input)
	_ = err // TODO: Gracefully handle 404s

	fmt.Println(input)

	d.Set("id", input.ID)
	d.Set("collector_id", input.CollectorID)
	d.Set("name", input.Name)
	d.Set("color", input.Color)
	d.Set("template", input.Template)

	return nil
}

func resourceGraylogSidecarConfigurationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)

	request, err := resourceGraylogSidecarConfigurationGenerateCreateRequest(d)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("/api/sidecar/configurations/%s", d.Id())
	err = client.Put(url, request, nil)
	if err != nil {
		return err
	}

	return resourceGraylogSidecarConfigurationRead(d, meta)
}

func resourceGraylogSidecarConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)
	url := fmt.Sprintf("/api/sidecar/configurations/%s", d.Id())
	return client.Delete(url)
}

func resourceGraylogSidecarConfigurationGenerateCreateRequest(d *schema.ResourceData) (*types.SidecarConfigurationCreateRequest, error) {
	request := &types.SidecarConfigurationCreateRequest{
		CollectorID: d.Get("collector_id").(string),
		Name:        d.Get("name").(string),
		Color:       d.Get("color").(string),
		Template:    d.Get("template").(string),
	}

	return request, nil
}
