package ipset

// Methods
// addEntry(s: ipset, s: entry) → as
// Add a new entry to ipset. The entry must match the type of the ipset. If the ipset is using the timeout option, it is not possible to see the entries, as they are timing out automatically in the kernel. For permanent operation see org.fedoraproject.FirewallD1.config.ipset.Methods.addEntry.
//
// Possible errors: INVALID_IPSET, IPSET_WITH_TIMEOUT
func addEntry() {}

// getEntries(s: ipset) → Nothing
// Get all entries added to the ipset. If the ipset is using the timeout option, it is not possible to see the entries, as they are timing out automatically in the kernel. Return value is a array of entry. For permanent operation see org.fedoraproject.FirewallD1.config.ipset.Methods.getEntries.
//
// Possible errors: INVALID_IPSET, IPSET_WITH_TIMEOUT
func getEntries() {}

// getSettings(s: ipset) → (ssssa{ss}as)
// Return runtime settings of given ipset. For getting permanent settings see org.fedoraproject.FirewallD1.config.ipset.Methods.getSettings. Settings are in format: version, name, description, type, dictionary of options and array of entries.
//
// version (s): see version attribute of ipset tag in firewalld.ipset(5).
// name (s): see short tag in firewalld.ipset(5).
// description (s): see description tag in firewalld.ipset(5).
// type (s): see type attribute of ipset tag in firewalld.ipset(5).
// options (a{ss}): dictionary of {option : value} . See options tag in firewalld.ipset(5).
// entries (as): array of entries, see entry tag in firewalld.ipset(5).
// Possible errors: INVALID_IPSET
func getSettings() {}

// getIPSets() → as
// Return array of ipset names (s) in runtime configuration. For permanent configuration see org.fedoraproject.FirewallD1.config.Methods.listIPSets.
func getIPSets() {}

// queryService(s: ipset, s: entry) → b
// Return whether entry has been added to ipset. For permanent operation see org.fedoraproject.FirewallD1.config.ipset.Methods.queryEntry.
//
// Possible errors: INVALID_IPSET
func queryService() {

}

// queryService(s: ipset) → b
// Return whether ipset is defined in runtime configuration.
//func (query bool) queryService(ipset string) bool {
//	query = false
//}

// removeEntry(s: ipset, s: entry) → as
// Removes an entry from ipset. For permanent operation see org.fedoraproject.FirewallD1.config.ipset.Methods.removeEntry.
//
// Possible errors: INVALID_IPSET, IPSET_WITH_TIMEOUT
func removeEntry() {}

// setEntries(as: entries) → Nothing
// Permanently set list of entries to entries. For permanent operation see org.fedoraproject.FirewallD1.config.ipset.Methods.setEntries. See entry tag in firewalld.ipset(5).
func setEntries() {}

// Signals
// EntryAdded(s: ipset, s: entry)
// Emitted when entry has been added to ipset.
func EntryAdded() {}

// EntryRemoved(s: ipset, s: entry)
// Emitted when entry has been removed from ipset.
func EntryRemoved() {}
