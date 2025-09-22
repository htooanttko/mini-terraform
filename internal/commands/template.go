package commands

import (
	"fmt"
	"os"
)

func TemplateCmd() {
	configContent := `{
	  "providers": {
	    "docker": {}
	  },
	  "variables": {
	    "image": "nginx:latest",
	    "container_name": "mini-terra-nginx"
	  },
	  "resources": [
	    {
	      "type": "docker_container",
	      "name": "${var.container_name}",
	      "provider": "docker",
	      "attributes": {
	        "image": "${var.image}",
	        "ports": ["8080:80"]
	      }
	    }
	  ]
	}`

	varsContent := `{
	  "image": "nginx:latest",
	  "container_name": "mini-terra-nginx"
	}`

	if err := os.WriteFile("config.json", []byte(configContent), 0644); err != nil {
		fmt.Printf("Error writing config.json: %v\n", err)
		return
	}

	if err := os.WriteFile("vars.json", []byte(varsContent), 0644); err != nil {
		fmt.Printf("Error writing vars.json: %v\n", err)
		return
	}

	fmt.Println("Template files generated: config.json and vars.json")
}
