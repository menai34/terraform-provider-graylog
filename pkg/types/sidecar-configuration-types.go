package types

// https://github.com/Graylog2/graylog2-server/blob/3913f709385b42b1043f8556e49cea681d7036d7/graylog2-server/src/main/java/org/graylog/plugins/sidecar/rest/models/Configuration.java

type SidecarConfigurationCreateRequest struct {
	CollectorID string `json:"collector_id"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Template    string `json:"template"`
}

// https://github.com/Graylog2/graylog2-server/blob/3913f709385b42b1043f8556e49cea681d7036d7/graylog2-server/src/main/java/org/graylog/plugins/sidecar/rest/models/Configuration.java

type SidecarConfigurationCreateResponse struct {
	ID          string `json:"id"`
	CollectorID string `json:"collector_id"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Template    string `json:"template"`
}
