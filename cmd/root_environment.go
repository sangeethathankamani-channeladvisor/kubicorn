// Copyright © 2017 The Kubicorn Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/kubicorn/kubicorn/pkg/cli"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	// Environment Variables non-prefixed
	envVarAwsProfile = "AWS_PROFILE"
	envVarEditor     = "EDITOR"

	// Environment Variables kubicorn prefixed
	envVarStateStore     = "KUBICORN_STATE_STORE"
	envVarStateStorePath = "KUBICORN_STATE_STORE_PATH"
	envVarSet            = "KUBICORN_SET"
	envVarMasterSet      = "KUBICORN_MASTER_SET"
	envVarNodeSet        = "KUBICORN_NODE_SET"
	envVarGitConfig      = "KUBICORN_GIT_CONFIG"
	envVarTrueColor      = "KUBICORN_TRUECOLOR"
	envVarProfile        = "KUBICORN_PROFILE"
	envVarCloudID        = "KUBICORN_CLOUDID"
	envVarName           = "KUBICORN_NAME"
	envVarKubicornOutput = "KUBICORN_OUTPUT"
	envVarS3AccessKey    = "KUBICORN_S3_ACCESS_KEY"
	envVarS3SecreteKey   = "KUBICORN_S3_SECRET_KEY"
	enVarS3Endpoint      = "KUBICORN_S3_ENDPOINT"
	envVarS3SSL          = "KUBICORN_S3_SSL"
	envVarS3Bucket       = "KUBICORN_S3_BUCKET"

	// --- keys for flags/EnvVars etc ---
	keyAwsProfile = "aws-profile"

	keyStateStore          = "state-store"
	keyStateStorePath      = "state-store-path"
	keyNoHeaders           = "no-headers"
	keyS3Access            = "s3-access"
	keyS3Secret            = "s3-secret"
	keyS3Endpoint          = "s3-endpoint"
	keyS3SSL               = "s3-ssl"
	keyS3Bucket            = "s3-bucket"
	keyPolicyAttachments   = "policy-attachment"
	keyKubicornSet         = "set"
	keyGitConfig           = "git-config"
	keyKubicornName        = "kubicorn-name"
	keyCloudID             = "cloudId"
	keyProfile             = "profile"
	keyKubeConfigLocalFile = "kubeconfig-local-file"
	keySet                 = "set"
	keyMasterSet           = "master-set"
	keyNodeSet             = "node-set"
	keyPurge               = "purge"
	keyEditor              = "editor"
	keyOutput              = "output"
	keyTrueColor           = "truecolor"

	// -- descriptions ---
	descStateStore          = "The state store type to use for the cluster."
	descStateStorePath      = "The state store path to use."
	descSet                 = "Set cluster level setting within the providerConfig"
	descMasterSet           = "Set master level setting within the providerConfig"
	descNodeSet             = "Set node level setting within the providerConfig"
	descAwsProfile          = "The profile to be used as defined in $HOME/.aws/credentials"
	descGitConfig           = "The git remote url to be used for saving the git state for the cluster."
	descS3AccessKey         = "The s3 access key."
	descS3SecretKey         = "The s3 secret key."
	descS3Endpoints         = "The s3 endpoint url."
	descS3SSL               = "The s3 bucket name to be used for saving the git state for the cluster."
	descS3Bucket            = "The s3 bucket name to be used for saving the git state for the cluster."
	descPolicyAttachments   = "Additional AWS IAM Policy ARNs to attach to the master and node roles."
	descCloudID             = "The cloud id."
	descProfile             = "The cluster profile to use."
	descKubeConfigLocalFile = "The path to the local file where to store the kubeconfig retrieved from the cluster."
	descPurge               = "Remove the API model from the state store after the resources are deleted."
	descEditor              = "The editor used to edit the state store."
	descOutput              = "Output format (currently only JSON supported)."
	desNoHeaders            = "Show the list containing names only."
)

func initEnvDefaults() {
	viper.SetDefault(keyAwsProfile, "")
	viper.SetDefault(keyEditor, "vi")

	viper.SetDefault(keyStateStore, "fs")
	viper.SetDefault(keyStateStorePath, "./_state")
	viper.SetDefault(keyKubicornSet, "")
	viper.SetDefault(keyGitConfig, "git")
	viper.SetDefault(keyTrueColor, "")
	viper.SetDefault(keyProfile, "google")
	viper.SetDefault(keyCloudID, "")
	viper.SetDefault(keyKubicornName, "")
	viper.SetDefault(keyOutput, "json")
	viper.SetDefault(keyS3Access, "")
	viper.SetDefault(keyS3Secret, "")
	viper.SetDefault(keyS3Endpoint, "")
	viper.SetDefault(keyS3SSL, true)
	viper.SetDefault(keyS3Bucket, "")

	viper.SetDefault(keySet, "")
	viper.SetDefault(keyMasterSet, "")
	viper.SetDefault(keyNodeSet, "")
	viper.SetDefault(keyPurge, false)
	viper.SetDefault(keyNoHeaders, false)

}

func bindEnvVars() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("kubicorn")

	viper.BindEnv(keyAwsProfile, envVarAwsProfile)
	viper.BindEnv(keyEditor, envVarEditor)
	viper.BindEnv(keyStateStore, envVarStateStore)
	viper.BindEnv(keyStateStorePath, envVarStateStorePath)
	viper.BindEnv(keyKubicornSet, envVarSet)
	viper.BindEnv(keyGitConfig, envVarGitConfig)
	viper.BindEnv(keyTrueColor, envVarTrueColor)
	viper.BindEnv(keyProfile, envVarProfile)
	viper.BindEnv(keyCloudID, envVarCloudID)
	viper.BindEnv(keyKubicornName, envVarName)
	viper.BindEnv(keyOutput, envVarKubicornOutput)
	viper.BindEnv(keyS3Access, envVarS3AccessKey)
	viper.BindEnv(keyS3Secret, envVarS3SecreteKey)
	viper.BindEnv(keyS3Endpoint, enVarS3Endpoint)
	viper.BindEnv(keyS3SSL, envVarS3SSL)
	viper.BindEnv(keyS3Bucket, envVarS3Bucket)
}

func bindCommonAwsFlags(o *cli.AwsOptions, fs *flag.FlagSet) {
	fs.StringVar(&o.S3AccessKey, keyS3Access, viper.GetString(keyS3Access), descS3AccessKey)
	fs.StringVar(&o.S3SecretKey, keyS3Secret, viper.GetString(keyS3Secret), descS3SecretKey)
	fs.StringVar(&o.BucketEndpointURL, keyS3Endpoint, viper.GetString(keyS3Endpoint), descS3Endpoints)
	fs.StringVar(&o.BucketName, keyS3Bucket, viper.GetString(keyS3Bucket), descS3Bucket)

	fs.BoolVar(&o.BucketSSL, keyS3SSL, viper.GetBool(keyS3SSL), descS3SSL)
}

func bindCommonStateStoreFlags(o *cli.StateStoreOptions, fs *flag.FlagSet) {
	fs.StringVarP(&o.StateStore, keyStateStore, "s", viper.GetString(keyStateStore), descStateStore)
	fs.StringVarP(&o.StateStorePath, keyStateStorePath, "S", viper.GetString(keyStateStorePath), descStateStorePath)
}
