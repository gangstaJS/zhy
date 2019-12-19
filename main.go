package main

import (
	"encoding/json"
	"net/http"
	"github.com/labstack/echo"
	"os/exec"
	"log"
)

var Battery struct {
	Health      string
	Percentage  int
	Status      string
	Temperature float64
}

func main() {



	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		cmd := exec.Command("termux-battery-status")

		//StdoutPipe returns a pipe that will be connected to
		//the command's standard output when the command starts
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}

		if err := json.NewDecoder(stdout).Decode(&Battery); err != nil {
			log.Fatal(err)
		}

		//Wait waits for the command to exit
		//It must have been started by Start
		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}

		return c.JSON(http.StatusOK, Battery)
	})


	e.Logger.Fatal(e.Start(":1323"))
}