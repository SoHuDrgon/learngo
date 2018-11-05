package maps

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int      //m3 = nil
	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println("Traversing map ,map 遍历操作")
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println(k)
	}
	for _, v := range m {
		fmt.Println(v)
	}
	fmt.Println("Getting Values 获取操作")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if couseName, ok := m["couse"]; ok {
		fmt.Println(couseName)
	} else {
		fmt.Println("key does not exist.")
	} //Zero value
	fmt.Println("Deleting values 删除操作")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
