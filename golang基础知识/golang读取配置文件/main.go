package golang读取配置文件
/*
import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	_ "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

1. json使用
JSON 应该比较熟悉，它是一种轻量级的数据交换格式。层次结构简洁清晰 ，易于阅读和编写，同时也易于机器解析和生成。

　　1. 创建 conf.json：

{
"enabled": true,
"path": "/usr/local"
}


　　2. 新建config_json.go：

package main

import (
"encoding/json"
"fmt"
"os"
)

type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	// 打开文件
	file, _ := os.Open("conf.json")

	// 关闭文件
	defer file.Close()

	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)

	conf := configuration{}
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("path:" + conf.Path)
}


　　启动运行后，输出如下：

D:\Go_Path\go\src\configmgr>go run config_json.go
path:/usr/local

##################################################################################################
2. ini的使用
INI文件格式是某些平台或软件上的配置文件的非正式标准，由节(section)和键(key)构成，比较常用于微软Windows操作系统中。这种配置文件的文件扩展名为INI。

　　1. 创建 conf.ini：

[Section]
enabled = true
path = /usr/local # another comment
　　2.下载第三方库：go get gopkg.in/gcfg.v1

　　3. 新建 config_ini.go：

package main

import (
"fmt"
gcfg "gopkg.in/gcfg.v1"
)

func main() {
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}

	err := gcfg.ReadFileInto(&config, "conf.ini")

	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
}


　　启动运行后，输出如下：

D:\Go_Path\go\src\configmgr>go run config_ini.go
true
/usr/local

##############################################################################################
3. yaml使用
yaml 可能比较陌生一点，但是最近却越来越流行。也就是一种标记语言。层次结构也特别简洁清晰 ，易于阅读和编写，同时也易于机器解析和生成。

golang的标准库中暂时没有给我们提供操作yaml的标准库，但是github上有很多优秀的第三方库开源给我们使用。

　　1. 创建 conf.yaml：

enabled: true
path: /usr/local
　　2. 下载第三方库：go get  gopkg.in/yaml.v2

　　3. 创建 config_yaml.go：

package main

import (
"fmt"
"io/ioutil"
"log"

"gopkg.in/yaml.v2"
)

type conf struct {
	Enabled bool   `yaml:"enabled"` //yaml：yaml格式 enabled：属性的为enabled
	Path    string `yaml:"path"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func main() {
	var c conf
	c.getConf()
	fmt.Println("path:" + c.Path)
}


　　启动运行后，输出如下：

D:\Go_Path\go\src\configmgr>go run config_yaml.go
path:/usr/local

####################################################################################
4. 读取.conf文件。类似读取ini文件
package main

import (
"fmt"
gcfg "gopkg.in/gcfg.v1"
)

func main() {
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}

	err := gcfg.ReadFileInto(&config, "conf.conf")

	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
}


　　启动运行后，输出如下：

D:\Go_Path\go\src\configmgr>go run config_ini.go
true
/usr/local

####################################################################################
5. 读其他常规文件，如.data文件，看里面是什么格式，有些.data文件里面是json格式，就可以使用json来读
file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
// decode the file
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(conf); err != nil {
		return "", err
	}

	// check config
	if err := HostTableConfCheck(*conf); err != nil {
		return "", err
	}

	return *(conf.Version), nil

##############################################################################
6.读其他常规的文本文件

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFileByte(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	//指定读的长度
	var tmp = make([]byte, 100)
	n, err := file.Read(tmp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(tmp[:n]))

}

func readFileLine(path string) {
	fi, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer func() {
		err := fi.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a))
	}

或者
scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Remove all leading and trailing spaces and tabs
		line := strings.Trim(scanner.Text(), " \t")
		//Line begins with "#" is considered as a comment
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}

		// Check line format
		startIP, endIP, err = checkLine(line)
		if err != nil {
			return nil, fmt.Errorf("checkLine(): line[%s] err[%s]", line, err.Error())
		}

		// insert start ip and end ip into dict
		if startIP.Equal(endIP) {
			singleIPCounter += 1
		} else {
			pairIPCounter += 1
		}
	}
	err = scanner.Err()
}

func readFileAll(path string) {
	fileObj, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(fileObj))
}


func main() {
	path := "yourpaht/xxxx.txt"
	// 指定字节读
	readFileByte(path)
	// 逐行读取文件
	readFileLine(path)
	// 全部读取
	readFileAll(path)
}
*/
