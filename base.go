package myuuid

import (
	"github.com/denisbrodbeck/machineid"
	uuid "github.com/satori/go.uuid"
	"strings"
)

func UUID() string {
	uuidFunc := uuid.NewV4()
	uuidStr := uuidFunc.String()
	uuidStr = strings.Replace(uuidStr, "-", "", -1)
	return uuidStr
}

func MachineId() (string, error) {
	return machineid.ID()
}

func MachineIdWith(appName string) (string, error) {
	return machineid.ProtectedID(appName)
}
