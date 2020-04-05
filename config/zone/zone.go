package zone

// Methods
// addForwardPort(s: port, s: protocol, s: toport, s: toaddr) → Nothing
// Permanently add (port, protocol, toport, toaddr) to list of forward ports of zone. See forward-port tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.addForwardPort.
//
// Possible errors: ALREADY_ENABLED
func addForwardPort() {}

// addIcmpBlock(s: icmptype) → Nothing
// Permanently add icmptype to list of icmp types blocked in zone. See icmp-block tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.addIcmpBlock.
//
// Possible errors: ALREADY_ENABLED
func addIcmpBlock() {}

// addIcmpBlock(s: icmptype) → Nothing
// Permanently add icmp block inversion to zone. See icmp-block-inversion tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.addIcmpBlockInversion.
//
// Possible errors: ALREADY_ENABLED
//func addIcmpBlock() {}

// addInterface(s: interface) → Nothing
// Permanently add interface to list of interfaces bound to zone. See interface tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.addInterface.
//
// Possible errors: ALREADY_ENABLED
func addInterface() {}

// addMasquerade() → Nothing
// Permanently enable masquerading in zone. See masquerade tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.addMasquerade.
//
// Possible errors: ALREADY_ENABLED
func addMasquerade() {}

// addPort(s: port, s: protocol) → Nothing
// Permanently add (port, protocol) to list of ports of zone. See port tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.addPort.
//
// Possible errors: ALREADY_ENABLED
func addPort() {}

// addProtocol(s: protocol) → Nothing
// Permanently add protocol into zone. The protocol can be any protocol supported by the system. Please have a look at /etc/protocols for supported protocols. For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.addProtocol.
//
// Possible errors: INVALID_PROTOCOL, ALREADY_ENABLED
func addProtocol() {}

// addRichRule(s: rule) → Nothing
// Permanently add rule to list of rich-language rules in zone. See rule tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.addRichRule.
//
// Possible errors: ALREADY_ENABLED
func addRichRule() {}

// addService(s: service) → Nothing
// Permanently add service to list of services used in zone. See service tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.addService.
//
// Possible errors: ALREADY_ENABLED
func addService() {}

// addSource(s: source) → Nothing
// Permanently add source to list of source addresses bound to zone. See source tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.addSource.
//
// Possible errors: ALREADY_ENABLED
func addSource() {}

// addSourcePort(s: port, s: protocol) → Nothing
// Permanently add (port, protocol) to list of source ports of zone. See source-port tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.addSourcePort.
//
// Possible errors: ALREADY_ENABLED
func addSourcePort() {}

// getDescription() → s
// Get description of zone. See description tag in firewalld.zone(5).
func getDescription() {}

// getForwardPorts() → a(ssss)
// Get list of (port, protocol, toport, toaddr) defined in zone. See forward-port tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.getForwardPorts.
func getForwardPorts() {}

// getIcmpBlockInversion() → b
// Get icmp block inversion flag of zone. See icmp-block-inversion tag in firewalld.zone(5).
func getIcmpBlockInversion() {}

// getIcmpBlocks() → as
// Get list of icmp type names blocked in zone. See icmp-block tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.getIcmpBlocks.
func getIcmpBlocks() {}

// getInterfaces() → as
// Get list of interfaces bound to zone. See interface tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.getInterfaces.
func getInterfaces() {}

// getMasquerade() → b
// Return whether masquerade is enabled in zone. This is the same as queryMasquerade() method. See masquerade tag in firewalld.zone(5).
func getMasquerade() {}

// getPorts() → a(ss)
// Get list of (port, protocol) defined in zone. See port tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.getPorts.
func getPorts() {}

// getProtocols() → as
// Return array of protocols (s) previously enabled in zone. For getting runtime settings see org.fedoraproject.FirewallD1.zone.Methods.getProtocols.
func getProtocols() {}

// getRichRules() → as
// Get list of rich-language rules in zone. See rule tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.getRichRules.
func getRichRules() {}

// getServices() → as
// Get list of service names used in zone. See service tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.getServices.
func getServices() {}

// getSettings() → (sssbsasa(ss)asba(ssss)asasasasa(ss))
// Return permanent settings of given zone. For getting runtime settings see org.fedoraproject.FirewallD1.Methods.getZoneSettings. Settings are in format: version, name, description, UNUSED, target, array of services, array of ports (port, protocol), array of icmp-blocks, masquerade, array of forward-ports (port, protocol, to-port, to-addr), array of interfaces, array of sources, array of rich rules, array of protocols and array of source-ports (port, protocol).
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
func getSettings() {}

// getShort() → s
// Get name of zone. See short tag in firewalld.zone(5).
func getShort() {}

