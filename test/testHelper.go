package testhelp

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type Cfdi struct {
	FechaEmision string `xml:"Fecha,attr"`
	Serie        string `xml:"Serie,attr"`
	Folio        string `xml:"Folio,attr"`
}
type BuildSettings struct {
	User     string
	Password string
	Token    string
}

func Init() *BuildSettings {
	_ = godotenv.Load("../../.env")
	settings := &BuildSettings{}
	settings.User = os.Getenv("SDK_USER")
	settings.Password = os.Getenv("SDK_PASSWORD")
	settings.Token = os.Getenv("SDK_TOKEN")
	return settings
}

func GetXml() string {
	byteXml, err := ioutil.ReadFile("../../test/resources/cfdi40.xml")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	serializedXml := &Cfdi{}
	err = xml.Unmarshal(byteXml, serializedXml)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	getXml := string(byteXml)
	setDate := time.Now().Add(-(time.Hour * 4)).Format("2006-01-02T15:04:05")
	setXml := strings.Replace(getXml, serializedXml.FechaEmision, setDate, 1)
	setXml = strings.Replace(setXml, serializedXml.Serie, "sw-sdk-go", 1)
	setXml = strings.Replace(setXml, serializedXml.Folio, uuid.New().String(), 1)
	return setXml
}
