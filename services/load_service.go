package service

import (
	dataSource "databahn-api/data_sources"
	util "databahn-api/utils"
	"fmt"
	"sync"
)

type LoadService struct {
	dataSource dataSource.DataSource
}

func NewLoadService(dataSource data.DataSource) *LoadService {
	return &LoadService{
		dataSource: dataSource,
	}
}

func (ls *LoadService) LoadData(directoryName, templateFileName string, count int) error {
	filePath := fmt.Sprintf("%s/%s/%s", ".", directoryName, templateFileName)
	templateContent, err := util.ReadTemplateFile(filePath)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, count)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			renderedContent, err := util.RenderTemplate(templateContent, nil)
			if err != nil {
				errChan <- err
				return
			}

			go func() {
				if err := ls.dataSource.PushData(renderedContent); err != nil {
					errChan <- err
				}
			}()
		}()
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	var errors []error
	for err := range errChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return fmt.Errorf("Encountered %d errors", len(errors))
	}

	return nil
}
