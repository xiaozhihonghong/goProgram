package main
//电话号码的字母组合
//方法，回溯或者递归
import "fmt"

var hashMap = map[byte][]string{
	'2':{"a", "b", "c"},
	'3':{"d", "e", "f"},
	'4':{"g", "h", "i"},
	'5':{"j", "k", "l"},
	'6':{"m", "n", "o"},
	'7':{"q", "p", "r", "s"},
	'8':{"t", "u", "v"},
	'9':{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	getLetterCombinations(digits, 0, "")
	return res
}

var res []string
func getLetterCombinations(digits string, index int, tmp string) {
	if len(digits) == index {
		//temp := make([]string, len(res))
		//copy(temp, res)
		res = append(res, tmp)
		return
	}
	digit := digits[index]
	chs := hashMap[digit]
	for j:=0;j<len(chs);j++ {
		tmp += chs[j]
		getLetterCombinations(digits, index+1, tmp)   //这种方法有问题，先将tmp字符串连接好，
		//后面return之后出来的是已经连接的字符串，这种就是回溯的解法，需要加上一个 tmp = tmp[:len(tmp)-1]
		tmp = tmp[:len(tmp)-1]
		//getLetterCombinations(digits, index+1, tmp+chs[j]) //这种return之后仍然是没有连接的tmp
	}
}

func main() {
	digit := "23"
	result := letterCombinations(digit)
	fmt.Println(result)
}
