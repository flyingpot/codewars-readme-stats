package main

import (
	"fmt"
	"strings"
)

var codewarsTemplate = `
==============================================
___  __  ____  ____  _  _   __   ____  ____ 
/ __)/  \(    \(  __)/ )( \ / _\ (  _ \/ ___)
( (__(  O )) D ( ) _) \ /\ //    \ )   /\___ \
\___)\__/(____/(____)(_/\_)\_/\_/(__\_)(____/
                        
%s

==============================================
`

var toCenterInfo = `Username: %s Rank: %s`

func render(userInfo *UserInfo) string {
	return fmt.Sprintf(codewarsTemplate, center(
		fmt.Sprintf(toCenterInfo, userInfo.Username, userInfo.Ranks.Overall.Name),
		getWidth(codewarsTemplate)))
}

func center(info string, length int) string {
	if len(info) > length {
		return info
	}
	diff := length - len(info)
	frontIndent := diff / 2
	return strings.Repeat(" ", frontIndent) + info
}

func getWidth(template string) int {
	lineArray := strings.Split(template, "\n")
	max := -1
	for _, s := range lineArray {
		if len(s) > max {
			max = len(s)
		}
	}
	return max
}
