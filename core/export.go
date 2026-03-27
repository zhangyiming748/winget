package core

import (
	"log"
	"os/exec"
	"path/filepath"
)

func Export(root string) {
	f := filepath.Join(root, "export.json")
	args := []string{}
	args = append(args, "export")
	args = append(args, "--output", f)
	cmd := exec.Command("winget", args...)
	log.Printf("导出文件的命令是:%s\n", cmd.String())
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("导出文件失败:%v,输出%s\n", err, string(out))
	} else {
		log.Printf("导出文件成功:%v,输出%s\n", err, string(out))
	}

}
