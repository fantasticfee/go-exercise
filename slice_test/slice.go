package main

import "fmt"

func main() {
	var a []string
	var c []string
	b := a
	c = a
	b = append(b, "abc")
	c = append(c, "abc")
	fmt.Println("b:", b)
	fmt.Println("a:", a)

	/*
			filePath := "./test1.ini"
			var hosts []Host = {{"Name":"node1","Ip":"192.168.122.12","Role":"master"},{"Name":"node1","Ip":"192.168.122.12","Role":"master"}, {"Name":"node2","Ip":"192.168.122.13","Role":"master"}}
		 		f, err := os.Create(filePath)
		        if err != nil {
		                fmt.Printf("create map file error: %v\n", err)
		                return err
		        }
		        defer f.Close()

				w := bufio.NewWriter(f)
				fmt.Fprintln(w, "[all]")
		        for k, v := range hosts {
		                lineStr := fmt.Sprintf("%s ansible_host=%s ip=%s", v.Name, v.Ip,v.Ip,)
		                fmt.Fprintln(w, lineStr)
		        }
			return w.Flush()
	*/

}
