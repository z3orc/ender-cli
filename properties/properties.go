package properties

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/z3orc/ender-cli/logger"
	"gopkg.in/yaml.v3"
)

var p *ServerProperties

type ParseError struct {
	key string
}

func init() {
	properties, err := New("./testing/server.properties")
	if err != nil {
		logger.Error.Fatalln(err)
	}
	p = properties
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("could not parse properties file, %s has wrong formatting", e.key)
}

type ServerProperties struct {
	OtherProperties map[string]any
}

func New(path string) (*ServerProperties, error) {
	properties := &ServerProperties{}
	otherProperties := make(map[string]any)

	reader, err := os.Open(path)
	if err != nil {
		return nil, errors.New("could not find properties file")
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		list := strings.Split(line, "=")
		if len(list) < 2 {
			continue
		}

		key := list[0]
		value := list[1]

		parsedValue, err := parseValue(value)
		if err != nil {
			log.Fatalln(err)
		}
		otherProperties[key] = parsedValue
	}

	properties.OtherProperties = otherProperties
	return properties, nil
}

func Set(key string, value any) error {
	if _, ok := p.OtherProperties[key]; ok {
		p.OtherProperties[key] = value
		return nil
	} else {
		return fmt.Errorf("could not find key: %s", key)
	}
}

func Get(key string) (string, error) {
	if value, ok := p.OtherProperties[key]; ok {
		return fmt.Sprintf("%+v", value), nil
	} else {
		return "", fmt.Errorf("could not find key: %s", key)
	}
}

func GetUnformatted(key string) (any, error) {
	if value, ok := p.OtherProperties[key]; ok {
		return value, nil
	} else {
		return "", fmt.Errorf("could not find key: %s", key)
	}
}
func Save(path string) error {

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for key, value := range p.OtherProperties {
		writer.WriteString(fmt.Sprintf("%s=%+v\n", key, value))

	}

	writer.Flush()

	return nil
}

func SaveToYaml(path string) error {
	yamlData, err := yaml.Marshal(p.OtherProperties)
	if err != nil {
		log.Fatalln("could not save properties to file")
	}

	err = os.WriteFile(path, yamlData, 0644)
	if err != nil {
		panic("Unable to write data into the file")
	}
	fmt.Println(string(yamlData))
	return nil
}

func parseValue(value string) (any, error) {
	boolValue, err := strconv.ParseBool(value)
	if err == nil {
		return boolValue, nil
	}

	intValue, err := strconv.ParseInt(value, 10, 0)
	if err == nil {
		return intValue, nil
	}

	return value, nil
}
