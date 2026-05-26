package envutil

// DefaultEnvFile default file name
const DefaultEnvFile = ".env"

// Dotenv load and parse dotenv files
type Dotenv struct {
	// Files dot env file paths, allow multi files.
	//  - filename support simple glob pattern. eg: ".env.*"
	//
	// default: [".env"]
	Files []string
	// BaseDir base dir for join files path
	//
	// default is workdir
	BaseDir string
	// UpperKey change key to upper on set ENV. default: true
	UpperKey bool
	// IgnoreNotExist only load exists.
	//
	// - default: false - will return error if not exists
	IgnoreNotExist bool
	// LoadFirstExist only load first exists env file on Files
	LoadFirstExist bool

	loadFiles []string
	loadData  map[string]string
}

// NewDotenv create a new dotenv config
func NewDotenv() *Dotenv { _ = "STUB: not implemented"; return nil }

// init fields

// LoadAndInit load dotenv files and parse to os.Environ
func (c *Dotenv) LoadAndInit() error { _ = "STUB: not implemented"; return nil }

// LoadFiles append load dotenv files
//   - filename support simple glob pattern. eg: ".env.*"
func (c *Dotenv) LoadFiles(files ...string) error { _ = "STUB: not implemented"; return nil }

// LoadText load dotenv contents and parse to os Env
func (c *Dotenv) LoadText(contents string) error { _ = "STUB: not implemented"; return nil }

// do load dotenv files
func (c *Dotenv) doLoadFiles(files []string) error { _ = "STUB: not implemented"; return nil }

// load and parse to ENV

func (c *Dotenv) loadFile(filePath string) error {
	_ = "STUB: not implemented"
	// filename support simple glob pattern.
	return nil
}

// Load single file

// parseFile load single file and parse to ENV
func (c *Dotenv) parseFile(filePath string) error { _ = "STUB: not implemented"; return nil }

// IgnoreNotExist: skip non-existent files

func (c *Dotenv) parseAndSetEnv(contents string) error {
	_ = "STUB: not implemented"
	// Parse ENV lines
	return nil
}

// Set to ENV

// UnloadEnv remove loaded dotenv data from os.Environ
func (c *Dotenv) UnloadEnv() bool { _ = "STUB: not implemented"; return false }

// LoadedData get loaded dotenv data map
func (c *Dotenv) LoadedData() map[string]string {
	_ = "STUB: not implemented"

	// LoadedFiles get loaded dotenv files
	return nil
}

func (c *Dotenv) LoadedFiles() []string {
	_ = "STUB: not implemented"

	// Reset unload all loaded ENV and reset data
	return nil
}

func (c *Dotenv) Reset() { _ = "STUB: not implemented"; return }

//
// region standard dotenv instance
//

var stdEnv = NewDotenv()

// StdDotenv get standard dotenv instance
func StdDotenv() *Dotenv {
	_ = "STUB: not implemented"

	// DotenvLoad load dotenv file and parse to ENV
	return nil
}

func DotenvLoad(fns ...func(cfg *Dotenv)) error { _ = "STUB: not implemented"; return nil }

// LoadEnvFiles load dotenv files and parse to ENV
func LoadEnvFiles(baseDir string, files ...string) error { _ = "STUB: not implemented"; return nil }

// LoadedEnvFiles get loaded dotenv files
func LoadedEnvFiles() []string { _ = "STUB: not implemented"; return nil }
