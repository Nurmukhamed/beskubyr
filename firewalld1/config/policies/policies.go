package policies

// Methods
// addLockdownWhitelistCommand(s: command) → Nothing
// Add command to whitelist. See command option in firewalld.lockdown-whitelist(5). For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.addLockdownWhitelistCommand.
//
// Possible errors: ALREADY_ENABLED, INVALID_TYPE
func addLockdownWhitelistCommand() {}

// addLockdownWhitelistContext(s: context) → Nothing
// Add context to whitelist. See selinux option in firewalld.lockdown-whitelist(5). For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.addLockdownWhitelistContext.
//
// Possible errors: ALREADY_ENABLED, INVALID_TYPE
func addLockdownWhitelistContext() {}

// addLockdownWhitelistUid(i: uid) → Nothing
// Add user id uid to whitelist. See user option in firewalld.lockdown-whitelist(5). For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.addLockdownWhitelistUid.
//
// Possible errors: ALREADY_ENABLED, INVALID_TYPE
func addLockdownWhitelistUID() {}

// addLockdownWhitelistUser(s: user) → Nothing
// Add user name to whitelist. See user option in firewalld.lockdown-whitelist(5). For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.addLockdownWhitelistUser.
//
// Possible errors: ALREADY_ENABLED, INVALID_TYPE
func addLockdownWhitelistUser() {}

// getLockdownWhitelist() → (asasasai)
// Get settings of permanent lockdown-whitelist configuration in format: commands, selinux contexts, users, uids
//
// commands (as): see command option in firewalld.lockdown-whitelist(5).
// selinux contexts (as): see selinux option in firewalld.lockdown-whitelist(5).
// users (as): see name attribute of user option in firewalld.lockdown-whitelist(5).
// uids (ai): see id attribute of user option in firewalld.lockdown-whitelist(5).
func getLockdownWhitelist() {}

// getLockdownWhitelistCommands() → as
// List all command lines (s) that are on whitelist. For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.getLockdownWhitelistCommands.
func getLockdownWhitelistCommands() {}

// getLockdownWhitelistContexts() → as
// List all contexts (s) that are on whitelist. For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.getLockdownWhitelistContexts.
func getLockdownWhitelistContexts() {}

// getLockdownWhitelistUids() → ai
// List all user ids (i) that are on whitelist. For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.getLockdownWhitelistUids.
func getLockdownWhitelistUids() {}

// getLockdownWhitelistUsers() → as
// List all users (s) that are on whitelist. For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.getLockdownWhitelistUsers.
func getLockdownWhitelistUsers() {}

// queryLockdownWhitelistCommand(s: command) → b
// Query whether command is on whitelist. For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.queryLockdownWhitelistCommand.
func queryLockdownWhitelistCommand() {}

// queryLockdownWhitelistContext(s: context) → b
// Query whether context is on whitelist. For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.queryLockdownWhitelistContext.
func queryLockdownWhitelistContext() {}

// queryLockdownWhitelistUid(i: uid) → b
// Query whether user id uid is on whitelist. For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.queryLockdownWhitelistUid.
func queryLockdownWhitelistUID() {}

// queryLockdownWhitelistUser(s: user) → b
// Query whether user is on whitelist. For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.queryLockdownWhitelistUser.
func queryLockdownWhitelistUser() {}

// removeLockdownWhitelistCommand(s: command) → Nothing
// Remove command from whitelist. For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.removeLockdownWhitelistCommand.
//
// Possible errors: NOT_ENABLED
func removeLockdownWhitelistCommand() {}

// removeLockdownWhitelistContext(s: context) → Nothing
// Remove context from whitelist. For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.removeLockdownWhitelistContext.
//
// Possible errors: NOT_ENABLED
func removeLockdownWhitelistContext() {}

// removeLockdownWhitelistUid(i: uid) → Nothing
// Remove user id uid from whitelist. For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.removeLockdownWhitelistUid.
//
// Possible errors: NOT_ENABLED
func removeLockdownWhitelistUID() {}

// removeLockdownWhitelistUser(s: user) → Nothing
// Remove user from whitelist. For runtime operation see org.fedoraproject.FirewallD1.policies.Methods.removeLockdownWhitelistUser.
//
// Possible errors: NOT_ENABLED
func removeLockdownWhitelistUser() {}

// setLockdownWhitelist((asasasai): settings) → Nothing
// Set permanent lockdown-whitelist configuration to settings. Settings are in format: commands, selinux contexts, users, uids
//
// commands (as): see command option in firewalld.lockdown-whitelist(5).
// selinux contexts (as): see selinux option in firewalld.lockdown-whitelist(5).
// users (as): see name attribute of user option in firewalld.lockdown-whitelist(5).
// uids (ai): see id attribute of user option in firewalld.lockdown-whitelist(5).
// Possible errors: INVALID_TYPE
func setLockdownWhitelist() {}

// Signals
// LockdownWhitelistUpdated()
// Emitted when permanent lockdown-whitelist configuration has been updated.
func LockdownWhitelistUpdated() {}
