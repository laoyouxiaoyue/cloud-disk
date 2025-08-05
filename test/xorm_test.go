package test

import (
	"cloud-disk/core/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)
import "xorm.io/xorm"

func TestXoraTest(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/cloud-disk?charset=utf8")
	assert.Equal(t, err, nil)
	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	assert.Equal(t, err, nil)
	fmt.Printf("%v", data[0])
}
