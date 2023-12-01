package utils

import (
	"os"
	"path"
	"strings"
)

func PathExists(path string) (isExist bool, isDir bool) {
	if path == "" {
		return false, false
	}
	if info, err := os.Stat(path); err == nil {
		return true, info.IsDir()
	}
	return false, false
}

func InsertSuffix(name, suffix string) string {
	if name == "" || suffix == "" {
		return name
	}
	dir, filename := path.Split(name)

	var prefix string
	var ext string
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			prefix = filename[:i]
			ext = filename[i:]
			break
		}
	}
	newName := prefix + suffix + ext
	if dir == "" {
		return newName
	}
	return path.Join(dir, newName)
}

func CurrentExecDir() string {
	p, err := os.Executable()
	if err != nil {
		return ""
	}
	p = strings.ReplaceAll(p, "\\", "/")
	return path.Dir(p)
}

func PathJoin(paths ...string) string {
	if len(paths) == 0 {
		return ""
	}
	var buff strings.Builder
	for _, p := range paths {
		if p == "" {
			continue
		}
		if !strings.HasPrefix(p, "/") {
			buff.WriteString("/")
		}
		if strings.HasSuffix(p, "/") {
			buff.WriteString(p[0 : len(p)-1])
		} else {
			buff.WriteString(p)
		}
	}
	return buff.String()
}
