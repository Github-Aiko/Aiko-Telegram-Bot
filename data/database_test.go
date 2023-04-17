package data

import (
	"testing"

	"github.com/Github-Aiko/Aiko-Telegram-Bot/config"
)

func TestNew(t *testing.T) {
	cfg := &config.Config{
		APPs: config.APPs{
			Database: config.Database{
				User: "123",
				Pass: "CnEEyfKF6fTtP6tD",
				IP:   "127.0.0.1",
				Port: 3306,
				Name: "123",
			},
		},
	}
	db := New(cfg)
	if db == nil {
		t.Errorf("Connection failed")
	}

	t.Log("Connection succeeded")

}
