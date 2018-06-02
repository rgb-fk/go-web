package config

import (
	"github.com/go-ini/ini"
)

func ReadConfigByKeys(cfgPath string, section string, keys []string) map[string]string {
	result := make(map[string]string, 1)
	cfg, _ := ini.Load(cfgPath)
	sec, _ := cfg.GetSection(section)
	for i := range keys {
		key := keys[i]
		if sec.HasKey(key) {
			result[key] = sec.Key(key).String()
		}
	}
	return result
}

func ReadConfigByKeyInt(cfgPath string, section string, key string) int64 {
	var result int64
	cfg, _ := ini.Load(cfgPath)
	sec, _ := cfg.GetSection(section)
	if sec.HasKey(key) {
		result = sec.Key(key).MustInt64(0)
	}
	return result
}

func ReadConfigByKey(cfgPath string, section string, key string) string {
	result := ""
	cfg, _ := ini.Load(cfgPath)
	sec, _ := cfg.GetSection(section)
	if sec.HasKey(key) {
		result = sec.Key(key).String()
	}
	return result
}
