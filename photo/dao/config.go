package dao

type Config struct {
	session *aws.Config
}

func getConfig() {
	switch env := os.Getenv("UP_STAGE"); env {
	case "production":
		return getProd()
	case "staging":
		return getTest()
	case "test":
		return getTest()
	default:
		return getDev()
	}
}

func getDev() Config {
	return Config{
		session: &aws.Config{
			Region:   aws.String("eu-west-3"),
			Endpoint: aws.String("http://localhost:8000")
		}
	}
}

func getStaging() Config {
	return Config{
		session: &aws.Config{
			Region:   aws.String("eu-west-3")
		}
	}
}

func getProd() Config {
	return Config{
		session: &aws.Config{
			Region:   aws.String("eu-west-3")
		}
	}
}

func getTest() Config {
	return Config{
		session: &aws.Config{
			Region:   aws.String("eu-west-3"),
			Endpoint: aws.String("http://localhost:8000")
		}
	}
}
