package main

import (
  "net/http"
  "log"
  "encoding/json"
  "io/ioutil"
  "io"
  "bytes"
)

type LoginRequest struct {
  Username string `json:"u"`
  Password string `json:"p"`
  Token    bool   `json:"t"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
  loginRequest := &LoginRequest{}
  log.Println(loginRequest)

  var reader io.Reader = r.Body
  log.Println(reader)
  err := json.NewDecoder(reader).Decode(loginRequest)
  if err != nil {
    log.Println(err)
  }
  log.Println(loginRequest.Username)
  url := "https://api-staging.vantagesports.com/users/v1/login"
  // var jsonStr = []byte(`{"u": "t-adam@vantagesports.com", "p": "t-adam", "t": true}`)
  // var request io.Reader = loginRequest
  b, err := json.Marshal(loginRequest)
  // req, err := http.NewRequest("POST", url, b)
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  log.Println("response Status:", resp.Status)
  log.Println("response Headers:", resp.Header)
  body, _ := ioutil.ReadAll(resp.Body)
  log.Println(body)
  log.Println("response Body:", string(body))

  w.Write(body)
  // json.NewEncoder(w).Encode(string(body))
}

func main() {
  http.HandleFunc("/login/", loginHandler)

  http.ListenAndServe(":8080", nil);
}
