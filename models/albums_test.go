package models

import (
	"testing"

	"github.com/alochym01/gin-website/config"
	_ "github.com/mattn/go-sqlite3"
)

func TestAlbum(t *testing.T) {
	config.DB = config.SqliteConn("../foo.db")
	defer config.DB.Close()

	PreparesqliteDB(config.DB)

	t.Run("Test Album Get Method", func(t *testing.T) {
		_, err := Album{}.Get()

		if err != nil {
			t.Errorf("Create function fail with err: %s", err.Error())
		}
	})

	t.Run("Test Album GetByID Method", func(t *testing.T) {
		album, err := Album{}.GetByID("1")

		if err != nil {
			t.Errorf("GetByID function fail with err: %s", err.Error())
		}

		// Check struct is empty
		if (album == Album{}) {
			t.Errorf("GetByID function fail with err: %s", "Not Found")
		}
	})

	t.Run("Test Album Create Method", func(t *testing.T) {
		err := Album{}.Create("Kubernetes with Rancher", "Alex Cuckooo", 7.99)

		if err != nil {
			t.Errorf("Create function fail with err: %s", err.Error())
		}
	})

	t.Run("Test Album Update Method", func(t *testing.T) {
		err := Album{}.Update("Kubernetes with Rancher", "Alex Cuckooo", 7.99, "1")

		if err != nil {
			t.Errorf("Update function fail with err: %s", err.Error())
		}
	})

	t.Run("Test Album Delete Method", func(t *testing.T) {
		err := Album{}.Delete("1")

		if err != nil {
			t.Errorf("Delete function fail with err: %s", err.Error())
		}
	})
}
