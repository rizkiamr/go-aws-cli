package model

import "github.com/aws/aws-sdk-go-v2/service/iam/types"

type SimulatePrincipalPolicyResponse struct {
	EvaluationResults []types.EvaluationResult `json:"EvaluationResults"`
}
