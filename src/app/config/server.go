package config

type Server struct {
    Port int `json:"port"`
    Shutdown `json:"shutdown"`
}

type Shutdown struct {
    URL string `json:"url"`
}
