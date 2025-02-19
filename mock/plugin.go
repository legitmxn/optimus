package mock

import (
	"context"

	"github.com/odpf/optimus/models"
	"github.com/stretchr/testify/mock"
)

type SupportedPluginRepo struct {
	mock.Mock
}

func (repo *SupportedPluginRepo) Add(plugin models.BasePlugin, mod models.CommandLineMod, mod2 models.DependencyResolverMod) error {
	return repo.Called(plugin, mod, mod2).Error(0)
}

func (repo *SupportedPluginRepo) GetByName(s string) (*models.Plugin, error) {
	args := repo.Called(s)
	return args.Get(0).(*models.Plugin), args.Error(1)
}

func (repo *SupportedPluginRepo) GetAll() []*models.Plugin {
	args := repo.Called()
	return args.Get(0).([]*models.Plugin)
}

func (repo *SupportedPluginRepo) GetTasks() []*models.Plugin {
	panic("implement me")
}

func (repo *SupportedPluginRepo) GetHooks() []*models.Plugin {
	panic("implement me")
}

func (repo *SupportedPluginRepo) GetCommandLines() []models.CommandLineMod {
	panic("implement me")
}

func (repo *SupportedPluginRepo) GetDependencyResolvers() []models.DependencyResolverMod {
	panic("implement me")
}

type BasePlugin struct {
	mock.Mock `hash:"-"`
}

func (repo *BasePlugin) PluginInfo() (*models.PluginInfoResponse, error) {
	args := repo.Called()
	return args.Get(0).(*models.PluginInfoResponse), args.Error(1)
}

type CLIMod struct {
	mock.Mock `hash:"-"`
}

func (repo *CLIMod) PluginInfo() (*models.PluginInfoResponse, error) {
	args := repo.Called()
	return args.Get(0).(*models.PluginInfoResponse), args.Error(1)
}

func (repo *CLIMod) DefaultConfig(ctx context.Context, inp models.DefaultConfigRequest) (*models.DefaultConfigResponse, error) {
	args := repo.Called(ctx, inp)
	return args.Get(0).(*models.DefaultConfigResponse), args.Error(1)
}

func (repo *CLIMod) DefaultAssets(ctx context.Context, inp models.DefaultAssetsRequest) (*models.DefaultAssetsResponse, error) {
	args := repo.Called(ctx, inp)
	return args.Get(0).(*models.DefaultAssetsResponse), args.Error(1)
}

func (repo *CLIMod) CompileAssets(ctx context.Context, inp models.CompileAssetsRequest) (*models.CompileAssetsResponse, error) {
	args := repo.Called(ctx, inp)
	return args.Get(0).(*models.CompileAssetsResponse), args.Error(1)
}

func (repo *CLIMod) GetQuestions(ctx context.Context, inp models.GetQuestionsRequest) (*models.GetQuestionsResponse, error) {
	args := repo.Called(ctx, inp)
	return args.Get(0).(*models.GetQuestionsResponse), args.Error(1)
}

func (repo *CLIMod) ValidateQuestion(ctx context.Context, inp models.ValidateQuestionRequest) (*models.ValidateQuestionResponse, error) {
	args := repo.Called(ctx, inp)
	return args.Get(0).(*models.ValidateQuestionResponse), args.Error(1)
}

type DependencyResolverMod struct {
	mock.Mock `hash:"-"`
}

func (repo *DependencyResolverMod) PluginInfo() (*models.PluginInfoResponse, error) {
	args := repo.Called()
	return args.Get(0).(*models.PluginInfoResponse), args.Error(1)
}

func (repo *DependencyResolverMod) GenerateDestination(ctx context.Context, inp models.GenerateDestinationRequest) (*models.GenerateDestinationResponse, error) {
	args := repo.Called(ctx, inp)
	return args.Get(0).(*models.GenerateDestinationResponse), args.Error(1)
}

func (repo *DependencyResolverMod) GenerateDependencies(ctx context.Context, inp models.GenerateDependenciesRequest) (*models.GenerateDependenciesResponse, error) {
	args := repo.Called(ctx, inp)
	return args.Get(0).(*models.GenerateDependenciesResponse), args.Error(1)
}
