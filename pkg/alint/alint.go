package alint

import (
	"fmt"
	db "github.com/greenpau/go-ansible-db/pkg/db"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// InventoryHost is a superset of Paul Greenberg's struct of the same name
type InventoryHost struct {
	DC	string
	db.InventoryHost
}
func (h InventoryHost) String() string {
	return fmt.Sprintf("InventoryHost struct {\n" +
		"\tName        string = %q\n" +
		"\tDataCentre  string = %q\n" +
		"\tParent      string = %q\n" +
		"\tVariables   map[string]string  = %v\n" +
		"\tGroups      []string  = %v\n" +
		"\tGroupChains []string  = %v\n}\n",
		h.Name, h.DC, h.Parent, h.Variables, h.Groups, h.GroupChains)
}



func LintHostFiles(inventoryDir string, verbose bool) {
	var err error

	err = filepath.Walk(inventoryDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		// find every hosts file there is
		if info.IsDir() {
			//log.Printf("dir: %q, %q\n", path, info.Name())
		} else { // it's a file
			//log.Printf("file: %q, %q\n", path, info.Name())
			if info.Name() == "hosts" {
				readIndividualHostsFile(path, verbose)
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("error seen walking the path %q: %v, ignored\n", inventoryDir, err)
	}

}

// readIndividualHostsFile reads files and calls NewInventory,
// which will report duplicates
func readIndividualHostsFile(path string, verbose bool){
	var h InventoryHost

	p := strings.Split(path, "/")
	dc := p[len(p)-2]

	inv := db.NewInventory()
	if err := inv.LoadFromFile(path); err != nil {
		fmt.Printf("%s: %s\n", path, err)
	}

	if !verbose {
		return
	}
	// for each name in inv, optionally print them in detail
	for _, v := range inv.Hosts {
		h.Name = v.Name
		h.DC = dc
		h.Parent = v.Parent
		h.Variables = v.Variables
		h.Groups = v.Groups
		h.GroupChains = v.GroupChains
		fmt.Printf("%s", h)
	}
}



