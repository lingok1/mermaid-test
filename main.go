package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/joshcarp/mermaid-go/mermaid"
	"github.com/mskrha/svg2png"
	"io/ioutil"
	"os"
)

const (
	/*
		Default inkscape binary path
	*/
	BINARY = "D:\\Program Files\\Inkscape\\bin\\inkscape.exe"
)

func main() {

	str := `
%%{init: {align: 'topDown'}}
	graph TD
	A(开始)
	B(发布开标通知)
	C(准备开标材料)
	D(组织开标人员)
	E(验证开标人员身份)
	F(开启开标工具)
	G(导入投标文件)
	H(解密投标文件)
	I(评审投标文件)
	J(生成评审结果)
	K(签署评审文件)
	L(结束)
	M(投标文件合格)
	N(投标文件不合格)
	
	A --> B
	B --> C
	C --> D
	D --> E
	E --> F
	F --> G
	G --> H
	H --> I
	I -->|投标文件合格| J
	    I -->|投标文件不合格| N
	N --> K
	M --> K
	J --> K
	K --> L
	
		`
	//str := `
	//graph TD
	//A(开始)
	//B(发布开标通知)
	//C(准备开标材料)
	//D(组织开标人员)
	//E(验证开标人员身份)
	//F(开启开标工具)
	//G(导入投标文件)
	//H(解密投标文件)
	//I(评审投标文件)
	//J(生成评审结果)
	//K(签署评审文件)
	//L(结束)
	//M(投标文件合格)
	//N(投标文件不合格)
	//
	//classDef topToBottom fill:#f96,stroke:#333,stroke-width:2px;
	//class A,B,C,D,E,F,G,H,I,J,K,L,M,N topToBottom;
	//
	//A --> B
	//B --> C
	//C --> D
	//D --> E
	//E --> F
	//F --> G
	//G --> H
	//H --> I
	//I -->|投标文件合格| J
	//    I -->|投标文件不合格| N
	//N --> K
	//M --> K
	//J --> K
	//K --> L
	//
	//	`

	//	str := `
	//%%{init: {align: 'topDown'}}
	//graph TD
	//    A(开始)
	//    B(发布开标通知)
	//    C(准备开标材料)
	//    D(组织开标人员)
	//    E(验证开标人员身份)
	//    F(开启开标工具)
	//    G(导入投标文件)
	//    H(解密投标文件)
	//    I(评审投标文件)
	//    J(生成评审结果)
	//    K(签署评审文件)
	//    L(结束)
	//
	//    classDef topToBottom fill:#f96,stroke:#333,stroke-width:2px;
	//    class A,B,C,D,E,F,G,H,I,J,K,L topToBottom;
	//
	//    A --> B
	//    B --> C
	//    C --> D
	//    D --> E
	//    E --> F
	//    F --> G
	//    G --> H
	//    H --> I
	//    I --> J
	//    J --> K
	//    K --> L
	//
	//    %% Branching nodes
	//    I1(投标文件符合要求)
	//    I2(投标文件不符合要求)
	//    J1(评审结果通过)
	//    J2(评审结果不通过)
	//
	//    H -->|投标文件符合要求| I1
	//    H -->|投标文件不符合要求| I2
	//    I1 --> I
	//    I2 -->|通知投标人修改| H
	//
	//    I -->|评审结果通过| J1
	//    I -->|评审结果不通过| J2
	//    J1 --> K
	//    J2 -->|通知投标人修改| I
	//
	//`
	svg := mermaid.Execute(str)
	if err := ioutil.WriteFile("mermaid.svg", []byte(svg), 0644); err != nil {
		panic(err)
	}
	//ConvertSVGToPNG([]byte(svg))
	//将svg string类型转为 io.Reader类型
	file, err2 := os.Open("mermaid.svg")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer file.Close()

	// 通过文件创建Scanner
	scanner := bufio.NewScanner(file)
	line := ""
	// 循环迭代文件的每一行
	for scanner.Scan() {
		line = scanner.Text()
		// 在这里对每一行进行处理
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时出错:", err)
	}

	reader := bytes.NewReader([]byte(svg))
	png, err := SVG2PNG(reader)
	if err != nil {
		fmt.Println(err)
	}
	if png == nil {
		fmt.Println("png is nil")
		return
	}
	fmt.Println("png:", png)
	all, err := ioutil.ReadAll(png)
	if err != nil {
		fmt.Println(err)
	}
	if err := ioutil.WriteFile("mermaid.png", all, 0644); err != nil {
		panic(err)
	}

}
func ConvertSVGToPNG(input []byte) error {

	converter := svg2png.New()
	converter.SetBinary(BINARY)

	output, err := converter.Convert(input)
	if err != nil {
		fmt.Println(err)
	}
	if err := ioutil.WriteFile("mermaid.png", output, 0644); err != nil {
		panic(err)
	}
	return nil
}
