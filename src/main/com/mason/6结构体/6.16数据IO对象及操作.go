package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("----------Reader----------")
	fmt.Println("testRead")
	testRead()
	fmt.Println("testReadByte")
	testReadByte()
	fmt.Println("testReadBytes")
	testReadBytes()
	fmt.Println("testReadLine")
	testReadLine()
	fmt.Println("testReadRune")
	testReadRune()
	fmt.Println("testReadSlice")
	testReadSlice()
	fmt.Println("testReadString")
	testReadString()
	fmt.Println("testReaderBuffered")
	testReaderBuffered()
	fmt.Println("testPeek")
	testPeek()

	fmt.Println("----------Writer----------")
	fmt.Println("testAvailable")
	testAvailable()
	fmt.Println("testWriterBuffered")
	testWriterBuffered()
	fmt.Println("testFlush")
	testFlush()
	fmt.Println("testWrite")
	testWrite()
	fmt.Println("testWriteByte")
	testWriteByte()
	fmt.Println("testWriteRune")
	testWriteRune()
	fmt.Println("testWriteString")
	testWriteString()
}

func testWriteString() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	s := "C语言中文网"
	n, err := w.WriteString(s)
	w.Flush()
	fmt.Println(string(wr.Bytes()), n, err)
}

func testWriteRune() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	var r rune = 'G'
	size, err := w.WriteRune(r)
	w.Flush()
	fmt.Println(string(wr.Bytes()), size, err)
}

func testWriteByte() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	var c byte = 'G'
	err := w.WriteByte(c)
	w.Flush()
	fmt.Println(string(wr.Bytes()), err)
}

func testWrite() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	p := []byte("C语言中文网")
	n, err := w.Write(p)
	w.Flush()
	fmt.Println(string(wr.Bytes()), n, err)
}

func testFlush() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	p := []byte("C语言中文网")
	w.Write(p)
	fmt.Printf("未执行 Flush 缓冲区输出 %q\n", string(wr.Bytes()))
	w.Flush()
	fmt.Printf("执行 Flush 后缓冲区输出 %q\n", string(wr.Bytes()))
}

func testWriterBuffered() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	p := []byte("C语言中文网")
	fmt.Println("写入前未使用的缓冲区为：", w.Buffered())
	w.Write(p)
	fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(p), w.Buffered())
	w.Flush()
	fmt.Println("执行 Flush 方法后，写入的字节数为：", w.Buffered())
}

func testAvailable() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	p := []byte("C语言中文网")
	fmt.Println("写入前未使用的缓冲区为：", w.Available())
	w.Write(p)
	fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(p), w.Available())
}

func testPeek() {
	data := []byte("Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	bl, err := r.Peek(8)
	fmt.Println(string(bl), err)
	bl, err = r.Peek(14)
	fmt.Println(string(bl), err)
	bl, err = r.Peek(20)
	fmt.Println(string(bl), err)
}

func testReaderBuffered() {
	data := []byte("Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [14]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
	rn := r.Buffered()
	fmt.Println(rn)
	n, err = r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
	rn = r.Buffered()
	fmt.Println(rn)
}

func testReadString() {
	data := []byte("C语言中文网, Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadString(delim)
	fmt.Println(line, err)
}

func testRead() {
	data := []byte("C语言中文网")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
}

func testReadByte() {
	data := []byte("C语言中文网")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	c, err := r.ReadByte()
	fmt.Println(string(c), err)
}

func testReadBytes() {
	data := []byte("C语言中文网，Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadBytes(delim)
	fmt.Println(string(line), err)
}

func testReadLine() {
	data := []byte("Golang is a beautiful language. \r\n I like it!")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	line, prefix, err := r.ReadLine()
	fmt.Println(string(line), prefix, err)
}

func testReadRune() {
	data := []byte("C语言中文网")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	ch, size, err := r.ReadRune()
	fmt.Println(string(ch), size, err)
}

func testReadSlice() {
	data := []byte("C语言中文网, Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadSlice(delim)
	fmt.Println(string(line), err)
	line, err = r.ReadSlice(delim)
	fmt.Println(string(line), err)
	line, err = r.ReadSlice(delim)
	fmt.Println(string(line), err)
}
