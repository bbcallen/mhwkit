package mhw

import (
	"fmt"
	"testing"
)

func TestLoadDataManager(t *testing.T) {
	dm := loadDataManager()
	if testing.Verbose() {
		fmt.Println("dataManager:", dm)
	}
}
