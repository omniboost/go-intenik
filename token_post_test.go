package intenik_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTokenPost(t *testing.T) {
	req := client.NewTokenPostRequest()
	req.RequestBody().Username = client.Username()
	req.RequestBody().Password = client.Password()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
