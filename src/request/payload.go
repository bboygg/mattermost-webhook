package request

type WebhookInfo struct {
	Sentry []struct {
		Channel string `json:"channel"`
		URL     string `json:"url"`
	} `json:"sentry"`
	Qase []struct {
		Channel string `json:"channel"`
		URL     string `json:"url"`
	} `json:"qase"`
}
