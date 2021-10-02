package cmd

import (
	"flag"
)

// Config will be the holder for our flags
type Config struct {
	subject      string
	isAwesome    bool
	howAwesome   int
}

// Setup initializes a config from flags that
// are passed in
func (c *Config) Setup() {
	// you can set a flag directly like so:
	// var someVar = flag.String("flag_name", "default_val", "description")
	// but in practice putting it in a struct is generally better

	// longhand
	flag.StringVar(&c.subject, "subject", "", "subject is a string, it defaults to empty")
	// shorthand
	flag.StringVar(&c.subject, "s", "", "subject is a string, it defaults to empty (shorthand)")

	flag.BoolVar(&c.isAwesome, "isawesome", false, "is it awesome or what?")
	flag.IntVar(&c.howAwesome, "howawesome", 10, "how awesome out of 10?")

	flag.Parse()
}

