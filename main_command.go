package main

import (
	"docker/container"
	"docker/cgroups/subsystems"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var runCommand = &cli.Command{
	Name: "run",
	Usage: `创建容器  with namespace and cgroups limit
			mydocker run -ti [command]`,
	Flags: []cli.Flag{
		&cli.BoolFlag {
			Name:	"ti",
			Usage:	"enable tty",
		},
		&cli.StringFlag{
			Name:        "m",
			Usage:       "memory limit",
		},
		&cli.StringFlag{
			Name:        "cpushare",
			Usage:       "cpushare limit",
		},
		&cli.StringFlag{
			Name:        "cpuset",
			Usage:       "cpuset limit",
		},
	},
	Action: func(context *cli.Context) error{
		if context.Args().Len() < 1{
			return fmt.Errorf("Missing container command")
		}
		var cmdArray []string
		for i:=0;i<context.Args().Len();i++{
			cmdArray = append(cmdArray,context.Args().Get(i))
		}
		fmt.Println("-----------------")
		fmt.Println(cmdArray)
		fmt.Println("-----------------")
		tty := context.Bool("ti")
		resConf := &subsystems.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuShare:    context.String("cpushare"),
			CpuSet:      context.String("cpuset"),
		}
		fmt.Println(resConf)
		Run(tty,cmdArray,resConf)
		return nil
	},
}

var initCommand = &cli.Command{
	Name: "init",
	Usage: "Init container process run user's process in container. Don't call it outside",
	Action: func(context *cli.Context) error{
		log.Infof("init come on")
		err := container.RunContainerInitProcess()
		return err
	},
}
