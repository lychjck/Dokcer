package main

import(
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

const usage = `
		Hello! This is my docker
		Then ,你可以试一试啦，哈哈哈！
`

func main(){
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = usage

	app.Commands = []*cli.Command{
		initCommand,
		runCommand,
	}
	app.Before = func(context *cli.Context) error{
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}
	
	if err := app.Run(os.Args);err != nil {
		log.Fatal(err)
	}
}
