package main

import (
"bufio"
"fmt"
"os"
"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 读取命令数量
	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	// 初始化路径为根目录，并记录上一次目录（如果有 cd - 的需求）
	currentPath := "/"
	lastPath := "/"

	for i := 0; i < n; i++ {
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if strings.HasPrefix(command, "cd") {
			pathArg := strings.TrimSpace(command[2:])

			// 如果没有参数，则进入 /home/user
			if pathArg == "" {
				lastPath = currentPath
				currentPath = "/home/user"
				continue
			}

			// 处理 `cd -` 的情况
			if pathArg == "-" {
				currentPath, lastPath = lastPath, currentPath
				continue
			}

			// 保存当前路径以备后续切换回来
			lastPath = currentPath

			if strings.HasPrefix(pathArg, "/") {
				// 绝对路径
				currentPath = pathArg
			} else {
				// 相对路径
				if currentPath != "/" {
					currentPath += "/"
				}
				currentPath += pathArg
			}

			// 规范化路径，处理 ".." 和 "." 以及多余的 "/"
			pathParts := strings.Split(currentPath, "/")
			var stack []string

			for _, part := range pathParts {
				if part == "" || part == "." {
					// 忽略空部分或当前目录标记
					continue
				} else if part == ".." {
					// 返回上级目录
					if len(stack) > 0 {
						// 如果当前路径是 `/home/user`，返回到 `/home`
						if len(stack) == 2 && stack[0] == "home" && stack[1] == "user" {
							stack = stack[:len(stack)-1]
						} else {
							stack = stack[:len(stack)-1]
						}
					}
				} else {
					// 添加新目录部分
					stack = append(stack, part)
				}
			}

			// 重新构建规范化路径
			if len(stack) == 0 {
				currentPath = "/"
			} else {
				currentPath = "/" + strings.Join(stack, "/")
			}
		}
	}

	// 输出最终路径
	fmt.Println(currentPath)
}
