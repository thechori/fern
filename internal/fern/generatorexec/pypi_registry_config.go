package generatorexec

type PypiRegistryConfig struct {
	RegistryUrl string `json:"registryUrl"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	PackageName string `json:"packageName"`
}
