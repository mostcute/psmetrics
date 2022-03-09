package psmetrics

import (
	"github.com/gin-gonic/gin"
	psDisk "github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	psMem "github.com/shirou/gopsutil/mem"
	psNet "github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/v3/cpu"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

func HandleDisk(c *gin.Context) {

	partitions, err := psDisk.Partitions(false)
	if err != nil {
		log.Println("err", err)
	}
	c.JSON(http.StatusOK, partitions)
}

func HandleDiskUsage(c *gin.Context) {
	path := c.GetHeader("path")
	stats, err := psDisk.Usage(path)
	if err != nil {
		log.Println("err", err)
	}
	c.JSON(http.StatusOK, stats)
}

func HandleDiskIOCounter(c *gin.Context) {
	path := c.GetHeader("path")
	stats, err := psDisk.IOCounters(path)
	if err != nil {
		log.Println("err", err)
	}
	c.JSON(http.StatusOK, stats)
}

func HandleNetIOCounter(c *gin.Context) {
	stats, err := psNet.IOCounters(true)
	if err != nil {
		log.Println("err", err)
	}
	c.JSON(http.StatusOK, stats)
}

func HandleSensorsTemperatures(c *gin.Context) {
	sensors, err := host.SensorsTemperatures()
	if err != nil {
		log.Println("err", err)
	}
	c.JSON(http.StatusOK, sensors)
}

func HandleCpuCount(c *gin.Context) {
	cpuCount, err := cpu.Counts(false)
	if err != nil {
		log.Println("err", err)
	}
	if cpuCount == 0 {
		is, err := cpu.Info()
		if err != nil {
			log.Println("err", err)
		}
		if is[0].Cores > 0 {
			c.JSON(http.StatusOK, len(is)/2)
			return
		}
		c.JSON(http.StatusOK, len(is))
		return
	}

	c.JSON(http.StatusOK, cpuCount)
	return
}

func HandleCpuPercent(c *gin.Context) {
	percpu := c.GetHeader("percpu")
	var per bool
	if percpu != "" {
		per = true
	}
	sensors, err := cpu.Percent(0, per)
	if err != nil {
		log.Println("err", err)
	}
	c.JSON(http.StatusOK, sensors)
}

func HandleMem(c *gin.Context) {
	mem, err := psMem.VirtualMemory()
	if err != nil {
		log.Println("err", err)
	}
	c.JSON(http.StatusOK, mem)
}

type Proc struct {
	Pid         int
	CommandName string
	FullCommand string
	CPU         float64
	Mem         float64
}

func Handleprocs(c *gin.Context) {
	output, err := exec.Command("ps", "-axo", "pid:10,comm:50,pcpu:5,pmem:5,args").Output()
	if err != nil {
		log.Println("err", err)
		return
	}

	// converts to []string, removing trailing newline and header
	linesOfProcStrings := strings.Split(strings.TrimSuffix(string(output), "\n"), "\n")[1:]

	procs := []Proc{}
	for _, line := range linesOfProcStrings {
		pid, err := strconv.Atoi(strings.TrimSpace(line[0:10]))
		if err != nil {
			log.Println("err", err)
		}
		cpu, err := strconv.ParseFloat(strings.TrimSpace(line[63:68]), 64)
		if err != nil {
			log.Println("err", err)
		}
		mem, err := strconv.ParseFloat(strings.TrimSpace(line[69:74]), 64)
		if err != nil {
			log.Println("err", err)
		}
		proc := Proc{
			Pid:         pid,
			CommandName: strings.TrimSpace(line[11:61]),
			FullCommand: line[74:],
			CPU:         cpu,
			Mem:         mem,
		}
		procs = append(procs, proc)
	}
	c.JSON(http.StatusOK, procs)
}

func HandleDevices(c *gin.Context) {
	sensors, err := host.SensorsTemperatures()
	if err != nil {
		log.Println("err", err)
	}
	c.JSON(http.StatusOK, sensors)
}

func HandleMemSwap(c *gin.Context) {
	memory, err := psMem.SwapMemory()
	if err != nil {
		log.Println("err", err)
	}
	c.JSON(http.StatusOK, memory)
}

func HandleNvidiaInfo(c *gin.Context) {
	bs, err := exec.Command(
		"nvidia-smi",
		"--query-gpu=name,index,temperature.gpu,utilization.gpu,memory.total,memory.used",
		"--format=csv,noheader,nounits").Output()
	if err != nil {
		log.Println("err", err)
	}
	c.JSON(http.StatusOK, bs)
}
