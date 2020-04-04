package direct

// Methods
// addChain(s: ipv, s: table, s: chain) → Nothing
// dd a new chain to table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). Make sure there's no other chain with this name already. There already exist basic chains to use with direct methods, for example INPUT_direct chain. These chains are jumped into before chains for zones, i.e. every rule put into INPUT_direct will be checked before rules in zones. For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.addChain.
//
// Possible errors: INVALID_IPV, INVALID_TABLE, ALREADY_ENABLED, COMMAND_FAILED
func addChain() {}

// addPassthrough(s: ipv, as: args) → Nothing
// Add a tracked passthrough rule with the arguments args for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). Valid commands in args are only -A/--append, -I/--insert and -N/--new-chain. This method is (unlike passthrough method) tracked, i.e. firewalld remembers it. It's useful with org.fedoraproject.FirewallD1.Methods.runtimeToPermanent For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.addPassthrough.
//
// Possible errors: INVALID_IPV, ALREADY_ENABLED, COMMAND_FAILED
func addPassthrough() {}

// addRule(s: ipv, s: table, s: chain, i: priority, as: args) → Nothing
// Add a rule with the arguments args to chain in table with priority for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). The priority is used to order rules. Priority 0 means add rule on top of the chain, with a higher priority the rule will be added further down. Rules with the same priority are on the same level and the order of these rules is not fixed and may change. If you want to make sure that a rule will be added after another one, use a low priority for the first and a higher for the following. For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.addRule.
//
// Possible errors: INVALID_IPV, INVALID_TABLE, ALREADY_ENABLED, COMMAND_FAILED
func addRule() {}

// getAllChains() → a(sss)
// Get all chains added to all tables in format: ipv, table, chain. This concerns only chains previously added with addChain. Return value is a array of (ipv, table, chain). For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.getAllChains.
//
// ipv (s): either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables).
// table (s): one of filter, mangle, nat, raw, security
// chain (s): name of a chain.
func getAllChains() {}

// getAllPassthroughs() → a(sas)
// Get all tracked passthrough rules added in all ipv types in format: ipv, rule. This concerns only rules previously added with addPassthrough. Return value is a array of (ipv, array of arguments). For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.getAllPassthroughs.
//
// ipv (s): either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables).
// arguments (as): array of commands, parameters and other iptables/ip6tables/ebtables command line options.
func getAllPassthroughs() {}

//getAllRules() → a(sssias)
// Get all rules added to all chains in all tables in format: ipv, table, chain, priority, rule. This concerns only rules previously added with addRule. Return value is a array of (ipv, table, chain, priority, array of arguments). For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.getAllRules.
//
// ipv (s): either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables).
// table (s): one of filter, mangle, nat, raw, security
// chain (s): name of a chain.
// priority (i): used to order rules.
// arguments (as): array of commands, parameters and other iptables/ip6tables/ebtables command line options.
func getAllRules() {}

// getChains(s: ipv, s: table) → as
// Return an array of chains (s) added to table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). This concerns only chains previously added with addChain. For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.getChains.
//
// Possible errors: INVALID_IPV, INVALID_TABLE
func getChains() {}

// getPassthroughs(s: ipv) → aas
// Get tracked passthrough rules added in either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). This concerns only rules previously added with addPassthrough. Return value is a array of (array of arguments). For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.getPassthroughs.
//
// arguments (as): array of commands, parameters and other iptables/ip6tables/ebtables command line options.
// getRules(s: ipv, s: table, s: chain) → a(ias)
// Get all rules added to chain in table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). This concerns only rules previously added with addRule. Return value is a array of (priority, array of arguments). For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.getRules.
//
// priority (i): used to order rules.
// arguments (as): array of commands, parameters and other iptables/ip6tables/ebtables command line options.
// Possible errors: INVALID_IPV, INVALID_TABLE
func getPassthroughs() {}

// passthrough(s: ipv, as: args) → s
// Pass a command through to the firewall. ipv can be either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). args can be all iptables, ip6tables and ebtables command line arguments. args can be all iptables, ip6tables and ebtables command line arguments. This command is untracked, which means that firewalld is not able to provide information about this command later on.
//
// Possible errors: COMMAND_FAILED
func passthrough() {}

// queryChain(s: ipv, s: table, s: chain) → b
// Return whether a chain exists in table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). This concerns only chains previously added with addChain. For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.queryChain.
//
// Possible errors: INVALID_IPV, INVALID_TABLE
func queryChain() {}

// queryPassthrough(s: ipv, as: args) → b
// Return whether a tracked passthrough rule with the arguments args exists for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). This concerns only rules previously added with addPassthrough. For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.queryPassthrough.
//
// Possible errors: INVALID_IPV
func queryPassthrough() {}

// queryRule(s: ipv, s: table, s: chain, i: priority, as: args) → b
// Return whether a rule with priority and the arguments args exists in chain in table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). This concerns only rules previously added with addRule. For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.queryRule.
//
// Possible errors: INVALID_IPV, INVALID_TABLE
func queryRules() {}

// removeAllPassthroughs() → Nothing
// Remove all passthrough rules previously added with addPassthrough.
func removeAllPassthroughs() {}

// removeChain(s: ipv, s: table, s: chain) → Nothing
// Remove a chain from table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). Only chains previously added with addChain can be removed this way. For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.removeChain.
//
// Possible errors: INVALID_IPV, INVALID_TABLE, NOT_ENABLED, COMMAND_FAILED
func removeChain() {}

// removePassthrough(s: ipv, as: args) → Nothing
// Remove a tracked passthrough rule with arguments args for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). Only rules previously added with addPassthrough can be removed this way. For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.removePassthrough.
//
// Possible errors: INVALID_IPV, NOT_ENABLED, COMMAND_FAILED
func removePassthrough() {}

// removeRule(s: ipv, s: table, s: chain, i: priority, as: args) → Nothing
// Remove a rule with priority and arguments args from chain in table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). Only rules previously added with addRule can be removed this way. For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.removeRule.
//
// Possible errors: INVALID_IPV, INVALID_TABLE, NOT_ENABLED, COMMAND_FAILED
func removeRule() {}

// removeRules(s: ipv, s: table, s: chain) → Nothing
// Remove all rules from chain in table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables). This concerns only rules previously added with addRule. For permanent operation see org.fedoraproject.FirewallD1.config.direct.Methods.removeRules.
//
// Possible errors: INVALID_IPV, INVALID_TABLE
func removeRules() {}

// Signals
// ChainAdded(s: ipv, s: table, s: chain)
// Emitted when chain has been added into table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables).
func ChainAdded() {}

// ChainRemoved(s: ipv, s: table, s: chain)
// Emitted when chain has been removed from table for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables).
func ChainRemoved() {}

// PassthroughAdded(s: ipv, as: args)
// Emitted when a tracked passthruogh rule with args has been added for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables).
func PassthroughAdded() {}

// PassthroughRemoved(s: ipv, as: args)
// Emitted when a tracked passthrough rule with args has been removed for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables).
func PassthroughRemoved() {}

// RuleAdded(s: ipv, s: table, s: chain, i: priority, as: args)
// Emitted when a rule with args has been added to chain in table with priority for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables).
func RuleAdded() {}

// RuleRemoved(s: ipv, s: table, s: chain, i: priority, as: args)
// Emitted when a rule with args has been removed from chain in table with priority for ipv being either ipv4 (iptables) or ipv6 (ip6tables) or eb (ebtables).
func RuleRemoved() {}
