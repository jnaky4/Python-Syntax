package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main(){
	//initPlugin()
	readInit()

	//var funcMap = make(map[string]map[string]string)
	//funcMap["put"] = map[string]string{}
	//funcMap["put"]["params"] = `40`
	//funcMap["put"]["results"] = `80`
	//for k, _ := range funcMap{
	//	for _, v := range funcMap[k]{
	//		println(v)
	//	}
	//
	//}


	//for k, v := range funcMap {
	//
	//}

	//funcMap["main"] = "fmt.Println(\"hello world\")"

	//funcMap["package"]["myplugin"] = "myplugin"
	//funcMap["import"]["git.target.com/davidgoldstein/target-plugin"] = "plg"
	//data := generateFuncStr(funcMap)

}

func initPlugin(){
	cwd, err := os.Getwd()
	if err != nil {
		println("Err, failed to get cwd")
		return
	}

	pkgPath := strings.Split(cwd, string(os.PathSeparator))
	initStr := `package ` + pkgPath[len(pkgPath)-1] +`

import (
	plg "git.target.com/davidgoldstein/target-plugin"
)
type KVService interface {
	plg.Plugin
	Put(key string, value []byte) error
	Get(key string) ([] byte, error)
}`
	//println(initStr)

	err = os.WriteFile(cwd +"/interface.go", []byte(initStr), 0777)
	if err != nil{
		fmt.Println("Err, failed to write file to " + cwd + string(os.PathSeparator) +"/test/interface.go")
	}

}
func readInit(){
	//stringHolder := ""
	cwd, err := os.Getwd()
	if err != nil {
		println("Err, failed to get cwd")
		return
	}

	file, err := os.Open(cwd +"/files/test/interface.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)


	//^ match begging of file
	//\\s* any number of white space
	//[a-zA-Z] character is letter
	//[a-zA-Z0-9]+ any number of alphanumeric characters
	//[(] matches Parenthesis
	//[\[\].,a-z\s]* matches any number of parameters
	//[)] matches Parenthesis

	functionDeclaration, _ := regexp.Compile("^\\s*[[:alpha:]][[:alnum:]]+[(][\\[\\].,a-z\\s]*[)]")
	functionName, _ := regexp.Compile("[[:alpha:]][[:alnum:]]+")
	parameters, _ := regexp.Compile("[(][\\[\\].,a-z\\s]*[)]")
	returnValues, _:= regexp.Compile("[)][.,a-z\\s]*")

	for scanner.Scan() {
		matchString := functionDeclaration.MatchString(scanner.Text())

		if matchString{
			name := functionName.Find([]byte(scanner.Text()))
			params := parameters.Find([]byte(scanner.Text()))
			returVals := returnValues.Find([]byte(scanner.Text()))
			returnValues.FindSubmatch([]byte(scanner.Text()))
			println("found " + scanner.Text())
			println("name " + string(name))
			println("params " + string(params))
			println("return values " + string(returVals))
		}

	}

	////todo fix temp hardcoded route
	//fileData, err  := os.ReadFile(cwd +"/files/test/interface.go")
	//if err != nil{
	//	fmt.Println("Err, failed to read file " + cwd + string(os.PathSeparator) +"interface.go")
	//}
	////todo read file by line
	//
	//

	//
	//
	//println(string(fileData))
}

//
//func generateProto() {
//
//}
//
//func generateTest(functionNames map[string]map[string]string){
//	cwd, err := os.Getwd()
//	if err != nil {
//		println("Err, failed to get cwd")
//		return
//	}
//
//	pkgPath := strings.Split(cwd, string(os.PathSeparator))
//	//function name, parameters, expected output
//	//["Put"]["Params"]
//	//["Put"]["Result"]
//	println(pkgPath)
////	testStr := `package ` + pkgPath[len(pkgPath)-1] +`
////import "testing"
////
////
////
////
////func Test`+ for k, v :=  + `(t *testing.T){
////
////
////}
////`
//
//
//
//	err = os.WriteFile(cwd +"/config_test.go", []byte(initStr), 0777)
//	if err != nil{
//		fmt.Println("Err, failed to write file to " + cwd)
//	}
//}

//func format(s string, v interface{}) string {
//	t, b := new(template.Template), new(strings.Builder)
//	err := template.Must(t.Parse(s)).Execute(b, v)
//	if err != nil {
//		return ""
//	}
//	return b.String()
//}
//
//func generateFuncStr(test map[string]string) string{
//	var data string
//	for k, v := range test{
//		data = "package main\nimport ( \"fmt\" )\n"
//		data = data + format("func {{.}}(){\n\t", k)
//		data = data + v + "\n}"
//	}
//
//	return data
//}



//func switchType(){
//}

//func strucGenerator(map[])