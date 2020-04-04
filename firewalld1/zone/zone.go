package zone

// Methods
// addForwardPort(s: zone, s: port, s: protocol, s: toport, s: toaddr, i: timeout) → s
// Add the IPv4 forward port into zone. If zone is empty, use default zone. The port can either be a single port number portid or a port range portid-portid. The protocol can either be tcp or udp. The destination address is a simple IP address. If timeout is non-zero, the operation will be active only for the amount of seconds. For permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.addForwardPort.
//
// Returns name of zone to which the forward port was added.
//
// Possible errors: INVALID_ZONE, INVALID_PORT, MISSING_PROTOCOL, INVALID_PROTOCOL, INVALID_ADDR, INVALID_FORWARD, ALREADY_ENABLED, INVALID_COMMAND
func addForwardPort() {}

// addIcmpBlock(s: zone, s: icmp, i: timeout) → s
// Add an ICMP block icmp into zone. The icmp is the one of the icmp types firewalld supports. To get a listing of supported icmp types use org.fedoraproject.FirewallD1.Methods.listIcmpTypes If zone is empty, use default zone. If timeout is non-zero, the operation will be active only for the amount of seconds. For permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.addIcmpBlock.
//
// Returns name of zone to which the ICMP block was added.
//
// Possible errors: INVALID_ZONE, INVALID_ICMPTYPE, ALREADY_ENABLED, INVALID_COMMAND
func addIcmpBlock() {}

// addIcmpBlockInversion(s: zone) → s
// Add ICMP block inversion to zone. If zone is empty, use default zone. For permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.addIcmpBlockInversion.
//
// Returns name of zone to which the ICMP block inversion was added.
//
// Possible errors: INVALID_ZONE, ALREADY_ENABLED, INVALID_COMMAND
func addIcmpBlockInversion() {}

// addInterface(s: zone, s: interface) → s
// Bind interface with zone. From now on all traffic going through the interface will respect the zone's settings. If zone is empty, use default zone. For permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.addInterface.
//
// Returns name of zone to which the interface was bound.
//
// Possible errors: INVALID_ZONE, INVALID_INTERFACE, ALREADY_ENABLED, INVALID_COMMAND
func addInterface() {}

// addMasquerade(s: zone, i: timeout) → s
// Enable masquerade in zone. If zone is empty, use default zone. If timeout is non-zero, masquerading will be active for the amount of seconds. For permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.addMasquerade.
//
// Returns name of zone in which the masquerade was enabled.
//
// Possible errors: INVALID_ZONE, ALREADY_ENABLED, INVALID_COMMAND
func addMasquerade() {}

// addPort(s: zone, s: port, s: protocol, i: timeout) → s
// Add port into zone. If zone is empty, use default zone. The port can either be a single port number or a port range portid-portid. The protocol can either be tcp or udp. If timeout is non-zero, the operation will be active only for the amount of seconds. For permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.addPort.
//
// Returns name of zone to which the port was added.
//
// Possible errors: INVALID_ZONE, INVALID_PORT, MISSING_PROTOCOL, INVALID_PROTOCOL, ALREADY_ENABLED, INVALID_COMMAND
func addPort() {}

// addProtocol(s: zone, s: protocol, i: timeout) → s
// Add protocol into zone. If zone is empty, use default zone. The protocol can be any protocol supported by the system. Please have a look at /etc/protocols for supported protocols. If timeout is non-zero, the operation will be active only for the amount of seconds. For permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.addProtocol.
//
// Returns name of zone to which the protocol was added.
//
// Possible errors: INVALID_ZONE, INVALID_PROTOCOL, ALREADY_ENABLED, INVALID_COMMAND
func addProtocol() {}

// addRichRule(s: zone, s: rule, i: timeout) → s
// Add rich language rule into zone. For the rich language rule syntax, please have a look at firewalld.direct(5). If zone is empty, use default zone. If timeout is non-zero, the operation will be active only for the amount of seconds. For permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.addRichRule.
//
// Returns name of zone to which the rich language rule was added.
//
// Possible errors: INVALID_ZONE, INVALID_RULE, ALREADY_ENABLED, INVALID_COMMAND
func addRichRule() {}

