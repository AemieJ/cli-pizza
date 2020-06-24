package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var app = cli.NewApp()
var finalStatement = []string{"Enjoy the pizza with the delicious", "toppings."}

func info() {
	app.Name = "Create Pizza To Your Own Design"
	app.Usage = "CLI To help you create pizza as per your toppings"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []*cli.Command{
		{
			Name:    "Veg Favs",
			Aliases: []string{"vf", "pt"},
			Usage:   "Add paneer tikka as the toppings on your pizza",
			Action: func(c *cli.Context) error {
				pt := "Paneer Tikka"
				paneer := append(finalStatement, "0")
				copy(paneer[2:], paneer[1:])
				paneer[1] = pt
				statement := strings.Join(paneer, " ")
				fmt.Println(statement)
				return nil
			},
		},
		{
			Name:    "Non-Veg Favs",
			Aliases: []string{"nvf", "bt"},
			Usage:   "Add butter chicken as the toppings on your pizza",
			Action: func(c *cli.Context) error {
				bc := "Butter Chicken"
				chicken := append(finalStatement, "0")
				copy(chicken[2:], chicken[1:])
				chicken[1] = bc
				statement := strings.Join(chicken, " ")
				fmt.Println(statement)
				return nil
			},
		},
		{
			Name:    "Cheese lovers",
			Aliases: []string{"cl", "ec"},
			Usage:   "Add extra cheese as the toppings on your pizza",
			Action: func(c *cli.Context) error {
				ec := "Extra Cheese"
				cheese := append(finalStatement, "0")
				copy(cheese[2:], cheese[1:])
				cheese[1] = ec
				statement := strings.Join(cheese, " ")
				fmt.Println(statement)
				return nil
			},
		},
	}
}

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
