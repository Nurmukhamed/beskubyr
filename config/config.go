package config

// Methods
// addIPSet(s: ipset, (ssssa{ss}as): settings) → o
// Add ipset with given settings into permanent configuration. Settings are in format: version, name, description, type, dictionary of options and array of entries.
//
// version (s): see version attribute of ipset tag in firewalld.ipset(5).
// name (s): see short tag in firewalld.ipset(5).
// description (s): see description tag in firewalld.ipset(5).
// type (s): see type attribute of ipset tag in firewalld.ipset(5).
// options (a{ss}): dictionary of {option : value} . See options tag in firewalld.ipset(5).
// entries (as): array of entries, see entry tag in firewalld.ipset(5).
// Possible errors: NAME_CONFLICT, INVALID_NAME, INVALID_TYPE
func addIPSet() {}

// addIcmpType(s: icmptype, (sssas): settings) → o
// Add icmptype with given settings into permanent configuration. Settings are in format: version, name, description, array of destinations. Returns object path of the new icmp type.
//
// version (s): see version attribute of icmptype tag in firewalld.icmptype(5).
// name (s): see short tag in firewalld.icmptype(5).
// description (s): see description tag in firewalld.icmptype(5).
// destinations (as): array, either empty or containing strings 'ipv4' or 'ipv6', see destination tag in firewalld.icmptype(5).
// Possible errors: NAME_CONFLICT, INVALID_NAME, INVALID_TYPE
func addIcmpType() {}

// addService(s: service, (sssa(ss)asa{ss}asa(ss)): settings) → o
// Add service with given settings into permanent configuration. Settings are in format: version, name, description, array of ports (port, protocol), array of module names, dictionary of destinations, array of protocols and array of source-ports (port, protocol). Returns object path of the new icmp type.
//
// version (s): see version attribute of service tag in firewalld.service(5).
// name (s): see short tag in firewalld.service(5).
// description (s): see description tag in firewalld.service(5).
// ports (a(ss)): array of port and protocol pairs. See port tag in firewalld.service(5).
// module names (as): array of kernel netfilter helpers, see module tag in firewalld.service(5).
// destinations (a{ss}): dictionary of {IP family : IP address} where 'IP family' key can be either 'ipv4' or 'ipv6'. See destination tag in firewalld.service(5).
// protocols (as): array of protocols. See protocol tag in firewalld.service(5).
// source-ports (a(ss)): array of port and protocol pairs. See source-port tag in firewalld.service(5).
// Possible errors: NAME_CONFLICT, INVALID_NAME, INVALID_TYPE
func addService() {}

// addZone(s: zone, (sssbsasa(ss)asba(ssss)asasasasa(ss)): settings) → o
// Add zone with given settings into permanent configuration. Settings are in format: version, name, description, UNUSED, target, array of services, array of ports (port, protocol), array of icmp-blocks, masquerade, array of forward-ports (port, protocol, to-port, to-addr), array of interfaces, array of sources, array of rich rules, array of protocols and array of source-ports (port, protocol).
//
// version (s): see version attribute of zone tag in firewalld.zone(5).
// name (s): see short tag in firewalld.zone(5).
// description (s): see description tag in firewalld.zone(5).
// UNUSED (b): this boolean value is no longer used for anything.
// target (s): see target attribute of zone tag in firewalld.zone(5).
// services (as): array of service names, see service tag in firewalld.zone(5).
// ports (a(ss)): array of port and protocol pairs. See port tag in firewalld.zone(5).
// icmp-blocks (as): array of icmp-blocks. See icmp-block tag in firewalld.zone(5).
// masquerade (b): see masquerade tag in firewalld.zone(5).
// forward-ports (a(ssss)): array of (port, protocol, to-port, to-addr). See forward-port tag in firewalld.zone(5).
// interfaces (as): array of interfaces. See interface tag in firewalld.zone(5).
// source addresses (as): array of source addresses. See source tag in firewalld.zone(5).
// rich rules (as): array of rich-language rules. See rule tag in firewalld.zone(5).
// protocols (as): array of protocols. See protocol tag in firewalld.zone(5).
// source-ports (a(ss)): array of port and protocol pairs. See source-port tag in firewalld.zone(5).
// Possible errors: NAME_CONFLICT, INVALID_NAME, INVALID_TYPE
func addZone() {}

// getHelperByName(s: helper) → o
// Return object path (permanent configuration) of helper with given name.
//
// Possible errors: INVALID_HELPER
func getHelperByName() {}

// getHelperNames() → as
// Return list of helper names (permanent configuration).
//
// getIPSetByName(s: ipset) → o
// Return object path (permanent configuration) of ipset with given name.
//
// Possible errors: INVALID_IPSET
func getHelperNames() {}

// getIPSetNames() → as
// Return list of ipset names (permanent configuration).
func getIPSetNames() {}

// getIcmpTypeByName(s: icmptype) → o
// Return object path (permanent configuration) of icmptype with given name.
//
// Possible errors: INVALID_ICMPTYPE
func getIcmpTypeByName() {}

