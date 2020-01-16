package main

import (
	"docker/cgroups/subsystems"
	"docker/container"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)


func Run(tty bool,comArray []string,res *subsystems.ResourceConfig){
	parent,writePipe := container.NewParentProcess(tty)
	if parent ==nil{
		log.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil{
		log.Error(err)
	}

	//cgroupmanager := cgroups.NewCgroupManager("Docker-cgroup")
	//defer cgroupmanager.Destroy()
	//cgroupmanager.Set(res)
	//cgroupmanager.Apply(parent.Process.Pid)

	sendInitCommand(comArray,writePipe)
	parent.Wait()
	mntURL := "../workspace/root/mnt/"
	rootURL := "../workspace/root/"
	container.DeleteWorkSpace(rootURL,mntURL)
	os.Exit(0)
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
