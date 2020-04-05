package firewald1

// authorizeAll() → Nothing
// Initiate authorization for the complete firewalld D-Bus interface. This method it mostly useful for configuration applications.
//
func authorizeAll() {

}
// completeReload() → Nothing
// Reload firewall completely, even netfilter kernel modules. This will most likely terminate active connections, because state information is lost. This option should only be used in case of severe firewall problems. For example if there are state information problems that no connection can be established with correct firewall rules.
//
func completeReload() {}

//disablePanicMode() → Nothing
//Disable panic mode. After disabling panic mode established connections might work again, if panic mode was enabled for a short period of time.
//
//Possible errors: NOT_ENABLED, COMMAND_FAILED
func disablePanicMode() {}

//enablePanicMode() → Nothing
//Enable panic mode. All incoming and outgoing packets are dropped, active connections will expire. Enable this only if there are serious problems with your network environment.
//
//Possible errors: ALREADY_ENABLED, COMMAND_FAILED

func enablePanicMode() {}

// getAutomaticHelpers() → s
// Return the AutomaticHelpers value. For the secure use of iptables and connection tracking helpers it is recommended to turn AutomaticHelpers off. But this might have side effects on other services using the netfilter helpers as the sysctl setting in /proc/sys/net/netfilter/nf_conntrack_helper will be changed. With the system setting, the default value set in the kernel or with sysctl will be used. Possible values are: yes, no and system. The default value is system.
func getAutomaticHelpers() {}

// getDefaultZone() → s
// Return default zone.

func getDefaultZone() {}

// getHelperSettings(s: helper) → (sssssa(ss))
// Return runtime settings of given helper. For getting permanent settings see org.fedoraproject.FirewallD1.config.helper.Methods.getSettings. Settings are in format: version, name, description, family, module and array of ports.
//
// version (s): see version attribute of helper tag in firewalld.helper(5).
// name (s): see short tag in firewalld.helper(5).
// description (s): see description tag in firewalld.helper(5).
// family (s): see family tag in firewalld.helper(5).
// module (s): see module tag in firewalld.helper(5).
// ports (a(ss)): array of port and protocol pairs. See port tag in firewalld.helper(5).
// Possible errors: INVALID_HELPER
func getHelperSettings() {}

// getHelpers() → as
// Return array of helper names (s) in runtime configuration. For permanent configuration see org.fedoraproject.FirewallD1.config.Methods.listHelpers.

func getHelpers() {}

// getIcmpTypeSettings(s: icmptype) → (sssas)
// Return runtime settings of given icmptype. For getting permanent settings see org.fedoraproject.FirewallD1.config.icmptype.Methods.getSettings. Settings are in format: version, name, description, array of destinations.
// 
// version (s): see version attribute of icmptype tag in firewalld.icmptype(5).
// name (s): see short tag in firewalld.icmptype(5).
// description (s): see description tag in firewalld.icmptype(5).
// destinations (as): array, either empty or containing strings 'ipv4' or 'ipv6', see destination tag in firewalld.icmptype(5).
// Possible errors: INVALID_ICMPTYPE

func getIcmpTypeSettings() {

}

// getLogDenied() → s
// Retruns the LogDenied value. If LogDenied is enabled, then logging rules are added right before reject and drop rules in the INPUT, FORWARD and OUTPUT chains for the default rules and also final reject and drop rules in zones. Possible values are: all, unicast, broadcast, multicast and off. The default value is off

func getLogDenied() {}

// getServiceSettings(s: service) → (sssa(ss)asa{ss}asa(ss))
// Return runtime settings of given service. For getting permanent settings see org.fedoraproject.FirewallD1.config.service.Methods.getSettings. Settings are in format: version, name, description, array of ports (port, protocol), array of module names, dictionary of destinations, array of protocols and array of source-ports (port, protocol).
//
// version (s): see version attribute of service tag in firewalld.service(5).
// name (s): see short tag in firewalld.service(5).
// description (s): see description tag in firewalld.service(5).
// ports (a(ss)): array of port and protocol pairs. See port tag in firewalld.service(5).
// module names (as): array of kernel netfilter helpers, see module tag in firewalld.service(5).
// destinations (a{ss}): dictionary of {IP family : IP address} where 'IP family' key can be either 'ipv4' or 'ipv6'. See destination tag in firewalld.service(5).
// protocols (as): array of protocols, see protocol tag in firewalld.service(5).
// source-ports (a(ss)): array of port and protocol pairs. See source-port tag in firewalld.service(5).
// Possible errors: INVALID_SERVICE

func getServiceSettings() {}

// getZoneSettings(s: zone) → (sssbsasa(ss)asba(ssss)asasasasa(ss))
// Return runtime settings of given zone. For getting permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.getSettings. Settings are in format: version, name, description, UNUSED, target, array of services, array of ports (port, protocol), array of icmp-blocks, masquerade, array of forward-ports (port, protocol, to-port, to-addr), array of interfaces, array of sources, array of rich rules, array of protocols and array of source-ports (port, protocol).

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
// protocols (as): array of protocols, see protocol tag in firewalld.zone(5).
// source-ports (a(ss)): array of port and protocol pairs. See source-port tag in firewalld.zone(5).
// Possible errors: INVALID_ZONE

