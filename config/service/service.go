package configservice

// Interface for permanent service configuration, see also firewalld.service(5).
// Methods
// addModule(s: module) → Nothing
// Permanently add module to list of modules (netfilter kernel helpers) used in service. See module tag in firewalld.service(5).
//
// Possible errors: ALREADY_ENABLED
func addModule() {}

// addPort(s: port, s: protocol) → Nothing
// Permanently add (port, protocol) to list of ports in service. See port tag in firewalld.service(5).
//
// Possible errors: ALREADY_ENABLED
func addPort() {}

// addProtocol(s: protocol) → Nothing
// Permanently add protocol into zone. The protocol can be any protocol supported by the system. Please have a look at /etc/protocols for supported protocols. See protocol tag in firewalld.service(5).
//
// Possible errors: INVALID_PROTOCOL, ALREADY_ENABLED
func addProtocol() {}

// addSourcePort(s: port, s: protocol) → Nothing
// Permanently add (port, protocol) to list of source ports in service. See source-port tag in firewalld.service(5).
//
// Possible errors: ALREADY_ENABLED
func addSourcePort() {}

// getDescription() → s
// Get description of service. See description tag in firewalld.service(5).
func getDescription() {}

// getDestination(s: family) → s
// Get destination for IP family being either 'ipv4' or 'ipv6'. See destination tag in firewalld.service(5).
//
// Possible errors: ALREADY_ENABLED
func getDestination() {}

// getDestinations() → a{ss}
// Get list of destinations. Return value is a dictionary of {IP family : IP address} where 'IP family' key can be either 'ipv4' or 'ipv6'. See destination tag in firewalld.service(5).
func getDestinations() {}

// getModules() → as
// Get list of modules (netfilter kernel helpers) used in service. See module tag in firewalld.service(5).
func getModules() {}

// getPorts() → a(ss)
// Get list of (port, protocol) defined in service. See port tag in firewalld.service(5).
func getPorts() {}

// getProtocols() → as
// Return array of protocols (s) defined in service. See protocol tag in firewalld.service(5).
func getProtocols() {}

// getSettings() → (sssa(ss)asa{ss}asa(ss))
// Return permanent settings of a service. For getting runtime settings see org.fedoraproject.FirewallD1.Methods.getServiceSettings. Settings are in format: version, name, description, array of ports (port, protocol), array of module names, dictionary of destinations, array of protocols and array of source-ports (port, protocol).
//
// version (s): see version attribute of service tag in firewalld.service(5).
// name (s): see short tag in firewalld.service(5).
// description (s): see description tag in firewalld.service(5).
// ports (a(ss)): array of port and protocol pairs. See port tag in firewalld.service(5).
// module names (as): array of kernel netfilter helpers, see module tag in firewalld.service(5).
// destinations (a{ss}): dictionary of {IP family : IP address} where 'IP family' key can be either 'ipv4' or 'ipv6'. See destination tag in firewalld.service(5).
// protocols (as): array of protocols. See protocol tag in firewalld.service(5).
// source-ports (a(ss)): array of port and protocol pairs. See source-port tag in firewalld.service(5).
func getSettings() {}

// getShort() → s
// Get name of service. See short tag in firewalld.service(5).
func getShort() {}

// getSourcePorts() → a(ss)
// Get list of (port, protocol) defined in service. See source-port tag in firewalld.service(5).
func getSourcePorts() {}

// getVersion() → s
// Get version of service. See version attribute of service tag in firewalld.service(5).
func getVersion() {}

// loadDefaults() → Nothing
// Load default settings for built-in service.
//
// Possible errors: NO_DEFAULTS
func loadDefaults() {}

// queryDestination(s: family, s: address) → b
// Return whether a destination is in dictionary of destinations of this service. destination is in format: (IP family, IP address) where IP family can be either 'ipv4' or 'ipv6'. See destination tag in firewalld.service(5).
func queryDestination() {}

// queryModule(s: module) → b
// Return whether module is in list of modules (netfilter kernel helpers) used in service. See module tag in firewalld.service(5).
func queryModule() {}

// queryPort(s: port, s: protocol) → b
// Return whether (port, protocol) is in list of ports in service. See port tag in firewalld.service(5).
func queryPort() {}

// queryProtocol(s: protocol) → b
// Return whether protocol is in list of protocols in service. See protocol tag in firewalld.service(5).
func queryProtocol() {}

// querySourcePort(s: port, s: protocol) → b
// Return whether (port, protocol) is in list of source ports in service. See source-port tag in firewalld.service(5).
func querySourcePort() {}

// remove() → Nothing
// Remove not built-in service.
//
// Possible errors: BUILTIN_SERVICE
func remove() {}