// addService(s: zone, s: service, i: timeout) → s
// Add service into zone. If zone is empty, use default zone. If timeout is non-zero, the operation will be active only for the amount of seconds. To get a list of supported services, use org.fedoraproject.FirewallD1.Methods.listServices. For permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.addService.
//
// Returns name of zone to which the service was added.
//
// Possible errors: INVALID_ZONE, INVALID_SERVICE, ALREADY_ENABLED, INVALID_COMMAND
func addService() {}

// addSource(s: zone, s: source) → s
// Bind source with zone. From now on all traffic going from this source will respect the zone's settings. A source address or address range is either an IP address or a network IP address with a mask for IPv4 or IPv6. For IPv4, the mask can be a network mask or a plain number. For IPv6 the mask is a plain number. Use of host names is not supported. If zone is empty, use default zone. For permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.addSource.
//
// Returns name of zone to which the source was bound.
//
// Possible errors: INVALID_ZONE, INVALID_ADDR, ALREADY_ENABLED, INVALID_COMMAND
func addSource() {}

// addSourcePort(s: zone, s: port, s: protocol, i: timeout) → s
// Add source port into zone. If zone is empty, use default zone. The port can either be a single port number or a port range portid-portid. The protocol can either be tcp or udp. If timeout is non-zero, the operation will be active only for the amount of seconds. For permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.addSourcePort.
//
// Returns name of zone to which the port was added.
//
// Possible errors: INVALID_ZONE, INVALID_PORT, MISSING_PROTOCOL, INVALID_PROTOCOL, ALREADY_ENABLED, INVALID_COMMAND
func addSourcePort() {}

// changeZone(s: zone, s: interface) → s
// This function is deprecated, use org.fedoraproject.FirewallD1.zone.Methods.changeZoneOfInterface instead.
func changeZone() {}

// changeZoneOfInterface(s: zone, s: interface) → s
// Change a zone an interface is bound to to zone. It's basically removeInterface(interface) followed by addInterface(zone, interface). If interface has not been bound to a zone before, it behaves like addInterface. If zone is empty, use default zone.
//
// Returns name of zone to which the interface was bound.
//
// Possible errors: INVALID_ZONE, ZONE_ALREADY_SET, ZONE_CONFLICT
func changeZoneOfInterface() {}

// changeZoneOfSource(s: zone, s: source) → s
// Change a zone an source is bound to to zone. It's basically removeSource(source) followed by addSource(zone, source). If source has not been bound to a zone before, it behaves like addSource. If zone is empty, use default zone.
//
// Returns name of zone to which the source was bound.
//
// Possible errors: INVALID_ZONE, ZONE_ALREADY_SET, ZONE_CONFLICT
func changeZoneOfSource() {}

// getActiveZones() → a{sa{sas}}
// Return dictionary of currently active zones altogether with interfaces and sources used in these zones. Active zones are zones, that have a binding to an interface or source.
//
// Return value is a dictionary where keys are zone names (s) and values are again dictionaries where keys are either 'interfaces' or 'sources' and values are arrays of interface names (s) or sources (s).
func getActiveZones() {}

// getForwardPorts(s: zone) → aas
// Return array of IPv4 forward ports previously added into zone. If zone is empty, use default zone. For getting permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.getForwardPorts.
//
// Return value is array of 4-tuples, where each 4-tuple consists of (port, protocol, to-port, to-addr). to-addr might be empty in case of local forwarding.
//
// Possible errors: INVALID_ZONE
func getForwardPorts() {}

// getIcmpBlocks(s: zone) → as
// Return array of ICMP type (s) blocks previously added into zone. If zone is empty, use default zone. For getting permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.getIcmpBlocks.
//
// Possible errors: INVALID_ZONE
func getIcmpBlocks() {}

// getIcmpBlockInversion(s: zone) → b
// Return whether ICMP block inversion was previously added to zone. If zone is empty, use default zone. For getting permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.getIcmpBlockInversion.
//
// Possible errors: INVALID_ZONE
func getIcmpBlockInversion() {}

// getInterfaces(s: zone) → as
// Return array of interfaces (s) previously bound with zone. If zone is empty, use default zone. For getting permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.getInterfaces.
//
// Possible errors: INVALID_ZONE
func getInterfaces() {}

