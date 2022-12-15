package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
	// "strings"
)

// go run main >> res.txt
func main() {
	arrfile := readFile("J:\\king\\file")
	arr := []string{}
	const key string = "v6ygEXyS.hncPruTEtojdl4vluGVjdwsemnEn7oI2"
	for _, val := range arrfile {
		arr = append(arr, fmt.Sprintf("java -jar J:\\king\\lfc-windows-slower.jar c J:\\king\\file\\%s J:\\king\\dealfile\\%s.slower slower %s 1g", val, val, key))
	}
	// 因为要输出 -> res.txt中，且为了保证不会抢占压缩工具的需要的cpu时间片，就不开协程了，全改为同步执行
	for _, val := range arr {
		fmt.Println(val)
		// command(val)
		// res1 := strings.Split(val, " ")
		// statFile(res1[5])
	}
}

func command(cmd string) error {
	startT := time.Now()
	c := exec.Command("cmd.exe", "/c", cmd)
	tc := time.Since(startT) //计算耗时
	fmt.Printf("time cost = %v\n", tc)
	output, err := c.CombinedOutput()
	fmt.Println(string(output))
	return err
}

func readFile(path string) []string {
	arr := []string{}
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		arr = append(arr, f.Name())
	}
	return arr
}

func statFile(fileName string) {
	info, err := os.Stat(fileName) //Stat获取文件属性
	if err != nil {
		fmt.Println("os.Stat err =", err)
		return
	}
	fmt.Println("name =", info.Name())
	fmt.Println("size =", float64(info.Size())/1024/1024, " MB")
}
