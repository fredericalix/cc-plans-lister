package fixtures

import "cc-plans-lister/pkg/clevercloud"

// TestAddonProviders provides test data for addon providers
func TestAddonProviders() []clevercloud.AddonProvider {
	return []clevercloud.AddonProvider{
		{
			ID:   "redis",
			Name: "Redis",
			Plans: []clevercloud.AddonPlan{
				{ID: "redis_small", Name: "Small Redis", Slug: "small"},
				{ID: "redis_large", Name: "Large Redis", Slug: "large"},
			},
		},
		{
			ID:   "postgresql",
			Name: "PostgreSQL",
			Plans: []clevercloud.AddonPlan{
				{ID: "pg_dev", Name: "Dev PostgreSQL", Slug: "dev"},
				{ID: "pg_prod", Name: "Production PostgreSQL", Slug: "prod"},
			},
		},
	}
}

// TestProductInstances provides test data for product instances
func TestProductInstances() []clevercloud.ProductInstance {
	return []clevercloud.ProductInstance{
		{
			Type:         "node",
			Version:      "20",
			Name:         "Node.js",
			Description:  "Node.js runtime",
			Enabled:      true,
			MaxInstances: 20,
			Tags:         []string{"runtime", "javascript"},
			Deployments:  []string{"git", "docker"},
			Flavors: []clevercloud.Flavor{
				{
					Name:            "nano",
					Mem:             256,
					Cpus:            1,
					Price:           0.02,
					Available:       true,
					Microservice:    true,
					MachineLearning: false,
					Memory: clevercloud.Memory{
						Unit:      "MB",
						Value:     256,
						Formatted: "256 MB",
					},
				},
				{
					Name:            "small",
					Mem:             512,
					Cpus:            1,
					Price:           0.04,
					Available:       true,
					Microservice:    false,
					MachineLearning: false,
					Memory: clevercloud.Memory{
						Unit:      "MB",
						Value:     512,
						Formatted: "512 MB",
					},
				},
			},
			DefaultFlavor: clevercloud.Flavor{
				Name: "nano",
			},
		},
		{
			Type:         "python",
			Version:      "3.11",
			Name:         "Python",
			Description:  "Python runtime",
			Enabled:      true,
			MaxInstances: 10,
			Tags:         []string{"runtime", "python"},
			Deployments:  []string{"git"},
			Flavors: []clevercloud.Flavor{
				{
					Name:            "small",
					Mem:             512,
					Cpus:            1,
					Price:           0.04,
					Available:       true,
					Microservice:    false,
					MachineLearning: true,
					Memory: clevercloud.Memory{
						Unit:      "MB",
						Value:     512,
						Formatted: "512 MB",
					},
				},
			},
			DefaultFlavor: clevercloud.Flavor{
				Name: "small",
			},
		},
	}
}