// getIcmpTypeNames() → as
// Return list of icmptype names (permanent configuration).
func getIcmpTypeNames() {}

// getServiceByName(s: service) → o
// Return object path (permanent configuration) of service with given name.
//
// Possible errors: INVALID_SERVICE
func getServiceByName() {}

// getServiceNames() → as
// Return list of service names (permanent configuration).
func getServiceNames() {}

// getZoneByName(s: zone) → o
// Return object path (permanent configuration) of zone with given name.
//
// Possible errors: INVALID_ZONE
func getZoneByName() {}

// getZoneNames() → as
// Return list of zone names (permanent configuration) of.
func getZoneNames() {}

// getZoneOfInterface(s: iface) → s
// Return name of zone the iface is bound to or empty string.
func getZoneOfInterface() {}

// getZoneOfSource(s: source) → s
// Return name of zone the source is bound to or empty string.
func getZoneOfSource() {}

// listHelpers() → ao
// Return array of object paths (o) of helper in permanent configuration. For runtime configuration see org.fedoraproject.FirewallD1.Methods.getHelpers.
func listHelpers() {}

// listIPSets() → ao
// Return array of object paths (o) of ipset in permanent configuration. For runtime configuration see org.fedoraproject.FirewallD1.ipset.Methods.getIPSets.
func listIPSets() {}

// listIcmpTypes() → ao
// Return array of object paths (o) of icmp types in permanent configuration. For runtime configuration see org.fedoraproject.FirewallD1.Methods.listIcmpTypes.
func listIcmpTypes() {}

// listServices() → ao
// Return array of objects paths (o) of services in permanent configuration. For runtime configuration see org.fedoraproject.FirewallD1.Methods.listServices.
func listServices() {}

// listZones() → ao
// List object paths of zones known to permanent environment. For list of zones known to runtime environment see org.fedoraproject.FirewallD1.zone.Methods.getZones. The lists (of zones known to runtime and permanent environment) will contain same zones in most cases, but might differ for example if org.fedoraproject.FirewallD1.config.Methods.addZone has been called recently, but firewalld has not been reloaded since then.
func listZones() {}

// Signals

// HelperAdded(s: helper)
// Emitted when helper has been added.
func HelperAdded() {}

// IPSetAdded(s: ipset)
// Emitted when ipset has been added.
func IPSetAdded() {}

// IcmpTypeAdded(s: icmptype)
// Emitted when icmptype has been added.
func IcmpTypeAdded() {}

// ServiceAdded(s: service)
// Emitted when service has been added.
func ServiceAdded() {}

// ZoneAdded(s: zone)
// Emitted when zone has been added.
func ZoneAdded() {}

// Properties
// AutomaticHelpers - s - (rw)
// Indicates whether automatic helper assignment in kernel should be used or not. With the system setting this is left to the kernel or system default.
func getAutomaticHelpersProperty() {}
func setAutomaticHelpersProperty() {}

// CleanupOnExit - s - (rw)
// If firewalld stops, it cleans up all firewall rules. Setting this option to no or false leaves the current firewall rules untouched.
func getCleanupOnExitProperty() {}
func setCleanupOnExitProperty() {}

// DefaultZone - s - (ro)
// Default zone for connections or interfaces if the zone is not selected or specified by NetworkManager, initscripts or command line tool.
func getDefaultZoneProperty() {}

// FirewallBackend - s - (rw)
// Selects the firewalld backend for all rules except the direct interface. Valid options are; nftables, iptables. Default in nftables.
func getFirewallBackendProperty() {}
func setFirewallBackendProperty() {}

// IPv6_rpfilter - s - (rw)
// Indicates whether the reverse path filter test on a packet for IPv6 is enabled. If a reply to the packet would be sent via the same interface that the packet arrived on, the packet will match and be accepted, otherwise dropped.
func getIPv6RPFilterProperty() {}
func setIPv6RPFilterProperty() {}

// IndividualCalls - s - (ro)
// Indicates whether individual calls combined -restore calls are used. If enabled, this increases the time that is needed to apply changes and to start the daemon, but is good for debugging.
func getIndividualCallsProperty() {}

// Lockdown - s - (rw)
// If this property is enabled, firewall changes with the D-Bus interface will be limited to applications that are listed in the lockdown whitelist.
func getLockdownProperty() {}
func setLockdownProperty() {}

// LogDenied - s - (rw)
// If LogDenied is enabled, then logging rules are added right before reject and drop rules in the INPUT, FORWARD and OUTPUT chains for the default rules and also final reject and drop rules in zones. Possible values are: all, unicast, broadcast, multicast and off.
func getLogDeniedProperty() {}
func setLogDeniedProperty() {}

// MinimalMark - i - (rw)
// For some firewall settings several rules are needed in different tables to be able to handle packets in the correct way. To achieve that these packets are marked using the MARK target. With the MinimalMark property a block of marks can be reserved for private use; only marks over this value are used.
func getMinimalMarkProperty() {}
func setMinimalMarkProperty() {}