// getPorts(s: zone) → aas
// Return array of ports (2-tuple of port and protocol) previously enabled in zone. If zone is empty, use default zone. For getting permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.getPorts.
//
// Possible errors: INVALID_ZONE
func getPorts() {}

// getProtocols(s: zone) → as
// Return array of protocols (s) previously enabled in zone. If zone is empty, use default zone. For getting permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.getProtocols.
//
// Possible errors: INVALID_ZONE
func getProtocols() {}

// getRichRules(s: zone) → as
// Return array of rich language rules (s) previously added into zone. If zone is empty, use default zone. For getting permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.getRichRules.
//
// Possible errors: INVALID_ZONE
func getRichRules() {}

// getServices(s: zone) → as
// Return array of services (s) previously enabled in zone. If zone is empty, use default zone. For getting permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.getServices.
//
// Possible errors: INVALID_ZONE
func getServices() {}

// getSourcePorts(s: zone) → aas
// Return array of source ports (2-tuple of port and protocol) previously enabled in zone. If zone is empty, use default zone. For getting permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.getSourcePorts.
//
// Possible errors: INVALID_ZONE
func getSourcePorts() {}

// getSources(s: zone) → as
// Return array of sources (s) previously bound with zone. If zone is empty, use default zone. For getting permanent settings see org.fedoraproject.FirewallD1.config.zone.Methods.getSources.
//
// Possible errors: INVALID_ZONE
func getSources() {}

// getZoneOfInterface(s: interface) → s
// Return name (s) of zone the interface is bound to or empty string.
func getZoneOfInterface() {}

// getZoneOfSource(s: source) → s
// Return name (s) of zone the source is bound to or empty string.
func getZoneOfSource() {}

// getZones() → as
// Return array of names (s) of predefined zones known to current runtime environment. For list of zones known to permanent environment see org.fedoraproject.FirewallD1.config.Methods.listZones. The lists (of zones known to runtime and permanent environment) will contain same zones in most cases, but might differ for example if org.fedoraproject.FirewallD1.config.Methods.addZone has been called recently, but firewalld has not been reloaded since then.
func getZones() {}

// isImmutable(s: zone) → b
// Deprecated.
func isImmutable() {}

// queryForwardPort(s: zone, s: port, s: protocol, s: toport, s: toaddr) → b
// Return whether the IPv4 forward port (port, protocol, toport, toaddr) has been added into zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.queryForwardPort.
//
// Possible errors: INVALID_ZONE, INVALID_PORT, MISSING_PROTOCOL, INVALID_PROTOCOL, INVALID_ADDR, INVALID_FORWARD
func queryForwardPort() {}

// queryIcmpBlock(s: zone, s: icmp) → b
// Return whether an ICMP block for icmp has been added into zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.queryIcmpBlock.
//
// Possible errors: INVALID_ZONE, INVALID_ICMPTYPE
func queryIcmpBlock() {}

// queryIcmpBlockInversion(s: zone) → b
// Return whether ICMP block inversion has been added to zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.queryIcmpBlockInversion.
//
// Possible errors: INVALID_ZONE, INVALID_ICMPTYPE
func queryIcmpBlockInversion() {}

// queryInterface(s: zone, s: interface) → b
// Query whether interface has been bound to zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.queryInterface.
//
// Possible errors: INVALID_ZONE, INVALID_INTERFACE
func queryInterface() {}

// queryMasquerade(s: zone) → b
// Return whether masquerading has been enabled in zone If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.queryMasquerade.
//
// Possible errors: INVALID_ZONE
func queryMasquerade() {}

// queryPort(s: zone, s: port, s: protocol) → b
// Return whether port/protocol has been added in zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.queryPort.
//
// Possible errors: INVALID_ZONE, INVALID_PORT, MISSING_PROTOCOL, INVALID_PROTOCOL
func queryPort() {}

// queryProtocol(s: zone, s: protocol) → b
// Return whether protocol has been added in zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.queryProtocol.
//
// Possible errors: INVALID_ZONE, INVALID_PROTOCOL
func queryProtocol() {}

// queryRichRule(s: zone, s: rule) → b
// Return whether rich rule rule has been added in zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.queryRichRule.
//
// Possible errors: INVALID_ZONE, INVALID_RULE
func queryRichRule() {}

