package internal

var (
	BuildVersion     string
	BuildDate        string
	BuildHash        string
	BuildEnvironment string
)

type Version struct {
    BuildVersion string `json:"version"`
    BuildDate string `json:"build_date"`
    BuildHash string `json:"build_hash"`
    BuildEnvironment string `json:"build_env"`
}

func GetVersion() *Version {
    return &Version{
        BuildVersion: BuildVersion,
        BuildDate: BuildDate,
        BuildHash: BuildHash,
        BuildEnvironment: BuildEnvironment,
    }
}


