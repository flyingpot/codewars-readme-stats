package main

import (
	"fmt"
	"strings"
)

var codewarsTemplate = `
===========================================================
_________            .___                                  
\_   ___ \  ____   __| _/______  _  _______ _______  ______
/    \  \/ /  _ \ / __ |/ __ \ \/ \/ /\__  \\_  __ \/  ___/
\     \___(  <_> ) /_/ \  ___/\     /  / __ \|  | \/\___ \ 
 \______  /\____/\____ |\___  >\/\_/  (____  /__|  /____  >
        \/            \/    \/             \/           \/ 
                        
%s

===========================================================
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
