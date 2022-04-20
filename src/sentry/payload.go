package sentry

type SentryTriggeredPayload struct {
	ID              string      `json:"id"`
	Project         string      `json:"project"`
	ProjectName     string      `json:"project_name"`
	ProjectSlug     string      `json:"project_slug"`
	Logger          interface{} `json:"logger"`
	Level           string      `json:"level"`
	Culprit         string      `json:"culprit"`
	Message         string      `json:"message"`
	URL             string      `json:"url"`
	TriggeringRules []string    `json:"triggering_rules"`
	Event           struct {
		EventID     string   `json:"event_id"`
		Level       string   `json:"level"`
		Version     string   `json:"version"`
		Type        string   `json:"type"`
		Fingerprint []string `json:"fingerprint"`
		Culprit     string   `json:"culprit"`
		Transaction string   `json:"transaction"`
		Logger      string   `json:"logger"`
		Platform    string   `json:"platform"`
		Timestamp   float64  `json:"timestamp"`
		Received    float64  `json:"received"`
		Environment string   `json:"environment"`
		Request     struct {
			URL                 string     `json:"url"`
			Method              string     `json:"method"`
			Data                string     `json:"data"`
			QueryString         [][]string `json:"query_string"`
			Headers             [][]string `json:"headers"`
			InferredContentType string     `json:"inferred_content_type"`
		} `json:"request"`
		Contexts struct {
			Runtime struct {
				Name    string `json:"name"`
				Version string `json:"version"`
				Type    string `json:"type"`
			} `json:"runtime"`
			Trace struct {
				TraceID string `json:"trace_id"`
				SpanID  string `json:"span_id"`
				Op      string `json:"op"`
				Status  string `json:"status"`
				Data    struct {
					BaseURL string `json:"baseUrl"`
					Query   struct {
						Version string `json:"version"`
					} `json:"query"`
					URL string `json:"url"`
				} `json:"data"`
				Tags struct {
					SentrySampleRate     string `json:"__sentry_sampleRate"`
					SentrySamplingMethod string `json:"__sentry_samplingMethod"`
					HTTPStatusCode       string `json:"http.status_code"`
				} `json:"tags"`
				Type string `json:"type"`
			} `json:"trace"`
		} `json:"contexts"`
		Exception struct {
			Values []struct {
				Type       string `json:"type"`
				Value      string `json:"value"`
				Stacktrace struct {
					Frames []struct {
						Function    string   `json:"function"`
						Module      string   `json:"module"`
						Filename    string   `json:"filename"`
						AbsPath     string   `json:"abs_path"`
						Lineno      int      `json:"lineno"`
						Colno       int      `json:"colno"`
						PreContext  []string `json:"pre_context"`
						ContextLine string   `json:"context_line"`
						PostContext []string `json:"post_context"`
						InApp       bool     `json:"in_app"`
					} `json:"frames"`
				} `json:"stacktrace"`
				Mechanism struct {
					Type    string `json:"type"`
					Handled bool   `json:"handled"`
				} `json:"mechanism"`
			} `json:"values"`
		} `json:"exception"`
		Tags [][]string `json:"tags"`
		Sdk  struct {
			Name         string   `json:"name"`
			Version      string   `json:"version"`
			Integrations []string `json:"integrations"`
			Packages     []struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"packages"`
		} `json:"sdk"`
		KeyID          string `json:"key_id"`
		Project        int    `json:"project"`
		GroupingConfig struct {
			Enhancements string `json:"enhancements"`
			ID           string `json:"id"`
		} `json:"grouping_config"`
		Metrics struct {
			BytesIngestedEvent int `json:"bytes.ingested.event"`
			BytesStoredEvent   int `json:"bytes.stored.event"`
		} `json:"_metrics"`
		Ref        int      `json:"_ref"`
		RefVersion int      `json:"_ref_version"`
		Hashes     []string `json:"hashes"`
		Location   string   `json:"location"`
		Metadata   struct {
			Filename string `json:"filename"`
			Function string `json:"function"`
			Type     string `json:"type"`
			Value    string `json:"value"`
		} `json:"metadata"`
		Title string `json:"title"`
		Meta  struct {
		} `json:"_meta"`
		ID string `json:"id"`
	} `json:"event"`
}
