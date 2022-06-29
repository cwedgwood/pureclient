/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Model28HostsBody struct {
	Chap *Api28hostsChap `json:"chap,omitempty"`
	// The iSCSI qualified name (IQN) associated with the host.
	Iqns []string `json:"iqns,omitempty"`
	// The NVMe Qualified Name (NQN) associated with the host.
	Nqns []string `json:"nqns,omitempty"`
	// Determines how the system tunes the array to ensure that it works optimally with the host. Set `personality` to the name of the host operating system or virtual memory system. Valid values are `aix`, `esxi`, `hitachi-vsp`, `hpux`, `oracle-vm-server`, `solaris`, and `vms`. If your system is not listed as one of the valid host personalities, do not set the option. By default, the personality is not set.
	Personality string `json:"personality,omitempty"`
	// For synchronous replication configurations, sets a host's preferred array to specify which array exposes active/optimized paths to that host. Enter multiple preferred arrays in comma-separated format. If a preferred array is set for a host, then the other arrays in the same pod will expose active/non-optimized paths to that host. If the host is in a host group, `preferred_arrays` cannot be set because host groups have their own preferred arrays. On a preferred array of a certain host, all the paths on all the ports (for both the primary and secondary controllers) are set up as A/O (active/optimized) paths, while on a non-preferred array, all the paths are A/N (Active/Non-optimized) paths.
	PreferredArrays []Api28hostsPreferredArrays `json:"preferred_arrays,omitempty"`
	// The Fibre Channel World Wide Name (WWN) associated with the host.
	Wwns []string `json:"wwns,omitempty"`
}