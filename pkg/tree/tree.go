package tree

import "log"

func AnsibleToText(inventoryDir, output string) {
	log.Printf("in ansibleToText(%s,%s)",inventoryDir, output)
	// as in sh/awk, walk the dirs to find hosts
	// build a tree of them
	// with an ordered list
	// walk the host_vars to get attributes to add
	// walk the group_vars to get more attributes
	// walk and print the tree in hostname order
}
