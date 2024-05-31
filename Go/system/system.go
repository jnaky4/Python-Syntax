package system

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	_ "syscall"
)

//func main(){
//	//runtimeOS := runtime.GOOS //get running OS
//	//fmt.Printf("%v\n", runtimeOS)
//	//info, err := cpu.Info()
//	//if err != nil {
//	//	return
//	//}
//	//fmt.Printf("%v\n", info)
//
//	cores := GetCores()
//	cpuAvg := GetUptimeAvgCPUMin(0)
//	fmt.Printf("uptime - avg cpu utilization: %f%%\n", (cpuAvg / float64(cores)) * 100 )
//	GetAVGCPUUsage()
//}

type Platform struct{
	Architecture string   `json:"architecture"`
	OS           string   `json:"os"`
	OSVersion    string   `json:"os.version,omitempty"`
	OSFeatures   []string `json:"os.features,omitempty"`
	Variant      string   `json:"variant,omitempty"`
}
type ComputerSpecs struct{
	AvgCPU0 float64
	AvgCPU5 float64
	AvgCPU15 float64
	AvgCPUPerCore []float64
	CPU int				  `json:"cpu"`
	MemUsed int
	MemAvail int
}

func GetPlatformSpecs() (p Platform){
	switch runtime.GOOS {

	case "Darwin":
		p.OS = "Unix"
		p.OSVersion = runtime.GOOS
	default:
		p.OS = runtime.GOOS
	}

	p.Architecture = runtime.Version()
	return
}

func GetUptimeAvgCPUMin(min int) float64{
	cmd := exec.Command("uptime")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	//fmt.Println(string(stdout))

	uptime := strings.Split(string(stdout), " ")
	//fmt.Printf("%s\n", uptime[len(uptime) -1])
	s := ""
	switch min {
		case 0:
			s = uptime[len(uptime) -3]
			break
		case 5:
			s = uptime[len(uptime) -2]
			break
		case 15:
			s = uptime[len(uptime) -1]
			break

	}

	fmt.Printf("uptime: %s\n", s)
	s = strings.Replace(s, "\n", "", -1)
	float, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return -1
	}
	return float
}

func GetAVGCPUUsage() float64{
	percent, err := cpu.Percent(0, true)
	if err != nil {
		return -1
	}

	sum := 0.0
	for _, corePercent := range percent {
		sum += corePercent
	}
	fmt.Printf("uptime: %v\n sum: %f\n", percent, sum)
	sum = sum / float64(len(percent))
	fmt.Printf("goputil - avg cpu utilization: %f%% \n", sum)
	return sum
}

func GetCores() int{
	return runtime.NumCPU()
	//if runtime.GOOS == "darwin" {
	//	cmd := exec.Command("sysctl", "-n", "hw.ncpu")
	//	stdout, err := cmd.Output()
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return -1
	//	}
	//	s := string(stdout)
	//	s = strings.Replace(s, "\n", "", -1)
	//	atoi, err := strconv.Atoi(s)
	//	if err != nil {
	//		fmt.Printf("%s", err.Error())
	//		return -1
	//	}
	//
	//	fmt.Printf("cores: %d\n", atoi)
	//	return atoi
	//}
	//return -1
}


