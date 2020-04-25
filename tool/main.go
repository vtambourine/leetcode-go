package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"text/template"
)

//fetch("https://leetcode.com/api/problems/all/", {
//"headers": {
//"accept": "application/json, text/javascript, */*; q=0.01",
//"accept-language": "en-US,en;q=0.9,ru;q=0.8,nl;q=0.7",
//"content-type": "application/json",
//"sec-fetch-dest": "empty",
//"sec-fetch-mode": "cors",
//"sec-fetch-site": "same-origin",
//"x-newrelic-id": "UAQDVFVRGwEAXVlbBAg=",
//"x-requested-with": "XMLHttpRequest",
//"cookie": "__cfduid=d9ecbb9bb3be2ef1e05da195d26e339e21585765021; csrftoken=wmB1RkwRfRgtjfCFMTDfuOVmoHLANlmDduvDJY9RTuTmEsZfGgpFd4moCwE87DaO; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiODEyMDgzIiwiX2F1dGhfdXNlcl9iYWNrZW5kIjoiZGphbmdvLmNvbnRyaWIuYXV0aC5iYWNrZW5kcy5Nb2RlbEJhY2tlbmQiLCJfYXV0aF91c2VyX2hhc2giOiJlN2MzMTc2MWUyOWUyYzBiNzAxOTY5NTcyMjM4MDBlZjRmNGUzMmMyIiwiaWQiOjgxMjA4MywiZW1haWwiOiJ2dGFtYm91cmluZUB5YS5ydSIsInVzZXJuYW1lIjoidnRhbWJvdXJpbmUiLCJ1c2VyX3NsdWciOiJ2dGFtYm91cmluZSIsImF2YXRhciI6Imh0dHBzOi8vd3d3LmdyYXZhdGFyLmNvbS9hdmF0YXIvN2VjOTQ3NjVhNDg5NzZlMDIxOTE5OWU4OGFkNzYwMDgucG5nP3M9MjAwIiwidGltZXN0YW1wIjoiMjAyMC0wNC0xOCAwNzozNzowMS4wNzg3OTkrMDA6MDAiLCJJUCI6Ijk1LjI0LjI5LjIwNSIsIklERU5USVRZIjoiNDVmNjFmNDcxNWNjMDIzMzZlNDFjOGU4ZThlMmQ2MTAiLCJfc2Vzc2lvbl9leHBpcnkiOjEyMDk2MDB9.mBVftmVG_Vmr1e4y7mYQMsgFKUPSkaf4ysArgsYXhwk; c_a_u=\"dnRhbWJvdXJpbmU=:1jSG4Z:JBLIrwJUzBkb_Z59Qkt3bH0OILo\"; __cf_bm=453457239c4a633432bf9e2525bae8f20beb0454-1587851995-1800-ASS7FmSbULGi9190gDVgCy23nSxqEsrhFCb7fMx/tvykwjQsiwDvIuQLHbUiM/lqrPtIzdCpzkahV0iKYXJk33k="
//},
//"referrer": "https://leetcode.com/problemset/all/",
//"referrerPolicy": "no-referrer-when-downgrade",
//"body": null,
//"method": "GET",
//"mode": "cors"
//});

type LeetCodeResponse struct {
	UserName string                  `json:"user_name"`
	Solved   float64                 `json:"num_solved"`
	Total    float64                 `json:"num_total"`
	Category string                  `json:"category_slug"`
	Problems []LeetCodeProblemStatus `json:"stat_status_pairs"`
}

type LeetCodeProblemStatus struct {
	Difficulty LeetCodeProblemDifficulty `json:"difficulty"`
	Problem    LeetCodeProblem           `json:"stat"`
}

type LeetCodeProblem struct {
	ID        float64 `json:"frontend_question_id"`
	Title     string  `json:"question__title"`
	TitleSlug string  `json:"question__title_slug"`
}

type LeetCodeProblemDifficulty struct {
	Level float64 `json:"level"`
}

func fetchAll() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://leetcode.com/api/problems/all/", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := ioutil.WriteFile("./problems.json", body, 0644); err != nil {
		log.Fatal(err)
	}
}

func readAll() map[int]LeetCodeProblem {
	data, err := ioutil.ReadFile("./problems.json")
	if err != nil {
		log.Fatal(err)
	}
	var problemsData LeetCodeResponse

	if err = json.Unmarshal(data, &problemsData); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", problemsData.Problems[0])

	index := make(map[int]LeetCodeProblem)
	for _, p := range problemsData.Problems {
		index[int(p.Problem.ID)] = p.Problem
	}
	return index
}

var data map[int]LeetCodeProblem

const readme = `{{- /* Tempalte for problem's readme file '*/ -}}
#### [{{.ID}}. {{.Title}}](https://leetcode.com/problems/{{.TitleSlug}})
`

var t = template.Must(template.New("readme").Parse(readme))

func formatProblem(file os.FileInfo, n int) {
	files, err := ioutil.ReadDir(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.Name() == "README.md" {
			return
		}
	}

	info := data[n]

	readmeFile, err := os.Create(file.Name() + "/README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer readmeFile.Close()
	t.Execute(readmeFile, info)
	fmt.Println("Done!", info.ID)
	//t.Execute(os.Stdout, info)
}

func commitAll() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	pr := regexp.MustCompile(`^(\d{4})`)
	for _, file := range files {
		if file.IsDir() && pr.MatchString(file.Name()) {
			n, _ := strconv.Atoi(pr.FindString(file.Name()))
			info := data[n]
			if n != 0 && n != 9999 {
				out, nil := exec.Command("git", "status", file.Name(), "--porcelain").Output()
				if err != nil {
					log.Fatal(err)
				}
				if len(out) == 0 {
					continue
				}
				msg := fmt.Sprintf("%.0f. %s", info.ID, info.Title)
				//fmt.Println(msg)
				err = exec.Command("git", "add", file.Name()).Run()
				err = exec.Command("git", "commit", "-m", msg).Run()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func main() {
	//fetchAll()
	data = readAll()

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	pr := regexp.MustCompile(`^(\d{4})`)
	for _, file := range files {
		if file.IsDir() && pr.MatchString(file.Name()) {
			n, _ := strconv.Atoi(pr.FindString(file.Name()))
			if n != 0 && n != 9999 {
				formatProblem(file, n)
			}
		}
	}
	commitAll()
}
