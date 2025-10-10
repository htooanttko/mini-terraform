package main

import (
	_ "embed"
	"os"

	"github.com/htooanttko/mini-terraform/internal/commands"
	"github.com/htooanttko/mini-terraform/internal/providers"
	awsProvider "github.com/htooanttko/mini-terraform/internal/providers/aws"
	dockerProvider "github.com/htooanttko/mini-terraform/internal/providers/docker"
	vpsProvider "github.com/htooanttko/mini-terraform/internal/providers/vps"
)

func main() {
	if len(os.Args) < 2 {
		commands.Usage()
		os.Exit(1)
	}
	// register providers
	providers.RegisterProvider("docker", dockerProvider.NewDockerProvider())
	providers.RegisterProvider("vps", vpsProvider.NewVPSProvider())
	providers.RegisterProvider("aws", awsProvider.NewAWSProvider())

	cmd := os.Args[1]
	switch cmd {
	case "init":
		commands.InitCmd()
	case "plan":
		commands.PlanCmd()
	case "apply":
		commands.ApplyCmd()
	case "destroy":
		commands.DestroyCmd()
	case "show":
		commands.ShowCmd()
	case "version":
		commands.Version()
	case "template":
		commands.TemplateCmd()
	default:
		commands.Usage()
	}
}
