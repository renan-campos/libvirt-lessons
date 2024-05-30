package main

import (
	"fmt"
	"os"

	"github.com/renan-campos/libvirt-lessons/pkg/libvirt"
)

func main() {
	// Connect to the local libvirtd instance
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to libvirtd: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	// List all VMs and their CPU/memory allocation
	domains, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing domains from libvirtd: %v", err)
		os.Exit(1)
	}
	for _, domain := range domains {
		name, err := domain.GetName()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get domain name: %v", err)
		}

		info, err := domain.GetInfo()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get info: %v", err)
		}
		fmt.Printf("%s: %d Total CPUs, %d MB Max memory\n", name, info.NrVirtCpu, info.MaxMem)
	}
}
