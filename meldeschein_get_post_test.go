package intenik_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMeldescheinGetPost(t *testing.T) {
	req := client.NewMeldescheinGetPostRequest()
	// req.PathParams().GUID = client.ToastRestaurantExternalID()
	req.RequestBody().Meldescheine = []int{12}
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
