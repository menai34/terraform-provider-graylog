package types

// https://github.com/Graylog2/graylog2-server/blob/334fd8633b9943eb6c95a25ff5b7562b89338c73/graylog2-server/src/main/java/org/graylog/plugins/sidecar/rest/models/ConfigurationVariable.java

type SidecarConfigurationVariableCreateResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
