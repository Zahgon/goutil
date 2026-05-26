package maputil

// Aliases implemented a simple string alias map.
//   - key: alias, value: real name
type Aliases map[string]string

// AddAlias to the Aliases map
func (as Aliases) AddAlias(alias, real string) { _ = "STUB: not implemented"; return }

// AddAliases to the Aliases map
func (as Aliases) AddAliases(real string, aliases []string) { _ = "STUB: not implemented"; return }

// AddAliasMap to the Aliases map
func (as Aliases) AddAliasMap(alias2real map[string]string) { _ = "STUB: not implemented"; return }

// HasAlias in the Aliases map
func (as Aliases) HasAlias(alias string) bool { _ = "STUB: not implemented"; return false }

// ResolveAlias by given name.
func (as Aliases) ResolveAlias(alias string) string { _ = "STUB: not implemented"; return "" }

// AliasesNames returns all sorted alias names.
func (as Aliases) AliasesNames() []string { _ = "STUB: not implemented"; return nil }

// GroupAliases groups aliases by real name.
//
// returns: {real name -> []aliases, ...}
func (as Aliases) GroupAliases() map[string][]string { _ = "STUB: not implemented"; return nil }
