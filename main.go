package main

import (
	"fmt"

	"github.com/beevik/etree"
)

func main() {
	fmt.Println("Xml demo!")

	/*创建xml*/
	// doc := etree.NewDocument()
	// doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	// doc.CreateProcInst("xml-stylesheet", `type="text/xsl" href="style.xsl"`)

	// people := doc.CreateElement("People")
	// people.CreateComment("These are all known people")

	// jon := people.CreateElement("Person")
	// jon.CreateAttr("name", "Jon")

	// sally := people.CreateElement("Person")
	// sally.CreateAttr("name", "Sally")

	// doc.Indent(2)
	// doc.WriteTo(os.Stdout)

	/*读xml文件*/
	doc := etree.NewDocument()
	err := doc.ReadFromString(`
	<bookstore xmlns:p="urn:schemas-books-com:prices">

	<book category="COOKING">
	  <title lang="en">Everyday Italian</title>
	  <author>Giada De Laurentiis</author>
	  <year>2005</year>
	  <p:price>30.00</p:price>
	</book>
  
	<book category="CHILDREN">
	  <title lang="en">Harry Potter</title>
	  <author>J K. Rowling</author>
	  <year>2005</year>
	  <p:price>29.99</p:price>
	</book>
  
	<book category="WEB">
	  <title lang="en">XQuery Kick Start</title>
	  <author>James McGovern</author>
	  <author>Per Bothner</author>
	  <author>Kurt Cagle</author>
	  <author>James Linn</author>
	  <author>Vaidyanathan Nagarajan</author>
	  <year>2003</year>
	  <p:price>49.99</p:price>
	</book>
  
	<book category="WEB">
	  <title lang="en">Learning XML</title>
	  <author>Erik T. Ray</author>
	  <year>2003</year>
	  <p:price>39.95</p:price>
	</book>
  
  </bookstore>
  
	`)
	if err != nil {
		panic(err)
	}
	/*访问元素*/
	root := doc.SelectElement("bookstore")
	fmt.Println("ROOT element:", root.Tag)

	for _, book := range root.SelectElements("book") {
		fmt.Println("CHILD element:", book.Tag)
		if title := book.SelectElement("title"); title != nil {
			lang := title.SelectAttrValue("lang", "unknown")
			fmt.Printf("  TITLE: %s (%s)\n", title.Text(), lang)
		}
		for _, attr := range book.Attr {
			fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
		}
	}
}
