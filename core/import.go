package core

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func Import(src string, download bool) {
	log.Println("Importing...")
	ids, err := Parse(src)
	if err != nil {
		log.Fatalf("解析失败：%v", err)
	}
	for i, id := range ids {
		log.Printf("包 %d: %s\n", i, id)
		args := []string{"install"}
		args = append(args, "--id", id)
		args = append(args, "--accept-package-agreements")
		args = append(args, "--accept-source-agreements")
		cmd := exec.Command("winget", args...)
		log.Printf("安装命令:%s\n", cmd.String())
		if out, err := cmd.CombinedOutput(); err != nil {
			log.Printf("安装%s失败:%v\t输出:%s\n", id, err, string(out))
		} else {
			log.Printf("安装%s成功\t输出:%s\n", id, string(out))
		}
	}
	if download {
		home, _ := os.UserHomeDir()
		dir := filepath.Join(home, "Downloads")
		Download(ids, dir)
	}
}

// Root represents the top-level structure of the exported winget JSON.
type Root struct {
	Schema        string    `json:"$schema"`
	CreationDate  time.Time `json:"CreationDate"`
	Sources       []Source  `json:"Sources"`
	WinGetVersion string    `json:"WinGetVersion"`
}

type Source struct {
	Packages      []Package     `json:"Packages"`
	SourceDetails SourceDetails `json:"SourceDetails"`
}

type Package struct {
	PackageIdentifier string `json:"PackageIdentifier"`
}

type SourceDetails struct {
	Argument   string `json:"Argument"`
	Identifier string `json:"Identifier"`
	Name       string `json:"Name"`
	Type       string `json:"Type"`
}

func Parse(fp string) ([]string, error) {
	var root Root

	// 读取文件内容
	data, err := os.ReadFile(fp)
	if err != nil {
		return nil, fmt.Errorf("读取文件失败：%w", err)
	}

	// 解析 JSON 到 Root 结构体
	if err := json.Unmarshal(data, &root); err != nil {
		return nil, fmt.Errorf("解析 JSON 失败：%w", err)
	}
	// 在这里提取出结构体每一个包的PackageIdentifier并打印出来
	var ids []string
	for _, source := range root.Sources {
		for i, pkg := range source.Packages {
			fmt.Printf("包 %d: %s\n", i+1, pkg.PackageIdentifier)
			ids = append(ids, pkg.PackageIdentifier)
		}
	}
	return ids, nil
}
