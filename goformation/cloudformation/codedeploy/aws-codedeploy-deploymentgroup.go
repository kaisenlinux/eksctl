package codedeploy

import (
	"bytes"
	"encoding/json"
	"fmt"

	"goformation/v4/cloudformation/types"

	"goformation/v4/cloudformation/policies"
)

// DeploymentGroup AWS CloudFormation Resource (AWS::CodeDeploy::DeploymentGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html
type DeploymentGroup struct {

	// AlarmConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-alarmconfiguration
	AlarmConfiguration *DeploymentGroup_AlarmConfiguration `json:"AlarmConfiguration,omitempty"`

	// ApplicationName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-applicationname
	ApplicationName *types.Value `json:"ApplicationName,omitempty"`

	// AutoRollbackConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-autorollbackconfiguration
	AutoRollbackConfiguration *DeploymentGroup_AutoRollbackConfiguration `json:"AutoRollbackConfiguration,omitempty"`

	// AutoScalingGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-autoscalinggroups
	AutoScalingGroups *types.Value `json:"AutoScalingGroups,omitempty"`

	// BlueGreenDeploymentConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-bluegreendeploymentconfiguration
	BlueGreenDeploymentConfiguration *DeploymentGroup_BlueGreenDeploymentConfiguration `json:"BlueGreenDeploymentConfiguration,omitempty"`

	// Deployment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-deployment
	Deployment *DeploymentGroup_Deployment `json:"Deployment,omitempty"`

	// DeploymentConfigName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-deploymentconfigname
	DeploymentConfigName *types.Value `json:"DeploymentConfigName,omitempty"`

	// DeploymentGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-deploymentgroupname
	DeploymentGroupName *types.Value `json:"DeploymentGroupName,omitempty"`

	// DeploymentStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-deploymentstyle
	DeploymentStyle *DeploymentGroup_DeploymentStyle `json:"DeploymentStyle,omitempty"`

	// ECSServices AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-ecsservices
	ECSServices []DeploymentGroup_ECSService `json:"ECSServices,omitempty"`

	// Ec2TagFilters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-ec2tagfilters
	Ec2TagFilters []DeploymentGroup_EC2TagFilter `json:"Ec2TagFilters,omitempty"`

	// Ec2TagSet AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-ec2tagset
	Ec2TagSet *DeploymentGroup_EC2TagSet `json:"Ec2TagSet,omitempty"`

	// LoadBalancerInfo AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-loadbalancerinfo
	LoadBalancerInfo *DeploymentGroup_LoadBalancerInfo `json:"LoadBalancerInfo,omitempty"`

	// OnPremisesInstanceTagFilters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-onpremisesinstancetagfilters
	OnPremisesInstanceTagFilters []DeploymentGroup_TagFilter `json:"OnPremisesInstanceTagFilters,omitempty"`

	// OnPremisesTagSet AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-onpremisestagset
	OnPremisesTagSet *DeploymentGroup_OnPremisesTagSet `json:"OnPremisesTagSet,omitempty"`

	// ServiceRoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-servicerolearn
	ServiceRoleArn *types.Value `json:"ServiceRoleArn,omitempty"`

	// TriggerConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-triggerconfigurations
	TriggerConfigurations []DeploymentGroup_TriggerConfig `json:"TriggerConfigurations,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *DeploymentGroup) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r DeploymentGroup) MarshalJSON() ([]byte, error) {
	type Properties DeploymentGroup
	return json.Marshal(&struct {
		Type                string
		Properties          Properties
		DependsOn           []string                     `json:"DependsOn,omitempty"`
		Metadata            map[string]interface{}       `json:"Metadata,omitempty"`
		DeletionPolicy      policies.DeletionPolicy      `json:"DeletionPolicy,omitempty"`
		UpdateReplacePolicy policies.UpdateReplacePolicy `json:"UpdateReplacePolicy,omitempty"`
		Condition           string                       `json:"Condition,omitempty"`
	}{
		Type:                r.AWSCloudFormationType(),
		Properties:          (Properties)(r),
		DependsOn:           r.AWSCloudFormationDependsOn,
		Metadata:            r.AWSCloudFormationMetadata,
		DeletionPolicy:      r.AWSCloudFormationDeletionPolicy,
		UpdateReplacePolicy: r.AWSCloudFormationUpdateReplacePolicy,
		Condition:           r.AWSCloudFormationCondition,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *DeploymentGroup) UnmarshalJSON(b []byte) error {
	type Properties DeploymentGroup
	res := &struct {
		Type                string
		Properties          *Properties
		DependsOn           []string
		Metadata            map[string]interface{}
		DeletionPolicy      string
		UpdateReplacePolicy string
		Condition           string
	}{}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields() // Force error if unknown field is found

	if err := dec.Decode(&res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = DeploymentGroup(*res.Properties)
	}
	if res.DependsOn != nil {
		r.AWSCloudFormationDependsOn = res.DependsOn
	}
	if res.Metadata != nil {
		r.AWSCloudFormationMetadata = res.Metadata
	}
	if res.DeletionPolicy != "" {
		r.AWSCloudFormationDeletionPolicy = policies.DeletionPolicy(res.DeletionPolicy)
	}
	if res.UpdateReplacePolicy != "" {
		r.AWSCloudFormationUpdateReplacePolicy = policies.UpdateReplacePolicy(res.UpdateReplacePolicy)
	}
	if res.Condition != "" {
		r.AWSCloudFormationCondition = res.Condition
	}
	return nil
}
