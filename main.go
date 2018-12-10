package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

// Checking Error panic, and logging it out.
func checkErr(err error, er *error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	if er != nil {
		fmt.Println("Error: Please Run the program again: ", *er)
	}
}

func main() {

	app := cli.NewApp()
	app.Name = "Traverses Internal directory"
	app.Usage = "Let's you traverse through your directory and run files"
	app.Version = "1.0.0"

	// Enter own source path
	dirname := "../../../Desktop/"

	// Kill Paths and ports
	killName := "kill -9 "

	myFlag := []cli.Flag{
		cli.StringFlag{
			Name:  "cmd",
			Value: dirname,
		},
	}
	killFlag := []cli.Flag{
		cli.StringFlag{
			Name:  "kill",
			Value: killName,
		},
	}

	// Creating Command TRC -- Reading Files
	app.Commands = []cli.Command{
		{
			Name:  "trc",
			Usage: "Reads Files in directories",
			Flags: myFlag,
			Action: func(c *cli.Context) error {
				var input string

				fmt.Println("Enter Path: ")

				_, fail := fmt.Scan(&input)

				trc, err := os.Open(dirname + input)

				files, err := trc.Readdir(-1)
				fmt.Println("Open: ", trc.Name())
				trc.Close()

				checkErr(fail, nil)
				checkErr(err, nil)
				checkErr(err, nil)

				for _, file := range files {
					fmt.Println(file.Name())
				}
				return nil
			},
		},
		{
			Name:  "pow",
			Usage: "Enter File name and start command for application",
			Flags: myFlag,
			Action: func(c *cli.Context) error {
				var paths string
				var runCmd1 string
				var runCmd2 string

				/* Todo:
				* Show The List of directories within the path
				* Command: terminates the port running exp: 3000
				* Use Command: kill -9 $(lsof -i:3000 -t) 2> /dev/null, to kill port
				 */

				fmt.Println("Enter First Part of The Command To Run: ")
				_, errcmd1 := fmt.Scan(&runCmd1)

				fmt.Println("Enter Second Part of The Command To Run: ")
				_, errcmd2 := fmt.Scan(&runCmd2)

				fmt.Println("Enter Desktop Dir Path: ")
				_, fail := fmt.Scan(&paths)

				cmd := exec.Command(runCmd1, runCmd2)
				cmd.Dir = dirname + paths

				log.Printf("Running command and waiting for it to finish...")

				err := cmd.Run()

				// Invoking Error Function:
				checkErr(fail, nil)
				checkErr(errcmd1, nil)
				checkErr(errcmd2, nil)
				checkErr(nil, &err)

				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "Command Kills All ports by using this command: kill -9 $(lsof -i:3000 -t) 2> /dev/null",
			Flags: killFlag,
			Action: func(c *cli.Context) error {
				var input string
				fmt.Println("Enter port Number: ")
				_, err := fmt.Scan(&input)

				newInput := killName + "$(lsof -i:" + input + " -t) 2> /dev/null"

				out, cmdError := exec.Command(newInput).Output()

				log.Print("Excuting Kill port... ", out)

				checkErr(err, nil)
				checkErr(nil, &cmdError)

				return nil
			},
		},
	}

	err := app.Run(os.Args)

	checkErr(err, nil)
}
