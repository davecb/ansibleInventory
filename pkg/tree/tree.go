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
	var x map[string]InventoryHost

	log.Printf("in ansibleToText(%s,%s)",inventoryDir, output)

	// as in sh/awk, walk the dirs to find hosts
	x = traverseHostsFiles(inventoryDir)
	log.Printf("x=%v\n", x)

	// build a tree of them
	// walk and print the tree in hostname order
}


// InventoryHost is a superset of Paul Greenberg's struct of the same name
type InventoryHost struct {
	DC	string
	db.InventoryHost
}
func (h InventoryHost) String() string {
	return fmt.Sprintf("type InventoryHost struct {\n" +
		"\tName        string = %q\n" +
		"\tDataCentre  string = %q\n" +
		"\tParent      string = %q\n" +
		"\tVariables   map[string]string  = %v\n" +
		"\tGroups      []string  = %v\n" +
		"\tGroupChains []string  = %v\n}\n",
		h.Name, h.DC, h.Parent, h.Variables, h.Groups, h.GroupChains)
}


// traverseHostFiles makes a map of hosts, each with an InventoryHost struct
func traverseHostsFiles(inventoryDir string) map[string]InventoryHost {
	var err error
	var hosts = make(map[string]InventoryHost)

	err = filepath.Walk(inventoryDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		// find every hosts file there is
		if info.IsDir() {
			log.Printf("dir: %q, %q\n", path, info.Name())
		} else { // it's a file
			log.Printf("file: %q, %q\n", path, info.Name())
			if info.Name() == "hosts" {
				readIndividualHostsFile(path, &hosts)
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("error walking the path %q: %v\n", inventoryDir, err)
		return hosts // FIXME
	}
	return hosts
}

// readIndividualHostsFile reads files and assigns map[hostname] with attributes
// this will include the DC and the groupings in the hosts file itself
// This does NOT propogate values, just hosts that end in .com
func readIndividualHostsFile(path string, hosts *map[string]InventoryHost){
	var h InventoryHost

	log.Printf("in readIndividualHostsFile(%s, map)\n", path)
	p := strings.Split(path, "/")
	dc := p[len(p)-2]

	inv := db.NewInventory()
	if err := inv.LoadFromFile(path); err != nil {
		log.Printf("ERROR reading inventory from %s: %s, ignored", path, err)
		return 
	}

	// for each name in inv, copy them into our extended struct
	for _, v := range inv.Hosts {
		h.Name = v.Name
		h.DC = dc
		h.Parent = v.Parent
		h.Variables = v.Variables
		h.Groups = v.Groups
		h.GroupChains = v.GroupChains

		if value, exists :=  (*hosts)[v.Name]; exists {
			log.Printf("ERROR, %v already exists, %v\n", value)
		}
		(*hosts)[v.Name] = h
		log.Printf("found %s", h)
	}
}



