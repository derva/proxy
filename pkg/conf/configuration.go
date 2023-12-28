package conf

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading dotenv files!")
	}

	return os.Getenv(key)
}

func LoadConfFiles() map[string]string {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading dotenv files!")
	}

	CertificatePath := os.Getenv("CERTIFICATE_PATH")
	EncodingAlgo := os.Getenv("PREFERRED_ENCODING_ALGORITHM")

	return map[string]string{"certificate": CertificatePath, "encoding": EncodingAlgo}
}
