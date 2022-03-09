package psmetrics

import "github.com/gin-gonic/gin"

func Register(r *gin.Engine) {
	metricsgroup := r.Group("/metric")

	metricsgroup.GET("/disk", HandleDisk)
	metricsgroup.GET("/diskusage", HandleDiskUsage)
	metricsgroup.GET("/diskiocounter", HandleDiskIOCounter)
	metricsgroup.GET("/sensorstemperatures", HandleSensorsTemperatures)
	metricsgroup.GET("/cpucount", HandleCpuCount)
	metricsgroup.GET("/cpupercent", HandleCpuPercent)
	metricsgroup.GET("/netiocounter", HandleNetIOCounter)
	metricsgroup.GET("/mem", HandleMem)
	metricsgroup.GET("/proc", Handleprocs)
	metricsgroup.GET("/memswap", HandleMemSwap)
	metricsgroup.GET("/nvidiaL", HandleMemSwap)
	metricsgroup.GET("/nvidiainfo", HandleNvidiaInfo)
}