// removeDestination(s: family) → Nothing
// Permanently remove a destination with family ('ipv4' or 'ipv6') from service. See destination tag in firewalld.service(5).
//
// Possible errors: NOT_ENABLED
func removeDestination() {}

// removeModule(s: module) → Nothing
// Permanently remove module from list of modules (netfilter kernel helpers) used in service. See module tag in firewalld.service(5).
//
// Possible errors: NOT_ENABLED
func removeModule() {}

// removePort(s: port, s: protocol) → Nothing
// Permanently remove (port, protocol) from list of ports in service. See port tag in firewalld.service(5).
//
// Possible errors: NOT_ENABLED
func removePort() {}

// removeProtocol(s: protocol) → Nothing
// Permanently remove protocol from list of protocols in service. See protocol tag in firewalld.service(5).
//
// Possible errors: NOT_ENABLED
func removeProtocol() {}

// removeSourcePort(s: port, s: protocol) → Nothing
// Permanently remove (port, protocol) from list of source ports in service. See source-port tag in firewalld.service(5).
//
// Possible errors: NOT_ENABLED
func removeSourcePort() {}

// rename(s: name) → Nothing
// Rename not built-in service to name.
//
// Possible errors: BUILTIN_SERVICE
func rename() {}

// setDescription(s: description) → Nothing
// Permanently set description of service to description. See description tag in firewalld.service(5).
func setDescription() {}

// setDestination(s: family, s: address) → Nothing
// Permanently set a destination address. destination is in format: (IP family, IP address) where IP family can be either 'ipv4' or 'ipv6'. See destination tag in firewalld.service(5).
//
// Possible errors: ALREADY_ENABLED
func setDestination() {}

// setDestinations(a{ss}: destinations) → Nothing
// Permanently set destinations of service to destinations, which is a dictionary of {IP family : IP address} where 'IP family' key can be either 'ipv4' or 'ipv6'. See destination tag in firewalld.service(5).
func setDestinations() {}

// setModules(as: modules) → Nothing
// Permanently set list of modules (netfilter kernel helpers) used in service to modules. See module tag in firewalld.service(5).
func setModules() {}

// setPorts(a(ss): ports) → Nothing
// Permanently set ports of service to list of (port, protocol). See port tag in firewalld.service(5).
func setPorts() {}

// setProtocols(as: protocols) → Nothing
// Permanently set protocols of service to list of protocols. See protocol tag in firewalld.service(5).
func setProtocols() {}

// setShort(s: short) → Nothing
// Permanently set name of service to short. See short tag in firewalld.service(5).
func setShort() {}

// setSourcePorts(a(ss): ports) → Nothing
// Permanently set source-ports of service to list of (port, protocol). See source-port tag in firewalld.service(5).
func setSourcePorts() {}

// setVersion(s: version) → Nothing
// Permanently set version of service to version. See version attribute of service tag in firewalld.service(5).
func setVersion() {}

// update((sssa(ss)asa{ss}asa(ss)): settings) → Nothing
// Update settings of service to settings. Settings are in format: version, name, description, array of ports (port, protocol), array of module names, dictionary of destinations, array of protocols and array of source-ports (port, protocol).
//
// version (s): see version attribute of service tag in firewalld.service(5).
// name (s): see short tag in firewalld.service(5).
// description (s): see description tag in firewalld.service(5).
// ports (a(ss)): array of port and protocol pairs. See port tag in firewalld.service(5).
// module names (as): array of kernel netfilter helpers, see module tag in firewalld.service(5).
// destinations (a{ss}): dictionary of {IP family : IP address} where 'IP family' key can be either 'ipv4' or 'ipv6'. See destination tag in firewalld.service(5).
// protocols (as): array of protocols. See protocol tag in firewalld.service(5).
// Possible errors: INVALID_TYPE
func update() {}

// Signals
// Removed(s: name)
// Emitted when service with name has been removed.
func removed() {}

// Renamed(s: name)
// Emitted when service has been renamed to name.
func renamed() {}

// Updated(s: name)
// Emitted when service with name has been updated.
func updated() {}

// Properties
// builtin - b - (ro)
// True if service is build-in, false else.
func isBuiltin() {}

// default - b - (ro)
// True if build-in service has default settings. False if it has been modified. Always False for not build-in services.
func isDefault() {}

// filename - s - (ro)
// Name (including .xml extension) of file where the configuration is stored.
func getFilename() {}

// name - s - (ro)
// Name of service.
func getName() {}

// path - s - (ro)
// Path to directory where the configuration is stored. Should be either /usr/lib/firewalld/services or /etc/firewalld/services.
func getPath() {}
