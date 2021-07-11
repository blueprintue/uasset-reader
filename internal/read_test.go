package internal

import (
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("\n" + "error on init TestRead")
	}

	type dataProvider struct {
		Name     string
		Filepath string
		Result   error
	}

	dataCases := []dataProvider{
		{
			Name:     "open file",
			Filepath: dir + "/../test/4.26/BP_BaseObject_empty.uasset",
			Result:   nil,
		},
	}

	for _, dataCase := range dataCases {
		if ReadUasset(dataCase.Filepath) != dataCase.Result {
			t.Errorf("\n"+"error on case: %s", dataCase.Name)
		}
	}
}
