package service

import (
	"errors"
	"testing"

	dataSources "databahn-api/data_sources"
	utils "databahn-api/utils"

	"github.com/stretchr/testify/mock"
)

func TestLoadDataSuccess(t *testing.T) {
	dataSource := &dataSources.MockDataSource{}
	utils := &utils.MockUtils{}
	ls := NewLoadService(dataSource, 10)
	readTemplateFile = utils.ReadTemplateFile
	renderTemplate = utils.RenderTemplate

	directoryName := "../templates"
	templateFileName := "timestamp_template.txt"
	count := 3

	utils.On("ReadTemplateFile", mock.Anything).Return("MockedTemplateContent", nil)
	utils.On("RenderTemplate", mock.Anything, mock.Anything).Return("MockedRenderedContent", nil)
	dataSource.On("PushData", mock.Anything).Return(nil)

	err := ls.LoadData(directoryName, templateFileName, count)
	if err != nil {
		t.Errorf("LoadData failed unexpectedly: %v", err)
	}

	wg.Wait()

	utils.AssertCalled(t, "ReadTemplateFile", mock.Anything)
	utils.AssertCalled(t, "RenderTemplate", mock.Anything, mock.Anything)
	dataSource.AssertExpectations(t)
	utils.AssertExpectations(t)
}

func TestLoadDataTemplateReadError(t *testing.T) {
	mockError := errors.New("template read error")
	dataSource := &dataSources.MockDataSource{}
	utils := &utils.MockUtils{}

	ls := NewLoadService(dataSource, 10)
	readTemplateFile = utils.ReadTemplateFile
	renderTemplate = utils.RenderTemplate

	utils.On("ReadTemplateFile", mock.Anything).Return("", mockError)

	directoryName := "../templates"
	templateFileName := "timestamp_template.txt"
	count := 1

	err := ls.LoadData(directoryName, templateFileName, count)
	if err != mockError {
		t.Errorf("Expected template read error, got: %v", err)
	}
	wg.Wait()

	utils.AssertCalled(t, "ReadTemplateFile", mock.Anything)
	dataSource.AssertExpectations(t)
	utils.AssertExpectations(t)
}

func TestLoadDataMaxCapacityExceeded(t *testing.T) {
	dataSource := &dataSources.MockDataSource{}
	utils := &utils.MockUtils{}
	ls := NewLoadService(dataSource, 1)

	readTemplateFile = utils.ReadTemplateFile
	renderTemplate = utils.RenderTemplate

	directoryName := "test_dir"
	templateFileName := "test_template"
	count := 2

	utils.On("ReadTemplateFile", mock.Anything).Return("MockedTemplateContent", nil)
	utils.On("RenderTemplate", mock.Anything, mock.Anything).Return("MockedRenderedContent", nil)
	dataSource.On("PushData", mock.Anything).Return(nil)

	err := ls.LoadData(directoryName, templateFileName, count)
	if err != nil {
		t.Errorf("LoadData failed unexpectedly: %v", err)
	}

	err = ls.LoadData(directoryName, templateFileName, count)
	if err == nil {
		t.Error("Expected max capacity exceeded error, but no error was returned")
	}
	wg.Wait()

	utils.AssertCalled(t, "ReadTemplateFile", mock.Anything)
	utils.AssertCalled(t, "RenderTemplate", mock.Anything, mock.Anything)
	dataSource.AssertExpectations(t)
	utils.AssertExpectations(t)
}
