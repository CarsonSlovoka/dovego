package config

type Server struct {
	Port     int `json:"port"`
	Shutdown `json:"shutdown"`
	StartURL string `json:"start_url"`
}

type Shutdown struct {
	URL string `json:"url"`
}
