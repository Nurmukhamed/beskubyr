package icmptypes

// Interface for permanent icmp type configuration, see also firewalld.icmptype(5).
// Methods
// addDestination(s: destination) → Nothing
// Permanently add a destination ('ipv4' or 'ipv6') to list of destinations of this icmp type. See destination tag in firewalld.icmptype(5).
//
// Possible errors: ALREADY_ENABLED
func addDestination() {}

// getDescription() → s
// Get description of icmp type. See description tag in firewalld.icmptype(5).
func getDescription() {}

// getDestinations() → as
// Get list of destinations. See destination tag in firewalld.icmptype(5).
func getDestinations() {}

// getSettings() → (sssas)
// Return permanent settings of icmp type. For getting runtime settings see org.fedoraproject.FirewallD1.Methods.getIcmpTypeSettings. Settings are in format: version, name, description, array of destinations.
//
// version (s): see version attribute of icmptype tag in firewalld.icmptype(5).
// name (s): see short tag in firewalld.icmptype(5).
// description (s): see description tag in firewalld.icmptype(5).
// destinations (as): array, either empty or containing strings 'ipv4' and/or 'ipv6', see destination tag in firewalld.icmptype(5).
func getSettings() {}

// getShort() → s
// Get name of icmp type. See short tag in firewalld.icmptype(5).
func getShort() {}

// getVersion() → s
// Get version of icmp type. See version attribute of icmptype tag in firewalld.icmptype(5).
func getVersion() {}

// loadDefaults() → Nothing
// Load default settings for built-in icmp type.
//
// Possible errors: NO_DEFAULTS
func loadDefaults() {}

// queryDestination(s: destination) → b
// Return whether a destination ('ipv4' or 'ipv6') is in list of destinations of this icmp type. See destination tag in firewalld.icmptype(5).
func queryDestination() {}

// remove() → Nothing
// Remove not built-in icmp type.
//
// Possible errors: BUILTIN_ICMPTYPE
func remove() {}

// removeDestination(s: destination) → Nothing
// Permanently remove a destination ('ipv4' or 'ipv6') from list of destinations of this icmp type. See destination tag in firewalld.icmptype(5).
//
// Possible errors: NOT_ENABLED
func removeDestination() {}

// rename(s: name) → Nothing
// Rename not built-in icmp type to name.
//
// Possible errors: BUILTIN_ICMPTYPE
func rename() {}

// setDescription(s: description) → Nothing
// Permanently set description of icmp type to description. See description tag in firewalld.icmptype(5).
func setDescription() {}

// setDestinations(as: destinations) → Nothing
// Permanently set destinations of icmp type to destinations, which is array, either empty or containing strings 'ipv4' and/or 'ipv6'. See destination tag in firewalld.icmptype(5).
func setDestinations() {}

// setShort(s: short) → Nothing
// Permanently set name of icmp type to short. See short tag in firewalld.icmptype(5).
func setShort() {}

// setVersion(s: version) → Nothing
// Permanently set version of icmp type to version. See version attribute of icmptype tag in firewalld.icmptype(5).
func setVersion() {}

// update((sssas): settings) → Nothing
// Update permanent settings of icmp type to settings. Settings are in format: version, name, description, array of destinations.
//
// version (s): see version attribute of icmptype tag in firewalld.icmptype(5).
// name (s): see short tag in firewalld.icmptype(5).
// description (s): see description tag in firewalld.icmptype(5).
// destinations (as): array, either empty or containing strings 'ipv4' and/or 'ipv6', see destination tag in firewalld.icmptype(5).
func update() {}

// Signals
// Removed(s: name)
// Emitted when icmp type with name has been removed.
func removed() {}

// Renamed(s: name)
// Emitted when icmp type has been renamed to name.
func renamed() {}

// Updated(s: name)
// Emitted when icmp type with name has been updated.
func updated() {}

// Properties
// builtin - b - (ro)
// True if icmptype is build-in, false else.
func isBuiltin() {}

// default - b - (ro)
// True if build-in icmp type has default settings. False if it has been modified. Always False for not build-in zones.
func isDefault() {}

// filename - s - (ro)
// Name (including .xml extension) of file where the configuration is stored.
func getFilename() {}

// name - s - (ro)
// Name of icmp type.
func getName() {}

// path - s - (ro)
// Path to directory where the icmp type configuration is stored. Should be either /usr/lib/firewalld/icmptypes or /etc/firewalld/icmptypes.
func getPath() {}
