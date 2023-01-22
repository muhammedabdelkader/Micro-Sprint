package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/google/go-github/v32/github"
	"github.com/spf13/viper"
	"gopkg.in/src-d/go-git.v4"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

/*
This script includes code to fetch the list of secrets for each platform using the respective SDKs or API calls. It also returns an error if any operation failed.
*/
type SecretLister interface {
	ListSecrets() ([]string, error)
}

type Buildkite struct {
	AccessToken  string
	Organization string
}

func (b *Buildkite) ListSecrets() ([]string, error) {
	// code to list secrets in Buildkite
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.buildkite.com/v2/organizations/"+b.Organization+"/secrets", nil)
	req.Header.Add("Authorization", "Bearer "+b.AccessToken)
	if err != nil {
		return nil, err

	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}
	var data map[string][]interface{}
	json.Unmarshal(body, &data)
	secrets := []string{}
	for _, secret := range data["secrets"] {
		name := secret.(map[string]interface{})["name"].(string)
		secrets = append(secrets, name)

	}
	return secrets, nil

}

type CircleCI struct {
	AccessToken string
	Username    string
}

func (c *CircleCI) ListSecrets() ([]string, error) {
	// code to list secrets in CircleCI
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://circleci.com/api/v2/project/github/"+c.Username+"/settings/env-vars", nil)
	req.Header.Add("Circle-Token", c.AccessToken)
	if err != nil {
		return nil, err

	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}
	var data map[string][]interface{}
	json.Unmarshal(body, &data)
	secrets := []string{}
	for _, secret := range data["items"] {
		name := secret.(map[string]interface{})["name"].(string)
		secrets = append(secrets, name)

	}
	return secrets, nil

}

type Jenkins struct {
	URL      string
	Username string
	Password string
}

func (j *Jenkins) ListSecrets() ([]string, error) {
	// code to list secrets in Jenkins
	client := &http.Client{}
	req, err := http.NewRequest("GET", j.URL+"/credentials/store/system/domain/_/credentials/", nil)
	req.SetBasicAuth(j.Username, j.Password)
	if err != nil {
		return nil, err

	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	credentials := data["credentials"].([]interface{})
	secrets := []string{}
	for _, cred := range credentials {
		secret := cred.(map[string]interface{})["description"].(string)
		secrets = append(secrets, secret)

	}
	return secrets, nil

}

type Bitbucket struct {
	Username string
	Password string
}

func (b *Bitbucket) ListSecrets() ([]string, error) {
	// code to list secrets in Bitbucket
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.bitbucket.org/2.0/repositories/OWNER/REPO/pipelines_config/variables", nil)
	req.SetBasicAuth(b.Username, b.Password)
	if err != nil {
		return nil, err

	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	values := data["values"].([]interface{})
	secrets := []string{}
	for _, val := range values {
		secret := val.(map[string]interface{})["key"].(string)
		secrets = append(secrets, secret)

	}
	return secrets, nil

}

type GitHub struct {
	Token  string
	client *github.Client
}

func (g *GitHub) ListSecrets() ([]string, error) {
	// code to list secrets in GitHub
	g.client = github.NewClient(nil)
	g.client.SetAuthToken(g.Token)
	secrets := []string{}
	repos, _, err := g.client.Repositories.List("", nil)
	if err != nil {
		return nil, err

	}
	for _, repo := range repos {
		name := *repo.Name
		secret, _, err := g.client.Actions.ListRepoSecrets(context.Background(), "OWNER", name, nil)
		if err != nil {
			return nil, err

		}
		for _, s := range secret {
			secrets = append(secrets, *s.Name)

		}

	}
	return secrets, nil

}

type GitLab struct {
	Token  string
	client *git.Client
}

func (g *GitLab) ListSecrets() ([]string, error) {
	// code to list secrets in GitLab
	g.client = git.NewClient(nil)
	g.client.SetAuthToken(g.Token)
	secrets := []string{}
	// list all gitlab projects
	projects, err := g.client.Projects.List()
	if err != nil {
		return nil, err

	}
	for _, p := range projects {
		secret, err := g.client.Secrets.List(p.ID, "")
		if err != nil {
			return nil, err

		}
		for _, s := range secret {
			secrets = append(secrets, s.Name)

		}

	}
	return secrets, nil

}

type Kubernetes struct {
	Cluster string
	client  *kubernetes.Clientset
}

func (k *Kubernetes) ListSecrets() ([]string, error) {
	// code to list secrets in Kubernetes
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err

	}
	k.client, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err

	}
	secrets, err := k.client.CoreV1().Secrets(k.Cluster).List(metav1.ListOptions{})
	if err != nil {
		return nil, err

	}
	secretList := []string{}
	for _, secret := range secrets.Items {
		secretList = append(secretList, secret.Name)

	}
	return secretList, nil

}

type ECS struct {
	Cluster string
	client  *ecs.ECS
}

func (e *ECS) ListSecrets() ([]string, error) {
	// code to list secrets in ECS
	secrets := []string{}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	e.client = ecs.New(sess, &aws.Config{Region: aws.String(e.Cluster)})

	input := &ecs.ListServicesInput{
		Cluster: aws.String(e.Cluster),
	}

	result, err := e.client.ListServices(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())

			}

		} else {
			// Print the error, cast err to awserr.Error to
			// get the Code and
			// Message from an error.
			fmt.Println(err.Error())

		}
		return nil, err

	}

	for _, service := range result.Services {
		secrets = append(secrets, *service.ServiceName)

	}
	return secrets, nil

}

type AKS struct {
	Cluster string
	client  *azure.AKS
}

