package main

import (
    "testing"
    "encoding/json"
    "reflect"
)

func TestJsonUnmarshal(t *testing.T) {
    t.Parallel()
    jsonStr := `
{
  "result" : "ok",
  "tuc" : [ {
    "phrase" : {
      "text" : "песочное печенье",
      "language" : "ru"
    },
    "meaningId" : -638808362965036511,
    "authors" : [ 2899 ]
  }, {
    "meanings" : [ {
      "language" : "en",
      "text" : "very rich thick margarine cookie"
    } ],
    "meaningId" : null,
    "authors" : [ 93369 ]
  }, {
    "meanings" : [ {
      "language" : "en",
      "text" : "blah blah blah"
    } ],
    "meaningId" : null,
    "authors" : [ 1 ]
  } ],
  "phrase" : "shortbread",
  "from" : "en",
  "dest" : "ru",
  "authors" : {
    "1" : {
      "U" : "http://en.wiktionary.org",
      "id" : 1,
      "N" : "en.wiktionary.org",
      "url" : "https://glosbe.com/source/1"
    },
    "2899" : {
      "U" : "http://sourceforge.net/projects/freedict/files/",
      "id" : 2899,
      "N" : "English-Russian",
      "url" : "https://glosbe.com/source/2899"
    },
    "93369" : {
      "U" : "http://plwordnet.pwr.wroc.pl/wordnet/",
      "id" : 93369,
      "N" : "plwordnet-defs",
      "url" : "https://glosbe.com/source/93369"
    }
  }
}
    `
    var actual Entry
    err := json.Unmarshal([]byte(jsonStr), &actual)
    if err != nil {
        t.Fail()
    }
    expected := Entry{
        []Option{
            Option{Phrase: Value{"песочное печенье", "ru"}},
            Option{Meanings: []Value{Value{"very rich thick margarine cookie", "en"}}},
            Option{Meanings: []Value{Value{"blah blah blah", "en"}}},
        },
    }
    if !reflect.DeepEqual(actual, expected) {
        t.Fail()
    }
}

func TestJsonUnmarshalEmpty(t *testing.T) {
    t.Parallel()
    jsonStr := `
{
  "result" : "ok",
  "tuc" : [ ],
  "phrase" : "idontexist",
  "from" : "en",
  "dest" : "ru",
  "authors" : { }
}
    `
    var actual Entry
    err := json.Unmarshal([]byte(jsonStr), &actual)
    if err != nil {
        t.Fail()
    }
    if !reflect.DeepEqual(actual, Entry{[]Option{}}) {
        t.Fail()
    }
}

func TestToSlice(t *testing.T) {
    t.Parallel()
    parsed := Entry{
        []Option{
            Option{Phrase: Value{"песочное печенье", "ru"}},
            Option{
                Meanings: []Value{
                    Value{"very rich <b>thick</b> margarine cookie", "en"},
                    Value{"&quot;blah&quot;", "fr"},
                },
            },
            Option{Meanings: []Value{Value{"blah blah blah&#39;s", "en"}}},
        },
    }
    expected := []string{"песочное печенье", "very rich *thick* margarine cookie", `"blah"`, "blah blah blah's"}
    if !reflect.DeepEqual(toSlice(&parsed), expected) {
        t.Fail()
    }
}

func TestToSliceEmpty(t *testing.T) {
    t.Parallel()
    if toSlice(&Entry{[]Option{}}) != nil {
        t.Fail()
    }
}

func TestToWord(t *testing.T) {
    t.Parallel()
    args := []string{"ignore", "love", "me", "do"}
    actual := toWord(&args)
    if !reflect.DeepEqual(actual, "love+me+do") {
        t.Fail()
    }
}

func TestToWordSingle(t *testing.T) {
    t.Parallel()
    args := []string{"ignore", "hate"}
    actual := toWord(&args)
    if !reflect.DeepEqual(actual, "hate") {
        t.Fail()
    }
}
