package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/menai34/terraform-provider-graylog/pkg/graylog"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRAYLOG_SERVER_URL", "http://localhost:9000"),
				Description: "URL to the Graylog API",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRAYLOG_USERNAME", "admin"),
				Description: "Username for the Graylog API",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRAYLOG_PASSWORD", "admin"),
				Description: "Password for the Graylog API",
			},
			"credentials": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRAYLOG_CREDENTIALS", ""),
				Description: "Credentials file for the Graylog API",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"graylog_input":                          resourceGraylogInput(),
			"graylog_sidecar_collector":              resourceGraylogSidecarCollector(),
			"graylog_sidecar_configuration":          resourceGraylogSidecarConfiguration(),
			"graylog_sidecar_configuration_variable": resourceGraylogSidecarConfigurationVariable(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return &graylog.Client{
		ServerURL:   d.Get("server_url").(string),
		Username:    d.Get("username").(string),
		Password:    d.Get("password").(string),
		Credentials: d.Get("credentials").(string),
	}, nil
}
