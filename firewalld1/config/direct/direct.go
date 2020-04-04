package direct

// Methods
// addChain(s: ipv, s: table, s: chain) → Nothing
// Add a new chain to table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). Make sure there's no other chain with this name already. There already exist basic chains to use with direct methods, for example INPUT_direct chain. These chains are jumped into before chains for zones, i.e. every rule put into INPUT_direct will be checked before rules in zones. For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.addChain.
//
// Possible errors: INVALID_IPV, INVALID_TABLE, ALREADY_ENABLED
func addChain() {}

// addPassthrough(s: ipv, as: args) → Nothing
// Add a passthrough rule with the arguments args for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.addPassthrough.
//
// Possible errors: INVALID_IPV, ALREADY_ENABLED
func addPassthrough() {}

// addRule(s: ipv, s: table, s: chain, i: priority, as: args) → Nothing
// Add a rule with the arguments args to chain in table with priority for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). The priority is used to order rules. Priority 0 means add rule on top of the chain, with a higher priority the rule will be added further down. Rules with the same priority are on the same level and the order of these rules is not fixed and may change. If you want to make sure that a rule will be added after another one, use a low priority for the first and a higher for the following. For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.addRule.
//
// Possible errors: INVALID_IPV, INVALID_TABLE, ALREADY_ENABLED
func addRule() {}

// getAllChains() → a(sss)
// Get all chains added to all tables in format: ipv, table, chain. This concerns only chains previously added with addChain. Return value is a array of (ipv, table, chain). For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.getAllChains.
//
// ipv (s): either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables).
// table (s): one of filter, mangle, nat, raw, security
// chain (s): name of a chain.
func getAllChains() {}

// getAllPassthroughs() → a(sas)
// Get all passthrough rules added in all ipv types in format: ipv, rule. This concerns only rules previously added with addPassthrough. Return value is a array of (ipv, array of arguments). For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.getAllPassthroughs.
//
// ipv (s): either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables).
// arguments (as): array of commands, parameters and other iptables/ip6tables/ebtables command line options.
func getAllPassthroughs() {}

// getAllRules() → a(sssias)
// Get all rules added to all chains in all tables in format: ipv, table, chain, priority, rule. This concerns only rules previously added with addRule. Return value is a array of (ipv, table, chain, priority, array of arguments). For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.getAllRules.
//
// ipv (s): either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables).
// table (s): one of filter, mangle, nat, raw, security
// chain (s): name of a chain.
// priority (i): used to order rules.
// arguments (as): array of commands, parameters and other iptables/ip6tables/ebtables command line options.
func getAllRules() {}

// getChains(s: ipv, s: table) → as
// Return an array of chains (s) added to table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). This concerns only chains previously added with addChain. For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.getChains.
//
// Possible errors: INVALID_IPV, INVALID_TABLE
func getChains() {}

// getPassthroughs(s: ipv) → aas
// Get tracked passthrough rules added in either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). This concerns only rules previously added with addPassthrough. Return value is a array of (array of arguments). For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.getPassthroughs.
//
// arguments (as): array of commands, parameters and other iptables/ip6tables/ebtables command line options.
func getPassthroughs() {}

// getRules(s: ipv, s: table, s: chain) → a(ias)
// Get all rules added to chain in table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). This concerns only rules previously added with addRule. Return value is a array of (priority, array of arguments). For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.getRules.
//
// priority (i): used to order rules.
// arguments (as): array of commands, parameters and other iptables/ip6tables/ebtables command line options.
// Possible errors: INVALID_IPV, INVALID_TABLE
func getRules() {}

// getSettings() → (a(sss)a(sssias)a(sas))
// Get settings of permanent direct configuration in format: array of chains, array of rules, array of passthroughs.
//
// chains (a(sss)): array of (ipv, table, chain), see 'chain' in firewalld.direct(5).
// rules (a(sssias)): array of (ipv, table, chain, priority, array of arguments), see 'rule' in firewalld.direct(5).
// passthroughs (a(sas)): array of (ipv, array of arguments), see passthrough in firewalld.direct(5).
func getSettings() {}

// queryChain(s: ipv, s: table, s: chain) → b
// Return whether a chain exists in table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). This concerns only chains previously added with addChain. For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.queryChain.
//
// Possible errors: INVALID_IPV, INVALID_TABLE
func queryChain() {}

// queryPassthrough(s: ipv, as: args) → b
// Return whether a tracked passthrough rule with the arguments args exists for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). This concerns only rules previously added with addPassthrough. For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.queryPassthrough.
//
// Possible errors: INVALID_IPV
func queryPassthrough() {}

// queryRule(s: ipv, s: table, s: chain, i: priority, as: args) → b
// Return whether a rule with priority and the arguments args exists in chain in table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). This concerns only rules previously added with addRule. For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.queryRule.
//
// Possible errors: INVALID_IPV, INVALID_TABLE
func queryRule() {}

// removeChain(s: ipv, s: table, s: chain) → Nothing
// Remove a chain from table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). Only chains previously added with addChain can be removed this way. For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.removeChain.
//
// Possible errors: INVALID_IPV, INVALID_TABLE, NOT_ENABLED
func removeChain() {}

// removePassthrough(s: ipv, as: args) → Nothing
// Remove a passthrough rule with arguments args for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). Only rules previously added with addPassthrough can be removed this way. For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.removePassthrough.
//
// Possible errors: INVALID_IPV, NOT_ENABLED
func removePassthrough() {}

// removeRule(s: ipv, s: table, s: chain, i: priority, as: args) → Nothing
// Remove a rule with priority and arguments args from chain in table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). Only rules previously added with addRule can be removed this way. For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.removeRule.
//
// Possible errors: INVALID_IPV, INVALID_TABLE, NOT_ENABLED
func removeRule() {}

// removeRules(s: ipv, s: table, s: chain) → Nothing
// Remove all rules from chain in table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). This concerns only rules previously added with addRule. For runtime operation see org.fedoraproject.FirewallD1.direct.Methods.removeRules.
//
// Possible errors: INVALID_IPV, INVALID_TABLE
func removeRules() {}

// update((a(sss)a(sssias)a(sas)): settings) → Nothing
// Update permanent direct configuration with given settings. Settings are in format: array of chains, array of rules, array of passthroughs.
//
// chains (a(sss)): array of (ipv, table, chain), see 'chain' in firewalld.direct(5).
// rules (a(sssias)): array of (ipv, table, chain, priority, array of arguments), see 'rule' in firewalld.direct(5).
// passthroughs (a(sas)): array of (ipv, array of arguments), see passthrough in firewalld.direct(5).
// Possible errors: INVALID_TYPE
func update() {}

// Signals
// Updated()
// Emitted when configuration has been updated.
func updated() {}
