package main

import (
	"fmt"
	"os"
	"plugin"
	"github.com/nkbai/blog/goplugin/anotherlib"
)

type Greeter interface {
	Greet()
	GetShareVariable() int
}

func main() {
	// load module
	// 1. open the so file to load the symbols
	plug, err := plugin.Open("./eng/eng.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)
	var greeter Greeter
	greeter, ok := symGreeter.(Greeter)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	// 4. use the module
	greeter.Greet()

	fmt.Println("anotherlib in main")
	fmt.Println(anotherlib.ShareVariable)
	fmt.Printf("plugin anotherlib =%d\n",greeter.GetShareVariable())
	fmt.Println("change anotherlib's variable")
	anotherlib.ShareVariable=5
	fmt.Printf("main share=%d,plugin share=%d\n",anotherlib.ShareVariable,greeter.GetShareVariable())
	//可以看到输出都是5

	//下面这种情况将会出现不一致的情况
	testpluginvendor()
}

func testpluginvendor(){
		// load module
	// 1. open the so file to load the symbols
	plug, err := plugin.Open("pluginwithvendor/eng.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)
	var greeter Greeter
	greeter, ok := symGreeter.(Greeter)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	// 4. use the module
	greeter.Greet()
	fmt.Println("call plugin withvendor")
	fmt.Println("anotherlib in main")
	fmt.Println(anotherlib.ShareVariable)
	fmt.Printf("plugin anotherlib =%d\n",greeter.GetShareVariable())
	fmt.Println("change anotherlib's variable")
	anotherlib.ShareVariable=5
	fmt.Printf("main share=%d,plugin share=%d\n",anotherlib.ShareVariable,greeter.GetShareVariable())
	//可以看到输出并不一致
}