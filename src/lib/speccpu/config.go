package speccpu

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

func ModfiyConfigFile(configPath string) {
	// 修改 config 内容，把 copies = 4 改成 copies = 1
	configFile, err := os.OpenFile(configPath, os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Panic("open config file failed", err)
	}
	defer configFile.Close()
	scanner := bufio.NewScanner(configFile)
	var newConfig []string
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`^*copies *= *\d+`)
		if re.MatchString(line) {
			line = "copies = 1"
		}
		newConfig = append(newConfig, line)
	}
	configFile.Truncate(0)
	configFile.Seek(0, 0)
	for _, v := range newConfig {
		configFile.WriteString(v + "\n")
	}
}
