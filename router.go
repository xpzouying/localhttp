package main

// Router is router information for http service
type Router struct {
	URI      string `yaml:"uri"`      // router uri name
	Response string `yaml:"response"` // yaml file for router URI
}

// Routers is Router manager
type Routers struct {
	Rs []Router `yaml:"routers"`
}
