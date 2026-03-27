package core

/*
这个函数主要批量下载提供的软件id列表
*/
import (
	"fmt"
	"os/exec"
)

func Download(softwareIds []string, src string) {
	for _, softwareId := range softwareIds {
		fmt.Println("正在下载软件：", softwareId)
		// 获取软件信息
		args := []string{"download"}
		args = append(args, "--id", softwareId)
		args = append(args, "--download-directory", src)
		args = append(args, "--accept-source-agreements")
		args = append(args, "--accept-package-agreements")
		cmd := exec.Command("winget", args...)
		if out, err := cmd.CombinedOutput(); err != nil {
			fmt.Printf("下载%s失败:%v,输出%s\n", softwareId, err, string(out))
		} else {
			fmt.Printf("下载%s成功:%v,输出%s\n", softwareId, err, string(out))
		}
	}
}
