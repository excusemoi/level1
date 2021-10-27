package task21

import "fmt"

//не очень полезный пример, но забавный

type User struct {

}

func (*User) WriteSomeStaticallyTypedCode(language StaticallyTypedLanguage) {
	language.GenerateStaticallyTypedCode()
}

type StaticallyTypedLanguage interface { //интерфейс языка со статической типизацией
	GenerateStaticallyTypedCode()
}

type Golang struct { //язык, поддерживающий статическую типизацию

}

func (Golang) GenerateStaticallyTypedCode(){ //создание статически типизированного кода
	println("Some chad cool code with huge muscular binary file")
}

type DynamicallyTypedLanguage interface { //интерфейс языка с динамической типизацией
	GenerateDynamicallyTypedCode()
}

type JavaScript struct { //язык, не поддерживающий статическую типизацию

}

func (JavaScript) GenerateDynamicallyTypedCode(){//создание динамически типизированного кода
	println("Oops...undefined again....ok....(")
}

type JavaScriptToTypeScriptAdapter struct {//адаптер javascript'a на typescript
	js *JavaScript
}

func (jsa JavaScriptToTypeScriptAdapter) GenerateStaticallyTypedCode(){//теперь можем адаптировать язык с динамической
	//типизацией к языку со статической типизацией
	jsa.js.GenerateDynamicallyTypedCode()
	println("Converting this uga buga Martin Heidegger text into real code...")
	println("Well done")
}

func Solve(){
	js := JavaScript{}
	jsa := JavaScriptToTypeScriptAdapter{js: &js}
	g := Golang{}
	u := User{}

	fmt.Printf("Case#1 Golang\n")
	u.WriteSomeStaticallyTypedCode(g)

	fmt.Printf("\nCase#2 JSAdapter\n")
	u.WriteSomeStaticallyTypedCode(jsa)
}