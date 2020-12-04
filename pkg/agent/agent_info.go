package agent

import (
	"encoding/json"
	"github.com/Klevry/klevr/pkg/common"
	"github.com/NexClipper/logger"
	"github.com/mackerelio/go-osstat/memory"
	"io/ioutil"
	netutil "k8s.io/apimachinery/pkg/util/net"
	"log"
	"runtime"
	"syscall"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// get local ip address
func Local_ip_add() string {
	// get Local IP address automatically
	default_ip, err := netutil.ChooseHostInterface()
	if err != nil {
		log.Fatalf("Failed to get IP address: %v", err)
	}

	return default_ip.String()
}

// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

// send agent info to manager
func SendMe(body *common.Body) {
	body.Me.IP = Local_ip_add()
	body.Me.Port = 18800
	body.Me.Version = "0.1.0"

	disk := DiskUsage("/")

	memory, _ := memory.Get()

	body.Me.Resource.Core = runtime.NumCPU()
	body.Me.Resource.Memory = int(memory.Total / MB)
	body.Me.Resource.Disk = int(disk.All / MB)
}

func JsonMarshal(a interface{}) []byte{
	b, err := json.Marshal(a)
	if err != nil{
		logger.Debugf("%v", string(b))
		logger.Error(err)
	}

	return b
}

func JsonUnmarshal(a []byte) common.Body{
	var body common.Body

	err := json.Unmarshal(a, &body)
	if err != nil{
		logger.Debugf("%v", string(a))
		logger.Error(err)
	}

	return body
}

func ReadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Error(err)
	}

	return data
}