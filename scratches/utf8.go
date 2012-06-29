package main

import "fmt"

func КНЦ(s string)          { fmt.Println("КНЦ, ЕПТ") }
func ВЫВОД(s string) string { fmt.Println(s); return s }
func ЕСЛИ(x int, s string, y int, to, last string, to2, last2 string) string {
	if x < y {
		return last
	} else {
		return last2
	}
	fmt.Println(last)
	return "ШТО"
}

func ЕСЛИИ(x, s, y, to, last, to2, last2 string) string {
	if x < y {
		return last
	} else {
		return last2
	}
	fmt.Println(last)
	return "ШТО"
}

func main() {
	ЗБС := "СВСЕМ ЗБС!"
	ССЗБ := "СВСЕМ ССЗБ!"
	НХЙ := "ПШЕЛ НХУЙ НРКМАН!"
//	КНЦ(ВЫВОД(ЕСЛИ(2, "МЕНЬШЕ", 1, "ТО", ЗБС, "ИНАЧЕ", НХЙ)))
	КНЦ(ЕСЛИИ((ВЫВОД(ЕСЛИ(2, "МЕНЬШЕ", 1, "ТО", ЗБС, "ИНАЧЕ", НХЙ))), "БОЛЬШЕ", (ВЫВОД(ЕСЛИ(2, "МЕНЬШЕ", 1, "ТО", ЗБС, "ИНАЧЕ", НХЙ))), "ТО", ССЗБ, "ИНАЧЕ", НХЙ))
	//	ВЫВОД(ЗБС)
}

//printf("IMAGE:base65 code")
