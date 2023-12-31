package {{resource}}_service
 
import (
	"context"
	"encoding/json"
	"io"
	"{{project_name}}/internal/adapters/out/http/client/{{resource}}_service/dtos"
	"{{project_name}}/internal/adapters/out/http/helpers"
	"{{project_name}}/internal/application/domain"
	"{{project_name}}/internal/application/ports/out"
	"net/http"
)

type GetAll{{custom_resource}}Client struct {
	client helpers.HttpClientRequest
}

func NewGetAll{{custom_resource}}Client() out.GetAll{{custom_resource}}ClientPort {
	return &GetAll{{custom_resource}}Client{}
}

func (g *GetAll{{custom_resource}}Client) GetAll{{custom_resource}}(ctx context.Context, id string) (*domain.{{custom_resource}}Domain, error) {

	req, _ := http.NewRequest(http.MethodGet, "url", nil)

	err, resp := g.client.Retry(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// TODO add logger
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)

	var response *dtos.GetAll{{custom_resource}}Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.ToDomain(), nil
}
