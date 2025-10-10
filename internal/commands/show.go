package commands

import (
	"fmt"

	"github.com/htooanttko/mini-terraform/internal/state"
)

func ShowCmd() {
	st, err := state.LoadState(".mini-terra/mini-terra.state.json")
	if err != nil {
		fmt.Println("No state found.")
		return
	}
	b, _ := state.PrettyJSON(st)
	fmt.Println(string(b))
}
