package models

type ConfigDBDriver struct {
	Host     string `yaml:"host" validate:"required"`
	Port     int    `yaml:"port"`
	Username string `yaml:"user" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	//InitialSchemaFile string `yaml:"initial_schema_file"`
}
