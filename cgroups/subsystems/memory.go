package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemorySubSystem struct{

}

//设置 cgroupPath 对应的 cgroup 的内存资源限制
func (s *MemorySubSystem) Set(cgroupPath string , res *ResourceConfig) error {
	//GetCgroupPath 的作用是获取当前 subsystem 在虚拟文件系统中的路径
	if subsysCgroupPath,err := GetCgroupPath(s.Name(),cgroupPath,true);err == nil{
		if res.MemoryLimit != ""{
			//设置这个 cgroup 的内存限制，即将限制写入到 cgroup 对应目录的 memory.limit in bytes文件中。
			if err := ioutil.WriteFile(path.Join(subsysCgroupPath,"memory.limit_in_bytes"),[]byte(res.MemoryLimit),0644);err != nil{
				return fmt.Errorf("set cgroup memory fail %v",err)
			}
		}
		return nil
	} else{
		return err
	}
}

//删除 cgroupPath 对应的 cgroup
func (s *MemorySubSystem)Remove(cgroupPath string) error {
	if subsysCgroupPath,err := GetCgroupPath(s.Name(),cgroupPath,false);err == nil{
		//删除 cgroup 便是删除对应的 cgroupPath 的目录
		os.Remove(subsysCgroupPath)
	}else{
		return err
	}
	return nil
}

//将1个迸程加入到 cgroupPath 对应的 cgroup
func (s *MemorySubSystem)Apply(cgroupPath string,pid int) error {
	if subsysCgroupPath,err := GetCgroupPath(s.Name(),cgroupPath,false);err == nil{
		//把进程的 PID 写到 cgroup 的虚拟文件系统对应目录下的” task ”文件中
		if err := ioutil.WriteFile(path.Join(subsysCgroupPath,"tasks"),[]byte(strconv.Itoa(pid)),0644); err != nil{
			return fmt.Errorf("set cgroup proc fail %v",err)
		}
		return nil
	}else{
		return fmt.Errorf("get cgroup %s error: %v",cgroupPath,err)
	}
}

func (s *MemorySubSystem) Name() string{
	return "memory"
}