func (a *AKS) ListSecrets() ([]string, error) {
	// code to list secrets in AKS
	secrets := []string{}
	a.client = azure.NewAKS(a.Cluster)
	secretList, err := a.client.ListSecrets()
	if err != nil {
		return nil, err

	}
	for _, secret := range secretList {
		secrets = append(secrets, secret.Name)

	}
	return secrets, nil

}

type GKE struct {
	Cluster string
	client  *gke.Client
}

func (g *GKE) ListSecrets() ([]string, error) {
	// code to list secrets in GKE
	secrets := []string{}
	g.client = gke.NewClient(g.Cluster)
	secretList, err := g.client.ListSecrets()
	if err != nil {
		return nil, err

	}
	for _, secret := range secretList {
		secrets = append(secrets, secret.Name)

	}
	return secrets, nil

}

type ACS struct {
	Cluster string
	client  *acs.Client
}

func (a *ACS) ListSecrets() ([]string, error) {
	// code to list secrets in ACS
	secrets := []string{}
	a.client = acs.NewClient(a.Cluster)
	secretList, err := a.client.ListSecrets()
	if err != nil {
		return nil, err

	}
	for _, secret := range secretList {
		secrets = append(secrets, secret.Name)

	}
	return secrets, nil

}

type AWS struct {
	Region string
	client *secretsmanager.SecretsManager
}

func (a *AWS) ListSecrets() ([]string, error) {
	// code to list secrets in AWS
	secrets := []string{}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	a.client = secretsmanager.New(sess, &aws.Config{Region: aws.String(a.Region)})
	input := &secretsmanager.ListSecretsInput{}
	result, err := a.client.ListSecrets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())

			}

		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())

		}
		return nil, err

	}
	for _, secret := range result.Secrets {
		secrets = append(secrets, *secret.Name)

	}
	return secrets, nil

}

func parseConfig(configFile string) {
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return

	}
	jenkinsConf := viper.Sub("jenkins")
	bitbucketConf := viper.Sub("bitbucket")
	buildkiteConf := viper.Sub("buildkite")
	circleCiConf := viper.Sub("circleci")
	githubConf := viper.Sub("github")
	gitlabConf := viper.Sub("gitlab")
	kubernetesConf := viper.Sub("kubernetes")
	ecsConf := viper.Sub("ecs")
	aksConf := viper.Sub("aks")
	gkeConf := viper.Sub("gke")
	acsConf := viper.Sub("acs")
	awsConf := viper.Sub("aws")

}

func main() {
	/*Config parser */
	var configFile string
	flag.StringVar(&configFile, "config", "", "./config.yaml")
	flag.Parse()

	parseConfig(configFile)

	jenkins := &Jenkins{
		URL:      jenkinsConf.GetString("url"),
		Username: jenkinsConf.GetString("username"),
		Password: jenkinsConf.GetString("password"),
	}
	bitbucket := &Bitbucket{
		Username: bitbucketConf.GetString("username"),
		Password: bitbucketConf.GetString("password"),
	}
	buildkite := &Buildkite{
		AccessToken:  buildkiteConf.GetString("access_token"),
		Organization: buildkiteConf.GetString("organization"),
	}
	circleci := &CircleCI{
		AccessToken: circleCiConf.GetString("access_token"),
		Username:    circleCiConf.GetString("username"),
	}
	github := &Github{
		Token:        githubConf.GetString("token"),
		Organization: githubConf.GetString("organization"),
	}
	gitlab := &Gitlab{
		URL:   gitlabConf.GetString("url"),
		Token: gitlabConf.GetString("token"),
	}
	kubernetes := &Kubernetes{
		URL:   kubernetesConf.GetString("url"),
		Token: kubernetesConf.GetString("token"),
	}
	ecs := &ECS{
		AccessKey: ecsConf.GetString("access_key"),
		SecretKey: ecsConf.GetString("secret_key"),
	}
	aks := &AKS{
		SubscriptionID: aksConf.GetString("subscription_id"),
		ClientID:       aksConf.GetString("client_id"),
		ClientSecret:   aksConf.GetString("client_secret"),
		TenantID:       aksConf.GetString("tenant_id"),
	}
	gke := &GKE{
		ProjectID:          gkeConf.GetString("project_id"),
		ServiceAccountFile: gkeConf.GetString("service_account_file"),
	}
	acs := &ACS{
		SubscriptionID: acsConf.GetString("subscription_id"),
		ClientID:       acsConf.GetString("client_id"),
		ClientSecret:   acsConf.GetString("client_secret"),
		TenantID:       acsConf.GetString("tenant_id"),
	}
	aws := &AWS{
		AccessKey: awsConf.GetString("access_key"),
		SecretKey: awsConf.GetString("secret_key"),
		Region:    awsConf.GetString("region"),
	}

	/**/
	/*
		platforms := []SecretLister{
			&GitHub{Token: "abc123"},
			&GitLab{Token: "def456"},
			&Kubernetes{Cluster: "cluster1"},
			&ECS{Cluster: "cluster2"},
			&AKS{Cluster: "cluster3"},
			&GKE{Cluster: "cluster4"},
			&ACS{Cluster: "cluster5"},
			&AWS{Region: "us-east-1"},
			&Jenkins{URL: "http://jenkins.example.com", Username: "admin", Password: "password"},
			&Bitbucket{Username: "username", Password: "password"},
			&Buildkite{AccessToken: "abc123", Organization: "myorg"},
			&CircleCI{AccessToken: "def456", Username: "myusername"},
		}
	*/
	for _, platform := range platforms {
		secrets, err := platform.ListSecrets()
		if err != nil {
			fmt.Println("Error fetching secrets from", platform, ":", err)
			continue

		}
		fmt.Println("Secrets in", platform, ":")
		for _, secret := range secrets {
			fmt.Println("-", secret)

		}

	}
}
