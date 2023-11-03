package service

import (
	dataSource "databahn-api/data_sources"
	util "databahn-api/utils"
	"fmt"
	"log"
	"sync"
)

type LoadService struct {
	dataSource dataSource.DataSource
}

func NewLoadService(dataSource dataSource.DataSource) *LoadService {
	return &LoadService{
		dataSource: dataSource,
	}
}

func (ls *LoadService) LoadData(directoryName, templateFileName string, count int) error {
	filePath := fmt.Sprintf("%s/%s/%s", ".", directoryName, templateFileName)
	log.Printf("Filepath: %s, count: %d", filePath, count)
	templateContent, err := util.ReadTemplateFile(filePath)
	if err != nil {
		log.Printf("Error reading template file: count=%d, template=%s, error=%v\n", count, templateFileName, err)
		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, count)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			renderedContent, err := util.RenderTemplate(templateContent, nil)
			log.Printf("Render contents: %s", renderedContent)
			if err != nil {
				errChan <- err
				log.Printf("Error rendering template: count=%d, template=%s, error=%v\n", count, templateFileName, err)
				return
			}

			go func() {
				if err := ls.dataSource.PushData(renderedContent); err != nil {
					errChan <- err
					log.Printf("Error pushing data to data source: count=%d, template=%s, error=%v\n", count, templateFileName, err)
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
		errMsg := fmt.Sprintf("Encountered %d errors: count=%d, template=%s", len(errors), count, templateFileName)
		log.Println(errMsg)
		return fmt.Errorf(errMsg)
	}

	return nil
}
