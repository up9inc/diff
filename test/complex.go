package test

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type PathsMap map[string]*PathItem
type PathItem struct {
	Summary     string     `json:"summary,omitempty" diff:"summary"`
	Description string     `json:"description,omitempty" diff:"description"`
	Get         *Operation `json:"get,omitempty" diff:"get"`
}

type Operation struct {
	Tags        []string     `json:"tags,omitempty" diff:"tags"`
	Summary     string       `json:"summary,omitempty" diff:"summary"`
	Description string       `json:"description,omitempty" diff:"description"`
	OperationID string       `json:"operationId,omitempty" diff:"operationId"`
	Responses   ResponsesMap `json:"responses,omitempty" diff:"responses"`
}

type ResponsesMap map[string]*Response
type Response struct {
	Description string     `json:"description,omitempty" diff:"description"`
	Content     ContentMap `json:"content,omitempty" diff:"content"`
}

type ContentMap map[string]*MediaType
type MediaType struct {
	Schema *Schema `json:"schema,omitempty" diff:"schema"`
}

type SchemasMap map[string]*Schema
type SchemasSlice []*Schema
type Schema struct {
	OneOf      SchemasSlice `json:"oneOf,omitempty" diff:"oneOf"`
	Properties SchemasMap   `json:"properties,omitempty" diff:"properties"`
	Items      *Schema      `json:"items,omitempty" diff:"items"`
	Enum       []string     `json:"enum,omitempty" diff:"enum"`

	Type        string   `json:"type,omitempty" diff:"type"`
	Title       string   `json:"title,omitempty" diff:"title"`
	Format      string   `json:"format,omitempty" diff:"format"`
	Description string   `json:"description,omitempty" diff:"description"`
	Required    []string `json:"required,omitempty" diff:"required"`
}

func ReadComplexTestFile(t *testing.T, path string) PathsMap {
	fData, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var data PathsMap
	err = json.Unmarshal(fData, &data)
	if err != nil {
		t.Fatal(err)
	}

	return data
}
