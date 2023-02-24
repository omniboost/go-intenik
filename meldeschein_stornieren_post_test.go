package intenik_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMeldescheinStornierenPost(t *testing.T) {
	req := client.NewMeldescheinStornierenPostRequest()
	// req.PathParams().GUID = client.ToastRestaurantExternalID()
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
