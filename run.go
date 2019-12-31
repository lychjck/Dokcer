package main

import(
	log "github.com/Sirupsen/logrus"
	"os"
	"./container"
	"./cgroups"
	"./cgroups/subsystems"
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

	cgroupmanager := cgroups.NewCgroupManager("Docker-cgroup")
	defer cgroupmanager.Destroy()
	cgroupmanager.Set(res)
	cgroupmanager.Apply(parent.Process.Pid)

	sendInitCommand(comArray,writePipe)
	parent.Wait()
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}