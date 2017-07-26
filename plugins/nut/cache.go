package nut

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

// CPut cache put
func CPut(key string, val interface{}, ttl time.Duration) {
	if err := cacheBm.Put(key, val, ttl); err != nil {
		beego.Error(err)
	}
}

// CGet cache get
func CGet(key string) interface{} {
	return cacheBm.Get(key)
}

var cacheBm cache.Cache

func init() {
	var err error
	if cacheBm, err = cache.
		NewCache(
			"redis",
			beego.AppConfig.String("cacheproviderconfig"),
		); err != nil {
		beego.Error(err)
	}
}
