package addPreview

import (
	"strings"
)

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func ToMiniCopy(imgSrc string) string {
	var sb strings.Builder
	reversed := Reverse(imgSrc)
	reversedAdditional := Reverse(string("_preview"))
	point := rune('.')
	for _, symbol := range reversed {
		sb.WriteString(string(symbol))
		if (symbol == point){
			sb.WriteString(reversedAdditional)
		}
	}
	return Reverse(sb.String())
}
