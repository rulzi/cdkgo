package environment

import (
	"cdkgo/lib/config"
	"cdkgo/lib/secretmanager"
	"cdkgo/lib/ssm"

	"github.com/aws/jsii-runtime-go"
)

// DefaultEnv create default environment
func DefaultEnv(props *config.StackProps) *map[string]*string {

	ssmsvc := ssm.NewSSMClient()
	sm := secretmanager.NewSecretManagerClient(*props.StackProps.Env.Region)

	domain, _ := ssmsvc.Param("/fortuneidn/web/publicDomainName", true).GetValue()

	mySQL := sm.Param(*jsii.String("staging-mysql8")).GetValueJSON()

	mySQLHost := mySQL.GetString("host")
	mySQLPassword := mySQL.GetString("password")
	mySQLPort := mySQL.GetInt("port")
	mySQLUser := mySQL.GetString("username")

	environments := &map[string]*string{
		"DB_MYSQL_HOST":     mySQLHost,
		"DB_MYSQL_PASSWORD": mySQLPassword,
		"DB_MYSQL_PORT":     mySQLPort,
		"DB_MYSQL_USER":     mySQLUser,
		"DOMAIN":            domain,
	}

	return environments
}

// MergeEnv merge env overwriting duplicate keys, you should handle that if there is a need
func MergeEnv(maps ...*map[string]*string) *map[string]*string {
	result := make(map[string]*string)
	for _, m := range maps {
		for k, v := range *m {
			result[k] = v
		}
	}
	return &result
}
