package builtin

import (
	"os"
	"text/template"
)

func TextTemplate() {
	//t1 := template.New("t1")
	//// 声明模板方式1
	//t1, err := t1.Parse("Value is {{.}}\n")
	//if err != nil {
	//	panic(err)
	//}
	// 声明模板的方式2
	//t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// 执行模板， 发送给os.Stdout 就是在控制台输出
	//t1.Execute(os.Stdout, "some text")
	//t1.Execute(os.Stdout, 5)
	//t1.Execute(os.Stdout, []string{
	//	"Go",
	//	"Rust",
	//	"C++",
	//	"C#",
	//})

	//Create := func(name, t string) *template.Template {
	//	return template.Must(template.New(name).Parse(t))
	//}
	////如果是struct或者map 可以用{{.FieldName}}来进行声明
	//t2 := Create("t2", "Name:{{.Name}}\n")
	//
	//t2.Execute(os.Stdout, struct {
	//	Name string
	//}{"Jane Done"})
	//t2.Execute(os.Stdout, map[string]string{
	//	"Name": "Mickey Mouse",
	//	"Nam":  "Test",
	//})

	//t3 := template.Must(template.New("t3").
	//	Parse("{{if . -}} yes {{else -}} no {{end}}\n"))
	//t3.Execute(os.Stdout, "not empty")
	//t3.Execute(os.Stdout, "")

	t4, err := template.New("t4").Parse("Range:{{range .}}{{.}} {{end}}\n")
	if err != nil {
		panic(err)
	}
	t4.Execute(os.Stdout, []string{
		"Go", "Rust", "C++", "C#",
	})

}
