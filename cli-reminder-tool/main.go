package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	markName  = "Golang-cli-reminder"
	markValue = "1"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage:%s <hh:mm> <text Message\n>", os.Args[0])
	}
	now := time.Now()
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	t, err := w.Parse(os.Args[1], now)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if t == nil {
		fmt.Println("Unable to parse time")
		os.Exit(2)
	}
	if now.After(t.Time) {
		fmt.Println("Set a future time")
		os.Exit(2)
	}

	diff := t.Time.Sub(now)

	if os.Getenv(markName) == markValue {
		time.Sleep(diff)
		err = beeep.Notify("Reminder", strings.Join(os.Args[2:], " "), "assets/information.png")
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(cmd.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
		fmt.Println("reminder will be disaplayed after", diff.Round(time.Second))
		os.Exit(0)
	}

}