// queryService(s: zone, s: service) → b
// Return whether service has been added for zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.queryService.
//
// Possible errors: INVALID_ZONE, INVALID_SERVICE
func queryService() {}

// querySource(s: zone, s: source) → b
// Query whether sourcehas been bound to zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.querySource.
//
// Possible errors: INVALID_ZONE, INVALID_ADDR
func querySource() {}

// querySourcePort(s: zone, s: port, s: protocol) → b
// Return whether port/protocol has been added in zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.querySourcePort.
//
// Possible errors: INVALID_ZONE, INVALID_PORT, MISSING_PROTOCOL, INVALID_PROTOCOL
func querySourcePort() {}

// removeForwardPort(s: zone, s: port, s: protocol, s: toport, s: toaddr) → s
// Remove IPv4 forward port ((port, protocol, toport, toaddr)) from zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.removeForwardPort.
//
// Returns name of zone from which the forward port was removed.
//
// Possible errors: INVALID_ZONE, INVALID_PORT, MISSING_PROTOCOL, INVALID_PROTOCOL, INVALID_ADDR, INVALID_FORWARD, NOT_ENABLED, INVALID_COMMAND
func removeForwardPort() {}

// removeIcmpBlock(s: zone, s: icmp) → s
// Remove ICMP block icmp from zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.removeIcmpBlock.
//
// Returns name of zone from which the ICMP block was removed.
//
// Possible errors: INVALID_ZONE, INVALID_ICMPTYPE, NOT_ENABLED, INVALID_COMMAND
func removeIcmpBlock() {}

// removeIcmpBlockInversion(s: zone) → s
// Remove ICMP block inversion from zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.removeIcmpBlockInversion.
//
// Returns name of zone from which the ICMP block inversion was removed.
//
// Possible errors: INVALID_ZONE, NOT_ENABLED, INVALID_COMMAND
func removeIcmpBlockInversion() {}

// removeInterface(s: zone, s: interface) → s
// Remove binding of interface from zone. If zone is empty, the interface will be removed from zone it belongs to. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.removeInterface.
//
// Returns name of zone from which the interface was removed.
//
// Possible errors: INVALID_ZONE, INVALID_INTERFACE, NOT_ENABLED, INVALID_COMMAND
func removeInterface() {}

// removeMasquerade(s: zone) → s
// Disable masquerade for zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.removeMasquerade.
//
// Returns name of zone for which the masquerade was disabled.
//
// Possible errors: INVALID_ZONE, NOT_ENABLED, INVALID_COMMAND
func removeMasquerade() {}

// removePort(s: zone, s: port, s: protocol) → s
// Remove port/protocol from zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.removePort.
//
// Returns name of zone from which the port was removed.
//
// Possible errors: INVALID_ZONE, INVALID_PORT, MISSING_PROTOCOL, INVALID_PROTOCOL, NOT_ENABLED, INVALID_COMMAND
func removePort() {}

// removeProtocol(s: zone, s: protocol) → s
// Remove protocol from zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.removeProtocol.
//
// Returns name of zone from which the protocol was removed.
//
// Possible errors: INVALID_ZONE, INVALID_PROTOCOL, NOT_ENABLED, INVALID_COMMAND
func removeProtocol() {}

// removeRichRule(s: zone, s: rule) → s
// Remove rich language rule from zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.removeRichRule.
//
// Returns name of zone from which the rich language rule was removed.
//
// Possible errors: INVALID_ZONE, INVALID_RULE, NOT_ENABLED, INVALID_COMMAND
func removeRichRule() {}

// removeService(s: zone, s: service) → s
// Remove service from zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.removeService.
//
// Returns name of zone from which the service was removed.
//
// Possible errors: INVALID_ZONE, INVALID_SERVICE, NOT_ENABLED, INVALID_COMMAND
func removeService() {}

// removeSource(s: zone, s: source) → s
// Remove binding of source from zone. If zone is empty, the source will be removed from zone it belongs to. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.removeSource.
//
// Returns name of zone from which the source was removed.
//
// Possible errors: INVALID_ZONE, INVALID_ADDR, NOT_ENABLED, INVALID_COMMAND
func removeSource() {}

