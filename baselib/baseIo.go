package baselib

import (
	"fmt"
	"io"
	"os"
)

//IO操作
func TestIo() {
	fmt.Println("########测试常见包IO")
	//var buf [16]byte
	//os.Stdin.Read(buf[:])
	//os.Stdin.WriteString(string(buf[:]))
	fmt.Println("创建并写入文件")
	file, err := os.Create("./hong.txt")
	if err != nil {
		fmt.Println("创建文件失败")
		return
	}
	defer file.Close()
	file.WriteString("hello world\n")
	file.Write([]byte("早上好\n"))

	fmt.Println("读文件")

	fd, err := os.Open("./hong.txt")
	var buf [128]byte
	var content []byte
	for {
		n, err := fd.Read(buf[:])
		if err == io.EOF {
			break
		}
		content = append(content, buf[:n]...)
	}
	defer fd.Close()
	fmt.Println(string(content))
}
