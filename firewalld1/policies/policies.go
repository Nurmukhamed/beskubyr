package policies

// Methods
// addLockdownWhitelistCommand(s: command) → Nothing
// Add command to whitelist. See command option in firewalld.lockdown-whitelist(5). For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.addLockdownWhitelistCommand.
//
// Possible errors: ALREADY_ENABLED, INVALID_COMMAND
func addLockdownWhitelistCommand() {}

// addLockdownWhitelistContext(s: context) → Nothing
// Add context to whitelist. See selinux option in firewalld.lockdown-whitelist(5). For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.addLockdownWhitelistContext.
//
// Possible errors: ALREADY_ENABLED, INVALID_COMMAND
func addLockdownWhitelistContext() {}

// addLockdownWhitelistUid(i: uid) → Nothing
// Add user id uid to whitelist. See user option in firewalld.lockdown-whitelist(5). For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.addLockdownWhitelistUid.
//
// Possible errors: ALREADY_ENABLED, INVALID_COMMAND
func addLockdownWhitelistUID() {}

// addLockdownWhitelistUser(s: user) → Nothing
// Add user name to whitelist. See user option in firewalld.lockdown-whitelist(5). For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.addLockdownWhitelistUser.
//
// Possible errors: ALREADY_ENABLED, INVALID_COMMAND
func addLockdownWhitelistUser() {}

// disableLockdown() → Nothing
// Disable lockdown. This is a runtime and permanent change.
//
// Possible errors: NOT_ENABLED
func disableLockdown() {}

// enableLockdown() → Nothing
// Enable lockdown. Be careful - if the calling application/user is not on lockdown whitelist when you enable lockdown you won't be able to disable it again with the application, you would need to edit firewalld.conf. This is a runtime and permanent change.
//
// Possible errors: ALREADY_ENABLED
func enableLockdown() {}

// getLockdownWhitelistCommands() → as
// List all command lines (s) that are on whitelist. For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.getLockdownWhitelistCommands.
func getLockdownWhitelistCommands() {}

// getLockdownWhitelistContexts() → as
// List all contexts (s) that are on whitelist. For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.getLockdownWhitelistContexts.
func getLockdownWhitelistContexts() {}

// getLockdownWhitelistUids() → ai
// List all user ids (i) that are on whitelist. For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.getLockdownWhitelistUids.
func getLockdownWhitelistUids() {}

// getLockdownWhitelistUsers() → as
// List all users (s) that are on whitelist. For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.getLockdownWhitelistUsers.
func getLockdownWhitelistUsers() {}

// queryLockdown() → b
// Query whether lockdown is enabled.
func queryLockdown() {}

// queryLockdownWhitelistCommand(s: command) → b
// Query whether command is on whitelist. For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.queryLockdownWhitelistCommand.
func queryLockdownWhitelistCommand() {}

// queryLockdownWhitelistContext(s: context) → b
// Query whether context is on whitelist. For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.queryLockdownWhitelistContext.
func queryLockdownWhitelistContext() {}

// queryLockdownWhitelistUid(i: uid) → b
// Query whether user id uid is on whitelist. For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.queryLockdownWhitelistUid.
func queryLockdownWhitelistUID() {}

// queryLockdownWhitelistUser(s: user) → b
// Query whether user is on whitelist. For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.queryLockdownWhitelistUser.
func queryLockdownWhitelistUser() {}

// removeLockdownWhitelistCommand(s: command) → Nothing
// Remove command from whitelist. For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.removeLockdownWhitelistCommand.
//
// Possible errors: NOT_ENABLED
func removeLockdownWhitelistCommand() {}

// removeLockdownWhitelistContext(s: context) → Nothing
// Remove context from whitelist. For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.removeLockdownWhitelistContext.
//
// Possible errors: NOT_ENABLED
func removeLockdownWhitelistContext() {}

// removeLockdownWhitelistUid(i: uid) → Nothing
// Remove user id uid from whitelist. For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.removeLockdownWhitelistUid.
//
// Possible errors: NOT_ENABLED
func removeLockdownWhitelistUID() {}

// removeLockdownWhitelistUser(s: user) → Nothing
// Remove user from whitelist. For permanent operation see org.fedoraproject.FirewallD1.config.policies.Methods.removeLockdownWhitelistUser.
//
// Possible errors: NOT_ENABLED
func removeLockdownWhitelistUser() {}

// Signals
// LockdownDisabled()
// Emitted when lockdown has been disabled.
func LockdownDisabled() {}

// LockdownEnabled()
// Emitted when lockdown has been enabled.
func LockdownEnabled() {}

// LockdownWhitelistCommandAdded(s: command)
// Emitted when command has been added to whitelist.
func LockdownWhitelistCommandAdded() {}

// LockdownWhitelistCommandRemoved(s: command)
// Emitted when command has been removed from whitelist.
func LockdownWhitelistCommandRemoved() {}

// LockdownWhitelistContextAdded(s: context)
// Emitted when context has been added to whitelist.
func LockdownWhitelistContextAdded() {}

// LockdownWhitelistContextRemoved(s: context)
// Emitted when context has been removed from whitelist.
func LockdownWhitelistContextRemoved() {}

// LockdownWhitelistUidAdded(i: uid)
// Emitted when user id uid has been added to whitelist.
func LockdownWhitelistUidAdded() {}

// LockdownWhitelistUidRemoved(i: uid)
// Emitted when user id uid has been removed from whitelist.
func LockdownWhitelistUidRemoved() {}

// LockdownWhitelistUserAdded(s: user)
// Emitted when user has been added to whitelist.
func LockdownWhitelistUserAdded() {}

// LockdownWhitelistUserRemoved(s: user)
// Emitted when user has been removed from whitelist.
func LockdownWhitelistUserRemoved() {}
