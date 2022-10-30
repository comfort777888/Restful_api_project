package user

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"rest_api_project/pkg/users"
	"testing"
)

func TestCRUD(t *testing.T) {
	//Test Create
	u1 := users.User{ID: 555, Data: "test_data"}
	bb, err := json.Marshal(u1)
	if err != nil {
		log.Fatalf("errrr:%v", err)
	}
	resp, err := http.Post("http://localhost:8080/user/555", "application/json", bytes.NewBuffer(bb))
	if err != nil {
		log.Fatalf("error in post - %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("error not the same status code")
	}

	//Test Read(Get)
	var u2 users.User
	resp2, err := http.Get("http://localhost:8080/user/555")
	if err != nil {
		log.Fatalf("error in get:%v", err)
	}

	if err = json.NewDecoder(resp2.Body).Decode(&u2); err != nil {
		log.Fatalf("error in read(get) unmarshal: %v", err)
	}

	if u1.ID != u2.ID || u1.Data != u2.Data {
		log.Fatalf("error in read(get): fields doesn't match")
	}

	//Test Delete
	req, err := http.NewRequest("DELETE", "http://localhost:8080/user/555", nil)
	if err != nil {
		log.Fatalf("error in deleting in new request - %v", err)
	}

	resp3, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error in deleting: client do - %v", err)
	}

	if resp3.StatusCode != http.StatusNoContent {
		log.Fatalf("error in deleting: status code doesn't match")
	}

	//TestUpdate
	u3 := users.User{ID: 188, Data: "test_update"}
	bb2, err := json.Marshal(u3)
	if err != nil {
		log.Fatalf("error in updating - marshal:%v", err)
	}
	resp4, err := http.NewRequest("PUT", "http://localhost:8080/user/188", bytes.NewBuffer(bb2))
	if err != nil {
		log.Fatalf("error in updating in post - %v", err)
	}

	respupd, err := http.DefaultClient.Do(resp4)
	if err != nil {
		log.Fatalf("error in updating in client do - %v", err)
	}

	if respupd.StatusCode != http.StatusOK {
		t.Errorf("error not the same status code in update")
	}

}
