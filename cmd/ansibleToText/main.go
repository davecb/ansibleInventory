package  main

import "gitlab.indexexchange.com/dave.collier-brown/inventoryTree/pkg/tree"

func main() {
	// parse options
	//expect input Directory
	// expect output file
	tree.AnsibleToText("/home/davecb/projects/at_index_exchange/IDB/ansible_inventory",
		"/home/davecb/projects/at_index_exchange/IDB/ansible.csv")
} 
