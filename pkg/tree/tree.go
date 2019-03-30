package tree

import (
	"fmt"
	db "github.com/greenpau/go-ansible-db/pkg/db"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func AnsibleToText(inventoryDir, output string) {
	var x map[string][]string

	log.Printf("in ansibleToText(%s,%s)",inventoryDir, output)

	// as in sh/awk, walk the dirs to find hosts
	x = traverseHostsFiles(inventoryDir)
	fmt.Printf("x=%v\n", x)



	// build a tree of them
	// with an ordered list
	// walk the host_vars to get attributes to add
	// walk the group_vars to get more attributes
	// walk and print the tree in hostname order
}



// traverseHostFiles makes a map of hosts, each with a slice of attributes
func traverseHostsFiles(inventoryDir string) map[string][]string {
	var err error
	var hosts = make(map[string][]string)

	fmt.Println("On Unix:")
	err = filepath.Walk(inventoryDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		// find every hosts file there is
		if info.IsDir() {
			fmt.Printf("dir: %q, %q\n", path, info.Name())
		} else { // it's a file
			fmt.Printf("file: %q, %q\n", path, info.Name())
			if info.Name() == "hosts" {
				readIndividualHostsFile(path, &hosts)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", inventoryDir, err)
		return hosts // FIXME
	}
	return hosts
}

// isDC -- returns true if the directory describes a data center XXX unused yet
func isDC(s string) bool {
	return s == "ansible_inventory" ||
		s == "host_vars" ||
		s == "group_vars" ||
		strings.HasPrefix(s,"DC")
}

// readIndividualHostsFile reads files and assigns map[hostname] with attributes
// this will include the DC and the groupings in the hosts file itself
// This does NOT propogate values, just hosts that end in .com
func readIndividualHostsFile(path string, hostAttrs *map[string][]string){

	log.Printf("in readIndividualHostsFile(%s, map)\n", path)

	// Create a new inventory file.
	inv := db.NewInventory()
	// Load the contents of the inventory from an input file.
	if err := inv.LoadFromFile(path); err != nil {
		panic(fmt.Errorf("error reading inventory: %s", err))
	}

	h := "10.5.2.47"
	host, err := inv.GetHost(h)
	if err != nil {
		panic(fmt.Errorf("error getting host %s from inventory: %s", h, err))
	}
	fmt.Printf("%s", printInventoryHost(host))
	os.Exit(0)
	
	 (*hostAttrs)[path] = []string{"seen"}   // junk assignment
}


func printInventoryHost(h *db.InventoryHost) string {
	return fmt.Sprintf("type InventoryHost struct {\n" +
    	"\tName        string = %q\n" +
		"\tParent      string = %q\n" +
		"\tVariables   map[string]string  = %v\n" +
    	"\tGroups      []string  = %v\n" +
    	"\tGroupChains []string  = %v\n}\n",
			h.Name, h.Parent, h.Variables, h.Groups, h.GroupChains)
}

