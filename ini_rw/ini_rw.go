package main

import (
	"fmt"
	"goconfig"
)

/*
type Host struct {
	Id          int64
	Ip          string
	Name        string
	Label       int
	Roles       string
	Description string
}

type HostIni struct {
	Hosts []Host
}
*/

func main() {

	cfg, err := goconfig.LoadConfigFile("host.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	value, err := cfg.GetValue(goconfig.ALL_SECTION, "ip")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
	/*
		_, error := os.OpenFile("hosts.ini", os.O_RDWR|os.O_CREATE, 0766)
		if error != nil {
			fmt.Println(error)
			return
		}

		cfg, _ := ini.LoadSources(ini.LoadOptions{
			AllowBooleanKeys: true,
		}, "hosts.ini")
		fmt.Println("ip:", cfg.Section("all").Key("ip").Value())
		fmt.Println("host:", cfg.Section("all").Key("ansible_host").Value())

			cfg, err := ini.Load("hosts.ini")
			if err != nil {
				fmt.Printf("Fail to read file: %v", err)
				os.Exit(1)
			}
	*/
	/*
		 cfg.NewSection("all")
		 hostIp := make(map[string]string)
		 hostIp["ansible_host"] =
		//file2.WriteString("node1")
		 cfg.Section("all").NewBooleanKey("node1")
		 cfg.Section("all").NewKey("ansible_host", "192.168.122.101")
		 cfg.Section("all").NewKey("ip", "192.168.122.101")
		 cfg.Section("all").NewKey("etcd_member_name", "etcd1")
	*/

	return
}

/*
func generateAllSection(alias, etcd string, keyValue map[string]string){
	 cfg.Section("all").NewBooleanKey(alias)
	 for key,val := range keyValue {
	 cfg.Section("all").NewKey("ansible_host", "192.168.122.101")

	 }

}

func generateHostIni() {
	var in HostIni
	json.Unmarshal(h.Ctx.Input.RequestBody, &in)

	_, error := os.OpenFile("hosts.ini", os.O_RDWR|os.O_CREATE, 0766)
	if error != nil {
		fmt.Println(error)
		return
	}

	cfg, err := ini.LoadSources(ini.LoadOptions{
		AllowBooleanKeys: true,
	}, "hosts.ini")
	generateDefaultSectionAll(cfg,)
	//generateHost(in)

	//
	return

}

func generateDefaultSectionAll(f *ini.File, in HostIni) {
	for i,host := range in.Hosts {
	 	cfg.Section("all").NewBooleanKey(host.Name)
		f.Section("all").NewKey("ansible_host", host.Ip)
		f.Section("all").NewKey("ip", host.Ip)
		if i < 3 {
		etcdNum := "etcd" + strconv.Itoa(i+1)
	 	f.Section("all").NewKey("etcd_member_name", etcdNum)
		}
	}
}
*/
