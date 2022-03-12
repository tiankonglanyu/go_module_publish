// @Time   : 2022/3/12 22:08    
// @Author : yxl      
// @File   : hello_world.go

package go_module_publish

import "fmt"


import "github.com/tiankonglanyu/go_module_publish/child_package"

func Hello(){
	fmt.Println("hello world, this is my first publish module")
	fmt.Println("this Hello func in top package")
	child_package_.Child()
}
