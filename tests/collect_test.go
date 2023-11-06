package tests

import (
	"github.com/stretchr/testify/assert"
	db "github.com/xiaoxuan6/go-package-db"
	"testing"
)

func TestCollect(t *testing.T) {
	db.Init("127.0.0.1", "3306", "root", "root", "go_package_db")
	defer db.Close()
	db.AutoMigrate()
	err := db.Insert(db.Collect{
		Name: "test",
		Url:  "example.com",
	})
	assert.Nil(t, err)

	collect, err := db.FindByName("test")
	assert.Nil(t, err)
	assert.Equal(t, "test", collect.Name)
	assert.Equal(t, "example.com", collect.Url)
}
