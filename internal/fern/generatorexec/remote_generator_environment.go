package generatorexec

type RemoteGeneratorEnvironment struct {
	CoordinatorUrl   string `json:"coordinatorUrl"`
	CoordinatorUrlV2 string `json:"coordinatorUrlV2"`
	Id               TaskId `json:"id"`
}
