package ipset

// Methods
// addEntry(s: entry) → Nothing
// Permanently add entry to list of entries of ipset. See entry tag in firewalld.ipset(5). For runtime operation see org.fedoraproject.FirewallD1.ipset.Methods.addEntry.
//
// Possible errors: ALREADY_ENABLED
func addEntry() {}

// addOption(s: key, s: value) → Nothing
// Permanently add (key, value) to the ipset. See option tag in firewalld.ipset(5).
//
// Possible errors: ALREADY_ENABLED
func addOption() {}

// getDescription() → s
// Get description of ipset. See description tag in firewalld.ipset(5).
func getDescription() {}

// getEntries() → as
// Get list of entries added to ipset. See entry tag in firewalld.ipset(5). For runtime operation see org.fedoraproject.FirewallD1.ipset.Methods.getEntries.
//
// Possible errors: IPSET_WITH_TIMEOUT
func getEntries() {}

// getOptions() → a{ss}
// Get dictionary of options set for ipset. See option tag in firewalld.ipset(5).
func getOptions() {}

// getSettings() → (ssssa{ss}as)
// Return permament settings of the ipset. For getting runtime settings see org.fedoraproject.FirewallD1.ipset.Methods.getIPSetSettings. Settings are in format: version, name, description, type, dictionary of options and array of entries.
//
// version (s): see version attribute of ipset tag in firewalld.ipset(5).
// name (s): see short tag in firewalld.ipset(5).
// description (s): see description tag in firewalld.ipset(5).
// type (s): see type attribute of ipset tag in firewalld.ipset(5).
// options (a{ss}): dictionary of {option : value} . See options tag in firewalld.ipset(5).
// entries (as): array of entries, see entry tag in firewalld.ipset(5).
func getSettings() {}

// getShort() → s
// Get name of ipset. See short tag in firewalld.ipset(5).
func getShort() {}

// getType() → s
// Get type of ipset. See type attribute of ipset tag in firewalld.ipset(5).
func getType() {}

// getVersion() → s
// Get version of ipset. See version attribute of ipset tag in firewalld.ipset(5).
func getVersion() {}

// loadDefaults() → Nothing
// Load default settings for built-in ipset.
//
// Possible errors: NO_DEFAULTS
func loadDefaults() {}

// queryEntry(s: entry) → b
// Return whether entry has been added to ipset. For runtime operation see org.fedoraproject.FirewallD1.ipset.Methods.queryEntry.
func queryEntry() {}

// queryOption(s: key, s: value) → b
// Return whether (key, value) has been added to options of the ipset.
func queryOption() {}

// remove() → Nothing
// Remove not built-in ipset.
//
// Possible errors: BUILTIN_IPSET
func remove() {}

// removeEntry(s: entry) → Nothing
// Permanently remove entry from ipset. See entry tag in firewalld.ipset(5). For runtime operation see org.fedoraproject.FirewallD1.ipset.Methods.removeEntry.
//
// Possible errors: NOT_ENABLED
func removeEntry() {}

// removeOption(s: key) → Nothing
// Permanently remove key from the ipset. See option tag in firewalld.ipset(5).
//
// Possible errors: NOT_ENABLED
func removeOption() {}

// rename(s: name) → Nothing
// Rename not built-in ipset to name.
//
// Possible errors: BUILTIN_IPSET
func rename() {}

// setDescription(s: description) → Nothing
// Permanently set description of ipset to description. See description tag in firewalld.ipset(5).
func setDescription() {}

// setEntries(as: entries) → Nothing
// Permanently set list of entries to entries. See entry tag in firewalld.ipset(5).
func setEntries() {}

// setOptions(a{ss}: options) → Nothing
// Permanently set dict of options to options. See option tag in firewalld.ipset(5).
func setOptions() {}

// setShort(s: short) → Nothing
// Permanently set name of ipset to short. See short tag in firewalld.ipset(5).
func setShort() {}

// setType(s: ipset_type) → Nothing
// Permanently set type of ipset to ipset_type. See type attribute of ipset tag in firewalld.ipset(5).
func setType() {}

// setVersion(s: version) → Nothing
// Permanently set version of ipset to version. See version attribute of ipset tag in firewalld.ipset(5).
func setVersion() {}

// update((ssssa{ss}as): settings) → Nothing
// Update settings of ipset to settings. Settings are in format: version, name, description, type, dictionary of options and array of entries.
//
// version (s): see version attribute of ipset tag in firewalld.ipset(5).
// name (s): see short tag in firewalld.ipset(5).
// description (s): see description tag in firewalld.ipset(5).
// type (s): see type attribute of ipset tag in firewalld.ipset(5).
// options (a{ss}): dictionary of {option : value} . See options tag in firewalld.ipset(5).
// entries (as): array of entries, see entry tag in firewalld.ipset(5).
// Possible errors: INVALID_TYPE
func update() {}

// Signals
// Removed(s: name)
// Emitted when ipset with name has been removed.
func removed() {}

// Renamed(s: name)
// Emitted when ipset has been renamed to name.
func renamed() {}

// Updated(s: name)
// mitted when ipset with name has been updated.
func updated() {}

// Properties
// builtin - b - (ro)
// True if ipset is build-in, false else.
func isBuiltin() {}

// default - b - (ro)
// True if build-in ipset has default settings. False if it has been modified. Always False for not build-in ipsets.
func isDefault() {}

// filename - s - (ro)
// Name (including .xml extension) of file where the configuration is stored.
func getFilename() {}

// name - s - (ro)
// Name of ipset.
func getName() {}

// path - s - (ro)
// Path to directory where the ipset configuration is stored. Should be either /usr/lib/firewalld/ipsets or /etc/firewalld/ipsets.
func getPath() {}
