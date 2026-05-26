package structs

// Aliases implemented a simple string alias map.
type Aliases struct {
	mapping map[string]string
	// Checker custom add alias name checker func
	Checker func(alias string) // should return bool OR error ??
}

// NewAliases create
func NewAliases(checker func(alias string)) *Aliases { _ = "STUB: not implemented"; return nil }

// AddAlias to the Aliases
func (as *Aliases) AddAlias(real, alias string) { _ = "STUB: not implemented"; return }

// AddAliases to the Aliases
func (as *Aliases) AddAliases(real string, aliases []string) { _ = "STUB: not implemented"; return }

// AddAliasMap to the Aliases
func (as *Aliases) AddAliasMap(alias2real map[string]string) { _ = "STUB: not implemented"; return }

// HasAlias in the Aliases
func (as *Aliases) HasAlias(alias string) bool { _ = "STUB: not implemented"; return false }

// ResolveAlias by given name. if not exists, return the alias self
func (as *Aliases) ResolveAlias(alias string) string { _ = "STUB: not implemented"; return "" }

// Mapping get all aliases mapping
func (as *Aliases) Mapping() map[string]string { _ = "STUB: not implemented"; return nil }
