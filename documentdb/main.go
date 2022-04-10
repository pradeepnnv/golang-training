package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/cockroachdb/pebble"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type server struct {
	port string
	db   *pebble.DB
}

func newServer(database string, port string) (*server, error) {
	s := server{db: nil, port: port}
	var err error
	s.db, err = pebble.Open(database, &pebble.Options{})
	return &s, err
}

func main() {
	s, err := newServer("docdb.data", "8080")
	if err != nil {
		log.Fatal(err)
	}
	defer s.db.Close()
	router := httprouter.New()
	router.POST("/docs", s.addDocument)
	router.GET("/docs", s.searchDocuments)
	router.GET("/docs/:id", s.getDocument)

	log.Println("Listening on :" + s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, router))

}

func (s server) addDocument(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	dec := json.NewDecoder(r.Body)
	var document map[string]any
	err := dec.Decode(&document)
	if err != nil {
		jsonResponse(w, nil, err)
		return
	}

	// New unique id for the document
	id := uuid.New().String()

	bs, err := json.Marshal(document)
	if err != nil {
		jsonResponse(w, nil, err)
		return
	}
	err = s.db.Set([]byte(id), bs, pebble.NoSync)
	if err != nil {
		jsonResponse(w, nil, err)
		return
	}

	jsonResponse(w, map[string]any{
		"id": id,
	}, nil)
}

func (s server) getDocumentById(id []byte) (map[string]any, error) {
	valBytes, closer, err := s.db.Get(id)

	if err != nil {
		return nil, err
	}
	defer closer.Close()

	var document map[string]any
	err = json.Unmarshal(valBytes, &document)
	return document, err
}

func jsonResponse(w http.ResponseWriter, body map[string]any, err error) {
	data := map[string]any{
		"body":   body,
		"status": "ok",
	}
	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		data["status"] = "error"
		data["error"] = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	err = enc.Encode(data)
	if err != nil {
		// TODO: setup panic handler?
		panic(err)
	}
}

func (s server) getDocument(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	document, err := s.getDocumentById([]byte(id))
	if err != nil {
		jsonResponse(w, nil, err)
		return
	}

	jsonResponse(w, map[string]any{
		"document": document,
	}, nil)
}

// Handles either quoted strings or unquoted strings of only continguous digits and letters

func lexString(input []rune, index int) (string, int, error) {
	if index >= len(input) {
		return "", index, nil
	}

	if input[index] == '"' {
		index++
		foundEnd := false
		var s []rune

		// TODO: handle nested quotes
		for index < len(input) {
			if input[index] == '"' {
				foundEnd = true
				break
			}
			s = append(s, input[index])
			index++
		}
		if !foundEnd {
			return "", index, fmt.Errorf("Expected end of quoted string")
		}

		return string(s), index + 1, nil
	}

	// if unquoted, read as much contiguous digits/letters as there are

	var s []rune
	var c rune

	// TODO: someone needs to validate there's not ...
	for index < len(input) {
		c = input[index]
		if !(unicode.IsLetter(c) || unicode.IsDigit(c) || c == '.') {
			break
		}
		s = append(s, c)
		index++
	}
	if len(s) == 0 {
		return "", index, fmt.Errorf("no string found")
	}

	return string(s), index, nil

}

type queryComparision struct {
	key   []string
	value string
	op    string
}

type query struct {
	ands []queryComparision
}

// E.g. q=a.b:12

func parseQuery(q string) (*query, error) {
	if q == "" {
		return &query{}, nil
	}
	i := 0
	var parsed query
	var qRune = []rune(q)
	for i < len(qRune) {
		// Eat whitespace
		for unicode.IsSpace(qRune[i]) {
			i++
		}

		key, nextIndex, err := lexString(qRune, i)

		if err != nil {
			return nil, fmt.Errorf("expected valid key, got [%s]: `%s`", err, q[nextIndex:])
		}

		// Expect some operator

		if q[nextIndex] != ':' {
			return nil, fmt.Errorf("expected colon at %d, got: `%s`", nextIndex, q[nextIndex:])
		}

		i = nextIndex + 1
		op := "="
		if q[i] == '>' || q[i] == '<' {
			i++
			op = string(q[i])
		}

		value, nextIndex, err := lexString(qRune, i)
		if err != nil {
			return nil, fmt.Errorf("expected valid value, got [%s]: `%s`", err, q[nextIndex:])
		}

		i = nextIndex
		argument := queryComparision{key: strings.Split(key, "."), value: value, op: op}

		parsed.ands = append(parsed.ands, argument)
	}
	return &parsed, nil
}

func (q query) match(doc map[string]any) bool {
	for _, argument := range q.ands {
		value, ok := getPath(doc, argument.key)
		if !ok {
			return false
		}

		// Handle equality

		if argument.op == "=" {
			match := fmt.Sprintf("%v", value) == argument.value
			if !match {
				return false
			}
			continue
		}

		// Handle <,>
		right, err := strconv.ParseFloat(argument.value, 64)
		if err != nil {
			return false
		}

		var left float64
		switch t := value.(type) {
		case float64:
			left = t
		case float32:
			left = float64(t)
		case uint:
			left = float64(t)
		case uint8:
			left = float64(t)
		case uint16:
			left = float64(t)
		case uint32:
			left = float64(t)
		case uint64:
			left = float64(t)
		case int:
			left = float64(t)
		case int8:
			left = float64(t)
		case int16:
			left = float64(t)
		case int32:
			left = float64(t)
		case int64:
			left = float64(t)
		case string:
			left, err = strconv.ParseFloat(t, 64)
			if err != nil {
				return false
			}
		default:
			return false
		}
		if argument.op == ">" {
			if left <= right {
				return false
			}
			continue
		}
		if left >= right {
			return false
		}

	}
	return true
}

func getPath(doc map[string]any, parts []string) (any, bool) {
	var docSegment any = doc
	for _, part := range parts {
		m, ok := docSegment.(map[string]any)
		if !ok {
			return nil, false
		}
		if docSegment, ok = m[part]; !ok {
			return nil, false
		}
	}
	return docSegment, true
}

func (s server) searchDocuments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	q, err := parseQuery(r.URL.Query().Get("q"))
	if err != nil {
		jsonResponse(w, nil, err)
		return
	}
	var documents []map[string]any

	iter := s.db.NewIter(nil)
	defer iter.Close()

	for iter.First(); iter.Valid(); iter.Next() {
		var document map[string]any
		err = json.Unmarshal(iter.Value(), &document)
		if err != nil {
			jsonResponse(w, nil, err)
			return
		}

		if q.match(document) {
			documents = append(documents, map[string]any{
				"id":   string(iter.Key()),
				"body": document,
			})
		}

	}
	jsonResponse(w, map[string]any{"documents": documents,
		"count": len(documents)}, nil)
}
