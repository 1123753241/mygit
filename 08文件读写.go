package main

/*
import "fmt"
import "os"
*/
import (
	"fmt"
	"os"
)

//写文件
func test_write(fileName string) {
	fout, err := os.Create(fileName) //默认w+
	if err != nil {
		//打开文件失败
		fmt.Println(err)
		return
	}

	//关闭文件
	defer fout.Close() //函数结束时调用

	//写信息到文件中
	for i := 0; i < 10; i++ {
		outstr := fmt.Sprintf("%s:%d\n", "hello go", i)
		fout.WriteString(outstr)
		fout.Write([]byte("abc\n")) //位的方式写入
	}
}

//读文件
func test_read(fileName string) {
	//打开文件
	f, err := os.Open(fileName)
	if err != nil {
		//打开文件失败
		fmt.Println(err)
		return
	}
	//关闭文件
	defer f.Close() //函数结束时调用

	//开辟空间，存储数据
	buf := make([]byte, 1024)
	for {
		num, _ := f.Read(buf) //返回行数和错误码
		if num == 0 {         //==0表示读完
			break
		}

		os.Stdout.Write(buf)
	}

}
func main() {
	//test_write("./test.txt")
	test_read("./test.txt")
}
