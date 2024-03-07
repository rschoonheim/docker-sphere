package deployment

import "context"

type DeploymentService struct {
	UnimplementedDeploymentServiceServer
}

func (s *DeploymentService) CreateDeployment(ctx context.Context, in *CreateDeploymentRequest) (*CreateDeploymentResponse, error) {
	println("Create deployment")
	return &CreateDeploymentResponse{}, nil
}

func (s *DeploymentService) GetDeployment(ctx context.Context, in *GetDeploymentRequest) (*GetDeploymentResponse, error) {
	return &GetDeploymentResponse{}, nil
}

func (s *DeploymentService) ListDeployments(ctx context.Context, in *ListDeploymentsRequest) (*ListDeploymentsResponse, error) {
	return &ListDeploymentsResponse{}, nil
}
