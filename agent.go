package dialogflow

// Agent Represents a conversational agent
type Agent struct {
	Parent                  string   `json:"parent"`
	DisplayName             string   `json:"displayName"`
	DefaultLanguageCode     string   `json:"defaultLanguageCode"`
	SupportedLanguageCodes  []string `json:"supportedLanguageCodes"`
	TimeZone                string   `json:"timeZone"`
	Description             string   `json:"description"`
	AvatarURI               string   `json:"avatarUri"`
	EnableLogging           bool     `json:"enableLogging"`
	MatchMode               string   `json:"matchMode"`
	ClassificationThreshold float32  `json:"classificationThreshold"`
}

func (a *Agent) String() string {
	return a.DisplayName
}

// GetAgent Retrieves the specified agent
func (c *Client) GetAgent() (*Agent, error) {
	var agent Agent
	err := c.get("/agent", &agent)
	return &agent, err
}
