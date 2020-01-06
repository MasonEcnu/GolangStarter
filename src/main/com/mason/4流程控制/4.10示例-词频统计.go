package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// 统计出一个文件里单词出现的频率
func main() {
	// 查看帮助文档
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "-help" {
		fmt.Printf("usage: %s <file1> [<file2> [... <fileN>]]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	// 从标准输入读取数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input the min frequency to show:")
	// 读取数据
	input, err := inputReader.ReadString('\n')
	if err != nil {
		goto exitWhenErrorOccurred
	} else {
		// 用切片操作删除最后的\n
		num, err := strconv.Atoi(input[:len(input)-1])
		if err != nil {
			num = 1
		}
		// 声明词频统计map
		frequencyForWord := make(map[string]int)
		for _, filename := range commandLineFiles(os.Args[1:]) {
			updateFrequencies(filename, frequencyForWord, num)
		}

		// 展示
		reportByWords(frequencyForWord)

		//reportByFrequency(invertStringIntMap(frequencyForWord))

	}

exitWhenErrorOccurred:
	fmt.Printf("An error occurred:%s\n", err)
	// 异常退出
	os.Exit(1)
}

func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name) // 无效模式
			} else {
				args = append(args, matches...)
			}
		}

		return args
	}
	return files
}

func updateFrequencies(filename string, frequencyForWord map[string]int, num int) {
	var file *os.File
	var err error
	if file, err = os.Open(filename); err != nil {
		log.Println("failed to open the file:", err)
		return
	}
	defer file.Close()
	readAndUpdateFrequencies(bufio.NewReader(file), frequencyForWord)

	// 删掉词频为1的单词
	needDelete := make([]string, 0)
	for key, value := range frequencyForWord {
		if value <= num {
			needDelete = append(needDelete, key)
		}
	}
	for _, value := range needDelete {
		delete(frequencyForWord, value)
	}
}

func readAndUpdateFrequencies(reader *bufio.Reader, frequencyForWord map[string]int) {
	for {
		line, err := reader.ReadString('\n')
		for _, word := range splitOnNonLetters(strings.TrimSpace(line)) {
			if len(word) > utf8.UTFMax || utf8.RuneCountInString(word) > 1 {
				frequencyForWord[strings.ToLower(word)] += 1
			}
		}

		if err != nil {
			if err != io.EOF {
				log.Println("failed to finish reading the file:", err)
			}
			break
		}
	}
}

func splitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	// todo 研究下strings.FieldsFunc
	return strings.FieldsFunc(s, notALetter)
}

// 同一词频的单词有多少个
func invertStringIntMap(intForString map[string]int) map[int][]string {
	stringsForInt := make(map[int][]string, len(intForString))
	for key, value := range intForString {
		stringsForInt[value] = append(stringsForInt[value], key)
	}
	return stringsForInt
}

func reportByWords(frequencyForWord map[string]int) {
	words := make([]string, 0, len(frequencyForWord))
	wordWidth, frequencyWidth := 0, 0
	for word, frequency := range frequencyForWord {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}

	sort.Strings(words)
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	// %*s 可以实现的打印固定长度的空白
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth, frequencyForWord[word])
	}
}

func reportByFrequency(wordsForFrequency map[int][]string) {
	frequencies := make([]int, 0, len(wordsForFrequency))
	for frequency := range wordsForFrequency {
		frequencies = append(frequencies, frequency)
	}

	sort.Ints(frequencies)

	width := len(fmt.Sprint(frequencies[len(frequencies)-1]))
	fmt.Println("Frequency-->Words")
	for _, frequency := range frequencies {
		words := wordsForFrequency[frequency]
		sort.Strings(words)
		fmt.Printf("%*d %s\n", width, frequency, strings.Join(words, ", "))
	}
}
