package aspera

type AuthSpec struct {
	RemoteHost     string
	SshPort        int
	RemoteUser     string
	RemotePassword string
	Token          string
}

type TransferSpec struct {
	AuthSpec
}
