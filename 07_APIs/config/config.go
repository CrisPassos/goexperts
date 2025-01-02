package configs

var cfg *conf

type conf struct {
	DBDriver      string
	DBHost        string
	DBPORT        string
	DBUser        string
	DBPassword    string
	DBName        string
	WebServerPort string
	JWTSecret     string
	JWTExperiesIn int
}

// vamos trabalhar com um arquivo .env
func LoadConfig(path string) (*conf, error) {

}