// removeSourcePort(s: zone, s: port, s: protocol) → s
// Remove port/protocol from zone. If zone is empty, use default zone. For permanent operation see org.fedoraproject.FirewallD1.config.zone.Methods.removeSourcePort.
//
// Returns name of zone from which the source port was removed.
//
// Possible errors: INVALID_ZONE, INVALID_PORT, MISSING_PROTOCOL, INVALID_PROTOCOL, NOT_ENABLED, INVALID_COMMAND
func removeSourcePort() {}

// Signals
// ForwardPortAdded(s: zone, s: port, s: protocol, s: toport, s: toaddr, i: timeout)
// Emitted when forward port has been added to zone with timeout.
func ForwardPortAdded() {}

// ForwardPortRemoved(s: zone, s: port, s: protocol, s: toport, s: toaddr)
// Emitted when forward port has been removed from zone.
func ForwardPortRemoved() {}

// IcmpBlockAdded(s: zone, s: icmp, i: timeout)
// Emitted when ICMP block for icmp has been added to zone with timeout.
func IcmpBlockAdded() {}

// IcmpBlockInversionAdded(s: zone)
// Emitted when ICMP block inversion has been added to zone.
func IcmpBlockInversionAdded() {}

// IcmpBlockInversionRemoved(s: zone)
// Emitted when ICMP block inversion has been removed from zone.
func IcmpBlockInversionRemoved() {}

// IcmpBlockRemoved(s: zone, s: icmp)
// Emitted when ICMP block for icmp has been removed from zone.
func IcmpBlockRemoved() {}

// InterfaceAdded(s: zone, s: interface)
// Emitted when interface has been added to zone.
func InterfaceAdded() {}

// InterfaceRemoved(s: zone, s: interface)
// Emitted when interface has been removed from zone.
func InterfaceRemoved() {}

// MasqueradeAdded(s: zone, i: timeout)
// Emitted when masquerade has been enabled for zone.
func MasqueradeAdded() {}

// MasqueradeRemoved(s: zone)
// Emitted when masquerade has been disabled for zone.
func MasqueradeRemoved() {}

// PortAdded(s: zone, s: port, s: protocol, i: timeout)
// Emitted when port/protocol has been added to zone with timeout.
func PortAdded() {}

// PortAdded(s: zone, s: port, s: protocol)
// Emitted when port/protocol has been removed from zone.
// func PortAdded() {}

// ProtocolAdded(s: zone, s: protocol, i: timeout)
// Emitted when protocol has been added to zone with timeout.
func ProtocolAdded() {}

// ProtocolRemoved(s: zone, s: protocol)
// Emitted when protocol has been removed from zone.
func ProtocolRemoved() {}

// RichRuleAdded(s: zone, s: rule, i: timeout)
// Emitted when rich language rule has been added to zone with timeout.
func RichRuleAdded() {}

// RichRuleRemoved(s: zone, s: rule)
// Emitted when rich language rule has been removed from zone.
func RichRuleRemoved() {}

// ServiceAdded(s: zone, s: service, i: timeout)
// Emitted when service has been added to zone with timeout.
func ServiceAdded() {}

// ServiceRemoved(s: zone, s: service)
// Emitted when service has been removed from zone.
func ServiceRemoved() {}

// SourceAdded(s: zone, s: source)
// Emitted when source has been added to zone.
func SourceAdded() {}

// SourcePortAdded(s: zone, s: port, s: protocol, i: timeout)
// Emitted when source-port/protocol has been added to zone with timeout.
func SourcePortAdded() {}

// SourcePortRemoved(s: zone, s: port, s: protocol)
// Emitted when source-port/protocol has been removed from zone.
func SourcePortRemoved() {}

// SourceRemoved(s: zone, s: source)
// Emitted when source has been removed from zone.
func SourceRemoved() {}

// ZoneChanged(s: zone, s: interface)
// Deprecated
func ZoneChanged() {}

// ZoneOfInterfaceChanged(s: zone, s: interface)
// Emitted when a zone an interface is part of has been changed to zone.
func ZoneOfInterfaceChanged() {}

// ZoneOfSourceChanged(s: zone, s: source)
// Emitted when a zone an source is part of has been changed to zone.
func ZoneOfSourceChanged() {}