// getSourcePorts() → a(ss)
// Get list of (port, protocol) defined in zone. See source-port tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.getSourcePorts.
func getSourcePorts() {}

// getSources() → as
// Get list of source addresses bound to zone. See source tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.getSources.
func getSources() {}

// getTarget() → s
// Get target of zone. See target attribute of zone tag in firewalld.zone(5).
func getTarget() {}

// getVersion() → s
// Get version of zone. See version attribute of zone tag in firewalld.zone(5).
func getVersion() {}

// loadDefaults() → Nothing
// Load default settings for built-in zone.
//
// Possible errors: NO_DEFAULTS
func loadDefaults() {}

// queryForwardPort(s: port, s: protocol, s: toport, s: toaddr) → b
// Return whether (port, protocol, toport, toaddr) is in list of forward ports of zone. See forward-port tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.queryForwardPort.
func queryForwardPort() {}

// queryIcmpBlock(s: icmptype) → b
// Return whether icmptype is in list of icmp types blocked in zone. See icmp-block tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.queryIcmpBlock.
func queryIcmpBlock() {}

// queryIcmpBlockInversion() → b
// Return whether icmp block inversion is in enabled in zone. See icmp-block-inversion tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.queryIcmpBlockInversion.
func queryIcmpBlockInversion() {}

// queryInterface(s: interface) → b
// Return whether interface is in list of interfaces bound to zone. See interface tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.queryInterface.
func queryInterface() {}

// queryMasquerade() → b
// Return whether masquerade is enabled in zone. This is the same as getMasquerade() method. See masquerade tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.queryMasquerade.
func queryMasquerade() {}

// queryPort(s: port, s: protocol) → b
// Return whether (port, protocol) is in list of ports of zone. See port tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.queryPort.
func queryPort() {}

// queryProtocol(s: protocol) → b
// Return whether protocol has been added in zone. For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.queryProtocol.
//
// Possible errors: INVALID_PROTOCOL
func queryProtocol() {}

// queryRichRule(s: rule) → b
// Return whether rule is in list of rich-language rules in zone. See rule tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.queryRichRule.
func queryRichRule() {}

// queryService(s: service) → b
// Return whether service is in list of services used in zone. See service tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.queryService.
func queryService() {}

// querySource(s: source) → b
// Return whether source is in list of source addresses bound to zone. See source tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.querySource.
func querySource() {}

// querySourcePort(s: port, s: protocol) → b
// Return whether (port, protocol) is in list of source ports of zone. See source-port tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.querySourcePort.
func querySourcePort() {}

// remove() → Nothing
// Remove not built-in zone.
//
// Possible errors: BUILTIN_ZONE
func remove() {}

// removeForwardPort(s: port, s: protocol, s: toport, s: toaddr) → Nothing
// Permanently remove (port, protocol, toport, toaddr) from list of forward ports of zone. See forward-port tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.removeForwardPort.
//
// Possible errors: NOT_ENABLED
func removeForwardPort() {}

// removeIcmpBlock(s: icmptype) → Nothing
// Permanently remove icmptype from list of icmp types blocked in zone. See icmp-block tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.removeIcmpBlock.
//
// Possible errors: NOT_ENABLED
func removeIcmpBlock() {}

// removeIcmpBlockInversion() → Nothing
// Permanently remove icmp block inversion from the zone. See icmp-block-inversion tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.removeIcmpBlockInversion.
//
// Possible errors: NOT_ENABLED
func removeIcmpBlockInversion() {}

// removeInterface(s: interface) → Nothing
// Permanently remove interface from list of interfaces bound to zone. See interface tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.removeInterface.
//
// Possible errors: NOT_ENABLED
func removeInterface() {}

// removeMasquerade() → Nothing
// Permanently disable masquerading in zone. See masquerade tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.removeMasquerade.
//
// Possible errors: NOT_ENABLED
func removeMasquerade() {}

// removePort(s: port, s: protocol) → Nothing
// Permanently remove (port, protocol) from list of ports of zone. See port tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.removePort.
//
// Possible errors: NOT_ENABLED
func removePort() {}

// removeProtocol(s: protocol) → Nothing
// Permanently remove protocol from zone. For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.removeProtocol.
//
// Possible errors: INVALID_PROTOCOL, NOT_ENABLED
func removeProtocol() {}

// removeRichRule(s: rule) → Nothing
// Permanently remove rule from list of rich-language rules in zone. See rule tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.removeRichRule.
//
// Possible errors: NOT_ENABLED
func removeRichRule() {}

