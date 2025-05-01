package exsort

import (
	"bufio"
	"fmt"
	"os"
)

func readFromFile(file string, cacheSize int) {
	// 打开文件
	fd, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// 确保文件在函数结束时关闭
	defer fd.Close()

	// 创建一个带缓冲的读取器，设置缓冲区大小为 4KB
	reader := bufio.NewReaderSize(fd, 4096)
	buffer := make([]byte, 4096)
	for {
		// 从文件中读取数据到缓冲区
		n, err := reader.Read(buffer)
		if n > 0 {
			// 处理读取到的数据
			fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
		}
		if err != nil {
			break
		}
	}
}

func writeToFile(file string, cacheSize int) {
	// 创建新文件
	fd, err := os.Create(file)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// 确保文件在函数结束时关闭
	defer fd.Close()

	// 创建一个带缓冲的写入器，设置缓冲区大小为 4KB
	writer := bufio.NewWriterSize(fd, 4096)
	// 模拟大文件数据，分块写入
	for i := 0; i < 1000; i++ {
		data := []byte("This is a block of data.\n")
		// 将数据写入缓冲区
		_, err := writer.Write(data)
		if err != nil {
			fmt.Println("Error writing to buffer:", err)
			return
		}
		// 手动刷新缓冲区，将数据写入文件
		if i%100 == 0 {
			err = writer.Flush()
			if err != nil {
				fmt.Println("Error flushing buffer:", err)
				return
			}
		}
	}
	// 最后一次刷新缓冲区，确保所有数据都写入文件
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
		return
	}
	fmt.Println("Data written to file successfully.")
}
