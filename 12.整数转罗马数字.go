package main

var luomap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func intToRoman(num int) string {
	res := ""
	convertList := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for i := 0; i < len(convertList); i++ {
		count := num / convertList[i]
		num = num - count * convertList[i]
		for count != 0 {
			res = res + luomap[convertList[i]]
			count--
		}
	}
	return res
}

func main(){

}