// removeService(s: service) → Nothing
// Permanently remove service from list of services used in zone. See service tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.removeService.
//
// Possible errors: NOT_ENABLED
func removeService() {}

// removeSource(s: source) → Nothing
// Permanently remove source from list of source addresses bound to zone. See source tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.removeSource.
//
// Possible errors: NOT_ENABLED
func removeSource() {}

// removeSourcePort(s: port, s: protocol) → Nothing
// Permanently remove (port, protocol) from list of source ports of zone. See source-port tag in firewalld.zone(5). For runtime operation see org.fedoraproject.FirewallD1.zone.Methods.removeSourcePort.
//
// Possible errors: NOT_ENABLED
func removeSourcePort() {}

// rename(s: name) → Nothing
// Rename not built-in zone to name.
//
// Possible errors: BUILTIN_ZONE
func rename() {}

// setDescription(s: description) → Nothing
// Permanently set description of zone to description. See description tag in firewalld.zone(5).
func setDescription() {}

// setForwardPorts(a(ssss): ports) → Nothing
// Permanently set forward ports of zone to list of (port, protocol, toport, toaddr). See forward-port tag in firewalld.zone(5).
func setForwardPorts() {}

// setIcmpBlockInversion(b: flag) → Nothing
// Permanently set icmp block inversion flag of zone to flag. See icmp-block-inversion tag in firewalld.zone(5).
func setIcmpBlockInversion() {}

// setIcmpBlocks(as: icmptypes) → Nothing
// Permanently set list of icmp types blocked in zone to icmptypes. See icmp-block tag in firewalld.zone(5).
func setIcmpBlocks() {}

// setInterfaces(as: interfaces) → Nothing
// Permanently set list of interfaces bound to zone to interfaces. See interface tag in firewalld.zone(5).
func setInterfaces() {}

// setMasquerade(b: masquerade) → Nothing
// Permanently set masquerading in zone to masquerade. See masquerade tag in firewalld.zone(5).
func setMasquerade() {}

// setPorts(a(ss): ports) → Nothing
// Permanently set ports of zone to list of (port, protocol). See port tag in firewalld.zone(5).
func setPorts() {}

// setProtocols(as: protocols) → Nothing
// Permanently set list of protocols used in zone to protocols. See protocol tag in firewalld.zone(5).
func setProtocols() {}

// setRichRules(as: rules) → Nothing
// Permanently set list of rich-language rules to rules. See rule tag in firewalld.zone(5).
func setRichRules() {}

// setServices(as: services) → Nothing
// Permanently set list of services used in zone to services. See service tag in firewalld.zone(5).
func setServices() {}

// setShort(s: short) → Nothing
// Permanently set name of zone to short. See short tag in firewalld.zone(5).
func setShort() {}

// setSourcePorts(a(ss): ports) → Nothing
// Permanently set source-ports of zone to list of (port, protocol). See source-port tag in firewalld.zone(5).
func setSourcePorts() {}

// setSources(as: sources) → Nothing
// Permanently set list of source addresses bound to zone to sources. See source tag in firewalld.zone(5).
func setSources() {}

// setTarget(s: target) → Nothing
// Permanently set target of zone to target. See target attribute of zone tag in firewalld.zone(5).
func setTarget() {}

// setVersion(s: version) → Nothing
// Permanently set version of zone to version. See version attribute of zone tag in firewalld.zone(5).
func setVersion() {}

// update((sssbsasa(ss)asba(ssss)asasasasa(ss)): settings) → Nothing
// Update settings of zone to settings. Settings are in format: version, name, description, UNUSED, target, array of services, array of ports (port, protocol), array of icmp-blocks, masquerade, array of forward-ports (port, protocol, to-port, to-addr), array of interfaces, array of sources, array of rich rules, array of protocols and array of source-ports (port, protocol).
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
// Possible errors: INVALID_TYPE
func update() {}

// Signals
// Removed(s: name)
// Emitted when zone with name has been removed.
func removed() {}

// Renamed(s: name)
// Emitted when zone has been renamed to name.
func renamed() {}

// Updated(s: name)
// Emitted when zone with name has been updated.
func updated() {}

// Properties
// builtin - b - (ro)
// True if zone is build-in, false else.
func isZoneBuiltin() {}

// default - b - (ro)
// True if build-in zone has default settings. False if it has been modified. Always False for not build-in zones.
func isZoneDefault() {}

// filename - s - (ro)
// Name (including .xml extension) of file where the configuration is stored.
func getFilename() {}

// name - s - (ro)
// Name of zone.
func getName() {}

// path - s - (ro)
// Path to directory where the zone configuration is stored. Should be either /usr/lib/firewalld/zones or /etc/firewalld/zones.
func getPath() {}
