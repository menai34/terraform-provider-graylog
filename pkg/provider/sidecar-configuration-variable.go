package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/menai34/terraform-provider-graylog/pkg/graylog"
	"github.com/menai34/terraform-provider-graylog/pkg/types"
)

func resourceGraylogSidecarConfigurationVariable() *schema.Resource {
	return &schema.Resource{
		Create: resourceGraylogSidecarConfigurationVariableCreate,
		Read:   resourceGraylogSidecarConfigurationVariableRead,
		Update: resourceGraylogSidecarConfigurationVariableUpdate,
		Delete: resourceGraylogSidecarConfigurationVariableDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGraylogSidecarConfigurationVariableCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)

	request, err := resourceGraylogSidecarConfigurationVariableGenerateCreateRequest(d)
	if err != nil {
		return err
	}

	response := new(types.SidecarConfigurationVariableActionResponse)
	err = client.Post("/api/sidecar/configuration_variables", request, response)
	if err != nil {
		return err
	}

	d.SetId(response.ID)

	return resourceGraylogSidecarConfigurationVariableRead(d, meta)
}

func resourceGraylogSidecarConfigurationVariableRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)

	input := new(types.SidecarConfigurationVariableActionResponse)
	url := fmt.Sprintf("/api/sidecar/configuration_variables")
	err := client.Get(url, input)
	_ = err // TODO: Gracefully handle 404s

	d.Set("id", input.ID)
	d.Set("name", input.Name)
	d.Set("description", input.Description)
	d.Set("content", input.Content)

	return nil
}

func resourceGraylogSidecarConfigurationVariableUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)

	request, err := resourceGraylogSidecarConfigurationVariableGenerateUpdateRequest(d)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("/api/sidecar/configuration_variables/%s", d.Id())
	err = client.Put(url, request, nil)
	if err != nil {
		return err
	}

	return resourceGraylogSidecarConfigurationVariableRead(d, meta)
}

func resourceGraylogSidecarConfigurationVariableDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*graylog.Client)
	url := fmt.Sprintf("/api/sidecar/configuration_variables/%s", d.Id())
	return client.Delete(url)
}

func resourceGraylogSidecarConfigurationVariableGenerateCreateRequest(d *schema.ResourceData) (*types.SidecarConfigurationVariableCreateRequest, error) {
	request := &types.SidecarConfigurationVariableCreateRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Content:     d.Get("content").(string),
	}

	return request, nil
}

func resourceGraylogSidecarConfigurationVariableGenerateUpdateRequest(d *schema.ResourceData) (*types.SidecarConfigurationVariableUpdateRequest, error) {
	request := &types.SidecarConfigurationVariableUpdateRequest{
		ID:          d.Id(),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Content:     d.Get("content").(string),
	}

	return request, nil
}
