package secretmanager

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// SecretsManager is a SecretsManager API client.
type SecretsManager struct {
	secret *secretsmanager.SecretsManager
}

// Sessions create session
func Sessions() (*session.Session, error) {
	sess, err := session.NewSession()
	svc := session.Must(sess, err)
	return svc, err
}

// NewSecretManagerClient create client Secret Manager
func NewSecretManagerClient(region string) *SecretsManager {
	sess := session.Must(session.NewSession())

	svc := secretsmanager.New(sess, aws.NewConfig().WithRegion(region))

	ssvc := &SecretsManager{svc}

	return ssvc
}

// Param struct param secret manager
type Param struct {
	SecretName string
	svc        *SecretsManager
}

// ResultJSON struct result json secret manager
type ResultJSON struct {
	jsonResult map[string]json.RawMessage
}

// Param creates the struct for querying the param store
func (s *SecretsManager) Param(name string) *Param {
	return &Param{
		SecretName: name,
		svc:        s,
	}
}

// GetValue get value secret manager
func (p *Param) GetValue() *secretsmanager.GetSecretValueOutput {
	result, err := p.svc.secret.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: &p.SecretName})
	if err != nil {
		log.Fatal(err.Error())
	}

	return result
}

// GetValueJSON get value secret manager
func (p *Param) GetValueJSON() *ResultJSON {
	result, err := p.svc.secret.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: &p.SecretName})
	if err != nil {
		log.Fatal(err.Error())
	}

	var objmap map[string]json.RawMessage
	json.Unmarshal([]byte(*result.SecretString), &objmap)

	return &ResultJSON{
		jsonResult: objmap,
	}
}

// GetString get value string from json secret manager
func (r *ResultJSON) GetString(param string) *string {
	var res string
	json.Unmarshal(r.jsonResult[param], &res)

	return &res
}

// GetInt get value int from json secret manager
func (r *ResultJSON) GetInt(param string) *string {
	var res int
	json.Unmarshal(r.jsonResult[param], &res)

	resString := strconv.Itoa(res)

	return &resString
}
