package main

import (
    "fmt"
    "os"
    "net/http"
    "encoding/json"
    "jaytaylor.com/html2text"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: whut <word to translate>")
        os.Exit(1)
    }
    url := "https://glosbe.com/gapi/translate?from=eng&dest=rus&format=json&phrase=" + toWord(&os.Args)
    client := &http.Client{}
    entry := &Entry{}
    getJson(client, url, entry)
    for idx, val := range toSlice(entry) {
        fmt.Printf("%d. %s\n", idx + 1, val)
    }
}

type Entry struct {
    Tuc []Option `json:"tuc"`
}

type Option struct {
    Phrase Value `json:"phrase"`
    Meanings []Value `json:"meanings"`
}

type Value struct {
    Text string `json:"text"`
    Language string `json:"language"`
}

func getJson(client *http.Client, url string, target interface{}) error {
    r, err := (*client).Get(url)
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()
    return json.NewDecoder(r.Body).Decode(target)
}

func toSlice(entry *Entry) []string {
    var result []string
    for _, val := range (*entry).Tuc {
        if (val.Phrase != Value{}) {
            result = appendText(result, val.Phrase.Text)
        } else {
            for _, val := range val.Meanings {
                result = appendText(result, val.Text)
            }
        }
    }
    return result
}

func appendText(result []string, text string) []string {
    txt, err := html2text.FromString(text, html2text.Options{})
    if err == nil {
        return append(result, txt)
    } else {
        return result
    }
}

func toWord(args *[]string) string {
    result := ""
    lng := len(*args)
    for i := 1; i < lng; i++ {
        result += (*args)[i]
        if i < lng - 1 {
            result += "+"
        }
    }
    return result
}
