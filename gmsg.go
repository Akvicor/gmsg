package gmsg

import (
	"github.com/Akvicor/gmsg/gclient"
	"github.com/Akvicor/gmsg/sync_str"
	"time"
)

var (
	AccessToken = sync_str.New("")
	ApiUrl      = sync_str.New("https://msg.example.com")
)

const (
	token       = "x"
	xToken      = "X"
	accessToken = "X-Access-Token"
	loginToken  = "X-Auth-Token"
)

var cli = gclient.NewClient(gclient.WithRetries(3), gclient.WithTimeout(10*time.Second))
