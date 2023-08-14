package myuuid

import (
	"testing"
)

func TestGetUUID(t *testing.T) {
	t.Run("uuid", func(t *testing.T) {
		// Parsing UUID from string input
		t.Log("gen uuid:", UUID())
	})

	t.Run("machine id", func(t *testing.T) {
		if machID, err := MachineId(); err != nil {
			t.Errorf("machine id err:%v", err)
		} else {
			t.Log("machine id :", machID)
		}
	})

	t.Run("machine id encode", func(t *testing.T) {
		if machID, err := MachineIdWith("luis1.jin"); err != nil {
			t.Error("machine encode id err:", err)
		} else {
			t.Log("machine encode id :", machID)
		}
	})
}
