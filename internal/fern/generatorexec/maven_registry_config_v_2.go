package generatorexec

type MavenRegistryConfigV2 struct {
	RegistryUrl string `json:"registryUrl"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Coordinate  string `json:"coordinate"`
}
