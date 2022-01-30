package main

import (
	"fmt"
    "log"
    "net/http"
	"io"
	"os"
	"encoding/json"
	"bytes"
	"bufio"
)

var bodyString string
var jsonStr    string
var token = os.Getenv("DOCKERHUBTOKEN")
var org = os.Getenv("DOCKERHUBORG")
var team = os.Getenv("DOCKERHUBTEAM")
var filename = os.Getenv("DOCKERHUBINVITEFILE")
var testmode = os.Getenv("DOCKERHUBTESTMODE")

// Person Type form json request body
type Person struct {
	Email  string `json:"member"`
}

func main() {
	fmt.Println("Token Check Started...")
	CheckVariables(token,org,team,filename)
	checkStatus := CheckToken(token,org)
	fmt.Println("Token Check",checkStatus)
    TestMode(testmode)
	fmt.Println("Start member invite...")
	InviteAll(filename)
	//members := GetMembers()
	//fmt.Println(members)
}

// GetMembers  helper function to check if api is working, not used during invite
func GetMembers() string {
	url := "https://hub.docker.com/v2/orgs/"+org+"/groups/"+team+"/members"
	client := &http.Client{
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", token )
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
	    log.Fatal(err)
    }
	defer resp.Body.Close()
  
	if resp.StatusCode == http.StatusOK {
	    bodyBytes, err := io.ReadAll(resp.Body)
	    if err != nil {
		    log.Fatal(err)
	    }
	    bodyString = string(bodyBytes)
	}
	return(bodyString)
}

// InviteMember function to form post request
func InviteMember(token string,team string, useremail string) string{
	url := "https://hub.docker.com/v2/orgs/"+org+"/groups/"+team+"/members"
	client := &http.Client{
	}
	person := Person{useremail}
	body, _ := json.Marshal(person)
	req, err := http.NewRequest("POST",url, bytes.NewBuffer(body) )
	req.Header.Add("Authorization", token )
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
	    log.Fatal(err)
    }
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			//Failed to read response.
			log.Fatal(err)
		}

		//Convert bytes to String and print
		jsonStr := string(body)
		fmt.Println("Response: ", jsonStr)

	}
	return(jsonStr)
}

// CheckToken function tests if token is not expired
func CheckToken(token string, org string) string {
	client := &http.Client{
	}
	url := "https://hub.docker.com/v2/orgs/"+org+"/groups/owners/members"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", token )
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
	    log.Fatal(err)
    }
	defer resp.Body.Close()
  
	if resp.StatusCode == http.StatusOK {
	    bodyBytes, err := io.ReadAll(resp.Body)
	    if err != nil {
		    log.Fatal(err)
	    }
	    bodyString = string(bodyBytes)
	} else {
	    log.Fatal("Token is invalid")
	}

	return(resp.Status)
}


// CheckVariables function tests if variables are not empty
func CheckVariables(token string, org string, team string, filename string) {
	if token == ""  {
	    log.Fatal("Problem with Docker Hub Token: env variable DOCKERHUBTOKEN is empty")
    }
	if org == ""  {
	    log.Fatal("Problem with Docker Hub Org: env variable DOCKERHUBORG is empty")
    }
	if team == ""  {
	    log.Fatal("Problem with Docker Hub Team:env variable DOCKERHUBTEAM is empty")
	}
	if team == ""  {
	    log.Fatal("Problem with Docker Hub Team:env variable DOCKERHUBINVITEFILE is empty")
	}
}

// InviteAll function reads file and call InviteMember
func InviteAll(filename string){
	file, err := os.Open(filename)
	if err != nil {
			panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
				break
		}
		email := string(line)
		fmt.Printf("Sending email to %s \n", email)
		InviteMember(token, team, email)
		fmt.Printf("------\n")
	}
}

//TestMode is readonly function to get list of members
func TestMode(testmode string){
	if testmode == "true"  {
		fmt.Printf("Getting list of members in %s organization and %s team\n", org, team) 
		members := GetMembers()
		fmt.Println("List of members:")
		fmt.Println(members)
		os.Exit(0)
	}
}