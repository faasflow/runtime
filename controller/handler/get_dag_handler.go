package handler

import (
	"fmt"
	"github.com/faasflow/runtime"
	"github.com/faasflow/sdk/executor"
	"github.com/faasflow/sdk/exporter"
	"log"
)

func GetDagHandler(response *runtime.Response, request *runtime.Request, ex executor.Executor) error {
	log.Printf("Exporting DAG of flow: %s\n", request.FlowName)

	flowExporter := exporter.CreateFlowExporter(ex)
	resp, err := flowExporter.Export()
	if err != nil {
		return fmt.Errorf("failed to export dag, error %v", err)
	}

	response.Body = resp
	return nil
}
