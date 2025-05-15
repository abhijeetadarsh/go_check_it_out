package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var raw = `
<!DOCTYPE html>
<html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
      <img src="xxx.jpg" width="104" height="142">
  </body>
</html>`

func main() {
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %s\n", err)
		os.Exit(-1)
	}

	// fmt.Printf("%T %v\n", doc, *doc)

	words, images := countWordsAndImages(doc)

	fmt.Printf("%d words and %d images\n", words, images)
}

func countWordsAndImages(doc *html.Node) (int, int) {
	var wc, ic int
	visit(doc, &wc, &ic)
	return wc, ic
}

func visit(node *html.Node, wc, ic *int) {
	// fmt.Println(node.Type, node.Data)
	if node.Type == html.TextNode {
		*wc += len(strings.Fields(node.Data))
	} else if node.Type == html.ElementNode && node.Data == "img" {
		*ic++
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		visit(child, wc, ic)
	}
}
