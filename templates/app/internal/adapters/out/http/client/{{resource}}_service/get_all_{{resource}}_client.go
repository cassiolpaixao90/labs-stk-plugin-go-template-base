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

type GetAll{{capitalize_resource}}Client struct {
	client helpers.HttpClientRequest
}

func NewGetAll{{capitalize_resource}}Client() out.GetAll{{capitalize_resource}}ClientPort {
	return &GetAll{{capitalize_resource}}Client{}
}

func (g *GetAll{{capitalize_resource}}Client) GetAll{{capitalize_resource}}(ctx context.Context, id string) (*domain.{{capitalize_resource}}Domain, error) {

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

	var response *dtos.GetAll{{capitalize_resource}}Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.ToDomain(), nil
}
