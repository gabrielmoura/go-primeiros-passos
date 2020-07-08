package main

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

const (
	// vcgencmd get_throttled result
	UnderVoltage        int64 = 1
	FreqCap                   = 1 << 1
	Throttling                = 1 << 2
	UnderVoltageOccured       = 1 << 16
	FreqCapOccured            = 1 << 17
	Throttled                 = 1 << 18
)

func getHostname() (string, error) {
	return os.Hostname()
}

func GetThrottled() (int64, error) {
	rawThrottled, err := exec.Command("vcgencmd", "get_throttled").Output()
	len := len(rawThrottled)
	if err != nil || len < 14 {
		return 0, errors.New("couldn't run vcgencmd")
	}
	rawThrottled = rawThrottled[12 : len-1]

	throttled, err := strconv.ParseInt(string(rawThrottled), 16, 32)
	if err != nil || len < 14 {
		return 0, errors.New("couldn't parse throttled output : " + string(rawThrottled))
	}
	return throttled, nil
}

//GetGPUTemp Retorna temperatura da GPU
func GetGPUTemp() (string, error) {
	temp, err := exec.Command("vcgencmd", "measure_temp").Output()

	if err != nil {
		return "", errors.New("couldn't run vcgencmd")
	}
	// /cpuTemp := cpu.Clean(cmd.Exec("vcgencmd", "measure_temp"), "temp=", "'C")
	ax := clean(string(temp), "temp=", "'C")

	return ax, nil
}

//GetCoreVolt Retorna Voltagem na CPU
func GetCoreVolt() (string, error) {
	volt, err := exec.Command("vcgencmd", "measure_volts").Output()

	if err != nil {
		return "", errors.New("couldn't run vcgencmd")
	}

	ax := clean(string(volt), "volt=", "V")

	return ax, nil
}

func Getmem() (string, string, error) {
	usageMem, err := exec.Command("vcgencmd", "get_mem arm").Output()
	gpuMem, err := exec.Command("vcgencmd", "get_mem gpu").Output()

	if err != nil {
		return "", "", errors.New("couldn't run vcgencmd")
	}

	//ax := clean(string(temp),"volt=", "V")

	return string(usageMem), string(gpuMem), nil
}

//GetCPUTemp Retorna temperatura da CPU
func GetCPUTemp() (string, error) {
	//temp, err := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp").Output()
	temp, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return "", errors.New("Permission Denied")
	}

	st, err := strconv.ParseInt(string(temp[:5]), 10, 64)
	if err != nil {
		return "", errors.New("Error to convert to int")
	}
	cat := float64(st) / float64(1000)

	return strconv.FormatFloat(cat, 'f', 2, 64) + "C", nil
}

//GetLoadAverage Retorna Load Average de 1min
func GetLoadAverage() (string, error) {
	const LOAD_AVERAGE_STRING = "load average:"

	out, err := exec.Command("uptime").Output()
	if err != nil {
		return "", err
	}
	uptimeResult := string(out)

	i := strings.Index(uptimeResult, LOAD_AVERAGE_STRING)
	loadValue := uptimeResult[i+len(LOAD_AVERAGE_STRING):]
	splitsAverage := strings.Split(loadValue, ",")
	oneMinuteLoad := strings.TrimSpace(splitsAverage[0])
	return oneMinuteLoad, nil
}

//GetMemoryUsage Retorna Uso da memória(RAM)
func GetMemoryUsage() (float64, error) {
	//system memory usage
	sysInfo := new(syscall.Sysinfo_t)
	err := syscall.Sysinfo(sysInfo)
	if err != nil {
		return 0.0, err
	}
	all := sysInfo.Totalram
	free := sysInfo.Freeram
	used := all - free

	usedPercent := float64(used) / float64(all)
	return usedPercent, nil
}

//GetDiskUsage Retorna uso do disco
func GetDiskUsage() (float64, float64, float64, error, error, error) {
	getBoot, err0 := getBootDiskUsage()
	getRoot, err1 := getRootDiskUsage()
	getHome, err2 := getHomeDiskUsage()
	return getBoot, getRoot, getHome, err0, err1, err2
}
func getBootDiskUsage() (float64, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/boot", &fs)
	if err != nil {
		return 0.0, err
	}
	all := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	used := all - free

	usedPercent := float64(used) / float64(all)
	return usedPercent, nil
}
func getRootDiskUsage() (float64, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/", &fs)
	if err != nil {
		return 0.0, err
	}
	all := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	used := all - free

	usedPercent := float64(used) / float64(all)
	return usedPercent, nil
}
func getHomeDiskUsage() (float64, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/home", &fs)
	if err != nil {
		return 0.0, err
	}
	all := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	used := all - free

	usedPercent := float64(used) / float64(all)
	return usedPercent, nil
}
func clean(str string, args ...string) string {
	for _, arg := range args {
		str = strings.Replace(str, arg, "", -1)
	}
	str = strings.TrimSpace(str)
	return str
}
