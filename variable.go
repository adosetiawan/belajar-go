package main

import "fmt"

func main() {
	//declare variable
	var name string

	name = "Ado Setiawan"
	fmt.Println(name)

	name = "Ado Raziq"
	fmt.Println(name)

	//auto declare
	var friendName = "Budi"
	fmt.Println(friendName)

	//short declarate
	kotaAsar := "Tasikmalaya"
	fmt.Println(kotaAsar)

	//multiple variable
	var (
		sd  = "SDN Sindangjaya"
		smp = "SMPN 1 Cikalong"
		smk = "SMKN 1 Cikalong"
	)
	fmt.Println(sd)
	fmt.Println(smp)
	fmt.Println(smk)
}
