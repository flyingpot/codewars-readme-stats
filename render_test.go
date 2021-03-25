package main

import (
	"testing"
)

var expectedResult = `
==============================================
___  __  ____  ____  _  _   __   ____  ____ 
/ __)/  \(    \(  __)/ )( \ / _\ (  _ \/ ___)
( (__(  O )) D ( ) _) \ /\ //    \ )   /\___ \
\___)\__/(____/(____)(_/\_)\_/\_/(__\_)(____/
                        
     Username: test_name Rank: test_rank

==============================================
`

func TestRender(t *testing.T) {
	var mockInfo UserInfo
	mockInfo.Username = "test_name"
	mockInfo.Ranks.Overall.Name = "test_rank"
	result := render(&mockInfo)
	if expectedResult != result {
		t.Errorf("%s to be expected, but got %s", expectedResult, result)
	}
}
