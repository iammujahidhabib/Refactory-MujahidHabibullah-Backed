package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Define data structure
type Posts struct {
	USERID int    `json:"userId"`
	ID     int    `json:"id"`
	TITLE  string `json:"title"`
	BODY   string `json:"body"`
}

type FinalPosts struct {
	USERID int         `json:"userId"`
	ID     int         `json:"id"`
	TITLE  string      `json:"title"`
	BODY   string      `json:"body"`
	USER   SecondPosts `json:"user"`
}

type SecondPosts struct {
	ID       int     `json:"id"`
	NAME     string  `json:"name"`
	USERNAME string  `json:"username"`
	EMAIL    string  `json:"email"`
	ADDRESS  Address `json:"address"`
	PHONE    string  `json:"phone"`
	WEBSITE  string  `json:"website"`
}
type Address struct {
	STREET  string `json:"street"`
	SUITE   string `json:"suite"`
	CITY    string `json:"city"`
	ZIPCODE string `json:"zipcode"`
	GEO     Geo    `json:"geo"`
}
type Geo struct {
	LAT string `json:"lat"`
	LNG string `json:"lng"`
}
type Company struct {
	NAME        string `json:"name"`
	CATCHPHRASE string `json:"catchPhrase"`
	BS          string `json:"bs"`
}

func readJSONFromUrl(url string) ([]Posts, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var postList []Posts
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &postList); err != nil {
		return nil, err
	}

	return postList, nil
}
func get_content() []Posts {
	// json data
	url := "https://jsonplaceholder.typicode.com/posts"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data []Posts

	// unmarshall
	json.Unmarshal(body, &data)
	// fmt.Printf("Results: %v\n", data)

	return data
}

func get_second_content() []SecondPosts {
	// json data
	url := "https://jsonplaceholder.typicode.com/users"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data []SecondPosts

	// unmarshall
	json.Unmarshal(body, &data)
	// fmt.Printf("Results: %v\n", data)

	return data
}

func main() {
	// url := "https://jsonplaceholder.typicode.com/posts"
	// postList, err := readJSONFromUrl(url)
	// if err != nil {
	// 	panic(err)
	// }

	// for idx, row := range postList {
	// 	// skip header
	// 	if idx == 0 {
	// 		continue
	// 	}

	// 	if idx == 6 {
	// 		break
	// 	}

	// 	fmt.Println(row.title)
	// }
	// fmt.Println(get_content())
	// fmt.Println(get_second_content())
	// fmt.Println(len(get_content()))
	// fmt.Println(len(get_second_content()))
	var first_content = get_content()
	var second_content = get_second_content()

	var data_ FinalPosts
	var final_data []FinalPosts

	for i := 0; i < len(second_content); i++ {
		for j := 0; j < len(first_content); j++ {
			if second_content[i].ID == first_content[j].USERID {
				data_.USERID = first_content[j].USERID
				data_.ID = first_content[j].ID
				data_.TITLE = first_content[j].TITLE
				data_.BODY = first_content[j].BODY
				data_.USER = second_content[i]
				final_data = append(final_data, data_)
			}
		}
	}
	// fmt.Println(final_data)

	// Convert to JSON
	// j, _ := json.Marshal(final_data)
	// log.Println(string(j))

	// Convert to JSON
	json_data, err := json.Marshal(final_data)
	if err != nil {
		fmt.Println(err)
		return
	}
	//print json data
	fmt.Println(string(json_data))

	//create json file
	json_file, err := os.Create("10.json")
	if err != nil {
		fmt.Println(err)
	}
	defer json_file.Close()

	json_file.Write(json_data)
	json_file.Close()
}