func getZoneSettings() {}

// listIcmpTypes() → as
// Return array of names (s) of icmp types in runtime configuration. For permanent configuration see org.fedoraproject.FirewallD1.config.Methods.listIcmpTypes.

func listIcmpTypes() {}

// listServices() → as
// Return array of service names (s) in runtime configuration. For permanent configuration see org.fedoraproject.FirewallD1.config.Methods.listServices.

func listServices() {}

// queryPanicMode() → b
// Return true if panic mode is enabled, false otherwise. In panic mode all incoming and outgoing packets are dropped.

func queryPanicMode() {}

// reload() → Nothing
// Reload firewall rules and keep state information. Current permanent configuration will become new runtime configuration, i.e. all runtime only changes done until reload are lost with reload if they have not been also in permanent configuration.

func reload() {}

// runtimeToPermanent() → Nothing
// Make runtime settings permanent. Replaces permanent settings with runtime settings for zones, services, icmptypes, direct and policies (lockdown whitelist).
//
// Possible errors: RT_TO_PERM_FAILED
func runtimeToPermanent() {}

// checkPermanentConfig() → Nothing
// Run checks on the permanent configuration. This is most useful if changes were made manually to configuration files.
//
//Possible errors: any

func checkPermanentConfig() {}

// setDefaultZone(s: zone) → Nothing
// Set default zone for connections and interfaces where no zone has been selected to zone. Setting the default zone changes the zone for the connections or interfaces, that are using the default zone. This is a runtime and permanent change.
//
// Possible errors: ZONE_ALREADY_SET, COMMAND_FAILED
func setDefaultZone() {}

// setLogDenied(s: value) → Nothing
//Set LogDenied value to value. If LogDenied is enabled, then logging rules are added right before reject and drop rules in the INPUT, FORWARD and OUTPUT chains for the default rules and also final reject and drop rules in zones. Possible values are: all, unicast, broadcast, multicast and off. The default value is off This is a runtime and permanent change.
//
//Possible errors: ALREADY_SET, INVALID_VALUE

func setLogDenied() {}

package firewalld1

// DefaultZoneChanged(s: zone)
// Emitted when default zone has been changed to zone.
func DefaultZoneChanged() {}

// LogDeniedChanged(s: value)
// Emitted when LogDenied value has been changed.
func LogDeniedChanged() {}

// PanicModeDisabled()
// Emitted when panic mode has been deactivated.
func PanicModeDisabled() {}

// PanicModeEnabled()
// Emitted when panic mode has been activated.
func PanicModeEnabled() {}

// Reloaded()
// Emitted when firewalld has been reloaded. Also emitted for a complete reload.
func Reloaded() {}

package firewalld1

// BRIDGE - b - (ro)
// Indicates whether the firewall has ethernet bridge support.
func getBridgeProperty() {}

// IPSet - b - (ro)
// Indicates whether the firewall has IPSet support.
func getIPSetProperty() {}

// IPSetTypes - as - (ro)
// The supported IPSet types by ipset and firewalld.
func getIPSetTypesProperties() {}

// IPv4 - b - (ro)
// Indicates whether the firewall has IPv4 support.
func getIPv4Property() {}

// IPv4ICMPTypes - as - (ro)
// The list of supported IPv4 ICMP types.
func getIPv4ICMPTypes() {}

// IPv6 - b - (ro)
// Indicates whether the firewall has IPv6 support.
func getIPv6Property() {}

// IPv6_rpfilter - b - (ro)
// Indicates whether the reverse path filter test on a packet for IPv6 is enabled. If a reply to the packet would be sent via the same interface that the packet arrived on, the packet will match and be accepted, otherwise dropped.
func getIPv6rpfilterProperty() {}

// IPv6ICMPTypes - as - (ro)
// The list of supported IPv6 ICMP types.
func getIPv6ICMPTypesProperties() {}

// nf_conntrach_helper_setting - b - (ro)
// Kernel nf_conntrack_helper setting.
func getNFConntrackHelperProperty() {}

// nf_conntrack_helpers - a{sas} - (ro)
// The list of conntrack helpers supported by the kernel.
func getNFConntrackHelpersProperties() {}

// nf_nat_helpers - a{sas} - (ro)
// The list of nat helpers supported by the kernel.
func getNFNatHelpersProperties() {}

// interface_version - s - (ro)
// firewalld D-Bus interface version string.
func getInterfaceVersionProperty() {}

// state - s - (ro)
// firewalld state. This can be either INIT or RUNNING. In INIT state, firewalld is starting up and initializing.
func getStateProperty() {}

// version - s - (ro)
// firewalld version string.
func getVersionProperty() {}
