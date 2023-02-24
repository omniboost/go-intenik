package intenik_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMeldescheinBuchenPost(t *testing.T) {
	req := client.NewMeldescheinBuchenPostRequest()
	// req.PathParams().GUID = client.ToastRestaurantExternalID()
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
