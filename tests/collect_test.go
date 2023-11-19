package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	db "github.com/xiaoxuan6/go-package-db"
	"testing"
)

func TestCollect(t *testing.T) {
	db.Init("127.0.0.1", "3306", "root", "root", "go_package_db")
	defer db.Close()
	db.AutoMigrate()

	err := db.Insert(db.Collect{
		Name:     "test",
		Url:      "example.com",
		Language: "go",
	})
	assert.Nil(t, err)

	collect, err := db.FindByUrl("example.com")
	assert.Nil(t, err)
	assert.Equal(t, "test", collect.Name)
	assert.Equal(t, "example.com", collect.Url)

	collects, err := db.FindByName("test")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(collects))
	assert.Equal(t, "test", collects[0].Name)

	err = db.DeleteAll()
	assert.Nil(t, err)

	err = db.Insert(db.Collect{
		Name:     "test",
		Url:      "example.com",
		Language: "go",
	})
	assert.Nil(t, err)

	err = db.Insert(db.Collect{
		Name:     "test",
		Url:      "example.com",
		Language: "go",
	})
	assert.Nil(t, err, fmt.Sprintf("重复插入数据，错误为：%s", err.Error()))
}
