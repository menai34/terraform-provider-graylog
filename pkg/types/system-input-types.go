package types

import "time"

// https://github.com/Graylog2/graylog2-server/blob/master/graylog2-server/src/main/java/org/graylog2/rest/models/system/inputs/requests/InputCreateRequest.java
type SystemInputCreateRequest struct {
	Title         string                 `json:"title"`
	Type          string                 `json:"type"`
	Global        bool                   `json:"global"`
	Configuration map[string]interface{} `json:"configuration"`
	Node          string                 `json:"node"`
}

// https://github.com/Graylog2/graylog2-server/blob/master/graylog2-server/src/main/java/org/graylog2/rest/models/system/inputs/responses/InputCreated.java

type SystemInputCreateResponse struct {
	ID string `json:"id"`
}

// https://github.com/Graylog2/graylog2-server/blob/master/graylog2-server/src/main/java/org/graylog2/rest/models/system/inputs/responses/InputSummary.java
type SystemInputSummary struct {
	Title         string
	Global        bool
	Name          string
	ContentPack   *string `json:"content_pack,omitempty"`
	ID            string
	CreatedAt     time.Time
	Type          string
	CreatorUserID string `json:"creator_user_id"`
	Attributes    map[string]interface{}
	StaticFields  map[string]string `json:"static_fields"`
	Node          *string           `json:",omitempty"`
}
