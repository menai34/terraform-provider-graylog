package types

// https://github.com/Graylog2/graylog2-server/blob/36438c2788e1b7ef4bd303ab85859a3ff758c5f1/graylog2-server/src/main/java/org/graylog/plugins/sidecar/rest/models/Collector.java

type SidecarCollectorCreateResponse struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	ServiceType         string `json:"service_type"`
	NodeOperatingSystem string `json:"node_operating_system"`
	ExecutablePath      string `json:"executable_path"`
	ExecuteParameters   string `json:"execute_parameters"`
	ValidationCommand   string `json:"validation_parameters"`
	DefaultTemplate     string `json:"default_template"`
}
