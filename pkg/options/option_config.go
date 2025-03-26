package options

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"k8sproxy/pkg/types"
	"log"
	"os"
	"path/filepath"
)

var clientCfg *types.ClientConfig

func LoadClientConfig(path, baseURL string) error {
	// 判断文件是否存在，如果不存在则创建
	if _, err := os.Stat(path); os.IsNotExist(err) {
		clientCfg = &types.ClientConfig{}
	} else {
		bytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if err = yaml.Unmarshal(bytes, &clientCfg); err != nil {
			return err
		}
	}

	if clientCfg.BaseURL == "" && baseURL == "" {
		fmt.Println("请输入BaseURL地址:")
		_, err := fmt.Scan(&baseURL)
		if err != nil {
			return fmt.Errorf("输入读取失败: %w", err)
		}
	}

	if clientCfg.BaseURL != baseURL {
		clientCfg.BaseURL = baseURL

		// 配置信息保存到文件
		configBytes, err := yaml.Marshal(&clientCfg)
		if err != nil {
			return err
		}

		// 1. 解析文件路径，获取目录部分
		dir := filepath.Dir(path)

		// 2. 创建所有缺失的目录（递归创建）
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Printf("创建目录失败: %v", err)
			return err
		}
		return os.WriteFile(path, configBytes, 0644)
	}

	return nil
}

func getBaseURL() string {
	if clientCfg != nil && clientCfg.BaseURL != "" {
		return clientCfg.BaseURL
	}
	return "http://127.0.0.1:8080"
}
