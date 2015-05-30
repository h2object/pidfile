pidfile package
---

# Quick Start

````
	import "github.com/h2object/pidfile"

	// insert into the process begin 
	pid, err := pidfile.New("h2object.pid")
	if err != nil {
		//....
	}
	defer pid.Kill()

````