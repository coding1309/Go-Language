package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

func main() {
	// Text file
	// textFile, err := ioutil.ReadFile("text.txt")
	textFile, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println("Error reading text file:", err)
		return
	}
	fmt.Println("Text file contents:")
	fmt.Println(string(textFile))

	// HTML file
	// htmlFile, err := ioutil.ReadFile("index.html")
	htmlFile, err := os.ReadFile("index.html")
	if err != nil {
		fmt.Println("Error reading HTML file:", err)
		return
	}
	fmt.Println("\nHTML file contents:")
	fmt.Println(string(htmlFile))

	// JSON template file
	// jsonTemplateFile, err := ioutil.ReadFile("data.json")
	jsonTemplateFile, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading JSON template file:", err)
		return
	}

	// Parse JSON data (example format)
	jsonData := `{"name": "Neel Patel", "age": 22}`
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error parsing JSON data:", err)
		return
	}

	// Create and parse the JSON template
	tmpl, err := template.New("jsonTemplate").Parse(string(jsonTemplateFile))
	if err != nil {
		fmt.Println("Error parsing JSON template:", err)
		return
	}

	// Execute the template with the data
	// err = tmpl.Execute(fmt.OsWriter, data)
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("Error executing JSON template:", err)
	}
}