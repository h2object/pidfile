pidfile package
---

# Quick Start

````
	import "github.com/h2object/pidfile"

	// create pid file 
	pid, err := pidfile.New("h2object.pid")
	if err != nil {
		//....
	}


	// stop pid process
	if err := pid.Kill(); err != nil {
		//....
	}

````