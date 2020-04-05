package confighelper

// Interface for permanent helper configuration, see also firewalld.helper(5).
// Methods
// addPort(s: port, s: protocol) → Nothing
// Permanently add (port, protocol) to list of ports in helper. See port tag in firewalld.helper(5).
//
// Possible errors: ALREADY_ENABLED
func addPort() {}

// getDescription() → s
// Get description of helper. See description tag in firewalld.helper(5).
func getDescription() {}

// getFamily() → s
// Get family being 'ipv4', 'ipv6' or empty for both. See family tag in firewalld.helper(5).
func getFamily() {}

// getModule() → s
// Get modules (netfilter kernel helpers) used in helper. See module tag in firewalld.helper(5).
func getModule() {}

// getPorts() → a(ss)
// Get list of (port, protocol) defined in helper. See port tag in firewalld.helper(5).
func getPorts() {}

// getSettings() → (sssssa(ss))
// Return permanent settings of a helper. For getting runtime settings see org.fedoraproject.FirewallD1.Methods.getHelperSettings. Settings are in format: version, name, description, family, module, array of ports (port, protocol).
//
// version (s): see version attribute of helper tag in firewalld.helper(5).
// name (s): see short tag in firewalld.helper(5).
// description (s): see description tag in firewalld.helper(5).
// family (s): see family tag in firewalld.helper(5).
// module (s): see module tag in firewalld.helper(5).
// ports (a(ss)): array of port and protocol pairs. See port tag in firewalld.helper(5).
func getSettings() {}

// getShort() → s
// Get name of helper. See short tag in firewalld.helper(5).
func getShort() {}

// getVersion() → s
// Get version of helper. See version attribute of helper tag in firewalld.helper(5).
func getVersion() {}

// loadDefaults() → Nothing
// Load default settings for built-in helper.
//
// Possible errors: NO_DEFAULTS
func loadDefaults() {}

// queryFamily(s: module) → b
// Return whether family is set for helper. See family tag in firewalld.helper(5).
func queryFamily() {}

// queryModule(s: module) → b
// Return whether module (netfilter kernel helpers) is used in helper. See module tag in firewalld.helper(5).
func queryModule() {}

// queryPort(s: port, s: protocol) → b
// Return whether (port, protocol) is in list of ports in helper. See port tag in firewalld.helper(5).
func queryPort() {}

// remove() → Nothing
// Remove not built-in helper.
//
// Possible errors: BUILTIN_HELPER
func remove() {}

// removePort(s: port, s: protocol) → Nothing
// Permanently remove (port, protocol) from list of ports in helper. See port tag in firewalld.helper(5).
//
// Possible errors: NOT_ENABLED
func removePort() {}

// rename(s: name) → Nothing
// Rename not built-in helper to name.
//
// Possible errors: BUILTIN_HELPER
func rename() {}

// setDescription(s: description) → Nothing
// Permanently set description of helper to description. See description tag in firewalld.helper(5).
func setDescription() {}

// setFamily(s: family) → Nothing
// Permanently set family of helper to family. See family tag in firewalld.helper(5).
func setFamily() {}

// setModule(s: module) → Nothing
// Permanently set module of helper to description. See module tag in firewalld.helper(5).
func setModule() {}

// setPorts(a(ss): ports) → Nothing
// Permanently set ports of helper to list of (port, protocol). See port tag in firewalld.helper(5).
func setPorts() {}

// setShort(s: short) → Nothing
// Permanently set name of helper to short. See short tag in firewalld.helper(5).
func setShort() {}

// setVersion(s: version) → Nothing
// Permanently set version of helper to version. See version attribute of helper tag in firewalld.helper(5).
func setVersion() {}

// update((sssssa(ss)): settings) → Nothing
// Update settings of helper to settings. Settings are in format: version, name, description, family, module and array of ports.
//
// version (s): see version attribute of helper tag in firewalld.helper(5).
// name (s): see short tag in firewalld.helper(5).
// description (s): see description tag in firewalld.helper(5).
// family (s): see family tag in firewalld.helper(5).
// module (s): see module tag in firewalld.helper(5).
// ports (a(ss)): array of port and protocol pairs. See port tag in firewalld.helper(5).
// Possible errors: INVALID_HELPER
func update() {}

// Signals
// Removed(s: name)
// Emitted when helper with name has been removed.
func removed() {}

// Renamed(s: name)
// Emitted when helper has been renamed to name.
func renamed() {}

// Updated(s: name)
// Emitted when helper with name has been updated.
func updated() {}

// Properties
// builtin - b - (ro)
// True if helper is build-in, false else.
func isBuiltin() {}

// default - b - (ro)
// True if build-in helper has default settings. False if it has been modified. Always False for not build-in helpers.
func isDefault() {}

// filename - s - (ro)
// Name (including .xml extension) of file where the configuration is stored.
func getFilename() {}

// name - s - (ro)
// Name of helper.
func getName() {}

// path - s - (ro)
// Path to directory where the configuration is stored. Should be either /usr/lib/firewalld/helpers or /etc/firewalld/helpers
func getPath() {}
