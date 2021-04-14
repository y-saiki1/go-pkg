package migration

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//sql and database info
const (
	Source = "file://./pkg/database/migration"
)

//declare command line options
var (
	Command = flag.String("exec", "", "set up or down as a argument")
	Force   = flag.Bool("f", false, "force exec fixed sql")
)

//available command list
var AvailableExecCommands = map[string]string{
	"up":      "Execute up sqls",
	"down":    "Execute down sqls",
	"version": "Just check current migrate version",
}

func main() {
	flag.Parse()
	if len(*Command) < 1 {
		fmt.Println("error: no argument")
		showUsageMessge()
		os.Exit(1)
		return
	}

	// dir, _ := os.Getwd()
	// if err := godotenv.Load(dir + "/cmd/migration/.env"); err != nil {
	// 	fmt.Println("error loading .env file")
	// 	os.Exit(1)
	// }

	m, err := migrate.New(Source, fmt.Sprintf("mysql://%s:%s@tcp(%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(os.Getenv("DB_NAME"))
	version, dirty, err := m.Version()
	showVersionInfo(version, dirty, err)

	fmt.Println("command: exec", *Command)
	applyQuery(m, version, dirty)
}

//exec up or down sqls
//with force option if needed
func applyQuery(m *migrate.Migrate, version uint, dirty bool) {
	if dirty && *Force {
		fmt.Println("force=true: force execute current version sql")
		m.Force(int(version))
	}

	var err error
	switch *Command {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "version":
		//do nothing
		return
	default:
		fmt.Println("error: invalid command '" + *Command)
		showUsageMessge()
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	} else {
		fmt.Println("success:", *Command+"\n")
		fmt.Println("updated version info")
		version, dirty, err := m.Version()
		showVersionInfo(version, dirty, err)
	}
}

func showUsageMessge() {
	fmt.Println("-------------------------------------")
	fmt.Println("Usage")
	fmt.Println("  go run migrate.go -exec <command>")
	fmt.Println("Available Exec Commands: ")
	for available_command, detail := range AvailableExecCommands {
		fmt.Println("  " + available_command + " : " + detail)
	}
	fmt.Println("-------------------------------------")
}

func showVersionInfo(version uint, dirty bool, err error) {
	fmt.Println("-------------------")
	fmt.Println("version  : ", version)
	fmt.Println("dirty    : ", dirty)
	fmt.Println("error    : ", err)
	fmt.Println("-------------------")
}
