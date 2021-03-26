package api_test

import (
	"testing"
)

var expectedResult = `
===========================================================
_________            .___                                  
\_   ___ \  ____   __| _/______  _  _______ _______  ______
/    \  \/ /  _ \ / __ |/ __ \ \/ \/ /\__  \\_  __ \/  ___/
\     \___(  <_> ) /_/ \  ___/\     /  / __ \|  | \/\___ \ 
 \______  /\____/\____ |\___  >\/\_/  (____  /__|  /____  >
        \/            \/    \/             \/           \/ 
                        
            Username: test_name Rank: test_rank

===========================================================
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
