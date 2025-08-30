package clevercloud

// AddonProvider represents an addon provider with its plans
type AddonProvider struct {
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Plans []AddonPlan `json:"plans"`
}

// AddonPlan represents a specific plan for an addon
type AddonPlan struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// ProductInstance represents an application type with its flavors (plans)
type ProductInstance struct {
	Type          string   `json:"type"`
	Version       string   `json:"version"`
	Name          string   `json:"name"`
	Variant       Variant  `json:"variant"`
	Description   string   `json:"description"`
	Enabled       bool     `json:"enabled"`
	ComingSoon    bool     `json:"comingSoon"`
	MaxInstances  int      `json:"maxInstances"`
	Tags          []string `json:"tags"`
	Deployments   []string `json:"deployments"`
	Flavors       []Flavor `json:"flavors"`
	DefaultFlavor Flavor   `json:"defaultFlavor"`
	BuildFlavor   Flavor   `json:"buildFlavor"`
}

// Variant represents application variant information
type Variant struct {
	ID         string `json:"id"`
	Slug       string `json:"slug"`
	Name       string `json:"name"`
	DeployType string `json:"deployType"`
	Logo       string `json:"logo"`
}

// Memory represents memory configuration
type Memory struct {
	Unit      string `json:"unit"`
	Value     int    `json:"value"`
	Formatted string `json:"formatted"`
}

// Flavor represents a specific flavor/plan for an application type
type Flavor struct {
	Name            string  `json:"name"`
	Mem             int     `json:"mem"`
	Cpus            int     `json:"cpus"`
	Gpus            int     `json:"gpus"`
	Disk            any     `json:"disk"`
	Price           float64 `json:"price"`
	Available       bool    `json:"available"`
	Microservice    bool    `json:"microservice"`
	MachineLearning bool    `json:"machine_learning"`
	Nice            int     `json:"nice"`
	PriceID         string  `json:"price_id"`
	Memory          Memory  `json:"memory"`
}
